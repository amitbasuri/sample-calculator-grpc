package main

import (
	"context"

	pb "calculator-server/pkg/gogen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type calculatorService struct {
	pb.UnimplementedCalculatorServiceServer
}

func newCalculatorService() pb.CalculatorServiceServer {
	return &calculatorService{}
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
