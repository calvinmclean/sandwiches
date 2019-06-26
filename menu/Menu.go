package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

var Menu []MenuItem

func main() {
	Menu = BuildMenu()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/menu", GetAllMenuItems).Methods("GET")
	router.HandleFunc("/menu/show", ShowMenu).Methods("GET")
	router.HandleFunc("/menu/{id}", GetSingleMenuItem).Methods("GET")
	router.HandleFunc("/menu", UpdateMenu).Methods("POST")
	http.ListenAndServe(":8080", router)
}

func GetSingleMenuItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	menuItem, _ := FindMenuItem(id)
	menuItemJson, _ := json.Marshal(menuItem)
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuItemJson)
}

func GetAllMenuItems(w http.ResponseWriter, r *http.Request) {
	var menuJson []byte
	menuJson, _ = json.Marshal(Menu)
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuJson)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	Menu = BuildMenu()
	var menuJson []byte
	menuJson, _ = json.Marshal(Menu)
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuJson)
}

func ShowMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, strings.Join(GetMenuString(), "\n"))
}

func GetMenuString() []string {
	if Menu == nil {
		Menu = BuildMenu()
	}
	var result []string
	result = append(result, "Welcome to the Sandwich Shop!\n")
	result = append(result, "Here is our menu:")
	for i, item := range Menu {
		result = append(result, fmt.Sprintf("%d. %s ($%.2f)", i+1, item.Name, item.Price))
	}
	return result
}

func BuildMenu() []MenuItem {
	recipes := GetRecipes() // Get all recipes
	menu := make([]MenuItem, len(recipes))

	// Add prices and names to the MenuItems
	for i, recipe := range recipes {
		menu[i] = MenuItem{
			Id:    recipe.Id,
			Price: CalcRecipePrice(recipe),
			Name:  recipe.Name,
		}
	}
	return menu
}

func CalcRecipePrice(recipe Recipe) float64 {
	price := 0.00
	price += CalcPriceFromSlice([]int{recipe.Bread})
	price += CalcPriceFromSlice(recipe.Meats)
	price += CalcPriceFromSlice(recipe.Cheeses)
	price += CalcPriceFromSlice(recipe.Toppings)
	return price
}

func CalcPriceFromSlice(items []int) float64 {
	price := 0.00
	for _, item := range items {
		price += GetIngredient(item).Price
	}
	return price
}

func GetIngredient(id int) Ingredient {
	var ingredient Ingredient
	var httpClient = &http.Client{}
	r, _ := httpClient.Get(fmt.Sprintf("http://ingredients:8080/ingredients/%d", id))
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredient)
	return ingredient
}

func GetRecipes() []Recipe {
	var recipes []Recipe
	var httpClient = &http.Client{}
	r, _ := httpClient.Get("http://recipes:8080/recipes")
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&recipes)
	return recipes
}

func FindMenuItem(id int) (MenuItem, error) {
	if Menu == nil {
		Menu = BuildMenu()
	}

	for _, menuItem := range Menu {
		if menuItem.Id == id {
			return menuItem, nil
		}
	}
	var fake MenuItem
	return fake, errors.New("No such recipe")
}
