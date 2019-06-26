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

type SandwichRequest struct {
	MenuItemId int
	Extras     []int
}

type MenuList struct {
	Items []MenuItem
}

type IngredientList struct {
	Items []Ingredient
}

type OrderInfo struct {
	Menu        bool
	Ingredients bool
	Sandwiches  MenuList
	Extras      IngredientList
}

func main() {
	fmt.Println("Clerk started...")
	http.HandleFunc("/clerk/purchase/", PurchaseSandwich)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/clerk/order", SandwichForm).Methods("GET")
	router.HandleFunc("/clerk/purchase", PurchaseSandwich).Methods("POST")
	http.ListenAndServe(":8080", router)
}

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
	info := OrderInfo{menuExist, ingExist, MenuList{menu}, IngredientList{ingredients}}
	tmpl, _ := template.ParseFiles("form.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, info)
}

func PurchaseSandwich(w http.ResponseWriter, r *http.Request) {
	sandwich := new(SandwichRequest)
	sandwich.MenuItemId, _ = strconv.Atoi(r.FormValue("item"))
	for k, v := range r.Form {
		if id, err := strconv.Atoi(v[0]); err == nil && strings.Contains(k, "extras") {
			sandwich.Extras = append(sandwich.Extras, id)
		}
	}
	menuItem, _ := GetMenuItem(sandwich.MenuItemId)

	fmt.Fprintf(w, "Price of your customized %s is $%.2f\n", menuItem.Name, CalculateSandwichPrice(*sandwich))
}

func CalculateSandwichPrice(sandwich SandwichRequest) float64 {
	menuItem, _ := GetMenuItem(sandwich.MenuItemId)
	price := menuItem.Price

	for _, id := range sandwich.Extras {
		ingredient, _ := GetIngredient(id)
		price += ingredient.Price
	}
	return price
}

func GetIngredient(id int) (Ingredient, error) {
	var ingredient Ingredient
	var httpClient = &http.Client{}
	r, err := httpClient.Get(fmt.Sprintf("http://ingredients:8080/ingredients/%d", id))
	if err != nil || r.StatusCode != 200 {
		return ingredient, errors.New(fmt.Sprintf("%d StatusCode from ingredients", r.StatusCode))
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredient)
	return ingredient, nil
}

func GetIngredients() ([]Ingredient, error) {
	var ingredients []Ingredient
	var httpClient = &http.Client{}
	r, err := httpClient.Get("http://ingredients:8080/ingredients")
	if err != nil || r.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%d StatusCode from ingredients", r.StatusCode))
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredients)
	return ingredients, nil
}

func GetMenuItem(id int) (MenuItem, error) {
	var menuItem MenuItem
	var httpClient = &http.Client{}
	r, err := httpClient.Get(fmt.Sprintf("http://menu:8080/menu/%d", id))
	if err != nil || r.StatusCode != 200 {
		return menuItem, errors.New(fmt.Sprintf("%d StatusCode from menu", r.StatusCode))
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&menuItem)
	return menuItem, nil
}

func GetMenu() ([]MenuItem, error) {
	var menu []MenuItem
	var httpClient = &http.Client{}
	r, err := httpClient.Get("http://menu:8080/menu")
	if err != nil || r.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%d StatusCode from menu", r.StatusCode))
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&menu)
	return menu, nil
}
