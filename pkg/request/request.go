package request

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Get(rawUrl string, headers map[string]string, params url.Values) ([]byte, error) {
	structUrl, err := url.Parse(rawUrl)
	if err != nil {
		log.Printf("url %s Parse error: %v", rawUrl, err)
		return nil, err
	}
	structUrl.RawQuery = params.Encode()
	request, err := http.NewRequest("GET", structUrl.String(), nil)
	if err != nil {
		log.Printf("NewRequest error: %v", err)
		return nil, err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil || resp == nil || resp.StatusCode != 200 {
		log.Printf("Client.Do error: %v, req: %v", err, request)
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func PostForm(url string, headers map[string]string, data url.Values) ([]byte, error) {
	request, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		log.Printf("NewRequest error: %v", err)
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil || resp == nil || resp.StatusCode != 200 {
		log.Printf("Client.Do error: %v, req: %v", err, request)
		return nil, err
	}
	log.Printf("%v", request)
	log.Printf("%v", resp)
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func PostJSON(url string, headers map[string]string, jsonStr []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("NewRequest error: %v", err)
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil || resp == nil || resp.StatusCode != 200 {
		log.Printf("Client.Do error: %v, req: %v", err, request)
		return nil, err
	}
	fmt.Printf("%v", request)
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
