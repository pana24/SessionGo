package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"session.com/config"
	"session.com/pkg/model"
)

var SessionWrt int

type input struct {
	Id string `json:"id"`
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input model.SessionInput
	json.NewDecoder(r.Body).Decode(&input)
	input.UserId = "user" + input.Token

	json.NewEncoder(w).Encode(input)
}

func FindSessions(w http.ResponseWriter, r *http.Request) {
	configuration := config.Configuration{}
	res, err := http.Get(configuration.Opa.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(body)
}

func FindSessionById(w http.ResponseWriter, r *http.Request) {

}

func DeleteSession(w http.ResponseWriter, r *http.Request) {

}
