package core

import "github.com/fajrirahmat/go-wallet/commons"

//Service implementation for core wallet service
type Service struct{}

//NewService create new service
func NewService() *Service {
	return &Service{}
}

//CreateAccount implement create account
func (s *Service) CreateAccount(request commons.CreateAccountRequest) (*commons.CreateAccountResponse, error) {
	if ok, err := request.Validate(); !ok {
		return nil, err
	}
	return &commons.CreateAccountResponse{
		AccountNo:   "90010062653",
		AccountType: commons.MainAccountType,
	}, nil
}
