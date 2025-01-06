package service

import (
	"github.com/GOAT-prod/goatcontext"
	"request-service/cluster/warehouse"
	"request-service/domain"
	"request-service/repository"
)

type Request interface {
	GetRequests(ctx goatcontext.Context) ([]domain.Request, error)
	UpdateStatus(ctx goatcontext.Context, id int, status domain.Status) error
}

type Impl struct {
	requestRepository repository.Request
	warehouseClient   *warehouse.Client
}

func NewRequestService(requestRepository repository.Request, client *warehouse.Client) Request {
	return &Impl{
		requestRepository: requestRepository,
		warehouseClient:   client,
	}
}

func (s Impl) GetRequests(ctx goatcontext.Context) ([]domain.Request, error) {
	dbRequests, err := s.requestRepository.GetPendingRequests(ctx)
	if err != nil {
		return nil, err
	}

	requests := make([]domain.Request, 0, len(dbRequests))

	for _, dbRequest := range dbRequests {
		request := domain.Request{
			Id:          dbRequest.Id,
			Status:      domain.Status(dbRequest.Status),
			Type:        domain.Type(dbRequest.Type),
			UpdatedDate: dbRequest.UpdateDate,
			Summary:     dbRequest.Summary,
		}

		product, err := s.warehouseClient.GetDetailedProduct(ctx, dbRequest.Items[0].ProductId)
		if err != nil {
			return nil, err
		}

		request.Item = product

		requests = append(requests, request)
	}

	return requests, nil
}

func (s Impl) UpdateStatus(ctx goatcontext.Context, id int, status domain.Status) error {
	return s.requestRepository.UpdateRequestStatus(ctx, id, string(status))
}
