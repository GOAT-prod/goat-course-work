package repository

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/jmoiron/sqlx"
	"user-service/database"
	"user-service/database/queries"
)

type User interface {
	GetUsers(ctx goatcontext.Context) ([]database.User, error)
	GetUserById(ctx goatcontext.Context, userId int) (database.User, error)
	GetUserByUsername(ctx goatcontext.Context, username string) (database.User, error)
	AddUser(ctx goatcontext.Context, user database.User) (int, error)
	UpdateUser(ctx goatcontext.Context, user database.User) error
	DeleteUser(ctx goatcontext.Context, userId int) error
}

type UserRepositoryImpl struct {
	postgres *sqlx.DB
}

func NewUserRepository(postgres *sqlx.DB) User {
	return &UserRepositoryImpl{
		postgres: postgres,
	}
}

func (r *UserRepositoryImpl) GetUsers(ctx goatcontext.Context) (users []database.User, err error) {
	return users, r.postgres.SelectContext(ctx, &users, queries.GetUsers)
}

func (r *UserRepositoryImpl) GetUserById(ctx goatcontext.Context, userId int) (user database.User, err error) {
	return user, r.postgres.GetContext(ctx, &user, queries.GetUserById, userId)
}

func (r *UserRepositoryImpl) GetUserByUsername(ctx goatcontext.Context, username string) (user database.User, err error) {
	return user, r.postgres.GetContext(ctx, &user, queries.GetUserByUsername, username)
}

func (r *UserRepositoryImpl) AddUser(ctx goatcontext.Context, user database.User) (id int, err error) {
	return id, r.postgres.GetContext(ctx, &id, queries.AddUser, user)
}

func (r *UserRepositoryImpl) UpdateUser(ctx goatcontext.Context, user database.User) error {
	_, err := r.postgres.NamedExecContext(ctx, queries.UpdateUser, user)
	return err
}

func (r *UserRepositoryImpl) DeleteUser(ctx goatcontext.Context, userId int) error {
	_, err := r.postgres.ExecContext(ctx, queries.DeleteUser, userId)
	return err
}
