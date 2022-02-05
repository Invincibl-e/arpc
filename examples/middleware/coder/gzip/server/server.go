package main

import (
	"github.com/Invincibl-e/arpc"
	"github.com/Invincibl-e/arpc/extension/middleware/coder/gzip"
	"github.com/Invincibl-e/arpc/log"
)

func main() {
	svr := arpc.NewServer()

	svr.Handler.UseCoder(gzip.New(1024))

	// register router
	svr.Handler.Handle("/echo", func(ctx *arpc.Context) {
		ctx.Write(ctx.Body())
		log.Info("/echo")
	})

	svr.Run("localhost:8888")
}
