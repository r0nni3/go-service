package controllers

import (
	"encoding/json"
	"net/http"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	responseMsg := &message{Error: false, Message: "Success"}
	response, err := json.Marshal(responseMsg)
	if err != nil {
		internalError(res, err)
		return
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write(response)
}
