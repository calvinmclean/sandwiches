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

var allMenu []MenuItem

func main() {
	allMenu = BuildMenu()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/menu", GetAllMenuItems).Methods("GET")
	router.HandleFunc("/menu/show", ShowMenu).Methods("GET")
	router.HandleFunc("/menu/{id}", GetSingleMenuItem).Methods("GET")
	router.HandleFunc("/menu", UpdateMenu).Methods("POST")
	http.ListenAndServe(":8080", router)
}

// GetSingleMenuItem responds to GET requests and returns a MenuItem from ID
func GetSingleMenuItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	menuItem, _ := FindMenuItem(id)
	menuItemJSON, _ := json.Marshal(menuItem)
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuItemJSON)
}

// GetAllMenuItems responds to GET requests and returns a list of all MenuItems
func GetAllMenuItems(w http.ResponseWriter, r *http.Request) {
	var menuJSON []byte
	menuJSON, _ = json.Marshal(allMenu)
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuJSON)
}

// UpdateMenu responds to POST requests and updates allMenu based on Recipe or
// Ingredient changes
func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	allMenu = BuildMenu()
	var menuJSON []byte
	menuJSON, _ = json.Marshal(allMenu)
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuJSON)
}

// ShowMenu responds to GET requests and writes the Menu string to HTTP
func ShowMenu(w http.ResponseWriter, r *http.Request) {
	var menuString []string
	menuString = append(menuString, "Welcome to the Sandwich Shop!\n")
	menuString = append(menuString, "Here is our menu:")
	for i, item := range allMenu {
		menuString = append(menuString, fmt.Sprintf("%d. %s ($%.2f)", i+1, item.Name, item.Price))
	}
	fmt.Fprintf(w, strings.Join(menuString, "\n"))
}

// BuildMenu gathers info from Recipes and Ingredients to create a list of
// MenuItems with calculated prices
func BuildMenu() []MenuItem {
	recipes := GetRecipes()
	menu := make([]MenuItem, len(recipes))

	// Add prices and names to the MenuItems
	for i, recipe := range recipes {
		menu[i] = MenuItem{
			ID:    recipe.ID,
			Price: calcRecipePrice(recipe),
			Name:  recipe.Name,
		}
	}
	return menu
}

func calcRecipePrice(recipe Recipe) float64 {
	price := 0.00
	price += calcPriceFromIngredients([]int{recipe.Bread})
	price += calcPriceFromIngredients(recipe.Meats)
	price += calcPriceFromIngredients(recipe.Cheeses)
	price += calcPriceFromIngredients(recipe.Toppings)
	return price
}

func calcPriceFromIngredients(items []int) float64 {
	price := 0.00
	for _, item := range items {
		price += GetIngredient(item).Price
	}
	return price
}

// GetIngredient makes a GET request to find a single Ingredient based on ID
func GetIngredient(id int) Ingredient {
	var ingredient Ingredient
	var httpClient = &http.Client{}
	r, _ := httpClient.Get(fmt.Sprintf("http://ingredients:8080/ingredients/%d", id))
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&ingredient)
	return ingredient
}

// GetRecipes makes a GET request to get all Recipes
func GetRecipes() []Recipe {
	var recipes []Recipe
	var httpClient = &http.Client{}
	r, _ := httpClient.Get("http://recipes:8080/recipes")
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&recipes)
	return recipes
}

// FindMenuItem returns a MenuItem from allMenu based on ID
func FindMenuItem(id int) (MenuItem, error) {
	for _, menuItem := range allMenu {
		if menuItem.ID == id {
			return menuItem, nil
		}
	}
	var fake MenuItem
	return fake, errors.New("No such recipe")
}
