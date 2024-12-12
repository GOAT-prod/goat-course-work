package mappings

import (
	"user-service/database"
	"user-service/domain"
)

func ToDomainUser(user database.User, role string) domain.User {
	return domain.User{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Status:   domain.UserStatus(user.Status),
		Role:     domain.UserRole(role),
		ClientId: user.ClientId,
	}
}

func ToDatabaseUser(user domain.User, roleId int) database.User {
	return database.User{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Status:   int(user.Status),
		RoleId:   roleId,
		ClientId: user.ClientId,
	}
}
