package internal

import (
	"context"
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
	"github.com/pkg/errors"
)

func (db *firestoreDatabase) storeAccount(ctx context.Context, account *domain.Account) error {
	// todo
	return errors.New("unimplemented")
}

func (db *firestoreDatabase) getAllAccounts(ctx context.Context, ) (accounts []*domain.Account, err error) {
	// todo
	return nil, errors.New("unimplemented")
}

func (db *firestoreDatabase) getAccountById(ctx context.Context, id domain.UniqueId) (*domain.Account, error) {
	// todo
	return nil, errors.New("unimplemented")
}
