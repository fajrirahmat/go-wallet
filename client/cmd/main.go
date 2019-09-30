package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"google.golang.org/grpc/credentials"

	"github.com/fajrirahmat/go-wallet/commons"

	"github.com/fajrirahmat/go-wallet/accounts"

	"google.golang.org/grpc"
)

func main() {
	serverAddr := flag.String("server", "localhost:8081", "Server address")
	flag.Parse()
	b, _ := ioutil.ReadFile("../../samples/cert/ca.cert")
	//cp, _ := x509.SystemCertPool()
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		fmt.Println("credentials: failed to append certificates")
		return
	}

	certificate, err := tls.LoadX509KeyPair("../../samples/cert/server.pem", "../../samples/cert/server.key")
	if err != nil {
		fmt.Println("could not load client key pair: ", err)
		return
	}

	config := &tls.Config{
		ServerName:         "localhost",
		InsecureSkipVerify: false,
		RootCAs:            cp,
		Certificates:       []tls.Certificate{certificate},
	}

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := accounts.NewAccountServiceClient(conn)
	ctx := context.Background()
	currentDate := time.Now()
	response, err := client.CreateAccount(ctx, &accounts.CreateAccountRequest{
		BirthDate:   currentDate.Format(commons.IsoDateLayout),
		BirthPlace:  "Jakarta",
		Email:       "fajri.rahmat@yahoo.com",
		FullName:    "Fajri Rahmat",
		PhoneNumber: "081908911861",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Account No: " + response.AccountNo)
}
