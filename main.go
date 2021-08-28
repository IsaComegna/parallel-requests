package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"

	"github.com/IsaComegna/parallel-requests/core/usecases"
	gateways "github.com/IsaComegna/parallel-requests/gateways/http"
)

var concurrencyDefaultValue = 0

func main() {
	concurrencyPtr := flag.Int("parallel", concurrencyDefaultValue, "concurrency")
	flag.Parse()

	var URLs []string

	if *concurrencyPtr != concurrencyDefaultValue {
		URLs = os.Args[3:]
	} else {
		URLs = os.Args[1:]
	}

	requestGateway := gateways.NewGETRequest()
	getHashResponsesUseCase := usecases.NewGetHashResponse(requestGateway)

	for _, URL := range URLs {
		resp, err := getHashResponsesUseCase.Execute(URL)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s %s", URL, resp)
	}
}
