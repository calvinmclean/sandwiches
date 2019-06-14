package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type SandwichRequest struct {
	MenuItemId int
	Extras     []int
}

func main() {
	http.HandleFunc("/clerk/order/", SandwichForm)
	http.HandleFunc("/clerk/purchase/", PurchaseSandwich)
	http.ListenAndServe(":8080", nil)
}

func SandwichForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "form.html")
}

func PurchaseSandwich(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sandwich := new(SandwichRequest)
		sandwich.MenuItemId, _ = strconv.Atoi(r.FormValue("item"))
		extrasStrings := strings.Split(r.FormValue("extras"), " ")
		for _, extra := range extrasStrings {
			id, _ := strconv.Atoi(extra)
			sandwich.Extras = append(sandwich.Extras, id)
		}
		fmt.Fprintf(w, "Price of your customized sandwich is $%.2f", CalculateSandwichPrice(*sandwich))
	} else {
		fmt.Fprintf(w, "Only POST requests are allowed at this endpoint")
	}
}

func CalculateSandwichPrice(sandwich SandwichRequest) float64 {
	menuItem := GetMenuItem(sandwich.MenuItemId)[0]
	price := menuItem.Price

	for _, id := range sandwich.Extras {
		price += GetIngredient(id)[0].Price
	}
	return price
}

func GetIngredient(id int) []Ingredient {
	var ingredient []Ingredient
	var httpClient = &http.Client{}
	r, _ := httpClient.Get(fmt.Sprintf("http://ingredients:8080/ingredients/%d", id))
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredient)
	return ingredient
}

func GetMenuItem(id int) []MenuItem {
	var menuItem []MenuItem
	var httpClient = &http.Client{}
	r, _ := httpClient.Get(fmt.Sprintf("http://menu:8080/menu/%d", id))
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&menuItem)
	return menuItem
}
