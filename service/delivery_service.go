package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//DeliverPostback is accessing the URL using the specified HTTP method
func DeliverPostback(httpMethod string, URL string) (string, error) {
	var err error
	var resp *http.Response

	switch httpMethod {
	case "GET":
		resp, err = http.Get(URL)
	case "POST":
		resp, err = http.Post(URL, "application/json", nil)
	default:
		err = fmt.Errorf("Invalid HTTP method: " + httpMethod)
		return "", err
	}

	if err != nil {
		log.Printf("Error accessing the URL %v : %v", URL, err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error decoding the response for URL %v : %v", URL, err)
		return "", err
	}

	log.Printf("Respose from %v with method %v: %v", URL, httpMethod, string(body))
	return string(body), err
}
