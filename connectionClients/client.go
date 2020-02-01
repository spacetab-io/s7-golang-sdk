package connectionClients

type Client interface {
	Request(soapAction string, request []byte, logAttributes map[string]string) ([]byte, error)
}
