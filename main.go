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
	"fmt"
	"visitor/core"
	"encoding/json"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"
	// получаем данные о посетителе
	coreVisitor := core.Visitor{Ua: ua, Ip: "79.104.42.249", Id: "f957a79770a7329a1b0f51780cc355b" + in.Name}
	visitor, err := coreVisitor.Identify()

	// если при определении информации о посетителе возникла ошибка
	if err != nil {
	}

	// упаковываем структуру в json
	jsonCode, err := json.Marshal(visitor)

	return &pb.HelloReply{Message: string(jsonCode)}, nil
}

//
func main() {

	// стартуем RPC сервер
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("failed start rpc-server: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: %v", err)
	}

	conf := config.New()
	runtime.GOMAXPROCS(conf.Cpu)

	// стартуем вебсервер
	apiHttp := api.Method{}
	http.HandleFunc("/api/visitor", apiHttp.Post)
	err = http.ListenAndServe(conf.Listen, nil)
	if err != nil {
		log.Fatalf("failed start web-server: %v", err)
	}
}