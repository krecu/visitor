package main

import (
	"runtime"
	"visitor/conf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "visitor/rpc"
	logger "visitor/log"
	"google.golang.org/grpc/reflection"
	"net"
	"visitor/core"
	"encoding/json"
)

type server struct{}

// основной метод получения данных пользователя
func (s *server) GetVisitor(ctx context.Context, in *pb.VisitorRequest) (*pb.VisitorReply, error) {

	// получаем данные о посетителе
	coreVisitor := core.Visitor{Ua: in.Ua, Ip: in.Ip, Id: in.Id}
	visitor, err := coreVisitor.Identify()

	// если при определении информации о посетителе возникла ошибка
	if err != nil {
	}

	// упаковываем структуру в json
	jsonCode, err := json.Marshal(visitor)

	if err != nil {
		return &pb.VisitorReply{Status: "false", Body: err.Error()}, nil
	}

	return &pb.VisitorReply{Status: "ok", Body: string(jsonCode)}, nil
}

//
func main() {

	conf := config.New()

	// юзаем заданное число процессоров
	runtime.GOMAXPROCS(conf.Cpu)

	// вешаем листнера на порт
	lis, err := net.Listen("tcp", conf.Grpc)
	if err != nil {

		logger.Notify(logger.Message{
			ShortMessage:"failed start rpc-server: " + err.Error(),
			State: "error",
		})

	}

	// стартуем RPC сервер
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {

		logger.Notify(logger.Message{
			ShortMessage:"failed start rpc-server: " + err.Error(),
			State: "error",
		})

	}
}