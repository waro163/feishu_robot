package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Body []byte
	Code int
}

func Request(url string, method string, payload interface{}, headers ...map[string]string) (Response, error) {

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("marshal payload occur error: %s\n", err)
		return Response{}, err
	}
	body := bytes.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return Response{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	if len(headers) > 0 {
		for key, value := range headers[0] {
			req.Header.Set(key, value)
		}
	}

	response, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}
	// log.Printf("%d, %s", response.StatusCode, string(resBody))
	return Response{resBody, response.StatusCode}, nil
}

func Get(url string, payload interface{}, headers ...map[string]string) (Response, error) {
	return Request(url, "GET", payload, headers...)
}

func Post(url string, payload interface{}, headers ...map[string]string) (Response, error) {
	return Request(url, "POST", payload, headers...)
}

func Put(url string, payload interface{}, headers ...map[string]string) (Response, error) {
	return Request(url, "PUT", payload, headers...)
}

func Patch(url string, payload interface{}, headers ...map[string]string) (Response, error) {
	return Request(url, "PATCH", payload, headers...)
}

func Delete(url string, payload interface{}, headers ...map[string]string) (Response, error) {
	return Request(url, "DELETE", payload, headers...)
}
