package cmds

import (
	"context"
	"fmt"
	bank_v1 "github.com/jmichiels/cloud-functions-tests/pkg/bank/protobuf/bank/v1"
	"github.com/spf13/cobra"
)

var (
	clientId        string
	clientFirstName string
	clientLastName  string
	clientAge       uint32
)

func init() {
	rootCmd.AddCommand(clientsCmd)
}

var clientsCmd = &cobra.Command{
	Use: "clients",
}

func init() {
	clientsCmd.AddCommand(clientsAddCmd)

	clientsAddCmd.Flags().StringVar(&clientFirstName, "first-name", "", "First name")
	clientsAddCmd.Flags().StringVar(&clientLastName, "last-name", "", "Last name")
	clientsAddCmd.Flags().Uint32Var(&clientAge, "age", 0, "Age")
}

var clientsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new client",
	RunE: func(cmd *cobra.Command, args []string) error {
		return withGrpcClient(func(grpcClt bank_v1.BankServiceClient) error {
			response, err := grpcClt.CreateClient(context.Background(), &bank_v1.CreateClientRequest{
				Client: &bank_v1.Client{
					FirstName: clientFirstName,
					LastName:  clientLastName,
					Age:       clientAge,
				},
			})
			if err != nil {
				return err
			}
			client := response.Client
			fmt.Printf("Created client %s: %s %s (age %d)\n", client.Id, client.FirstName, client.LastName, client.Age)
			return nil
		})
	},
}

func init() {
	clientsCmd.AddCommand(clientsListCmd)
}

var clientsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all clients",
	RunE: func(cmd *cobra.Command, args []string) error {
		return withGrpcClient(func(grpcClt bank_v1.BankServiceClient) error {
			response, err := grpcClt.GetAllClients(context.Background(), &bank_v1.GetAllClientsRequest{})
			if err != nil {
				return err
			}
			for idx, client := range response.Clients {
				fmt.Printf("%03d %s: %s %s (age %d)\n", idx, client.Id, client.FirstName, client.LastName, client.Age)
			}
			return nil
		})
	},
}

func init() {
	clientsCmd.AddCommand(clientsGetCmd)

	clientsGetCmd.Flags().StringVar(&clientId, "id", "", "Client ID")
}

var clientsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a single clients",
	RunE: func(cmd *cobra.Command, args []string) error {
		return withGrpcClient(func(grpcClt bank_v1.BankServiceClient) error {
			response, err := grpcClt.GetClientById(context.Background(), &bank_v1.GetClientByIdRequest{
				Id: clientId,
			})
			if err != nil {
				return err
			}
			client := response.Client
			fmt.Printf("%s: %s %s (age %d)\n", client.Id, client.FirstName, client.LastName, client.Age)
			return nil
		})
	},
}
