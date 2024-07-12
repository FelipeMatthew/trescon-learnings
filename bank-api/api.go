package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

type apiFunc func (http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}


func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handdleAccount))

	log.Println("JSON API running on: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}


func (s *APIServer) handdleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handdleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handdleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handdleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handdleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handdleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handdleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handdleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}