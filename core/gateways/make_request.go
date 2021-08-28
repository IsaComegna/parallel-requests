package gateways

type MakeRequest interface {
	Request(string) ([]byte, error)
}
