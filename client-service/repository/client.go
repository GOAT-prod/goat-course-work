package repository

import (
	"client-service/database"
	"client-service/database/queries"
	"github.com/GOAT-prod/goatcontext"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Client interface {
	GetClients(ctx goatcontext.Context) (clients []database.Client, err error)
	GetClient(ctx goatcontext.Context, id int) (client database.Client, err error)
	AddClient(ctx goatcontext.Context, client database.Client) (id int, err error)
	UpdateClient(ctx goatcontext.Context, client database.Client) error
	DeleteClient(ctx goatcontext.Context, id int) error
	GetClientsByIds(ctx goatcontext.Context, ids []int) (clients []database.Client, err error)
}

type ClientRepositoryImpl struct {
	postgres *sqlx.DB
}

func NewClientRepository(postgres *sqlx.DB) Client {
	return &ClientRepositoryImpl{
		postgres: postgres,
	}
}

func (r *ClientRepositoryImpl) GetClients(ctx goatcontext.Context) (clients []database.Client, err error) {
	return clients, r.postgres.SelectContext(ctx, &clients, queries.GetClients)
}

func (r *ClientRepositoryImpl) GetClient(ctx goatcontext.Context, id int) (client database.Client, err error) {
	return client, r.postgres.GetContext(ctx, &client, queries.GetClientById, id)
}

func (r *ClientRepositoryImpl) GetClientsByIds(ctx goatcontext.Context, ids []int) (clients []database.Client, err error) {
	return clients, r.postgres.SelectContext(ctx, &clients, queries.GetClientsByIds, pq.Array(ids))
}

func (r *ClientRepositoryImpl) AddClient(ctx goatcontext.Context, client database.Client) (id int, err error) {
	return id, r.postgres.GetContext(ctx, &id, queries.AddClient, client.Name, client.INN, client.Address)
}

func (r *ClientRepositoryImpl) UpdateClient(ctx goatcontext.Context, client database.Client) error {
	_, err := r.postgres.NamedExecContext(ctx, queries.UpdateClient, client)
	return err
}

func (r *ClientRepositoryImpl) DeleteClient(ctx goatcontext.Context, id int) error {
	_, err := r.postgres.ExecContext(ctx, queries.DeleteClient, id)
	return err
}
