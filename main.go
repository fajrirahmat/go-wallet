package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/fajrirahmat/go-wallet/cert"

	"github.com/fajrirahmat/go-wallet/config"

	"github.com/fajrirahmat/go-wallet/accounts"

	"google.golang.org/grpc"

	walletGrpc "github.com/fajrirahmat/go-wallet/transports/grpc"

	"github.com/fajrirahmat/go-wallet/endpoints"

	"github.com/fajrirahmat/go-wallet/core"
	"github.com/fajrirahmat/go-wallet/services"
)

func main() {
	//load and initiate configuration
	config.Init()

	//initiate certificate
	cert.Init()

	var svc services.WalletService
	svc = core.NewService()

	errChan := make(chan error)

	endpoint := endpoints.New(svc)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.App.Server.Host, config.App.Server.Port))
		if err != nil {
			errChan <- err
			return
		}
		handler := walletGrpc.NewWalletGRPCServer(*endpoint)
		grpcServer := grpc.NewServer(grpc.Creds(cert.App.GRPCCredentials))
		accounts.RegisterAccountServiceServer(grpcServer, handler)
		fmt.Println("Listen on :", fmt.Sprintf("%s:%d", config.App.Server.Host, config.App.Server.Port))
		errChan <- grpcServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
