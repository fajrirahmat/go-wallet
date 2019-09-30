package cert

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/fajrirahmat/go-wallet/config"

	"google.golang.org/grpc/credentials"
)

type certificate struct {
	Certificate     tls.Certificate
	GRPCCredentials credentials.TransportCredentials
}

//App application certificate
var App *certificate

//Init initiate value of application certificate
func Init() {
	App = &certificate{}
	certificate, err := tls.LoadX509KeyPair(config.App.Certificate.ServerCertificatePath, config.App.Certificate.ServerPrivateKeyPath)
	if err != nil {
		log.Fatalf("Failed to read certificate %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(config.App.Certificate.CACertificatePath)
	if err != nil {
		log.Fatalf("Failed to read certificate %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("Failed to read certificate %v", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
		RootCAs:      certPool,
	})

	App.Certificate = certificate
	App.GRPCCredentials = creds

}
