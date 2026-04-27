package app

import (
	"github.com/fruty-gw/lemon/internal/app/grpcapp"
	"github.com/fruty-gw/lemon/internal/service"
	grpchandler "github.com/fruty-gw/lemon/internal/transport/grpc"
)

type App struct {
	GRPCapp *grpcapp.App
}

func New() *App {
	a := &App{}

	lemonService := service.NewLemonService()
	lemonHandler := grpchandler.NewLemonHandler(lemonService)

	a.GRPCapp = grpcapp.New(
		lemonHandler,
		8982,
	)

	return a
}
