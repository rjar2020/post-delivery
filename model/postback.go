package model

import (
	"encoding/json"
	"log"
)

//PostBack maps a postback message comming from the ingestors
type PostBack struct {
	Endpoint Endpoint
	Data     []map[string]interface{}
}

//FromJSONtoPostback is parsing a JSON payload into Postback object
func FromJSONtoPostback(payload []byte) (PostBack, error) {
	var postBack PostBack
	err := json.Unmarshal(payload, &postBack)
	if err != nil {
		log.Printf("Error decoding kafka message: %s", err)
	}
	return postBack, err
}
