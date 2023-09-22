package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-type", "application/json")
}

type apiFunc func(http.ResponseWriter, *http.Request) error 
type ApiError struct {
	Error string 
}

func makehttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w,r); err != nil {
				WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
} 


type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenable string) *APIServer {
	return &APIServer{
		listenAddr: listenable,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makehttpHandleFunc(s.handleAccount))

	log.Println("JSON API Server running on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) getAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
