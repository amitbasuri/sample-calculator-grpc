package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "lovoo-assignment/pkg/gogen"
)


type stringFlag struct {
	set   bool
	value string
}

func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

func (sf *stringFlag) String() string {
	return sf.value
}

var method stringFlag
var opA stringFlag
var opB stringFlag
var a int
var b int

var addr = flag.String("addr", "localhost:8080", "the address to connect to")

func init() {
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
	var handler = map[string]func(ctx context.Context, in *pb.DoubleRequest, opts ...grpc.CallOption) (*pb.SingleResponse, error){
		"add": c.Add,
		"sub": c.Sub,
		"mul": c.Mul,
		"div": c.Div,
		"mod": c.Mod,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := handler[method.value](ctx, &pb.DoubleRequest{A: int32(a), B: int32(b)})
	if err != nil {
		log.Printf("%v\n", err)
	} else {
		log.Printf("res: %v\n", r.GetRes())
	}
}
