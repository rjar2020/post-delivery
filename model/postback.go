package model

import (
	"encoding/json"
	"log"
)

//Postback maps a postback message comming from the ingestors
type Postback struct {
	Endpoint Endpoint
	Data     []map[string]interface{}
}

//FromJSONtoPostback is parsing a JSON payload into Postback object
func FromJSONtoPostback(payload []byte) (Postback, error) {
	var postBack Postback
	err := json.Unmarshal(payload, &postBack)
	if err != nil {
		log.Printf("Error decoding kafka message: %s", err)
	}
	return postBack, err
}
