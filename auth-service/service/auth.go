package service

import (
	"auth-service/cluster/userservice"
	"auth-service/domain"
	"context"
	"errors"
	"fmt"
	"log"
)

type AuthService interface {
	Login(username, password string) (domain.TokenResponse, error)
	Logout(refreshToken string)
	RefreshToken(token string) (domain.TokenResponse, error)
	UpdatePassword(request domain.UpdatePasswordRequest) error
}

type AuthServiceImpl struct {
	hasher      PasswordHasher
	jwtService  JwtService
	userService *userservice.Client
}

func NewAuthService(hasher PasswordHasher, jwtService JwtService, userClient *userservice.Client) AuthService {
	return &AuthServiceImpl{
		hasher:      hasher,
		jwtService:  jwtService,
		userService: userClient,
	}
}

func (a *AuthServiceImpl) Login(username, password string) (domain.TokenResponse, error) {
	//получить пользователя по userName
	user, err := a.userService.GetUserByUserName(context.TODO(), username)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	//вычислить новый хэш пароля
	passwordHash, err := a.hasher.Hash(password)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	//сравинить пароли
	if passwordHash != user.Password {
		return domain.TokenResponse{}, errors.New("invalid password")
	}

	//сгененировать токены
	accessToken, refreshToken, err := a.jwtService.Generate(user)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	//отдать токены на фронт
	return domain.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *AuthServiceImpl) Logout(refreshToken string) {
	//удалить рефрешь токен из редиса
	if err := a.jwtService.RetrieveRefreshToken(refreshToken); err != nil {
		log.Printf("failed to delete refresh token: %v", err)
		return
	}
}

func (a *AuthServiceImpl) RefreshToken(token string) (domain.TokenResponse, error) {
	accessToken, refreshToken, err := a.jwtService.Refresh(token)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	return domain.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *AuthServiceImpl) UpdatePassword(request domain.UpdatePasswordRequest) error {
	user, err := a.userService.GetUserByUserName(context.TODO(), request.Username)
	if err != nil {
		return fmt.Errorf("не удалось получить пользователя: %w", err)
	}

	passwordHash, err := a.hasher.Hash(request.Password)
	if err != nil {
		return fmt.Errorf("не удалось захешировать новый пароль: %w", err)
	}

	if passwordHash == user.Password {
		return fmt.Errorf("новый и старый пароли совпадают")
	}

	if err = a.userService.UpdateUserPassword(context.TODO(), request); err != nil {
		return fmt.Errorf("не удалось сохранить новый пароль: %w", err)
	}

	return nil
}
