package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/COOPSPROFI/GoLangProject/pkg/models"
	"github.com/COOPSPROFI/GoLangProject/pkg/utils"
	"github.com/gorilla/mux"
)

var NewEvent models.Event

func GetEvents(w http.ResponseWriter, r *http.Request) {
	newEvents := models.GetEvents()
	res, _ := json.Marshal(newEvents)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEventById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["eventID"]
	ID, err := strconv.ParseInt(eventID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	eventDetails, _ := models.GetEventById(ID)
	res, _ := json.Marshal(eventDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	CreateEvent := &models.Event{}
	utils.ParseBody(r, CreateEvent)
	e := CreateEvent.CreateEvent()
	res, _ := json.Marshal(e)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["eventID"]
	ID, err := strconv.ParseInt(eventID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	event := models.DeleteEvent(ID)
	res, _ := json.Marshal(event)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var updateEvent = &models.Event{}
	utils.ParseBody(r, updateEvent)
	vars := mux.Vars(r)
	eventID := vars["eventID"]
	ID, err := strconv.ParseInt(eventID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	eventDetails, db := models.GetEventById(ID)
	if updateEvent.Name != "" {
		eventDetails.Name = updateEvent.Name
	}
	if updateEvent.Description != "" {
		eventDetails.Description = updateEvent.Description
	}
	if updateEvent.Img != "" {
		eventDetails.Img = updateEvent.Img
	}
	// if updateEvent.Date != time.Date() {
	// 	eventDetails.Date = updateEvent.Date
	// }
	db.Save(&eventDetails)
	res, _ := json.Marshal(eventDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
