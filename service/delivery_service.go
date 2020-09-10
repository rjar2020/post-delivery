package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//DeliverPostback is accessing the URL using the specified HTTP method
func DeliverPostback(httpMethod string, URL string) (string, error) {
	var response string
	var err error
	var resp *http.Response

	switch httpMethod {
	case "GET":
		resp, err = http.Get(URL)
	case "POST":
		resp, err = http.Post(URL, "application/json", nil)
	default:
		response = "Invalid HTTP method: " + httpMethod
		err = fmt.Errorf(response)
		return response, err
	}

	if err != nil {
		response = "Error accessing the URL: %v"
		log.Printf(response, URL, err)
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response = "Error decoding the response for URL %v: "
		log.Printf(response, URL, err)
		return response, err
	}

	log.Printf("Respose from %v: %v", URL, body)
	return string(body), err
}
