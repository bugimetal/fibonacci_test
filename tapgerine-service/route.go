package tapgerine

import (
	"net/http/pprof"

	"goji.io"
	"goji.io/pat"
)

func CreateRouter() *goji.Mux {
	router := goji.NewMux()
	router.HandleFunc(NewNamedPattern("fib", pat.Get("/fib/number/:number")), fibonacciHandler)
	router.HandleFunc(NewNamedPattern("fibevensum", pat.Get("/fib/evensum/:number")), fibonacciEvenSumHandler)

	// Register pprof handlers
	router.HandleFunc(NewNamedPattern("/debug/pprof/", pat.Get("/debug/pprof/")), pprof.Index)
	router.HandleFunc(NewNamedPattern("/debug/pprof/cmdline", pat.Get("/debug/pprof/cmdline")), pprof.Cmdline)
	router.HandleFunc(NewNamedPattern("/debug/pprof/profile", pat.Get("/debug/pprof/profile")), pprof.Profile)
	router.HandleFunc(NewNamedPattern("/debug/pprof/symbol", pat.Get("/debug/pprof/symbol")), pprof.Symbol)
	router.HandleFunc(NewNamedPattern("/debug/pprof/trace", pat.Get("/debug/pprof/trace")), pprof.Trace)

	return router
}
