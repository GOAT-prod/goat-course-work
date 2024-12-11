package repository

import "github.com/jmoiron/sqlx"

type Request interface{}

type Impl struct {
	postgres *sqlx.DB
}

func NewRequestRepository(postgres *sqlx.DB) Request {
	return &Impl{
		postgres: postgres,
	}
}
