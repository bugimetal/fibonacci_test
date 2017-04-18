package tapgerine

import (
	"net/http"

	"fibonacci_test/tapgerine-service/fibonacci"
)

var (
	cache       = fibonacci.CreateCache()
	cacheMatrix = fibonacci.CreateMatrixCache()
)

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	number, err := getNumberFromRequest(r)
	if err != nil {
		sendResponseWithStatus(w, nil, http.StatusBadRequest)
		return
	}

	result := fibonacci.Qmatrix(int(number), cacheMatrix)

	sendResponseWithStatus(w, result.String(), http.StatusOK)
}

func fibonacciEvenSumHandler(w http.ResponseWriter, r *http.Request) {
	number, err := getNumberFromRequest(r)
	if err != nil {
		sendResponseWithStatus(w, nil, http.StatusBadRequest)
		return
	}

	sequence := fibonacci.Sequence(int(number))
	var result int
	for _, item := range sequence {
		if item%2 == 0 {
			result += item
		}
	}

	sendResponseWithStatus(w, fibonacciResponse{N: number, Value: string(result)}, http.StatusOK)
}
