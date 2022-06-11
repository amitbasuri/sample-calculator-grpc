package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	pb "calculator-server/pkg/gogen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	method stringFlag
	opA    stringFlag
	opB    stringFlag
	a      int
	b      int
	addr   *string
)

type HandlerFunc func(ctx context.Context, in *pb.DoubleRequest, opts ...grpc.CallOption) (*pb.SingleResponse, error)

func init() {
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
	flag.Var(&method, "method", "options: add,sub,mul,div,mod")
	flag.Var(&opA, "a", "first operand")
	flag.Var(&opB, "b", "second operand")

	flag.Parse()

	if !method.set {
		log.Fatal("please enter method")
	}
	if !opA.set {
		log.Fatal("please operand a")
	}
	if !opB.set {
		log.Fatal("please operand b")
	}
	var err error
	a, err = strconv.Atoi(opA.value)
	if err != nil {
		log.Fatal("operand a not integer")
	}

	b, err = strconv.Atoi(opB.value)
	if err != nil {
		log.Fatal("operand b not integer")
	}
}

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	var handlers = map[string]HandlerFunc{
		"add": c.Add,
		"sub": c.Sub,
		"mul": c.Mul,
		"div": c.Div,
		"mod": c.Mod,
	}

	handler, ok := handlers[method.value]
	if !ok {
		log.Printf("%v\n", "invalid method")
		return
	}

	handleRequest(handler, a, b)
}

func handleRequest(handler HandlerFunc, a, b int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := handler(ctx, &pb.DoubleRequest{A: int32(a), B: int32(b)})
	if err != nil {
		log.Printf("%v\n", err)
	} else {
		log.Printf("res: %v\n", r.GetRes())
	}
}
