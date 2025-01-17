package router

import (
	"time"

	"github.com/Invincibl-e/arpc"
	"github.com/Invincibl-e/arpc/log"
)

// Logger returns the logger middleware.
func Logger() arpc.HandlerFunc {
	return func(ctx *arpc.Context) {
		t := time.Now()

		ctx.Next()

		cmd := ctx.Message.Cmd()
		method := ctx.Message.Method()
		addr := ctx.Client.Conn.RemoteAddr()
		cost := time.Since(t).Milliseconds()

		switch cmd {
		case arpc.CmdRequest, arpc.CmdNotify:
			log.Info("'%v',\t%v,\t%v ms cost", method, addr, cost)
			break
		default:
			log.Error("invalid cmd: %d,\tdropped", cmd)
			ctx.Done()
			break
		}
	}
}
