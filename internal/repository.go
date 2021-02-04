package internal

import (
	"errors"
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
)

// Returned by the repository when a requested entity is not found.
var errNotFound = errors.New("not found")

type repository interface {
	// Starts a transaction with the repository.
	runTransaction(func(tx transaction) error) error
	transaction
}

type transaction interface {
	// Persists the client.
	storeClient(client *domain.Client) error
	// Returns all the persisted clients.
	getAllClients() ([]*domain.Client, error)
	// Returns the persisted client with the specified id.
	getClientById(id domain.UniqueId) (*domain.Client, error)
	// Persists the account.
	storeAccount(account *domain.Account) error
	// Returns all the persisted accounts.
	getAllAccounts() ([]*domain.Account, error)
	// Returns the persisted account with the specified id.
	getAccountById(id domain.UniqueId) (*domain.Account, error)
}
