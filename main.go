package main

import (
	"net/http"
	"visitor/api"
	"runtime"
	"visitor/conf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "visitor/rpc"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

//
func main() {

	conf := config.New()
	runtime.GOMAXPROCS(conf.Cpu)

	// стартуем вебсервер
	apiHttp := api.Method{}
	http.HandleFunc("/api/visitor", apiHttp.Post)
	err := http.ListenAndServe(conf.Listen, nil)
	if err != nil {
		log.Fatalf("failed start web-server: %v", err)
	}

	// стартуем RPC сервер
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed start rpc-server: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}