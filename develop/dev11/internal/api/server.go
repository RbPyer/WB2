package api

import (
	"dev11/internal/config"
	"dev11/internal/storage"
	"log"
	"net"
	"net/http"
)

type IServer interface {
	Start() error
	SetupHandlers()
	API
}

type API interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	EventsForDay(w http.ResponseWriter, r *http.Request)
	EventsForWeek(w http.ResponseWriter, r *http.Request)
	EventsForMonth(w http.ResponseWriter, r *http.Request)
}

type CurrentServer struct {
	config  *config.ServerConfig
	storage storage.IStorage
}

func (s *CurrentServer) Start() error {
	log.Println("trying to start dev11 http-server")
	err := http.ListenAndServe(net.JoinHostPort(s.config.Host, s.config.Port), nil)
	return err
}

func (s *CurrentServer) SetupHandlers() {
	http.HandleFunc("/create_event", LoggerMiddleware(s.CreateEvent))
	http.HandleFunc("/update_event", LoggerMiddleware(s.UpdateEvent))
	http.HandleFunc("/delete_event", LoggerMiddleware(s.DeleteEvent))
	http.HandleFunc("/events_for_day", LoggerMiddleware(s.EventsForDay))
	http.HandleFunc("/events_for_week", LoggerMiddleware(s.EventsForWeek))
	http.HandleFunc("/events_for_month", LoggerMiddleware(s.EventsForMonth))
}

func NewServer() IServer {
	return &CurrentServer{
		storage: storage.NewStorage(),
		config:  config.NewConfig(),
	}
}
