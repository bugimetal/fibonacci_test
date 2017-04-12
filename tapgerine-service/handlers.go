package tapgerine

import (
	"net/http"

	"tapgerine/tapgerine-service/fibonacci"
)

var (
	cache = fibonacci.CreateCache()
)

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	number, err := getNumberFromRequest(r)
	if err != nil {
		sendResponseWithStatus(w, nil, http.StatusBadRequest)
		return
	}
	result := fibonacci.WithCache(number, cache)

	sendResponseWithStatus(w, fibonacciResponse{N: number, Value: result}, http.StatusOK)
}

func fibonacciEvenSumHandler(w http.ResponseWriter, r *http.Request) {
	number, err := getNumberFromRequest(r)
	if err != nil {
		sendResponseWithStatus(w, nil, http.StatusBadRequest)
		return
	}

	sequence := fibonacci.Sequence(number)
	var result int
	for _, item := range sequence {
		if item%2 == 0 {
			result += item
		}
	}

	sendResponseWithStatus(w, fibonacciResponse{N: number, Value: result}, http.StatusOK)
}
