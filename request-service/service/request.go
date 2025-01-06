package service

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/samber/lo"
	"log"
	"request-service/cluster/notifier"
	"request-service/cluster/warehouse"
	"request-service/database"
	"request-service/domain"
	"request-service/repository"
	"strings"
)

type Request interface {
	GetRequests(ctx goatcontext.Context) ([]domain.Request, error)
	UpdateStatus(ctx goatcontext.Context, id int, status domain.Status) error
}

type Impl struct {
	requestRepository repository.Request
	warehouseClient   *warehouse.Client
	notifierClient    *notifier.Client
}

func NewRequestService(requestRepository repository.Request, warehouseClient *warehouse.Client, notifierClient *notifier.Client) Request {
	return &Impl{
		requestRepository: requestRepository,
		warehouseClient:   warehouseClient,
		notifierClient:    notifierClient,
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
	if err := s.requestRepository.UpdateRequestStatus(ctx, id, string(status)); err != nil {
		return err
	}

	if status == domain.ApprovedStatus {
		go func(ctx goatcontext.Context, requestId int) {
			if err := s.sendMessage(ctx, id); err != nil {
				log.Printf("failed to send message: %v", err)
			}
		}(ctx, id)
	}

	return nil
}

func (s Impl) sendMessage(ctx goatcontext.Context, requestId int) error {
	request, err := s.requestRepository.GetRequestById(ctx, requestId)
	if err != nil {
		return err
	}

	product, err := s.warehouseClient.GetDetailedProduct(ctx, request.Items[0].ProductId)
	if err != nil {
		return err
	}

	mail := notifier.Mail{
		Subject: fmt.Sprintf(notifier.SubjectByRequestType[domain.Type(request.Type)], product.Name),
		Body:    getMailMessageBody(request, product),
	}

	return s.notifierClient.SendMail(ctx, mail)
}

func getMailMessageBody(request database.Request, product domain.Product) string {
	switch domain.Type(request.Type) {
	case domain.SupplyType:
		bodyBuilder := strings.Builder{}
		for _, requestItem := range request.Items {
			productItem, ok := lo.Find(product.Items, func(item domain.ProductItem) bool {
				return item.Id == requestItem.ProductId
			})

			if !ok {
				continue
			}

			bodyBuilder.WriteString(fmt.Sprintf(notifier.SupplyBodyTemplate, productItem.Color, productItem.Size, requestItem.ProductItemCount))
		}

		return bodyBuilder.String()
	case domain.ApproveType:
		return fmt.Sprintf(notifier.SubjectByRequestType[domain.Type(request.Type)], product.Name)
	default:
		return ""
	}
}
