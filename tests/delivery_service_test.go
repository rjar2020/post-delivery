package tests

import (
	"testing"

	"github.com/rjar2020/post-delivery/service"
)

func TestDeliverPostbackGETHappyPath(t *testing.T) {
	response, err := service.DeliverPostback("GET", "https://www.google.com")
	if err != nil {
		t.Errorf("Error delivering postback. Method: GET - Error: %v", err)
	} else {
		t.Log(response)
	}
}

func TestDeliverPostbackWhenInvalidHttpMethod(t *testing.T) {
	response, err := service.DeliverPostback("BAD", "https://www.google.com")
	if err == nil {
		t.Errorf("It should be an error when delivering postback. Method: BAD - Response: %v", response)
	} else {
		t.Log(err)
	}
}

func TestDeliverPostbackWhenNonExistentSite(t *testing.T) {
	response, err := service.DeliverPostback("GET", "https://www.googlemeami.com")
	if err == nil {
		t.Errorf("It should be an error when delivering postback. Method: BAD - Response: %v", response)
	} else {
		t.Log(err)
	}
}

func TestDeliverPostbackPOSTHappyPath(t *testing.T) {
	response, err := service.DeliverPostback("POST", "https://www.google.com")
	if err != nil {
		t.Errorf("Error delivering postback. Method: POST - Error: %v", err)
	} else {
		t.Log(response)
	}
}
