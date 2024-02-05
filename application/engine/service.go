package engine

import (
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/network/rpc"
	srv "github.com/alioth-center/restoration/application/service"
	"github.com/alioth-center/restoration/proto/pbg"
	"google.golang.org/grpc"
)

type service struct {
	srv *srv.RestorationService
}

func (s *service) Initialization() {
	if s.srv == nil {
		// service instance is nil, panic immediately
		panic("restoration service instance is nil")
	}
}

func (s *service) BindEngine(conn *grpc.Server) {
	pbg.RegisterRestorationServiceServer(conn, s.srv)
}

func NewService(log logger.Logger) rpc.Service {
	return &service{
		srv: srv.NewRestorationService(log),
	}
}
