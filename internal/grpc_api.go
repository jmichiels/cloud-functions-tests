package internal

import (
	"context"
	"github.com/jmichiels/cloud-functions-tests/internal/domain"
	bank_v1 "github.com/jmichiels/cloud-functions-tests/pkg/bank/protobuf/bank/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcApi struct {
	bank_v1.UnimplementedBankServiceServer
	srv *service
}

func newGrpcApi(srv *service) *grpcApi {
	return &grpcApi{
		srv: srv,
	}
}

// Returns a new bank_v1.BankServiceServer implementation.
func NewBankServiceServer() bank_v1.BankServiceServer {
	db := newMockDatabase()
	service := newService(db)
	grpcApi := newGrpcApi(service)
	return grpcApi
}

func (api *grpcApi) CreateClient(ctx context.Context, request *bank_v1.CreateClientRequest) (*bank_v1.CreateClientResponse, error) {
	client := &domain.Client{
		Id:        domain.GenerateRandomUniqueId(),
		FirstName: request.Client.FirstName,
		LastName:  request.Client.LastName,
		Age:       uint(request.Client.Age),
	}
	if err := client.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := api.srv.createClient(client); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &bank_v1.CreateClientResponse{
		Client: &bank_v1.Client{
			Id:        client.Id.String(),
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Age:       uint32(client.Age),
		},
	}, nil
}

func (api *grpcApi) GetAllClients(ctx context.Context, request *bank_v1.GetAllClientsRequest) (*bank_v1.GetAllClientsResponse, error) {
	clients, err := api.srv.getAllClients()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &bank_v1.GetAllClientsResponse{
		Clients: make([]*bank_v1.Client, len(clients)),
	}
	for i, clt := range clients {
		response.Clients[i] = &bank_v1.Client{
			Id:        clt.Id.String(),
			FirstName: clt.FirstName,
			LastName:  clt.LastName,
			Age:       uint32(clt.Age),
		}
	}
	return response, err
}

func (api *grpcApi) GetClientById(ctx context.Context, request *bank_v1.GetClientByIdRequest) (*bank_v1.GetClientByIdResponse, error) {
	id, err := domain.ParseUniqueIdFromString(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	client, err := api.srv.getClientById(id)
	if err != nil {
		var code codes.Code
		if err == errNotFound {
			code = codes.NotFound
		} else {
			code = codes.Internal
		}
		return nil, status.Error(code, err.Error())
	}
	return &bank_v1.GetClientByIdResponse{
		Client: &bank_v1.Client{
			Id:        client.Id.String(),
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Age:       uint32(client.Age),
		},
	}, err
}

func (api *grpcApi) CreateAccount(ctx context.Context, request *bank_v1.CreateAccountRequest) (*bank_v1.CreateAccountResponse, error) {
	clientId, err := domain.ParseUniqueIdFromString(request.Account.ClientId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	account := &domain.Account{
		Id:       domain.GenerateRandomUniqueId(),
		ClientId: clientId,
		Balance:  domain.NewMoneyFromUsd(request.Account.BalanceUsd),
	}
	if err := account.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := api.srv.createAccount(account); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &bank_v1.CreateAccountResponse{
		Account: &bank_v1.Account{
			Id:         account.Id.String(),
			ClientId:   account.ClientId.String(),
			BalanceUsd: account.Balance.Usd(),
		},
	}, nil
}

func (api *grpcApi) GetAllAccounts(ctx context.Context, request *bank_v1.GetAllAccountsRequest) (*bank_v1.GetAllAccountsResponse, error) {
	accounts, err := api.srv.getAllAccounts()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &bank_v1.GetAllAccountsResponse{
		Accounts: make([]*bank_v1.Account, len(accounts)),
	}
	for i, clt := range accounts {
		response.Accounts[i] = &bank_v1.Account{
			Id:         clt.Id.String(),
			ClientId:   clt.ClientId.String(),
			BalanceUsd: clt.Balance.Usd(),
		}
	}
	return response, err
}

func (api *grpcApi) GetAccountById(ctx context.Context, request *bank_v1.GetAccountByIdRequest) (*bank_v1.GetAccountByIdResponse, error) {
	id, err := domain.ParseUniqueIdFromString(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	account, err := api.srv.getAccountById(id)
	if err != nil {
		var code codes.Code
		if err == errNotFound {
			code = codes.NotFound
		} else {
			code = codes.Internal
		}
		return nil, status.Error(code, err.Error())
	}
	return &bank_v1.GetAccountByIdResponse{
		Account: &bank_v1.Account{
			Id:         account.Id.String(),
			ClientId:   account.ClientId.String(),
			BalanceUsd: account.Balance.Usd(),
		},
	}, err
}

func (api *grpcApi) Transfer(ctx context.Context, request *bank_v1.TransferRequest) (*bank_v1.TransferResponse, error) {
	originAccountId, err := domain.ParseUniqueIdFromString(request.OriginAccountId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	destinationAccountId, err := domain.ParseUniqueIdFromString(request.DestinationAccountId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := api.srv.transfer(domain.NewMoneyFromUsd(request.AmountUsd), originAccountId, destinationAccountId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &bank_v1.TransferResponse{}, nil
}
