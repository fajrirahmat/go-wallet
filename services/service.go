package services

import (
	"github.com/fajrirahmat/go-wallet/commons"
)

//WalletService wallet service interface
type WalletService interface {
	CreateAccount(commons.CreateAccountRequest) (*commons.CreateAccountResponse, error)
}
