package kafka

import (
	"encoding/json"
	"github.com/GOAT-prod/goatcontext"
	"request-service/database"
	"request-service/repository"
)

type MessageHandler interface {
	HandleMessage(ctx goatcontext.Context, message []byte) error
}

type Impl struct {
	requestRepository repository.Request
}

func NewMessageHandler(requestRepository repository.Request) MessageHandler {
	return &Impl{
		requestRepository: requestRepository,
	}
}

func (h *Impl) HandleMessage(ctx goatcontext.Context, message []byte) error {
	var request database.Request
	if err := json.Unmarshal(message, &request); err != nil {
		return err
	}

	if err := h.requestRepository.AddRequest(ctx, request); err != nil {
		return err
	}

	return nil
}
