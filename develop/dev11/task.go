package main

/*
=== HTTP server ===
Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.
В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

var memory cache

func main() {
	memory = *GetCache()
	s := NewServer()
	s.ListenAndServe()
}

// Сервер
type handlefunc func(w http.ResponseWriter, r *http.Request)
type decorator struct {
	h handlefunc
}

func (d *decorator) processMethod(method string) handlefunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			data, _ := json.Marshal(map[string]string{"error": "wrong method"})
			w.Write(data)
			w.WriteHeader(400)
			return
		}
		d.h(w, r)
	}
}
func MethodValidate(fn handlefunc, method string) handlefunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/json")
		if r.Method != method {
			w.WriteHeader(400)
			data, _ := json.Marshal(map[string]string{"error": "wrong method"})
			w.Write(data)
			return
		}
		fn(w, r)
	}
}
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
func NewServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", MethodValidate(createEvent, "POST"))
	mux.HandleFunc("/update_event", MethodValidate(updateEvent, "POST"))
	mux.HandleFunc("/delete_event", MethodValidate(deleteEvent, "POST"))
	mux.HandleFunc("/events_for_day", MethodValidate(eventsForDay, "GET"))
	mux.HandleFunc("/events_for_week", MethodValidate(eventsForWeek, "GET"))
	mux.HandleFunc("/events_for_month", MethodValidate(eventsForMonth, "GET"))
	handler := Logging(mux)
	s := http.Server{
		Addr:         "localhost:8000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      handler,
	}
	return &s
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	uid, date, desc, err := validateParams(r)
	if err != nil {
		w.WriteHeader(400)
		data, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(data)
		return
	}
	var e = event{
		UID:         uid,
		Date:        date,
		Description: desc,
	}
	memory.Set(uid, e)
}

func validateParams(r *http.Request) (int, time.Time, string, error) {
	err := r.ParseForm()
	if err != nil {
		log.Println("cannot parse form", err)
		return 0, time.Time{}, "", err
	}

	uid, err := validateID(r)
	if err != nil {
		return 0, time.Time{}, "", err
	}
	date, err := validateDate(r)
	if err != nil {
		return 0, time.Time{}, "", err
	}
	desc := r.Form.Get("description")

	return uid, date, desc, nil
}

// Валидация id
func validateID(r *http.Request) (int, error) {
	uid, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		log.Println("invalid id")
		return 0, err
	}
	return uid, nil
}

// Валидация даты
func validateDate(r *http.Request) (time.Time, error) {
	date, err := time.Parse("2006-01-02", r.Form.Get("date"))
	if err != nil {
		log.Println("no date specified")
		log.Println(err)
		return time.Time{}, err
	}
	return date, nil
}

func updateEvent(w http.ResponseWriter, r *http.Request) {

}
func deleteEvent(w http.ResponseWriter, r *http.Request) {

}
func eventsForDay(w http.ResponseWriter, r *http.Request) {

	uid, date, _, err := validateParams(r)
	if err != nil {
		w.WriteHeader(400)
		errData, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(errData)
		return
	}
	events, _ := memory.Get(uid, date, date.AddDate(0, 0, 1))

	data, err := formDoc(events)
	if err != nil {
		w.WriteHeader(503)
		errData, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(errData)
		return
	}
	w.WriteHeader(200)
	w.Write(data)

}
func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	uid, date, _, err := validateParams(r)
	if err != nil {
		w.WriteHeader(400)
		errData, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(errData)
		return
	}

	events, _ := memory.Get(uid, date, date.AddDate(0, 0, 7))

	data, err := formDoc(events)
	if err != nil {
		w.WriteHeader(503)
		errData, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(errData)
		return
	}
	w.WriteHeader(200)
	w.Write(data)

}
func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	uid, date, _, err := validateParams(r)
	if err != nil {
		w.WriteHeader(400)
		errData, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(errData)
		return
	}

	events, _ := memory.Get(uid, date, date.AddDate(0, 1, 0))

	data, err := formDoc(events)
	if err != nil {
		w.WriteHeader(503)
		errData, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(errData)
		return
	}
	w.WriteHeader(200)
	w.Write(data)
}

func createAnswer(status int) []byte {
	switch status {
	case 400:
	case 500:
	case 503:
	case 200:
	}
	return nil
}

// Сервис
type event struct {
	UID         int
	Date        time.Time
	Description string
}
type result struct {
	Result []event
}

func formDoc(evs []event) ([]byte, error) {
	if len(evs) == 0 {
		data, _ := json.Marshal([]string{""})
		return data, nil
	}
	r := result{evs}
	data, err := json.Marshal(r)
	if err != nil {
		data, _ = json.Marshal([]string{err.Error()})
	}
	return data, err
}

// Память
type cache struct {
	storage map[int][]event
	mutex   sync.RWMutex
}

func GetCache() *cache {
	c := cache{
		storage: make(map[int][]event),
		mutex:   sync.RWMutex{},
	}
	return &c
}

func (c *cache) Set(key int, value event) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, exists := c.storage[key]
	if !exists {
		c.storage[key] = make([]event, 0)
	}
	c.storage[key] = append(c.storage[key], value)
	sort.SliceIsSorted(c.storage[key], func(i, j int) bool {
		return c.storage[key][i].Date.Before(c.storage[key][j].Date)
	})
}
func (c *cache) Get(key int, start, end time.Time) ([]event, error) {
	c.mutex.RLock()

	defer c.mutex.RUnlock()
	val, ok := c.storage[key]
	if !ok {
		return nil, errors.New("No users with this id")
	}
	var i, j int
	for i = 0; i < len(val) && val[i].Date.Before(start); i++ {
	}
	j = i
	for j < len(val) && val[i].Date.Before(end) {
		j++
	}
	return val[i:j], nil
}

func (c *cache) Delete(key int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.storage, key)
}
