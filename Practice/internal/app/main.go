package app

import (
	"Goprojects/Practice/internal/controller/http/router"
	"Goprojects/Practice/pkj/httpServer"
	"fmt"
	"os"
	"os/signal"
)

func Run() {
	UseCase := usecase.NewUseCase()
	rout := router.NewRouter()

	httpServ := httpServer.NewServer(rout)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	select {
	case s := <-interrupt:
		fmt.Println("Got signal:", s)
	}
}
