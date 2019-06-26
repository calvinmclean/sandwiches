package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
)

type MenuItemWithIngredients struct {
	MenuItem MenuItem
	Bread    Ingredient
	Meats    []Ingredient
	Cheeses  []Ingredient
	Toppings []Ingredient
}

type MenuInfo struct {
	Menu        []MenuItemWithIngredients
	Ingredients []Ingredient
}

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
	ingredients, _ := GetIngredients()

	var menuWithIngredients []MenuItemWithIngredients
	for _, menuItem := range Menu {
		recipe := GetRecipe(menuItem.Id)

		var meats []Ingredient
		for _, id := range recipe.Meats {
			meats = append(meats, GetIngredient(id))
		}
		var cheeses []Ingredient
		for _, id := range recipe.Cheeses {
			cheeses = append(cheeses, GetIngredient(id))
		}
		var toppings []Ingredient
		for _, id := range recipe.Toppings {
			toppings = append(toppings, GetIngredient(id))
		}
		menuWithIngredients = append(menuWithIngredients, MenuItemWithIngredients{menuItem, GetIngredient(recipe.Bread), meats, cheeses, toppings})
	}

	info := MenuInfo{menuWithIngredients, ingredients}
	tmpl, _ := template.ParseFiles("menu.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, info)
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

func GetRecipes() []Recipe {
	var recipes []Recipe
	var httpClient = &http.Client{}
	r, _ := httpClient.Get("http://recipes:8080/recipes")
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&recipes)
	return recipes
}

func GetRecipe(id int) Recipe {
	var recipe Recipe
	var httpClient = &http.Client{}
	r, _ := httpClient.Get(fmt.Sprintf("http://recipes:8080/recipes/%d", id))
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&recipe)
	return recipe
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
