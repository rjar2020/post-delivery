package service

import (
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/rjar2020/post-delivery/model"
)

//ToURL takes the endpoint and replace Data key tokens in it by Data values
func ToURL(postback model.Postback) string {
	urlTemplate := postback.Endpoint.URL
	for _, tokenValueMap := range postback.Data {
		for k, v := range tokenValueMap {
			v = url.PathEscape(v.(string))
			v = strings.ReplaceAll(v.(string), ":", "%3A")
			urlTemplate = strings.ReplaceAll(urlTemplate, "{"+k+"}", v.(string))
			log.Printf("Key: %v - Value: %v", k, v)
		}
	}
	re := regexp.MustCompile(`{.*}`)
	urlTemplate = re.ReplaceAllString(urlTemplate, "")
	log.Printf("Decoded URL: %#v", urlTemplate)
	return urlTemplate
}
