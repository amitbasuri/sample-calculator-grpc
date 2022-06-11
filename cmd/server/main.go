package main

import (
	"log"
	"net"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "lovoo-assignment/pkg/gogen"
)

var (
	logger   = log.New(os.Stdout, "", 0)
	grpcPort = "8080"
)

type calculatorService struct {
	pb.UnimplementedCalculatorServiceServer
}

// Add takes two integers as input and returns their sum.
func (s *calculatorService) Add(ctx context.Context, req *pb.DoubleRequest) (*pb.SingleResponse, error) {
	return &pb.SingleResponse{
		Res: req.A + req.B,
	}, nil
}

// Sub takes two integers as input and returns their difference.
func (s *calculatorService) Sub(ctx context.Context, req *pb.DoubleRequest) (*pb.SingleResponse, error) {
	return &pb.SingleResponse{
		Res: req.A - req.B,
	}, nil
}

// Mul takes two integers as input and returns their product.
func (s *calculatorService) Mul(ctx context.Context, req *pb.DoubleRequest) (*pb.SingleResponse, error) {
	return &pb.SingleResponse{
		Res: req.A * req.B,
	}, nil
}

// Div takes two integers as input and returns their quotient.
// It returns error if the divisor is zero.
func (s *calculatorService) Div(ctx context.Context, req *pb.DoubleRequest) (*pb.SingleResponse, error) {
	if req.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot divide by zero")
	}
	return &pb.SingleResponse{
		Res: req.A / req.B,
	}, nil
}

// Mod takes two integers as input and returns the remainder.
// It returns error if the divisor is zero.
func (s *calculatorService) Mod(ctx context.Context, req *pb.DoubleRequest) (*pb.SingleResponse, error) {
	if req.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot divide by zero")
	}
	return &pb.SingleResponse{
		Res: req.A % req.B,
	}, nil
}

func startGRPCServer(hostPort string) error {
	listener, err := net.Listen("tcp", hostPort)
	if err != nil {
		return errors.Wrapf(err, "Failed to listen on %s: %v", hostPort, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, &calculatorService{})
	return grpcServer.Serve(listener)
}

func main() {

	grpcHostPort := net.JoinHostPort("0.0.0.0", grpcPort)

	go func() {
		err := startGRPCServer(grpcHostPort)
		if err != nil {
			logger.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	logger.Println("Server started...")
	select {}
}
