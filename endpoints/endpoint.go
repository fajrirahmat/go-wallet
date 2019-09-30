package endpoints

import (
	"github.com/fajrirahmat/go-wallet/services"
	"github.com/go-kit/kit/endpoint"
)

//Endpoints wallet endpoint
type Endpoints struct {
	CreateAccount endpoint.Endpoint
}

//New create new endpoints object
func New(svc services.WalletService) *Endpoints {
	return &Endpoints{
		CreateAccount: MakeCreateAccountEndpoint(svc),
	}
}
