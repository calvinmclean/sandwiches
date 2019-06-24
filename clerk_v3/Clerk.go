package main

import (
	"encoding/json"
	"errors"
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
	Menu bool
	Ingredients bool
	Sandwiches MenuList
	Extras     IngredientList
}

func main() {
	fmt.Println("Clerk started...")
	http.HandleFunc("/clerk/order/", SandwichForm)
	http.HandleFunc("/clerk/purchase/", PurchaseSandwich)
	http.ListenAndServe(":8080", nil)
}

func SandwichForm(w http.ResponseWriter, r *http.Request) {
	menu, menuErr := GetMenuItem(0)
	ingredients, ingErr := GetIngredient(0)
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
	if r.Method == "POST" {
		sandwich := new(SandwichRequest)
		sandwich.MenuItemId, _ = strconv.Atoi(r.FormValue("item"))
		for k, v := range r.Form {
			if id, err := strconv.Atoi(v[0]); err == nil && strings.Contains(k, "extras") {
				sandwich.Extras = append(sandwich.Extras, id)
			}
		}
		menu, _ := GetMenuItem(sandwich.MenuItemId)
		menuItem := menu[0]

		fmt.Fprintf(w, "Price of your customized %s is $%.2f\n", menuItem.Name, CalculateSandwichPrice(*sandwich))
		fmt.Fprintf(w, "You added the following ingredients:\n")
		for _, id := range sandwich.Extras {
			ingredients, _ := GetIngredient(id)
			ingredient := ingredients[0]
			fmt.Fprintf(w, "  * %s ($%.2f)\n", ingredient.Name, ingredient.Price)
		}
	} else {
		fmt.Fprintf(w, "Only POST requests are allowed at this endpoint")
	}
}

func CalculateSandwichPrice(sandwich SandwichRequest) float64 {
	menu, _ := GetMenuItem(sandwich.MenuItemId)
	menuItem := menu[0]
	price := menuItem.Price

	for _, id := range sandwich.Extras {
		ingredients, _ := GetIngredient(id)
		price += ingredients[0].Price
	}
	return price
}

func GetIngredient(id int) ([]Ingredient, error) {
	var ingredient []Ingredient
	var httpClient = &http.Client{}
	r, err := httpClient.Get(fmt.Sprintf("http://ingredients:8080/ingredients/%d", id))
	if err != nil || r.StatusCode != 200 {
		fmt.Println("Ingredient Error!")
		return nil, errors.New(fmt.Sprintf("%d StatusCode from ingredients", r.StatusCode))
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredient)
	return ingredient, nil
}

func GetMenuItem(id int) ([]MenuItem, error) {
	var menuItem []MenuItem
	var httpClient = &http.Client{}
	r, err := httpClient.Get(fmt.Sprintf("http://menu:8080/menu/%d", id))
	if err != nil || r.StatusCode != 200 {
		fmt.Println("Menu Error!")
		return nil, errors.New(fmt.Sprintf("%d StatusCode from menu", r.StatusCode))
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&menuItem)
	return menuItem, nil
}
