package repository

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/jmoiron/sqlx"
	"report-service/database"
	"report-service/database/queries"
	"time"
)

type Report interface {
	AddReportItems(ctx goatcontext.Context, items []database.ReportItem) error
	GetFactoryReportItems(ctx goatcontext.Context, factoryId int, startDate, endDate time.Time) (items []database.ReportItem, err error)
	GetUserReportItems(ctx goatcontext.Context, userId int, startDate, endDate time.Time) (items []database.ReportItem, err error)
}

type ReportRepository struct {
	postgres *sqlx.DB
}

func NewReportRepository(postgres *sqlx.DB) Report {
	return &ReportRepository{
		postgres: postgres,
	}
}

func (r *ReportRepository) AddReportItems(ctx goatcontext.Context, items []database.ReportItem) error {
	tx, err := r.postgres.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	for _, item := range items {
		if _, err = tx.NamedExecContext(ctx, queries.AddReportItems, item); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *ReportRepository) GetFactoryReportItems(ctx goatcontext.Context, factoryId int, startDate, endDate time.Time) (items []database.ReportItem, err error) {
	return items, r.postgres.SelectContext(ctx, &items, queries.GetFactoryReportItems, factoryId, startDate, endDate)
}

func (r *ReportRepository) GetUserReportItems(ctx goatcontext.Context, userId int, startDate, endDate time.Time) (items []database.ReportItem, err error) {
	return items, r.postgres.SelectContext(ctx, &items, queries.GetUserReportItems, userId, startDate, endDate)
}
