package router

import (
	"github.com/Invincibl-e/arpc"
	"github.com/Invincibl-e/arpc/util"
)

// Recover returns the recovery middleware handler.
func Recover() arpc.HandlerFunc {
	return func(ctx *arpc.Context) {
		defer util.Recover()
		ctx.Next()
	}
}
