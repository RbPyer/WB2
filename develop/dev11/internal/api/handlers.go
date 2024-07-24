package api

import (
	"dev11/internal/models"
	"dev11/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (s *CurrentServer) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	userId := r.FormValue("user_id")
	date := r.FormValue("date")

	if userId == "" || date == "" {
		http.Error(w, "no user-data or date-data", http.StatusBadRequest)
		return
	}

	event := models.Event{}

	userIdNum, err := strconv.Atoi(userId)

	if !utils.CheckValidDate(date, &event) || err != nil || userIdNum < 1 {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	err = s.storage.CreateEvent(userIdNum, event)
	if err != nil {
		utils.MakeJsonRespond(w, http.StatusInternalServerError, utils.JsonError(err.Error()))
		return
	}

	utils.MakeJsonRespond(w, 200, utils.JsonResult("new event successfully added"))
}

func (s *CurrentServer) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	userId := r.FormValue("user_id")
	date := r.FormValue("date")
	eventId := r.FormValue("event_id")

	if userId == "" || date == "" || eventId == "" {
		http.Error(w, "no user-data or event-data", http.StatusBadRequest)
		return
	}

	event := models.Event{UUID: eventId}

	userIdNum, err := strconv.Atoi(userId)

	if !utils.CheckValidDate(date, &event) || err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	err = s.storage.UpdateEvent(userIdNum, event)
	if err != nil {
		utils.MakeJsonRespond(w, http.StatusInternalServerError, utils.JsonError(err.Error()))
		return
	}

	utils.MakeJsonRespond(w, 200, utils.JsonResult(fmt.Sprintf("event #%s successfully updated", eventId)))
}

func (s *CurrentServer) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	userId := r.FormValue("user_id")
	eventId := r.FormValue("event_id")

	if userId == "" || eventId == "" {
		http.Error(w, "no user-data or event-data", http.StatusBadRequest)
		return
	}

	userIdNum, err := strconv.Atoi(userId)

	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	err = s.storage.DeleteEvent(userIdNum, eventId)

	if err != nil {
		utils.MakeJsonRespond(w, http.StatusInternalServerError, utils.JsonError(err.Error()))
		return
	}

	utils.MakeJsonRespond(w, 200, utils.JsonResult(fmt.Sprintf("event #%s successfully deleted", eventId)))
}

func (s *CurrentServer) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	userId := r.FormValue("user_id")
	day := r.FormValue("date")

	if userId == "" || day == "" {
		http.Error(w, "no user-data or date-data", http.StatusBadRequest)
	}
	tmpEvent := models.Event{}

	userIdNum, err := strconv.Atoi(userId)

	if err != nil || !utils.CheckValidDate(day, &tmpEvent) {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	response, err := s.storage.GetEventsDay(userIdNum, tmpEvent.Date)
	if err != nil {
		utils.MakeJsonRespond(w, http.StatusInternalServerError, utils.JsonError(err.Error()))
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.MakeJsonRespond(w, http.StatusInternalServerError, utils.JsonError(err.Error()))
		return
	}

	utils.MakeJsonRespond(w, 200, utils.JsonResult(string(jsonResponse)))
}

func (s *CurrentServer) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	userId := r.FormValue("user_id")

	if userId == "" {
		http.Error(w, "no user-data or month-data", http.StatusBadRequest)
	}

	userIdNum, err := strconv.Atoi(userId)

	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	response, err := s.storage.GetEventsWeek(userIdNum)
	if err != nil {
		utils.MakeJsonRespond(w, http.StatusInternalServerError, utils.JsonError(err.Error()))
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.MakeJsonRespond(w, http.StatusInternalServerError, utils.JsonError(err.Error()))
		return
	}

	utils.MakeJsonRespond(w, 200, utils.JsonResult(string(jsonResponse)))
}

func (s *CurrentServer) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	userId := r.FormValue("user_id")
	month := r.FormValue("month")

	if userId == "" || month == "" {
		http.Error(w, "no user-data or month-data", http.StatusBadRequest)
	}

	userIdNum, err := strconv.Atoi(userId)
	var monthT time.Time

	if err != nil || !utils.CheckValidMonth(month, &monthT) {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	response, err := s.storage.GetEventsMonth(userIdNum, monthT)
	if err != nil {
		utils.MakeJsonRespond(w, http.StatusInternalServerError, utils.JsonError(err.Error()))
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.MakeJsonRespond(w, http.StatusInternalServerError, utils.JsonError(err.Error()))
		return
	}

	utils.MakeJsonRespond(w, 200, utils.JsonResult(string(jsonResponse)))
}
