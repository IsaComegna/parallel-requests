package usecases

import (
	"crypto/md5"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/IsaComegna/parallel-requests/mocks"
)

func responseFixture() []byte {
	return []byte("Some response")
}

func TestGetHashResponseConcrete_Execute(t *testing.T) {
	t.Run("when there is no error making request it returns the hash", func(t *testing.T) {
		response := responseFixture()
		requestGatewayMock := &mocks.MakeRequest{}
		requestGatewayMock.
			On("Request", mock.AnythingOfType("string")).
			Return(response, nil)
		agreementUseCase := NewGetHashResponse(requestGatewayMock)

		hashResponse, err := agreementUseCase.Execute("http://some-url.com")

		expectedResponse := fmt.Sprintf("%x", md5.Sum(response))
		assert.Nil(t, err)
		requestGatewayMock.AssertCalled(t, "Request", "http://some-url.com")
		assert.Equal(t, expectedResponse, hashResponse)
	})

	t.Run("when there is an error making request it returns the wrapped error", func(t *testing.T) {
		response := responseFixture()
		err := errors.New("foo")
		requestGatewayMock := &mocks.MakeRequest{}
		requestGatewayMock.
			On("Request", mock.AnythingOfType("string")).
			Return(response, err)
		agreementUseCase := NewGetHashResponse(requestGatewayMock)

		hashResponse, err := agreementUseCase.Execute("http://some-url.com")

		assert.Equal(t, "", hashResponse)
		assert.NotNil(t, err)
		assert.Equal(t, "error making request for URL http://some-url.com: foo", err.Error())
		requestGatewayMock.AssertCalled(t, "Request", "http://some-url.com")
	})
}
