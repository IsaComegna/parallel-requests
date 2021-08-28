package http

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const timeoutTime = 10

type GETRequest struct {
}

func NewGETRequest() *GETRequest {
	return &GETRequest{}
}

func (r GETRequest) Request(URL string) (response []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	netClient := &http.Client{
		Timeout: time.Second * timeoutTime,
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	response, err = ioutil.ReadAll(resp.Body)

	if resp.StatusCode > http.StatusCreated {
		fmt.Printf("API response code is not 200|201")

		return nil, errors.New("error in request")
	}

	return response, nil
}
