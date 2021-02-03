package internal

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
)

// Firestore implementation of repository.
type firestoreRepository struct {
	clt *firestore.Client
}

const (
	clientsCollection  = "clients"
	accountsCollection = "accounts"
)

func newFirestoreRepository(firestoreClient *firestore.Client) *firestoreRepository {
	return &firestoreRepository{
		clt: firestoreClient,
	}
}

func (repo *firestoreRepository) storeClient(client *domain.Client) error {
	dto := newClientDtoFromDomain(client)
	_, err := repo.clt.Collection(clientsCollection).Doc(client.Id.String()).Set(context.Background(), &dto)
	return err
}

func (repo *firestoreRepository) getAllClients() ([]*domain.Client, error) {
	panic("implement me")
}

func (repo *firestoreRepository) getClientById(id domain.UniqueId) (*domain.Client, error) {
	panic("implement me")
}

func (repo *firestoreRepository) storeAccount(account *domain.Account) error {
	panic("implement me")
}

func (repo *firestoreRepository) getAllAccounts() ([]*domain.Account, error) {
	snapshots, err := repo.clt.Collection(accountsCollection).Documents(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}
	accounts := make([]*domain.Account, len(snapshots))
	for i, snapshot := range snapshots {
		var dto AccountDto
		if err := snapshot.DataTo(&dto); err != nil {
			return nil, err
		}
		dto.Id = snapshot.Ref.ID
		account, err := dto.ToDomain()
		if err != nil {
			return nil, err
		}
		accounts[i] = account
	}
	return accounts, err
}

func (repo *firestoreRepository) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	panic("implement me")
}

func (repo firestoreRepository) runTransaction(f2 func(tx transaction) error) error {
	panic("implement me")
}

type firestoreTransaction struct {
	tx *firestore.Transaction
}

func (tx *firestoreTransaction) storeClient(client *domain.Client) error {
	panic("implement me")
}

func (tx *firestoreTransaction) getAllClients() ([]*domain.Client, error) {
	panic("implement me")
}

func (tx *firestoreTransaction) getClientById(id domain.UniqueId) (*domain.Client, error) {
	panic("implement me")
}

func (tx *firestoreTransaction) storeAccount(account *domain.Account) error {
	panic("implement me")
}

func (tx *firestoreTransaction) getAllAccounts() ([]*domain.Account, error) {
	panic("implement me")
}

func (tx *firestoreTransaction) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	panic("implement me")
}
