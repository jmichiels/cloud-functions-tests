package internal

import (
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
	"github.com/pkg/errors"
)

type ClientDto struct {
	Id        string `json:"id" firestore:"-"`
	FirstName string `json:"firstName" firestore:"firstName"`
	LastName  string `json:"lastName" firestore:"lastName"`
	Age       uint   `json:"age" firestore:"age"`
}

func newClientDtoFromDomain(client *domain.Client) *ClientDto {
	return &ClientDto{
		Id:        client.Id.String(),
		FirstName: client.FirstName,
		LastName:  client.LastName,
		Age:       client.Age,
	}
}

func (dto *ClientDto) ToDomain() (*domain.Client, error) {
	clientId, err := domain.ParseUniqueIdFromString(dto.Id)
	if err != nil {
		return nil, errors.Wrap(err, "parse client id")
	}
	client := &domain.Client{
		Id:        clientId,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Age:       dto.Age,
	}
	return client, client.Validate()
}
