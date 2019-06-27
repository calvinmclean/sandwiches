package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type sandwichOrder struct {
	MenuItemID int
	Extras     []int
}

type orderInfo struct {
	Menu        bool
	Ingredients bool
	Sandwiches  []MenuItem
	Extras      []Ingredient
}

func main() {
	fmt.Println("Clerk started...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/clerk/order", SandwichForm).Methods("GET")
	router.HandleFunc("/clerk/purchase", PurchaseSandwich).Methods("POST")
	http.ListenAndServe(":8080", router)
}

// SandwichForm gets the Ingredients and Menu from other services and writes an
// HTML template to HTTP to present a form for ordering a sandwich
func SandwichForm(w http.ResponseWriter, r *http.Request) {
	menu, menuErr := GetMenu()
	ingredients, ingErr := GetIngredients()
	menuExist, ingExist := true, true
	if menuErr != nil {
		menuExist = false
		fmt.Println(menuErr)
	}
	if ingErr != nil {
		ingExist = false
		fmt.Println(ingErr)
	}
	info := orderInfo{menuExist, ingExist, menu, ingredients}
	tmpl, _ := template.ParseFiles("form.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, info)
}

// PurchaseSandwich responds to POST requests containing form data from the
// SandwichForm. It calculates the price of the sandwich with extras and writes
// it to HTTP
func PurchaseSandwich(w http.ResponseWriter, r *http.Request) {
	sandwich := new(sandwichOrder)
	sandwich.MenuItemID, _ = strconv.Atoi(r.FormValue("item"))
	for k, v := range r.Form {
		if id, err := strconv.Atoi(v[0]); err == nil && strings.Contains(k, "extras") {
			sandwich.Extras = append(sandwich.Extras, id)
		}
	}
	menuItem, _ := GetMenuItem(sandwich.MenuItemID)

	fmt.Fprintf(w, "Price of your customized %s is $%.2f\n", menuItem.Name, CalculateSandwichPrice(*sandwich))
}

// CalculateSandwichPrice returns the final price of a sandwich with addons
func CalculateSandwichPrice(sandwich sandwichOrder) float64 {
	menuItem, _ := GetMenuItem(sandwich.MenuItemID)
	price := menuItem.Price

	for _, id := range sandwich.Extras {
		ingredient, _ := GetIngredient(id)
		price += ingredient.Price
	}
	return price
}

// GetIngredient makes a GET request to find an Ingredient from ID
func GetIngredient(id int) (Ingredient, error) {
	var ingredient Ingredient
	var httpClient = &http.Client{}
	r, err := httpClient.Get(fmt.Sprintf("http://ingredients:8080/ingredients/%d", id))
	if err != nil || r.StatusCode != 200 {
		return ingredient, fmt.Errorf("%d StatusCode from ingredients", r.StatusCode)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredient)
	return ingredient, nil
}

// GetIngredients makes a GET request to return a list of all Ingredients
func GetIngredients() ([]Ingredient, error) {
	var ingredients []Ingredient
	var httpClient = &http.Client{}
	r, err := httpClient.Get("http://ingredients:8080/ingredients")
	if err != nil || r.StatusCode != 200 {
		return nil, fmt.Errorf("%d StatusCode from ingredients", r.StatusCode)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredients)
	return ingredients, nil
}

// GetMenuItem makes a GET request to find a MenuItem from ID
func GetMenuItem(id int) (MenuItem, error) {
	var menuItem MenuItem
	var httpClient = &http.Client{}
	r, err := httpClient.Get(fmt.Sprintf("http://menu:8080/menu/%d", id))
	if err != nil || r.StatusCode != 200 {
		return menuItem, fmt.Errorf("%d StatusCode from menu", r.StatusCode)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&menuItem)
	return menuItem, nil
}

// GetMenu makes a GET request to return a list of all MenuItems
func GetMenu() ([]MenuItem, error) {
	var menu []MenuItem
	var httpClient = &http.Client{}
	r, err := httpClient.Get("http://menu:8080/menu")
	if err != nil || r.StatusCode != 200 {
		return nil, fmt.Errorf("%d StatusCode from menu", r.StatusCode)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&menu)
	return menu, nil
}
