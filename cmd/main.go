package main

import (
	todo "github.com/hadelfi/auth-app"
	"github.com/hadelfi/auth-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s ", err.Error())
	}
}
