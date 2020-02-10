package clientsImplementation

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tmconsulting/s7-golang-sdk/configuration"
	"github.com/tmconsulting/s7-golang-sdk/publisher"
)

type TransportClient struct {
	LogsPublisher publisher.Publisher
	Config        configuration.Config
}

func NewTransportClient(p publisher.Publisher, c configuration.Config) *TransportClient {
	s := new(TransportClient)

	s.LogsPublisher = p
	s.Config = c

	return s
}

// Request sends request to S7 and returns response body
func (s *TransportClient) Request(soapAction string, request []byte, logAttributes map[string]string) ([]byte, error) {

	buffIO := bytes.NewReader(request)

	req, err := http.NewRequest(http.MethodPost, s.Config.Client.URL, buffIO)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(s.Config.Client.Timeout)*time.Second,
	)

	defer cancel()

	req.SetBasicAuth(s.Config.Secret.Login, s.Config.Secret.Password)
	req.Header.Add("Content-Type", "text/request")
	req.Header.Add("SOAPAction", fmt.Sprintf(`"http://api.s7.ru/%s"`, soapAction))
	req.Header.Set("X-API-Version", s.Config.Client.APIVersion)
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{s.Config.Certificate},
		},
	}

	client := &http.Client{Transport: tr}

	chanEnd := make(chan error)

	var response []byte

	go func() {
		res, err := client.Do(req)

		if err != nil {
			chanEnd <- err
			return
		}

		defer res.Body.Close()

		response, err = ioutil.ReadAll(res.Body)
		if err != nil {
			chanEnd <- err
			return
		}
		chanEnd <- nil

		if logAttributes != nil {
			err = s.LogsPublisher.PublishLogs(logAttributes, request, response)
			if err != nil {
				chanEnd <- err
			}
		}

		chanEnd <- err
	}()

	select {
	case <-ctx.Done():
		{
			<-chanEnd // Wait gorutine return.
			err = ctx.Err()
		}
	case <-chanEnd:
		{
			//return nil, err
		}
	}

	return response, nil
}
