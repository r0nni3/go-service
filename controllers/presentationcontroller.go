package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/r0nni3/go-service/models"
)

func PresentationsHandler(res http.ResponseWriter, req *http.Request) {
	dbCon := models.GetConString()
	var response []byte
	if req.Method == "GET" {
		var data dataMessage
		tmp, err := dbCon.GetPresentations()
		if err != nil {
			internalError(res, err)
			return
		}

		data.Data = tmp
		response, err = json.Marshal(data)
		if err != nil {
			internalError(res, err)
			return
		}
	} else if req.Method == "POST" {
		// Gets Presentation from request Body
		decoder := json.NewDecoder(req.Body)
		var p models.Presentation
		err := decoder.Decode(&p)
		if err != nil {
			internalError(res, err)
			return
		}

		// Inserts presentation to DB
		result, err := dbCon.InsertPresentation(&p)
		if err != nil {
			internalError(res, err)
			return
		}
		p.Id = result

		// Parses result to formal response to consumer
		response, err = json.Marshal(p)
		if err != nil {
			internalError(res, err)
			return
		}
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write(response)
}

func PresentationHandler(res http.ResponseWriter, req *http.Request) {
	dbCon := models.GetConString()
	var response []byte
	if req.Method == "GET" {
		var data dataMessage
		vars := mux.Vars(req)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		tmp, err := dbCon.GetPresentation(id)
		if err != nil {
			internalError(res, err)
			return
		}

		data.Data = append(data.Data, tmp)
		response, err = json.Marshal(data)
		if err != nil {
			internalError(res, err)
			return
		}
	} else if req.Method == "POST" {
		// Gets Presentation data from request Body
		decoder := json.NewDecoder(req.Body)
		var p models.Presentation
		err := decoder.Decode(&p)
		if err != nil {
			internalError(res, err)
			return
		}

		// Updates presentation
		err = dbCon.UpdatePresentation(&p)
		if err != nil {
			internalError(res, err)
			return
		}

		// Parses result to formal response to consumer
		response, err = json.Marshal(p)
		if err != nil {
			internalError(res, err)
			return
		}
	} else if req.Method == "DELETE" {
		vars := mux.Vars(req)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		err := dbCon.DeletePresentation(id)
		if err != nil {
			internalError(res, err)
			return
		}

		response, err = json.Marshal(&message{
			Error: false, Message: "Deleted product " + string(id)})
		if err != nil {
			internalError(res, err)
			return
		}
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write(response)
}
