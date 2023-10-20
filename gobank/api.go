package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenable string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenable,
		store: 			 store,	
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makehttpHandleFunc(s.handleAccount))
	
	router.HandleFunc("/account/{id}", makehttpHandleFunc(s.handleGetAccountByID))

	log.Println("JSON API Server running on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w,r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w,r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w,r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}
// GET /account
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts,err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(w,http.StatusOK,accounts)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id,err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid id given %s", s)
	}
	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err 
	}
	fmt.Println(id)
	// account := NewAccount("jason", "momoa")
	return WriteJSON(w,http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}
	account := NewAccount(createAccountReq.FirstName,createAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}
	return WriteJSON(w,http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

//don't hardcode your storage(database) into your handlers 

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error 
type ApiError struct {
	Error string `json:"error"`
}

func makehttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w,r); err != nil {
				WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
} 




