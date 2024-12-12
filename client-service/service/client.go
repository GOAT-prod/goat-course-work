package service

import (
	"client-service/domain"
	"client-service/domain/mappings"
	"client-service/repository"
	"github.com/GOAT-prod/goatcontext"
)

type Client interface {
	GetClients(ctx goatcontext.Context) ([]domain.Client, error)
	GetClient(ctx goatcontext.Context, id int) (domain.Client, error)
	AddClient(ctx goatcontext.Context, client domain.Client) (domain.Client, error)
	UpdateClient(ctx goatcontext.Context, client domain.Client) error
	DeleteClient(ctx goatcontext.Context, id int) error
}

type ClientServiceImpl struct {
	clientRepository repository.Client
}

func NewClientService(clientRepository repository.Client) Client {
	return &ClientServiceImpl{
		clientRepository: clientRepository,
	}
}

func (s *ClientServiceImpl) GetClients(ctx goatcontext.Context) ([]domain.Client, error) {
	dbClient, err := s.clientRepository.GetClients(ctx)
	if err != nil {
		return nil, err
	}

	return mappings.ToDomainClients(dbClient), nil
}

func (s *ClientServiceImpl) GetClient(ctx goatcontext.Context, id int) (domain.Client, error) {
	dbClient, err := s.clientRepository.GetClient(ctx, id)
	if err != nil {
		return domain.Client{}, err
	}

	return mappings.ToDomainClient(dbClient), nil
}

func (s *ClientServiceImpl) AddClient(ctx goatcontext.Context, client domain.Client) (domain.Client, error) {
	id, err := s.clientRepository.AddClient(ctx, mappings.ToDatabaseClient(client))
	if err != nil {
		return domain.Client{}, err
	}

	client.Id = id

	return client, nil
}

func (s *ClientServiceImpl) UpdateClient(ctx goatcontext.Context, client domain.Client) error {
	return s.clientRepository.UpdateClient(ctx, mappings.ToDatabaseClient(client))
}

func (s *ClientServiceImpl) DeleteClient(ctx goatcontext.Context, id int) error {
	return s.clientRepository.DeleteClient(ctx, id)
}
