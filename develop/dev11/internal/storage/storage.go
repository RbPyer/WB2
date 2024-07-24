package storage

import (
	"dev11/internal/models"
	"dev11/internal/processor"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type IStorage interface {
	CreateEvent(userId int, event models.Event) error
	UpdateEvent(userId int, event models.Event) error
	DeleteEvent(userId int, eventId string) error
	GetEventsDay(userId int, day time.Time) (map[string]models.Event, error)
	GetEventsWeek(userId int) (map[string]models.Event, error)
	GetEventsMonth(userId int, month time.Time) (map[string]models.Event, error)
	checkUser(userId int) bool
}

func NewStorage() IStorage {
	return &Storage{
		data: make(map[int]map[string]models.Event),
		rw:   new(sync.RWMutex),
		proc: processor.NewProcessor(),
	}
}

type Storage struct {
	data map[int]map[string]models.Event
	rw   *sync.RWMutex
	proc processor.IProcessor
}

func (s *Storage) CreateEvent(userId int, event models.Event) error {
	s.rw.Lock()
	defer s.rw.Unlock()
	event.UUID = uuid.New().String()
	if !s.checkUser(userId) {
		s.data[userId] = make(map[string]models.Event)
	}
	if _, ok := s.data[userId][event.UUID]; ok {
		return errors.New(fmt.Sprintf("event with id %s already exists", event.UUID))
	}
	s.data[userId][event.UUID] = event
	log.Printf("New event with id %s created\n", event.UUID)
	return nil
}

func (s *Storage) UpdateEvent(userId int, event models.Event) error {
	s.rw.Lock()
	defer s.rw.Unlock()
	if !s.checkUser(userId) {
		return errors.New("user not found")
	}
	if _, ok := s.data[userId][event.UUID]; !ok {
		return errors.New("event not found")
	}
	s.data[userId][event.UUID] = event
	return nil
}

func (s *Storage) DeleteEvent(userId int, eventId string) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if !s.checkUser(userId) {
		return errors.New("user not found")
	}
	delete(s.data[userId], eventId)
	return nil
}

func (s *Storage) GetEventsDay(userId int, day time.Time) (map[string]models.Event, error) {
	s.rw.RLock()
	defer s.rw.RUnlock()

	if !s.checkUser(userId) {
		return nil, errors.New("user not found")
	}

	return s.proc.FilterByDay(s.data[userId], day), nil
}

func (s *Storage) GetEventsWeek(userId int) (map[string]models.Event, error) {
	s.rw.RLock()
	defer s.rw.RUnlock()

	if !s.checkUser(userId) {
		return nil, errors.New("user not found")
	}

	return s.proc.FilterByWeek(s.data[userId]), nil
}

func (s *Storage) GetEventsMonth(userId int, month time.Time) (map[string]models.Event, error) {
	s.rw.RLock()
	defer s.rw.RUnlock()

	if !s.checkUser(userId) {
		return nil, errors.New("user not found")
	}

	response, err := s.proc.FilterByMonth(s.data[userId], month)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Storage) checkUser(userId int) bool {
	if _, ok := s.data[userId]; ok {
		return true
	}
	return false
}
