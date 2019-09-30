package grpc

import (
	"context"
	"time"

	"github.com/fajrirahmat/go-wallet/accounts"
	"github.com/fajrirahmat/go-wallet/commons"
)

func decodeGRPCCreateAccountRequest(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(*accounts.CreateAccountRequest)
	birtDate, err := time.Parse(commons.IsoDateLayout, request.BirthDate)
	if err != nil {
		return nil, err
	}
	return commons.CreateAccountRequest{
		BirthDate:   birtDate,
		BirthPlace:  request.BirthPlace,
		Email:       request.Email,
		FullName:    request.FullName,
		PhoneNumber: request.PhoneNumber,
	}, nil
}

func encodeGRPCCreateAccountResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	response := resp.(*commons.CreateAccountResponse)
	return &accounts.CreateAccountResponse{
		AccountNo:   response.AccountNo,
		AccountType: response.AccountType,
	}, nil
}
