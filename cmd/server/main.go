package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "calculator-server/pkg/gogen"
	"google.golang.org/grpc"
)

var (
	logger   = log.New(os.Stdout, "", 0)
	grpcPort = "8080"
)

func main() {
	grpcHostPort := net.JoinHostPort("0.0.0.0", grpcPort)

	listener, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", grpcHostPort, err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(grpcServer, newCalculatorService())

	go func() {
		exitSignal := make(chan os.Signal, 1)
		signal.Notify(exitSignal, os.Interrupt, syscall.SIGTERM)
		<-exitSignal
		logger.Println("Server exiting...")
		grpcServer.GracefulStop()
		os.Exit(1)
	}()

	logger.Println("Server started...")
	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Fatalf("failed to start gRPC server %s", err)

	}
}
