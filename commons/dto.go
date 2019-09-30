package commons

import (
	"time"
)

const (
	//MainAccountType primary account
	MainAccountType = "MAIN_ACCOUNT"
)

//Request and response
type (
	//CreateAccountRequest create account request payload
	CreateAccountRequest struct {
		FullName    string
		BirthPlace  string
		BirthDate   time.Time
		Email       string
		PhoneNumber string
	}
	//CreateAccountResponse response from service for create account process
	CreateAccountResponse struct {
		AccountNo   string
		AccountType string
	}
)

//Validate validate CreateAccountRequest
func (c CreateAccountRequest) Validate() (bool, error) {
	if c.FullName == "" {
		return false, ErrorFieldEmpty
	}
	if c.BirthPlace == "" {
		return false, ErrorFieldEmpty
	}
	if c.Email == "" {
		return false, ErrorFieldEmpty
	}
	if c.PhoneNumber == "" {
		return false, ErrorFieldEmpty
	}
	return true, nil
}
