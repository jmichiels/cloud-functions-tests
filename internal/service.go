package internal

import (
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

func (srv *service) createClient(client *domain.Client) error {
	return srv.db.runTransaction(func(tx repositories) error {
		// Check that the client does not already exists.
		if _, err := tx.getClientById(client.Id); err != errNotFound {
			if err == nil {
				return errors.New("a client with this id already exists")
			}
			return err
		}
		return tx.storeClient(client)
	})
}

func (srv *service) getAllClients() ([]*domain.Client, error) {
	return srv.db.getAllClients()
}

func (srv *service) getClientById(id domain.UniqueId) (*domain.Client, error) {
	return srv.db.getClientById(id)
}

func (srv *service) createAccount(account *domain.Account) error {
	return srv.db.runTransaction(func(tx repositories) error {
		// Check that the account does not already exists.
		if _, err := tx.getAccountById(account.Id); err != errNotFound {
			if err == nil {
				return errors.New("a account with this id already exists")
			}
			return err
		}
		// Check that the client exists (this is not part of the validation done in the domain).
		if _, err := tx.getClientById(account.ClientId); err != nil {
			return errors.Wrap(err, "get account client")
		}
		return tx.storeAccount(account)
	})
}

func (srv *service) getAllAccounts() ([]*domain.Account, error) {
	return srv.db.getAllAccounts()
}

func (srv *service) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	return srv.db.getAccountById(id)
}

func (srv *service) transfer(
	amount domain.Money,
	originAccountId domain.UniqueId,
	destinationAccountId domain.UniqueId,
) error {
	return srv.db.runTransaction(func(tx repositories) error {
		// First step: fetch all the data that we need from the repositories.
		originAccount, err := tx.getAccountById(originAccountId)
		if err != nil {
			return errors.Wrap(err, "get origin account")
		}
		destinationAccount, err := tx.getAccountById(destinationAccountId)
		if err != nil {
			return errors.Wrap(err, "get destination account")
		}
		// Second step: use our domain to apply our business logic.
		if err := originAccount.Transfer(amount, destinationAccount); err != nil {
			return err
		}
		// Third step: persist everything that has changed.
		if err := tx.storeAccount(originAccount); err != nil {
			return errors.Wrap(err, "store origin account")
		}
		if err := tx.storeAccount(destinationAccount); err != nil {
			return errors.Wrap(err, "store destination account")
		}
		return nil
	})
}
