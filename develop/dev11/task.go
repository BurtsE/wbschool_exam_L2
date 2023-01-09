package main

import (
	"errors"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

var memory cache

func main() {}

// Сервер
func NewServer() *http.Server {
	http.HandleFunc("/create_event", createEvent)
	http.HandleFunc("/update_event", updateEvent)
	http.HandleFunc("/delete_event", deleteEvent)
	http.HandleFunc("/events_for_day", eventsForDay)
	http.HandleFunc("/events_for_week", eventsForWeek)
	http.HandleFunc("/events_for_month", eventsForMonth)
	s := http.Server{
		Addr:         "localhost:8000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	return &s
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating event")
	err := r.ParseForm()
	if err != nil {
		log.Println("cannot parse form", err)
		return
	}

	uid, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		log.Println("invalid id")
		return
	}

	date, err := time.Parse("YYYY-MM-DD", r.Form.Get("date"))
	if err != nil {
		log.Println("no date specified")
		return
	}
	desc := r.Form.Get("description")
	var e = event{
		UID:        uid,
		date:        date,
		description: desc,
	}
	memory.Set(uid, e)
	log.Println("event was added")
}
func updateEvent(w http.ResponseWriter, r *http.Request) {

}
func deleteEvent(w http.ResponseWriter, r *http.Request) {

}
func eventsForDay(w http.ResponseWriter, r *http.Request) {

}
func eventsForWeek(w http.ResponseWriter, r *http.Request) {

}
func eventsForMonth(w http.ResponseWriter, r *http.Request) {

}

// Сервис

type event struct {
	UID        int
	date        time.Time
	description string
}

// Память
type cache struct {
	storage map[int][]event
	mutex   sync.RWMutex
}

func GetCache() *cache {
	if memory.storage == nil {
		memory = cache{
			storage: make(map[int][]event),
			mutex:   sync.RWMutex{},
		}
	}
	return &memory
}

func (c *cache) Set(key int, value event) {
	log.Println("adding message to cache", key)
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, exists := c.storage[key]
	if !exists {
		c.storage[key] = make([]event, 0)
	}
	c.storage[key] = append(c.storage[key], value)
	sort.SliceIsSorted(c.storage[key], func(i, j int) bool {
		return c.storage[key][i].date.Before(c.storage[key][j].date)
	})
}
func (c *cache) Get(key int) ([]event, error) {
	log.Println("getting events from cache")
	c.mutex.RLock()

	defer c.mutex.RUnlock()
	val, ok := c.storage[key]
	if !ok {
		return nil, errors.New("events not found")
	}
	return val, nil
}

func (c *cache) Delete(key int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.storage, key)
}