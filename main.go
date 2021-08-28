package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"sync"

	"github.com/IsaComegna/parallel-requests/core/usecases"
	gateways "github.com/IsaComegna/parallel-requests/gateways/http"
)

var concurrencyDefaultValue = 10
var concurrencyFlagNotPresent = 0
var concurrencyMaxValue = 50

func main() {
	var (
		concurrency int
		URLs        []string
	)

	concurrencyPtr := flag.Int("parallel", concurrencyFlagNotPresent, "concurrency")
	flag.Parse()

	if *concurrencyPtr == concurrencyFlagNotPresent {
		URLs = os.Args[1:]
		concurrency = concurrencyDefaultValue
	} else {
		URLs = os.Args[3:]
		concurrency = *concurrencyPtr
	}

	if len(URLs) == 0 {
		return
	}

	requestGateway := gateways.NewGETRequest()
	getHashResponsesUseCase := usecases.NewGetHashResponse(requestGateway)

	if concurrency == 1 {
		runSequentially(URLs, getHashResponsesUseCase)
		return
	}

	if concurrency > len(URLs) {
		concurrency = len(URLs)
	}

	if concurrency > concurrencyMaxValue {
		concurrency = concurrencyMaxValue
	}

	var (
		ch = make(chan int, concurrencyMaxValue)
		wg sync.WaitGroup
	)

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for {
				a, ok := <-ch
				if !ok {
					wg.Done()
					return
				}

				makeRequests(URLs[a], getHashResponsesUseCase)
			}
		}()
	}

	for i := 0; i < len(URLs); i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}

func runSequentially(URLs []string, getHashResponsesUseCase usecases.GetHashResponse) {
	for i := 0; i < len(URLs); i++ {
		makeRequests(URLs[i], getHashResponsesUseCase)
	}
}

func makeRequests(URL string, getHashResponsesUseCase usecases.GetHashResponse) {
	resp, err := getHashResponsesUseCase.Execute(URL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(URL, resp)
}
