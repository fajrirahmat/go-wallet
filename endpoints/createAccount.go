package endpoints

import (
	"context"

	"github.com/fajrirahmat/go-wallet/commons"

	"github.com/fajrirahmat/go-wallet/services"
	"github.com/go-kit/kit/endpoint"
)

//MakeCreateAccountEndpoint function to create endpoint to create new account
func MakeCreateAccountEndpoint(svc services.WalletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(commons.CreateAccountRequest)
		res, err := svc.CreateAccount(req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
