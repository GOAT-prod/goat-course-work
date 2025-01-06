package kafka

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goatlogger"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	_sessionTimeout     = 7000 // ms
	_autoCommitInterval = 5000
	_consumeTimeout     = -1
)

type Consumer struct {
	consumer       *kafka.Consumer
	messageHandler MessageHandler
	stop           bool
}

func NewConsumer(handler MessageHandler, address, topic string) (*Consumer, error) {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":        address,
		"group.id":                 "requests",
		"session.timeout.ms":       _sessionTimeout,
		"enable.auto.offset.store": false,
		"enable.auto.commit":       true,
		"auto.commit.interval.ms":  _autoCommitInterval,
	}

	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	if err = c.Subscribe(topic, nil); err != nil {
		return nil, err
	}

	return &Consumer{consumer: c, messageHandler: handler}, nil
}

func (c *Consumer) Consume(ctx goatcontext.Context, logger goatlogger.Logger) {
	for {
		if c.stop {
			break
		}

		kafkaMsg, err := c.consumer.ReadMessage(_consumeTimeout)
		if err != nil {
			logger.Error(fmt.Errorf("не удалось прочитать сообщение из кафки: %w", err).Error())
			continue
		}

		if kafkaMsg == nil {
			continue
		}

		if err = c.messageHandler.HandleMessage(ctx, kafkaMsg.Value); err != nil {
			logger.Error(fmt.Errorf("не удалось обработать сообщение: %w", err).Error())
			continue
		}

		if _, err = c.consumer.StoreMessage(kafkaMsg); err != nil {
			logger.Error(fmt.Errorf("не удалось сохранить информацию о прочтении: %w", err).Error())
			continue
		}
	}
}

func (c *Consumer) Stop() error {
	c.stop = true
	return c.consumer.Close()
}
