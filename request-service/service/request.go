package service

import "request-service/repository"

type Request interface{}

type Impl struct {
	requestRepository repository.Request
}

func NewRequestService(requestRepository repository.Request) Request {
	return &Impl{
		requestRepository: requestRepository,
	}
}
