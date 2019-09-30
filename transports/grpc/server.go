package grpc

import (
	"context"
	"os"

	"github.com/fajrirahmat/go-wallet/accounts"
	"github.com/fajrirahmat/go-wallet/endpoints"
	"github.com/go-kit/kit/log"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	createAccount grpcTransport.Handler
}

//NewWalletGRPCServer initiate wallet server
func NewWalletGRPCServer(endpoint endpoints.Endpoints) accounts.AccountServiceServer {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "instance_id", "WalletService")
	/*options := []grpcTransport.ServerOption{
		grpcTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}*/
	return &grpcServer{
		createAccount: grpcTransport.NewServer(
			endpoint.CreateAccount,
			decodeGRPCCreateAccountRequest,
			encodeGRPCCreateAccountResponse,
			//options,
		),
	}
}

//CreateAccount ....
func (s *grpcServer) CreateAccount(ctx context.Context, request *accounts.CreateAccountRequest) (*accounts.CreateAccountResponse, error) {
	_, resp, err := s.createAccount.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*accounts.CreateAccountResponse), nil
}
