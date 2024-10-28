package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"2.12/internal/calendar"
)

type Server struct {
	Calendar *calendar.Calendar
}

func writeJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func parseEventParams(r *http.Request) (calendar.Event, error) {
	var event calendar.Event
	err := r.ParseForm()
	if err != nil {
		return event, fmt.Errorf("could not parse form: %v", err)
	}

	userId, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return event, fmt.Errorf("invalid user_id: %v", err)
	}

	dateStr := r.FormValue("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return event, fmt.Errorf("invalid date: %v", err)
	}

	eventText := r.FormValue("event")
	if eventText == "" {
		return event, fmt.Errorf("empty filed event")
	}

	event = calendar.Event{
		UserID: userId,
		Date:   date,
		Event:  eventText,
	}
	return event, nil
}

func (s *Server) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "method not allowed"})
	}

	event, err := parseEventParams(r)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	eventId := s.Calendar.CreateEvent(event)

	writeJson(w, http.StatusOK, map[string]string{"result": fmt.Sprintf("event cteated with ID: %d", eventId)})
}

func (s *Server) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "method not allowed"})
	}

	event, err := parseEventParams(r)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}
	event.ID = id

	err = s.Calendar.UpdateEvent(event)
	if err != nil {
		writeJson(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
	}

	writeJson(w, http.StatusOK, map[string]string{"result": "event updated"})
}

func (s *Server) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "method not allowed"})
	}

	err := r.ParseForm()
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "could not parse form"})
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	err = s.Calendar.DeleteEvent(id)
	if err != nil {
		writeJson(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]string{"result": "event deleted"})
}

func (s *Server) EventsForPeriodHandler(w http.ResponseWriter, r *http.Request, duration time.Duration) {
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "method not allowed"})
	}

	err := r.ParseForm()
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "could not parse form"})
	}

	userIDStr := r.FormValue("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid user_id"})
		return
	}

	dateStr := r.FormValue("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid date"})
		return
	}

	events := s.Calendar.GetEventForPeriod(userID, date, duration)
	writeJson(w, http.StatusOK, map[string]interface{}{"result": events})
}

func (s *Server) EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	s.EventsForPeriodHandler(w, r, 24*time.Hour)
}

func (s *Server) EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	s.EventsForPeriodHandler(w, r, 7*24*time.Hour)
}

func (s *Server) EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	s.EventsForPeriodHandler(w, r, 30*24*time.Hour)
}
