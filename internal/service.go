package internal

import (
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
	"github.com/pkg/errors"
)

type service struct {
	repo repository
}

func newService(repo repository) *service {
	return &service{
		repo: repo,
	}
}

func (srv *service) storeClient(client *domain.Client) error {
	return srv.repo.storeClient(client)
}

func (srv *service) getAllClients() ([]*domain.Client, error) {
	return srv.repo.getAllClients()
}

func (srv *service) GetClientById(id domain.UniqueId) (*domain.Client, error) {
	return srv.repo.getClientById(id)
}

func (srv *service) storeAccount(account *domain.Account) error {
	return srv.repo.runTransaction(func(tx transaction) error {
		// Check that the client exists (this is not part of the validation done in the domain).
		if _, err := tx.getClientById(account.ClientId); err != nil {
			return errors.Wrap(err, "get account client")
		}
		return tx.storeAccount(account)
	})
}

func (srv *service) getAllAccounts() ([]*domain.Account, error) {
	return srv.repo.getAllAccounts()
}

func (srv *service) getAccountById(id domain.UniqueId) (*domain.Account, error) {
	return srv.repo.getAccountById(id)
}

func (srv *service) transfer(
	amount domain.Money,
	originAccountId domain.UniqueId,
	destinationAccountId domain.UniqueId,
) error {
	return srv.repo.runTransaction(func(tx transaction) error {
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
