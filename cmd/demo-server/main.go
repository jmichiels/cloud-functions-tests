package main

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/jmichiels/cloud-functions-tests/internal"
	bank_v1 "github.com/jmichiels/cloud-functions-tests/pkg/bank/protobuf/bank/v1"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	log.Println("starting server...")
	// Determine port.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s\n", port)
	}
	// Create a listener on the port.
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v\n", err)
	}
	mux := cmux.New(listener)
	// Create a grpc listener that matches grpc header.
	grpcListener := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	// All the rest is assumed to be HTTP
	httpListener := mux.Match(cmux.Any())
	// Create the grpc server.
	grpcServer := grpc.NewServer()
	// Create the bank service.
	bankServiceServer := internal.NewBankServiceServer()
	// Attach the bank service to the server
	bank_v1.RegisterBankServiceServer(grpcServer, bankServiceServer)
	// Wrap the grpc server with its grpc-web proxy.
	wrappedGrpcServer := grpcweb.WrapServer(grpcServer)
	// Create the http server.
	httpServer := &http.Server{
		Handler: wrappedGrpcServer,
	}
	g := errgroup.Group{}
	g.Go(func() error {
		return grpcServer.Serve(grpcListener)
	})
	g.Go(func() error {
		return httpServer.Serve(httpListener)
	})
	g.Go(func() error {
		return mux.Serve()
	})
	// Start listening.
	log.Printf("listening on port %s\n", port)
	// Wait for any error.
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
