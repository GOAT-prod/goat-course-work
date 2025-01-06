package service

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"log"
	"user-service/cluster/notifier"
	"user-service/domain"
	"user-service/domain/mappings"
	"user-service/repository"
)

type User interface {
	GetUsers(ctx goatcontext.Context) ([]domain.User, error)
	GetUserById(ctx goatcontext.Context, userId int) (domain.User, error)
	GetUserByUsername(ctx goatcontext.Context, userName string) (domain.User, error)
	AddUser(ctx goatcontext.Context, user domain.User) (domain.User, error)
	UpdateUser(ctx goatcontext.Context, user domain.User) error
	DeleteUser(ctx goatcontext.Context, userId int) error
	UpdateUserPassword(ctx goatcontext.Context, request domain.UpdatePasswordRequest) error
}

type UserServiceImpl struct {
	userRepository repository.User
	roleRepository repository.Role
	notifierClient *notifier.Client
}

func NewUserService(userRepository repository.User, roleRepository repository.Role, notifierClient *notifier.Client) User {
	return &UserServiceImpl{
		userRepository: userRepository,
		roleRepository: roleRepository,
		notifierClient: notifierClient,
	}
}

func (s *UserServiceImpl) GetUsers(ctx goatcontext.Context) ([]domain.User, error) {
	dbUsers, err := s.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]domain.User, 0, len(dbUsers))
	for _, dbUser := range dbUsers {
		userRole, rErr := s.roleRepository.GetRoleById(ctx, dbUser.RoleId)
		if rErr != nil {
			continue
		}

		users = append(users, mappings.ToDomainUser(dbUser, userRole))
	}

	return users, nil
}

func (s *UserServiceImpl) GetUserById(ctx goatcontext.Context, userId int) (domain.User, error) {
	dbUser, err := s.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}

	userRole, err := s.roleRepository.GetRoleById(ctx, dbUser.RoleId)
	if err != nil {
		return domain.User{}, err
	}

	return mappings.ToDomainUser(dbUser, userRole), nil
}

func (s *UserServiceImpl) GetUserByUsername(ctx goatcontext.Context, userName string) (domain.User, error) {
	dbUser, err := s.userRepository.GetUserByUsername(ctx, userName)
	if err != nil {
		return domain.User{}, err
	}

	if dbUser.Id == 0 {
		return domain.User{}, nil
	}

	userRole, err := s.roleRepository.GetRoleById(ctx, dbUser.RoleId)
	if err != nil {
		return domain.User{}, err
	}

	return mappings.ToDomainUser(dbUser, userRole), nil
}

func (s *UserServiceImpl) AddUser(ctx goatcontext.Context, user domain.User) (domain.User, error) {
	roleId, err := s.roleRepository.GetRoleIdByName(ctx, string(user.Role))
	if err != nil {
		return domain.User{}, err
	}

	user.Id, err = s.userRepository.AddUser(ctx, mappings.ToDatabaseUser(user, roleId))
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserServiceImpl) UpdateUser(ctx goatcontext.Context, user domain.User) error {
	return s.userRepository.UpdateUser(ctx, mappings.ToDatabaseUser(user, 0))
}

func (s *UserServiceImpl) DeleteUser(ctx goatcontext.Context, userId int) error {
	return s.userRepository.DeleteUser(ctx, userId)
}

func (s *UserServiceImpl) UpdateUserPassword(ctx goatcontext.Context, request domain.UpdatePasswordRequest) error {
	user, err := s.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return err
	}

	user.Password = request.Password

	if err = s.UpdateUser(ctx, user); err != nil {
		return err
	}

	go s.sendMessage(ctx, user)
	return nil
}

func (s *UserServiceImpl) sendMessage(ctx goatcontext.Context, user domain.User) {
	msg := notifier.MailMessage{
		Subject: fmt.Sprintf("Пароль пользователя %s был изменен", user.Username),
		Body:    fmt.Sprintf("Пароль пользователя %s был изменен", user.Username),
	}

	if err := s.notifierClient.SendMessage(ctx, msg); err != nil {
		log.Println(err)
	}
}
