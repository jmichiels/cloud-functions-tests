package cmds

import (
	"context"
	"fmt"
	bank_v1 "github.com/jmichiels/cloud-functions-tests/pkg/bank/protobuf/bank/v1"
	"github.com/spf13/cobra"
)

var (
	transferOriginId      string
	transferDestinationId string
	transferAmountUsd     float32
)

func init() {
	rootCmd.AddCommand(transferCmd)

	transferCmd.Flags().StringVar(&transferOriginId, "origin-id", "", "Origin account ID")
	transferCmd.Flags().StringVar(&transferDestinationId, "destination-id", "", "Destination account ID")
	transferCmd.Flags().Float32Var(&transferAmountUsd, "amount", 0, "Amount in USD")
}

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfers money between accounts.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return withGrpcClient(func(grpcClt bank_v1.BankServiceClient) error {
			_, err := grpcClt.Transfer(context.Background(), &bank_v1.TransferRequest{
				OriginAccountId:      transferOriginId,
				DestinationAccountId: transferDestinationId,
				AmountUsd:            transferAmountUsd,
			})
			if err != nil {
				return err
			}
			fmt.Printf("Transfered %.2f USD from account %s to %s\n", transferAmountUsd, transferOriginId, transferDestinationId)
			return nil
		})
	},
}
