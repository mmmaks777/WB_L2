package calendar

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Event  string    `json:"event"`
}

type Calendar struct {
	events map[int]Event
	nextId int
	mutex  *sync.Mutex
}

func NewCalendar() *Calendar {
	return &Calendar{
		events: make(map[int]Event),
		nextId: 1,
		mutex:  &sync.Mutex{},
	}
}

func (c *Calendar) CreateEvent(event Event) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event.ID = c.nextId
	c.nextId++
	c.events[event.ID] = event

	return event.ID
}

func (c *Calendar) UpdateEvent(event Event) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.events[event.ID]; exists {
		return fmt.Errorf("Event with ID %d not found", event.ID)
	}

	c.events[event.ID] = event
	return nil
}

func (c *Calendar) DeleteEvent(id int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.events[id]; !exists {
		return fmt.Errorf("Event with ID %d not found", id)
	}

	delete(c.events, id)
	return nil
}

func (c *Calendar) GetEventForPeriod(userId int, startDate time.Time, duration time.Duration) []Event {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	endDate := startDate.Add(duration)
	var result []Event

	for _, event := range c.events {
		if event.UserID == userId && (event.Date.Equal(startDate) || (event.Date.After(startDate) && event.Date.Before(endDate))) {
			result = append(result, event)
		}
	}

	return result
}
