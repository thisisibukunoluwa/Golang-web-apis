package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thisisibukunoluwa/Golang-web-apis/coffeester/helpers"
	"github.com/thisisibukunoluwa/Golang-web-apis/coffeester/services"
)

var coffee services.Coffee

// GET/coffees
func GetAllCoffees (w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return 
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
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
 

