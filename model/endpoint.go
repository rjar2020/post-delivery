package model

//Endpoint it's part of Postback, which maps a postback message comming from the ingestors
type Endpoint struct {
	Method string
	URL    string
}
