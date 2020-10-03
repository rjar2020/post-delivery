package service

import (
	"testing"

	"github.com/rjar2020/post-delivery/model"
)

func TestToURLForHappyPath(t *testing.T) {

	happyPathExample := `{
		"endpoint":{
		"method":"GET", 
		"url":"http://sample_domain_endpoint.com/data?tittle={icon}&image={coordinates}&foo={bar}"
		}, 
		"data":[
		{
			"icon":"Gopher", "coordinates":"https://blog.golang.org/gopher/gopher.png"
		},  
		{
			"blah":"bleh", "blih":"https://blog.hola.org/gopher/gopher.png"
		}]
		}`

	happyPathExpectedURL := "http://sample_domain_endpoint.com/data?tittle=Gopher&image=https%3A%2F%2Fblog.golang.org%2Fgopher%2Fgopher.png&foo="

	assertParsingFromPostbackRequestToExpectedResults(t, happyPathExample, happyPathExpectedURL)
}

func TestToURLWhenNoData(t *testing.T) {
	noDataExpectedExample := `{
		"endpoint":{
		"method":"GET", 
		"url":"http://sample_domain_endpoint.com/data?tittle={icon}&image={coordinates}&foo={bar}"
		}, 
		"data":[]
		}`

	noDataExpectedURL := "http://sample_domain_endpoint.com/data?tittle=&image=&foo="

	assertParsingFromPostbackRequestToExpectedResults(t, noDataExpectedExample, noDataExpectedURL)
}

func TestToURLWhenNothingToReplace(t *testing.T) {
	nothingToReplaceExample := `{
		"endpoint":{
		"method":"POST", 
		"url":"http://www.google.com"
		}, 
		"data":[
		{
			"icon":"Gopher", "coordinates":"https://blog.golang.org/gopher/gopher.png"
		},  
		{
			"blah":"bleh", "blih":"https://blog.hola.org/gopher/gopher.png"
		}]
		}`

	nothingToReplaceExpectedURL := "http://www.google.com"
	assertParsingFromPostbackRequestToExpectedResults(t, nothingToReplaceExample, nothingToReplaceExpectedURL)
}

func assertParsingFromPostbackRequestToExpectedResults(t *testing.T, example string, expectedResult string) {
	postBack, err := model.FromJSONtoPostback([]byte(example))
	if err == nil {
		url := ToURL(postBack)
		if url != expectedResult {
			t.Errorf("Parsed URL different to expected one. Expected: %v - Result: %v", expectedResult, url)
		}
	} else {
		t.Log("An error ocurred", err)
		t.Error(err)
	}
}
