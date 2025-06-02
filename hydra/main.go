package main

import (
	"fmt"
	"hydra/hlogger"
	"net/http"
)

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	logger.Println("Received an http request on root resource")

	fmt.Fprintf(w, "Welcome to the hydra software system")
}

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web server")

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}
