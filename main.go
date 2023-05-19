package main

import (
	"go-learinng/basics"
	"go-learinng/concurrency"
	"go-learinng/simpleHttp"
	"go-learinng/tools"
	"log"
	"net/http"
	"time"
)

func main() {
	basics.Examples()
	concurrency.Examples()
	simpleHttp.Examples()
	err := tools.MakeLoad(1000000, 250, 1*time.Second, "http://localhost:8090/api/v1/symbols/AAPL", http.MethodGet, nil)
	if err != nil {
		log.Panic("Error on load testing!", err)
	}
}
