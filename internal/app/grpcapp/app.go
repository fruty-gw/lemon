package grpcapp

import (
	"fmt"
	"log"
	"net"

	lemonv1 "github.com/fruty-gw/lemon/gen/go/api/lemon/v1"
	grpchandler "github.com/fruty-gw/lemon/internal/transport/grpc"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(
	handlers *grpchandler.LemonHandler,
	port int,
) *App {

	gRPCServer := grpc.NewServer()
	lemonv1.RegisterLemonServiceServer(gRPCServer, handlers)

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	log.Println("Running on port", a.port)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	return a.gRPCServer.Serve(l)
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
