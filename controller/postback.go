package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/rjar2020/post-delivery/env"
	"github.com/rjar2020/post-delivery/producer"

	"github.com/rjar2020/post-delivery/model"
)

type postbackController struct {
	postbackPattern *regexp.Regexp
}

func (pc postbackController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/postback" {
		switch r.Method {
		case http.MethodPost:
			log.Printf("New postback received")
			pb, err := pc.parseRequest(r)
			if err != nil {
				processBadRequest(w)
				return
			}
			log.Printf("Request: %#v", pb)
			jsonData, err := json.Marshal(pb)
			if err != nil {
				processBadRequest(w)
				return
			}
			err = producer.Produce(string(jsonData), os.Getenv(env.KafkaPostBackTopic))
			if err != nil {
				processBadRequest(w)
				return
			}
			w.Write([]byte("Postback delivered"))
			w.WriteHeader(http.StatusCreated)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func newPostbackController() *postbackController {
	return &postbackController{
		postbackPattern: regexp.MustCompile(`^/postback/?`),
	}
}

func (pc postbackController) parseRequest(r *http.Request) (model.PostBack, error) {
	buf, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
	}
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	log.Printf("BODY: %q", rdr1)
	dec := json.NewDecoder(rdr1)
	var pb model.PostBack
	err := dec.Decode(&pb)
	if err != nil {
		log.Printf("Could not parse postback object %#v", err)
		return model.PostBack{}, err
	}
	return pb, nil
}

func processBadRequest(w http.ResponseWriter) {
	message := "Could not parse postback object"
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
	log.Printf(message)
}
