package model

import (
	"encoding/json"
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
	return postBack, err
}
