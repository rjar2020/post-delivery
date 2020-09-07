package model

//Postback maps a postback message comming from the ingestors
type Postback struct {
	Endpoint Endpoint
	Data     []map[string]interface{}
}
