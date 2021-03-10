package internal

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
)

type firestoreDatabase struct {
	firestoreClt *firestore.Client
}

func newFirestoreDatabase(clt *firestore.Client) *firestoreDatabase {
	return &firestoreDatabase{
		firestoreClt: clt,
	}
}

func (db *firestoreDatabase) storeClient(client *domain.Client) error {
	// todo db.firestoreClt.Doc("")
	return
}

func (db *firestoreDatabase) getAllClients() (clients []*domain.Client, err error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.clientRepo.getAllClients()
}

func (db *firestoreDatabase) getClientById(id domain.UniqueId) (*domain.Client, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.clientRepo.getClientById(id)
}

func (db *firestoreDatabase) storeAccount(account *domain.Account) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.accountRepo.storeAccount(account)
}

func (db *firestoreDatabase) getAllAccounts() ([]*domain.Account, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.accountRepo.getAllAccounts()
}

func (db *firestoreDatabase) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.accountRepo.getAccountById(id)
}

func (db *firestoreDatabase) runTransaction(f func(tx transaction) error) error {
	return db.firestoreClt.RunTransaction(context.TODO(), func(ctx context.Context, tx *firestore.Transaction) error {
		return f(&struct {
			*firestoreClientRepository
			*firestoreAccountRepository
		}{
			&firestoreClientRepository{tx},
			&firestoreAccountRepository{tx},
		})
	})
}

type firestoreClientRepository struct {
	firestoreTx *firestore.Transaction
}

func (repo *firestoreClientRepository) storeClient(client *domain.Client) error {
	repo.firestoreTx.
		repo.clients[client.Id.String()] = client
	return nil
}

func (repo *firestoreClientRepository) getAllClients() ([]*domain.Client, error) {
	clients := make([]*domain.Client, 0, len(repo.clients))
	for _, client := range repo.clients {
		clients = append(clients, client)
	}
	return clients, nil
}

func (repo *firestoreClientRepository) getClientById(id domain.UniqueId) (*domain.Client, error) {
	client, ok := repo.clients[id.String()]
	if !ok {
		return nil, errNotFound
	}
	return client, nil
}

type firestoreAccountRepository struct {
	firestoreTx *firestore.Transaction
}

func (repo *firestoreAccountRepository) storeAccount(account *domain.Account) error {
	repo.accounts[account.Id.String()] = account
	return nil
}

func (repo *firestoreAccountRepository) getAllAccounts() ([]*domain.Account, error) {
	accounts := make([]*domain.Account, 0, len(repo.accounts))
	for _, account := range repo.accounts {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (repo *firestoreAccountRepository) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	account, ok := repo.accounts[id.String()]
	if !ok {
		return nil, errNotFound
	}
	return account, nil
}
