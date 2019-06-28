package main

import (
	"encoding/json"
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
	Sandwiches []MenuItem
	Extras     []Ingredient
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
	sandwiches, err := GetMenu()
	if err != nil {
		fmt.Println(err)
	}
	extras, err := GetIngredients()
	if err != nil {
		fmt.Println(err)
	}
	info := orderInfo{sandwiches, extras}
	tmpl, _ := template.ParseFiles("form.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, info)
}

// PurchaseSandwich responds to POST requests containing form data from the
// SandwichForm. It calculates the price of the sandwich with extras and writes
// it, and the extras, to HTTP
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
	if len(sandwich.Extras) > 0 {
		fmt.Fprintf(w, "You added the following ingredients:\n")
		for _, id := range sandwich.Extras {
			ingredient, _ := GetIngredient(id)
			fmt.Fprintf(w, "  * %s ($%.2f)\n", ingredient.Name, ingredient.Price)
		}
	}
}

// CalculateSandwichPrice returns the final price of a sandwich with addons
func CalculateSandwichPrice(sandwich sandwichOrder) (price float64) {
	menuItem, _ := GetMenuItem(sandwich.MenuItemID)
	price += menuItem.Price

	for _, id := range sandwich.Extras {
		ingredient, _ := GetIngredient(id)
		price += ingredient.Price
	}
	return
}

// GetIngredient makes a GET request to find an Ingredient from ID
func GetIngredient(id int) (ingredient Ingredient, err error) {
	var httpClient = &http.Client{}
	r, err := httpClient.Get(fmt.Sprintf("http://ingredients:8080/ingredients/%d", id))
	if err != nil || r.StatusCode != 200 {
		err = fmt.Errorf("%d StatusCode from ingredients", r.StatusCode)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredient)
	return
}

// GetIngredients makes a GET request to return a list of all Ingredients
func GetIngredients() (ingredients []Ingredient, err error) {
	var httpClient = &http.Client{}
	r, err := httpClient.Get("http://ingredients:8080/ingredients")
	if err != nil || r.StatusCode != 200 {
		err = fmt.Errorf("%d StatusCode from ingredients", r.StatusCode)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredients)
	return
}

// GetMenuItem makes a GET request to find a MenuItem from ID
func GetMenuItem(id int) (menuItem MenuItem, err error) {
	var httpClient = &http.Client{}
	r, err := httpClient.Get(fmt.Sprintf("http://menu:8080/menu/%d", id))
	if err != nil || r.StatusCode != 200 {
		err = fmt.Errorf("%d StatusCode from menu", r.StatusCode)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&menuItem)
	return
}

// GetMenu makes a GET request to return a list of all MenuItems
func GetMenu() (menu []MenuItem, err error) {
	var httpClient = &http.Client{}
	r, err := httpClient.Get("http://menu:8080/menu")
	if err != nil || r.StatusCode != 200 {
		err = fmt.Errorf("%d StatusCode from menu", r.StatusCode)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&menu)
	return
}
