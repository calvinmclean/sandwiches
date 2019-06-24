package main

import (
	"encoding/json"
	"fmt"
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
	Sandwiches MenuList
	Extras     IngredientList
}

func main() {
	http.HandleFunc("/clerk/order/", SandwichForm)
	http.HandleFunc("/clerk/purchase/", PurchaseSandwich)
	http.ListenAndServe(":8080", nil)
}

func SandwichForm(w http.ResponseWriter, r *http.Request) {
	info := OrderInfo{MenuList{GetMenuItem(0)}, IngredientList{GetIngredient(0)}}
	tmpl, _ := template.ParseFiles("form.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, info)
}

func PurchaseSandwich(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sandwich := new(SandwichRequest)
		sandwich.MenuItemId, _ = strconv.Atoi(r.FormValue("item"))
		for k, v := range r.Form {
			if id, err := strconv.Atoi(v[0]); err == nil && strings.Contains(k, "extras") {
				sandwich.Extras = append(sandwich.Extras, id)
			}
		}
		menuItem := GetMenuItem(sandwich.MenuItemId)[0]

		fmt.Fprintf(w, "Price of your customized %s is $%.2f\n", menuItem.Name, CalculateSandwichPrice(*sandwich))
		fmt.Fprintf(w, "You added the following ingredients:\n")
		for _, id := range sandwich.Extras {
			ingredient := GetIngredient(id)[0]
			fmt.Fprintf(w, "  * %s ($%.2f)\n", ingredient.Name, ingredient.Price)
		}
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
