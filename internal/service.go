package internal

import (
	"context"
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
	"github.com/pkg/errors"
)

type service struct {
	db database
}

func newService(db database) *service {
	return &service{
		db: db,
	}
}

func (srv *service) createClient(ctx context.Context, client *domain.Client) error {
	return srv.db.runTransaction(ctx, func(ctx context.Context, tx repositories) error {
		// Check that the client does not already exists.
		if _, err := tx.getClientById(ctx, client.Id); err != errNotFound {
			if err == nil {
				return errors.New("a client with this id already exists")
			}
			return err
		}
		return tx.storeClient(ctx, client)
	})
}

func (srv *service) getAllClients(ctx context.Context, ) ([]*domain.Client, error) {
	return srv.db.getAllClients(ctx)
}

func (srv *service) getClientById(ctx context.Context, id domain.UniqueId) (*domain.Client, error) {
	return srv.db.getClientById(ctx, id)
}

func (srv *service) createAccount(ctx context.Context, account *domain.Account) error {
	return srv.db.runTransaction(ctx, func(ctx context.Context, tx repositories) error {
		// Check that the account does not already exists.
		if _, err := tx.getAccountById(ctx, account.Id); err != errNotFound {
			if err == nil {
				return errors.New("a account with this id already exists")
			}
			return err
		}
		// Check that the client exists (this is not part of the validation done in the domain).
		if _, err := tx.getClientById(ctx, account.ClientId); err != nil {
			return errors.Wrap(err, "get account client")
		}
		return tx.storeAccount(ctx, account)
	})
}

func (srv *service) getAllAccounts(ctx context.Context, ) ([]*domain.Account, error) {
	return srv.db.getAllAccounts(ctx)
}

func (srv *service) getAccountById(ctx context.Context, id domain.UniqueId) (*domain.Account, error) {
	return srv.db.getAccountById(ctx, id)
}

func (srv *service) transfer(
	ctx context.Context,
	amount domain.Money,
	originAccountId domain.UniqueId,
	destinationAccountId domain.UniqueId,
) error {
	return srv.db.runTransaction(ctx, func(ctx context.Context, tx repositories) error {
		// First step: fetch all the data that we need from the repositories.
		originAccount, err := tx.getAccountById(ctx, originAccountId)
		if err != nil {
			return errors.Wrap(err, "get origin account")
		}
		destinationAccount, err := tx.getAccountById(ctx, destinationAccountId)
		if err != nil {
			return errors.Wrap(err, "get destination account")
		}
		// Second step: use our domain to apply our business logic.
		if err := originAccount.Transfer(amount, destinationAccount); err != nil {
			return err
		}
		// Third step: persist everything that has changed.
		if err := tx.storeAccount(ctx, originAccount); err != nil {
			return errors.Wrap(err, "store origin account")
		}
		if err := tx.storeAccount(ctx, destinationAccount); err != nil {
			return errors.Wrap(err, "store destination account")
		}
		return nil
	})
}
