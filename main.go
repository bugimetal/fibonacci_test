package main

import (
	"net/http"
	"runtime"

	"tapgerine/tapgerine-service"

	"github.com/Sirupsen/logrus"
)

// Main Program
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logger := logrus.New()
	logger.Info("Application working on port 8080")
	logger.Fatal(http.ListenAndServe(":8080", tapgerine.CreateRouter()))
}
