package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thisisibukunoluwa/Golang-web-apis/coffeester/helpers"
	"github.com/thisisibukunoluwa/Golang-web-apis/coffeester/services"
)

var coffee services.Coffee

// GET/coffees
func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return 
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}


//GET//coffees/coffee/{id}
func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	coffee,err := coffee.GetCoffeeById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return 
	}
	helpers.WriteJSON(w,http.StatusOK, coffee)
}



//POST/coffee/
func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err) 
		helpers.MessageLogs.ErrorLog.Println("HERE 1st err check")
		return 
	}
	// helpers.WriteJSON(w,http.StatusOK, coffeeData)
	coffeeCreated, err := coffee.CreateCoffee(coffeeData)
	if err != nil  {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.MessageLogs.ErrorLog.Println("HERE 2nd err check")
		return
	}
	helpers.WriteJSON(w, http.StatusOK, coffeeCreated)
}
 
//PUT/ccoffees/coffee/{id}

func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	id := chi.URLParam(r, "id")
	var coffee services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	coffeeUpdated, err := coffee.UpdateCoffee(id, coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, coffeeUpdated)
}


//DELETE/coffees/coffee/{id}

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := coffee.DeleteCoffee(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	msg := "successful deletion"
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message" : msg})
}

