package repository

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"order-service/database"
	"order-service/database/queries"
)

type Finance interface {
	GetOrderOperation(ctx goatcontext.Context, orderId uuid.UUID) (operation database.Operation, err error)
	CreateOrderOperation(ctx goatcontext.Context, operation database.Operation) error
	GetOperationDetails(ctx goatcontext.Context, operationId uuid.UUID) (details []database.OperationDetail, err error)
	CreateOperationDetail(ctx goatcontext.Context, operation database.OperationDetail) error
}

type FinanceRepositoryImpl struct {
	postgres *sqlx.DB
}

func NewFinanceRepository(postgres *sqlx.DB) Finance {
	return &FinanceRepositoryImpl{
		postgres: postgres,
	}
}

func (r *FinanceRepositoryImpl) GetOrderOperation(ctx goatcontext.Context, orderId uuid.UUID) (operation database.Operation, err error) {
	return operation, r.postgres.GetContext(ctx, &operation, queries.GetOperation, orderId)
}

func (r *FinanceRepositoryImpl) CreateOrderOperation(ctx goatcontext.Context, operation database.Operation) error {
	_, err := r.postgres.NamedExecContext(ctx, queries.CreateOperation, operation)
	return err
}

func (r *FinanceRepositoryImpl) GetOperationDetails(ctx goatcontext.Context, operationId uuid.UUID) (details []database.OperationDetail, err error) {
	return details, r.postgres.SelectContext(ctx, &details, queries.GetOperationDetails, operationId)
}

func (r *FinanceRepositoryImpl) CreateOperationDetail(ctx goatcontext.Context, operation database.OperationDetail) error {
	_, err := r.postgres.NamedExecContext(ctx, queries.CreateOperationDetail, operation)
	return err
}
