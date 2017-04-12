package tapgerine

import (
	"encoding/json"
	"errors"
	"goji.io/pat"
	"net/http"
	"strconv"
)

func sendResponseWithStatus(w http.ResponseWriter, resp interface{}, statusCode int, headers ...map[string]string) {
	for _, header := range headers {
		for key, value := range header {
			w.Header().Add(key, value)
		}
	}
	if resp != nil {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(statusCode)
		bytes, _ := json.Marshal(resp)
		w.Write(bytes)
	} else {
		w.WriteHeader(statusCode)
	}
}

func getNumberFromRequest(r *http.Request) (int, error) {
	number := pat.Param(r, "number")
	result, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}
	if result <= 0 {
		return 0, errors.New("Number should be greater that 0")
	}
	return result, nil
}
