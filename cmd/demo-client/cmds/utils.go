package cmds

import (
	bank_v1 "github.com/jmichiels/cloud-functions-tests/pkg/bank/protobuf/bank/v1"
	"google.golang.org/grpc"
)

func withGrpcClient(f func(grpcClt bank_v1.BankServiceClient) error) error {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := bank_v1.NewBankServiceClient(conn)
	return f(client)
}
