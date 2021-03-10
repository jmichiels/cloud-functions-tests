package internal

import (
	"context"
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
)

func (db *firestoreDatabase) storeAccount(ctx context.Context, account *domain.Account) error {
	// todo
}

func (db *firestoreDatabase) getAllAccounts(ctx context.Context, ) (accounts []*domain.Account, err error) {
	// todo
}

func (db *firestoreDatabase) getAccountById(ctx context.Context, id domain.UniqueId) (*domain.Account, error) {
	// todo
}
