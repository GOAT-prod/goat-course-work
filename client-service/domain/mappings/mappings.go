package mappings

import (
	"client-service/database"
	"client-service/domain"
	"github.com/samber/lo"
)

func ToDomainClient(client database.Client) domain.Client {
	return domain.Client{
		Id:      client.Id,
		Name:    client.Name,
		INN:     client.INN,
		Address: client.Address,
	}
}

func ToDatabaseClient(client domain.Client) database.Client {
	return database.Client{
		Id:      client.Id,
		Name:    client.Name,
		INN:     client.INN,
		Address: client.Address,
	}
}

func ToDomainClients(clients []database.Client) []domain.Client {
	return lo.Map(clients, func(item database.Client, _ int) domain.Client {
		return ToDomainClient(item)
	})
}

func ToDatabaseClients(clients []domain.Client) []database.Client {
	return lo.Map(clients, func(item domain.Client, _ int) database.Client {
		return ToDatabaseClient(item)
	})
}
