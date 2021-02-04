package internal

import (
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
	"sync"
)

type mockRepository struct {
	mutex  sync.Mutex
	unsafe unsafeRepository
}

func newMockRepository() *mockRepository {
	return &mockRepository{
		unsafe: unsafeRepository{
			clients:  make(map[string]*domain.Client),
			accounts: make(map[string]*domain.Account),
		},
	}
}

func (repo *mockRepository) runTransaction(f func(tx transaction) error) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return f(&repo.unsafe)
}

type unsafeRepository struct {
	clients  map[string]*domain.Client
	accounts map[string]*domain.Account
}

func (repo *unsafeRepository) storeClient(client *domain.Client) error {
	repo.clients[client.Id.String()] = client
	return nil
}

func (repo *unsafeRepository) getAllClients() ([]*domain.Client, error) {
	clients := make([]*domain.Client, 0, len(repo.clients))
	for _, client := range repo.clients {
		clients = append(clients, client)
	}
	return clients, nil
}

func (repo *unsafeRepository) getClientById(id domain.UniqueId) (*domain.Client, error) {
	client, ok := repo.clients[id.String()]
	if !ok {
		return nil, errNotFound
	}
	return client, nil
}

func (repo *unsafeRepository) storeAccount(account *domain.Account) error {
	repo.accounts[account.Id.String()] = account
	return nil
}

func (repo *unsafeRepository) getAllAccounts() ([]*domain.Account, error) {
	accounts := make([]*domain.Account, 0, len(repo.accounts))
	for _, account := range repo.accounts {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (repo *unsafeRepository) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	account, ok := repo.accounts[id.String()]
	if !ok {
		return nil, errNotFound
	}
	return account, nil
}

func (repo *mockRepository) storeClient(client *domain.Client) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.unsafe.storeClient(client)
}

func (repo *mockRepository) getAllClients() (clients []*domain.Client, err error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.unsafe.getAllClients()
}

func (repo *mockRepository) getClientById(id domain.UniqueId) (*domain.Client, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.unsafe.getClientById(id)
}

func (repo *mockRepository) storeAccount(account *domain.Account) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.unsafe.storeAccount(account)
}

func (repo *mockRepository) getAllAccounts() ([]*domain.Account, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.unsafe.getAllAccounts()
}

func (repo *mockRepository) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.unsafe.getAccountById(id)
}
