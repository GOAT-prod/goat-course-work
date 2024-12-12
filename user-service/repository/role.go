package repository

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/jmoiron/sqlx"
	"user-service/database/queries"
)

type Role interface {
	GetRoleIdByName(ctx goatcontext.Context, name string) (int, error)
	GetRoleById(ctx goatcontext.Context, id int) (name string, err error)
}

type RoleRepository struct {
	postgres *sqlx.DB
}

func NewRoleRepository(postgres *sqlx.DB) Role {
	return &RoleRepository{
		postgres: postgres,
	}
}

func (r *RoleRepository) GetRoleIdByName(ctx goatcontext.Context, name string) (id int, err error) {
	return id, r.postgres.GetContext(ctx, &id, queries.GetRoleIdByName, name)
}

func (r *RoleRepository) GetRoleById(ctx goatcontext.Context, id int) (name string, err error) {
	return name, r.postgres.GetContext(ctx, &name, queries.GetRoleById, id)
}
