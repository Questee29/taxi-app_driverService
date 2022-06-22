package grpcServer

import (
	"net"

	pb "github.com/Questee29/taxi-app_driverService/proto/protob"
	"google.golang.org/grpc"
)

type Deps struct {
	OrderHandler pb.OrderGrpcServer
}
type Server struct {
	deps Deps
	srv  *grpc.Server
}

func NewServer(deps Deps) *Server {
	return &Server{
		srv:  grpc.NewServer(),
		deps: deps,
	}
}

func (s *Server) ListenAndServe(port string) error {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	pb.RegisterOrderGrpcServer(s.srv, s.deps.OrderHandler)
	if err := s.srv.Serve(lis); err != nil {
		return err
	}
	return nil
}
