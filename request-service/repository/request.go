package repository

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"request-service/database"
	"request-service/database/queries"
)

type Request interface {
	GetPendingRequests(ctx goatcontext.Context) ([]database.Request, error)
	AddRequest(ctx goatcontext.Context, request database.Request) error
	UpdateRequestStatus(ctx goatcontext.Context, id int, status string) error
}

type Impl struct {
	postgres *sqlx.DB
}

func NewRequestRepository(postgres *sqlx.DB) Request {
	return &Impl{
		postgres: postgres,
	}
}

func (r Impl) GetPendingRequests(ctx goatcontext.Context) ([]database.Request, error) {
	var requests []database.Request
	if err := r.postgres.SelectContext(ctx, &requests, queries.GetPendingRequests); err != nil {
		return nil, err
	}

	for i := range requests {
		if err := r.postgres.SelectContext(ctx, &requests[i].Items, queries.GetRequestItemsByRequestId, requests[i].Id); err != nil {
			return nil, err
		}
	}

	return requests, nil
}

func (r Impl) AddRequest(ctx goatcontext.Context, request database.Request) error {
	var requestId int
	if err := r.postgres.GetContext(ctx, &requestId, queries.AddRequest, request); err != nil {
		return err
	}

	request.Items = lo.Map(request.Items, func(item database.RequestItem, _ int) database.RequestItem {
		item.Id = requestId
		return item
	})

	for _, requestItem := range request.Items {
		if _, err := r.postgres.NamedExecContext(ctx, queries.AddRequestItem, requestItem); err != nil {
			return err
		}
	}

	return nil
}

func (r Impl) UpdateRequestStatus(ctx goatcontext.Context, id int, status string) error {
	_, err := r.postgres.ExecContext(ctx, queries.UpdateRequestStatus, status, id)
	return err
}
