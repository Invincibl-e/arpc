package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	// "sync"
	"syscall"
	"time"

	"github.com/Invincibl-e/arpc"
	"github.com/Invincibl-e/arpc/extension/middleware/router"
)

func main() {
	server := arpc.NewServer()

	graceful := &router.Graceful{}

	// step 1: register graceful middleware
	server.Handler.Use(graceful.Handler())

	server.Handler.Handle("/echo", func(ctx *arpc.Context) {
		// delay 5s for you to shutdown server by `ctrl + c`
		time.Sleep(time.Second * 5)
		str := ""
		err := ctx.Bind(&str)
		ctx.Write(str)
		log.Printf("/echo: \"%v\", error: %v", str, err)
	}, true)

	go server.Run("localhost:8888")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// step 2: shutdown by graceful middleware
	graceful.Shutdown()

	// step 3: shutdown arpc server
	server.Shutdown(context.Background())
}
