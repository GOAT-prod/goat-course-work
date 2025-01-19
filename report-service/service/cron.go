package service

import (
	"context"
	"github.com/GOAT-prod/goatcontext"
	"github.com/robfig/cron"
	"log"
	"report-service/cluster/order"
	"report-service/repository"
	"report-service/settings"
	"time"
)

type CronService struct {
	cron *cron.Cron
}

type refreshJob struct {
	orderService     *order.Client
	reportRepository repository.Report
	scheduleChannel  chan time.Time
	settings         settings.Config
}

func (r *refreshJob) Run() {
	var expires time.Time

	defer func() {
		r.scheduleChannel <- expires
	}()

	expires = time.Now().Add(10 * time.Minute)

	go func() {
		items, err := r.orderService.GetLatestOrders(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		if err = r.reportRepository.AddReportItems(goatcontext.Context{}, items); err != nil {
			log.Println(err)
			return
		}
	}()
}

type refreshSchedule struct {
	channel     chan time.Time
	initialized bool
}

func (r *refreshSchedule) Next(_ time.Time) time.Time {
	if !r.initialized {
		r.initialized = true
		return time.Now()
	}

	return <-r.channel
}

func NewCronService(order *order.Client, repository repository.Report, settings settings.Config) *CronService {
	c := cron.New()

	ch := make(chan time.Time)
	c.Schedule(&refreshSchedule{
		channel: ch,
	}, &refreshJob{
		orderService:     order,
		reportRepository: repository,
		scheduleChannel:  ch,
		settings:         settings,
	})

	return &CronService{
		cron: c,
	}
}

func (r *CronService) Run() {
	go r.cron.Start()
}

func (r *CronService) Stop() {
	r.cron.Stop()
}
