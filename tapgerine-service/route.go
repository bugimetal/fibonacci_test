package tapgerine

import (
	"goji.io"
	"goji.io/pat"
)

func CreateRouter() *goji.Mux {
	router := goji.NewMux()
	router.HandleFunc(NewNamedPattern("fib", pat.Get("/fib/number/:number")), fibonacciHandler)
	router.HandleFunc(NewNamedPattern("fibevensum", pat.Get("/fib/evensum/:number")), fibonacciEvenSumHandler)
	return router
}
