package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/r0nni3/go-service/models"
)

type message struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type dataMessage struct {
	message
	Data []*models.Presentation `json:"data"`
}

// This is the value of a valid API key, for testing.
const api_KEY string = "<API_KEY_HERE>"

func HandleRoute(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		api_key := r.Header.Get("APP_ID")
		var responseMsg *message
		if api_key == api_KEY {
			fn(w, r)
		} else {
			responseMsg = &message{Error: true, Message: "Invalid Request"}
			response, err := json.Marshal(responseMsg)
			if err != nil {
				internalError(w, err)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
		}
	}
}

func internalError(res http.ResponseWriter, err error) {
	log.Println(err)
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(err.Error()))
}
