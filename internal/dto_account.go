package internal

import (
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
	"github.com/pkg/errors"
)

type AccountDto struct {
	Id         string  `json:"id" firestore:"-"`
	ClientId   string  `json:"clientId" firestore:"clientId"`
	BalanceUSD float32 `json:"balanceUsd" firestore:"balanceUsd"`
}

func newAccountDtoFromDomain(account *domain.Account) *AccountDto {
	return &AccountDto{
		Id:         account.Id.String(),
		ClientId:   account.ClientId.String(),
		BalanceUSD: account.Balance.Usd(),
	}
}

func (dto *AccountDto) ToDomain() (*domain.Account, error) {
	accountId, err := domain.ParseUniqueIdFromString(dto.Id)
	if err != nil {
		return nil, errors.Wrap(err, "parse account id")
	}
	clientId, err := domain.ParseUniqueIdFromString(dto.ClientId)
	if err != nil {
		return nil, errors.Wrap(err, "parse client id")
	}
	balance := domain.NewMoneyFromUsd(dto.BalanceUSD)
	account := &domain.Account{
		Id:       accountId,
		ClientId: clientId,
		Balance:  balance,
	}
	return account, account.Validate()
}
