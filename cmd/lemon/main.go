package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fruty-gw/lemon/internal/app"
)

func main() {
	application := app.New()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		application.GRPCapp.MustRun()
	}()

	<-stop

	application.GRPCapp.Stop()
}
