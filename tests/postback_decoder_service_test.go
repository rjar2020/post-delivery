package tests

import (
	"testing"

	"github.com/rjar2020/post-delivery/service"

	"github.com/rjar2020/post-delivery/model"
)

const happyPathExample = `{
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

const happyPathExpectedURL = "http://sample_domain_endpoint.com/data?tittle=Gopher&image=https%3A%2F%2Fblog.golang.org%2Fgopher%2Fgopher.png&foo="

const noDataExample = `{
	"endpoint":{
	"method":"GET", 
	"url":"http://sample_domain_endpoint.com/data?tittle={icon}&image={coordinates}&foo={bar}"
	}, 
	"data":[]
	}`

const noDataExpectedURL = "http://sample_domain_endpoint.com/data?tittle=&image=&foo="

const nothingToReplaceExample = `{
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

const nothingToReplaceExpectedURL = "http://www.google.com"

func TestToURLForHappyPath(t *testing.T) {
	assertPostbackParsing(t, happyPathExample, happyPathExpectedURL)
}

func TestToURLWhenNoData(t *testing.T) {
	assertPostbackParsing(t, noDataExample, noDataExpectedURL)
}

func TestToURLWhenNothingToReplace(t *testing.T) {
	assertPostbackParsing(t, nothingToReplaceExample, nothingToReplaceExpectedURL)
}

func assertPostbackParsing(t *testing.T, example string, expectedResult string) {
	postBack, err := model.FromJSONtoPostback([]byte(example))
	if err == nil {
		url := service.ToURL(postBack)
		if url != expectedResult {
			t.Errorf("Parsed URL different to expected one. Expected: %v - Result: %v", expectedResult, url)
		}
	} else {
		t.Log("An error ocurred", err)
		t.Error(err)
	}
}
