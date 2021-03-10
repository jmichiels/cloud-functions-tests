package internal

import (
	"context"
	"errors"
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
)

// Returned by the database when a requested entity is not found.
var errNotFound = errors.New("not found")

type database interface {
	// A database groups all the repositories.
	repositories
	// Runs a transaction with the database. All the calls to the repositories via the tx argument of the callback
	// will be made in a single transaction, committed once the callback returns. All reads must happen before any
	// write operation.
	runTransaction(context.Context, func(tx repositories) error) error
}

// Groups all the repositories together.
type repositories interface {
	clientRepository
	accountRepository
}

type clientRepository interface {
	// Persists the client.
	storeClient(ctx context.Context, client *domain.Client) error
	// Returns all the persisted clients.
	getAllClients(ctx context.Context) ([]*domain.Client, error)
	// Returns the persisted client with the specified id.
	getClientById(ctx context.Context, id domain.UniqueId) (*domain.Client, error)
}

type accountRepository interface {
	// Persists the account.
	storeAccount(ctx context.Context, account *domain.Account) error
	// Returns all the persisted accounts.
	getAllAccounts(ctx context.Context) ([]*domain.Account, error)
	// Returns the persisted account with the specified id.
	getAccountById(ctx context.Context, id domain.UniqueId) (*domain.Account, error)
}
