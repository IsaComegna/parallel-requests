package usecases

import (
	"crypto/md5"
	"fmt"

	"github.com/IsaComegna/parallel-requests/core/gateways"
)

type GetHashResponse interface {
	Execute(string) (string, error)
}

type GetHashResponseConcrete struct {
	requestGateway gateways.MakeRequest
}

func NewGetHashResponse(requestGateway gateways.MakeRequest) GetHashResponseConcrete {
	return GetHashResponseConcrete{
		requestGateway: requestGateway,
	}
}

func (useCase GetHashResponseConcrete) Execute(URL string) (string, error) {
	response, err := useCase.requestGateway.Request(URL)
	if err != nil {
		return "", fmt.Errorf("error making request for URL %s: %w", URL, err)
	}

	hash := md5.Sum(response)

	return fmt.Sprintf("%x", hash), nil
}
