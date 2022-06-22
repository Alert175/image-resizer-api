package utils

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

type Header struct {
	Field string
	Value string
}

type RequestArgs struct {
	Url           string
	Method        string
	SkipTlsVerify bool
	Body          *[]byte
	Headers       []Header
}

func Request(arg RequestArgs) ([]byte, error) {
	var reqBody []byte

	if arg.Body != nil {
		reqBody = *arg.Body
	}

	request, err := http.NewRequest(arg.Method, arg.Url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if arg.Headers != nil {
		for _, header := range arg.Headers {
			request.Header.Add(header.Field, header.Value)
		}
	}

	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
