package cmds

import (
	"context"
	"fmt"
	bank_v1 "github.com/jmichiels/cloud-functions-tests/pkg/bank/protobuf/bank/v1"
	"github.com/spf13/cobra"
)

var (
	accountId         string
	accountClientId   string
	accountBalanceUsd float32
)

func init() {
	rootCmd.AddCommand(accountsCmd)
}

var accountsCmd = &cobra.Command{
	Use: "accounts",
}

func init() {
	accountsCmd.AddCommand(accountsAddCmd)

	accountsAddCmd.Flags().StringVar(&accountClientId, "client-id", "", "Client ID")
	accountsAddCmd.Flags().Float32Var(&accountBalanceUsd, "balance", 0, "Balance in USD")
}

var accountsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new account",
	RunE: func(cmd *cobra.Command, args []string) error {
		return withGrpcClient(func(grpcClt bank_v1.BankServiceClient) error {
			response, err := grpcClt.CreateAccount(context.Background(), &bank_v1.CreateAccountRequest{
				Account: &bank_v1.Account{
					ClientId:   accountClientId,
					BalanceUsd: accountBalanceUsd,
				},
			})
			if err != nil {
				return err
			}
			account := response.Account
			fmt.Printf("Created account %s: clientId=%s, balanceUsd=%.2f\n", account.Id, account.Id, account.BalanceUsd)
			return nil
		})
	},
}

func init() {
	accountsCmd.AddCommand(accountsListCmd)
}

var accountsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return withGrpcClient(func(grpcClt bank_v1.BankServiceClient) error {
			response, err := grpcClt.GetAllAccounts(context.Background(), &bank_v1.GetAllAccountsRequest{})
			if err != nil {
				return err
			}
			for idx, account := range response.Accounts {
				fmt.Printf("%03d %s: clientId=%s, balanceUsd=%.2f\n", idx, account.Id, account.ClientId, account.BalanceUsd)
			}
			return nil
		})
	},
}

func init() {
	accountsCmd.AddCommand(accountsGetCmd)

	accountsGetCmd.Flags().StringVar(&accountId, "id", "", "Account ID")
}

var accountsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a single accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return withGrpcClient(func(grpcClt bank_v1.BankServiceClient) error {
			response, err := grpcClt.GetAccountById(context.Background(), &bank_v1.GetAccountByIdRequest{
				Id: accountId,
			})
			if err != nil {
				return err
			}
			account := response.Account
			fmt.Printf("%s: clientId=%s, balanceUsd=%.2f\n", account.Id, account.ClientId, account.BalanceUsd)
			return nil
		})
	},
}
