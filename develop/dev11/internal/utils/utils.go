package utils

import (
	"dev11/internal/models"
	"fmt"
	"log"
	"net/http"
	"time"
)

func JsonResult(respondText string) []byte {
	return []byte(fmt.Sprintf(`{"result": %s}`, respondText))
}

func JsonError(respondText string) []byte {
	return []byte(fmt.Sprintf(`{"error": "%s"}`, respondText))
}

func MakeJsonRespond(w http.ResponseWriter, code int, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func CheckValidDate(date string, event *models.Event) bool {
	value, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}
	event.Date = value
	return true
}

func CheckValidMonth(month string, result *time.Time) bool {
	monthTime, err := time.Parse("January", month)
	if err != nil {
		return false
	}
	*result = monthTime
	return true
}
