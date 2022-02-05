package main

import (
	"github.com/Invincibl-e/arpc"
	"github.com/Invincibl-e/arpc/extension/middleware/coder/msgpack"
	"github.com/Invincibl-e/arpc/log"
)

func main() {
	svr := arpc.NewServer()

	svr.Handler.UseCoder(msgpack.New())

	// register router
	svr.Handler.Handle("/echo", func(ctx *arpc.Context) {
		ctx.Write(ctx.Body())
		log.Info("/echo, %v", ctx.Values())
	})

	svr.Run("localhost:8888")
}
