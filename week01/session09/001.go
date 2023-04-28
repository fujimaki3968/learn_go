package main

import (
	"fmt"
	"net/http"
)

func LogOutput(message string) {
	fmt.Println(message)
}

// Simple Data Store
type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Alice",
			"2": "Bob",
			"3": "Charlie",
		},
	}
}

// Buisiness Logic
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("SayHello(" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", fmt.Errorf("user not found")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("SayGoodbye(" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", fmt.Errorf("user not found")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) HandleGreeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("SayHello: ")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	controller := NewController(l, logic)
	http.HandleFunc("/hello", controller.HandleGreeting)
	http.ListenAndServe(":8080", nil)
}
