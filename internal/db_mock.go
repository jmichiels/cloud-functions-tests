package internal

import (
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
	"sync"
)

type mockDatabase struct {
	mutex sync.Mutex

	clientRepo  mockClientRepository
	accountRepo mockAccountRepository
}

func newMockDatabase() *mockDatabase {
	return &mockDatabase{
		clientRepo: mockClientRepository{
			clients: make(map[string]*domain.Client),
		},
		accountRepo: mockAccountRepository{
			accounts: make(map[string]*domain.Account),
		},
	}
}

func (db *mockDatabase) runTransaction(f func(tx repositories) error) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return f(&struct {
		*mockClientRepository
		*mockAccountRepository
	}{
		&db.clientRepo,
		&db.accountRepo,
	})
}

type mockClientRepository struct {
	clients map[string]*domain.Client
}

func (repo *mockClientRepository) storeClient(client *domain.Client) error {
	repo.clients[client.Id.String()] = client
	return nil
}

func (repo *mockClientRepository) getAllClients() ([]*domain.Client, error) {
	clients := make([]*domain.Client, 0, len(repo.clients))
	for _, client := range repo.clients {
		clients = append(clients, client)
	}
	return clients, nil
}

func (repo *mockClientRepository) getClientById(id domain.UniqueId) (*domain.Client, error) {
	client, ok := repo.clients[id.String()]
	if !ok {
		return nil, errNotFound
	}
	return client, nil
}

type mockAccountRepository struct {
	accounts map[string]*domain.Account
}

func (repo *mockAccountRepository) storeAccount(account *domain.Account) error {
	repo.accounts[account.Id.String()] = account
	return nil
}

func (repo *mockAccountRepository) getAllAccounts() ([]*domain.Account, error) {
	accounts := make([]*domain.Account, 0, len(repo.accounts))
	for _, account := range repo.accounts {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (repo *mockAccountRepository) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	account, ok := repo.accounts[id.String()]
	if !ok {
		return nil, errNotFound
	}
	return account, nil
}

func (db *mockDatabase) storeClient(client *domain.Client) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.clientRepo.storeClient(client)
}

func (db *mockDatabase) getAllClients() (clients []*domain.Client, err error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.clientRepo.getAllClients()
}

func (db *mockDatabase) getClientById(id domain.UniqueId) (*domain.Client, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.clientRepo.getClientById(id)
}

func (db *mockDatabase) storeAccount(account *domain.Account) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.accountRepo.storeAccount(account)
}

func (db *mockDatabase) getAllAccounts() ([]*domain.Account, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.accountRepo.getAllAccounts()
}

func (db *mockDatabase) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.accountRepo.getAccountById(id)
}
