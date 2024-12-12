package service

import (
	"auth-service/cluster/clientservice"
	"auth-service/cluster/userservice"
	"auth-service/domain"
	"context"
	"errors"
)

type Registration interface {
	SignUp(registerData domain.RegisterData) (domain.TokenResponse, error)
}

type RegistrationServiceImpl struct {
	hasher        PasswordHasher
	jwtService    JwtService
	userService   *userservice.Client
	clientService *clientservice.Client
}

func NewRegistrationService(hasher PasswordHasher, jwtService JwtService, userClient *userservice.Client, clientService *clientservice.Client) Registration {
	return &RegistrationServiceImpl{
		hasher:        hasher,
		jwtService:    jwtService,
		userService:   userClient,
		clientService: clientService,
	}
}

func (r *RegistrationServiceImpl) SignUp(registerData domain.RegisterData) (domain.TokenResponse, error) {
	//сформировать хэш пароля
	passwordHash, err := r.hasher.Hash(registerData.Password)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	//отправить пользователя на сохранение
	if userCheck, checkUserErr := r.userService.GetUserByUserName(context.TODO(), registerData.UserName); checkUserErr != nil || userCheck.Id != 0 {
		return domain.TokenResponse{}, errors.New("user already exists")
	}

	clientData := domain.ClientData{
		Name:    registerData.ClientName,
		INN:     registerData.ClientINN,
		Address: registerData.ClientAddress,
	}

	clientData, err = r.clientService.AddClientData(context.TODO(), clientData)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	user := domain.User{
		UserName: registerData.UserName,
		Password: passwordHash,
		Role:     registerData.Role,
		Status:   domain.Active,
		ClientId: clientData.Id,
	}

	registeredUser, err := r.userService.RegisterUser(context.TODO(), user)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	//проверить ответ от сервиса пользователей
	if registeredUser.Id == 0 {
		return domain.TokenResponse{}, errors.New("user registration failed")
	}

	//сгенерировать токен или отправить в кафку ивент о подтвержлениии пользователя
	accessToken, refreshToken, err := r.jwtService.Generate(registeredUser)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	//вернуть на фронт ответ с токенами или с текстом что пользователь ушел на подтверждение
	return domain.TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
