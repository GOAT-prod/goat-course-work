package service

import (
	"github.com/GOAT-prod/goatcontext"
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
}

type UserServiceImpl struct {
	userRepository repository.User
	roleRepository repository.Role
}

func NewUserService(userRepository repository.User, roleRepository repository.Role) User {
	return &UserServiceImpl{
		userRepository: userRepository,
		roleRepository: roleRepository,
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
