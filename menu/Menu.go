package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var Menu []MenuItem

func main() {
	Menu = BuildMenu()
	http.HandleFunc("/menu/", GetMenu)
	http.HandleFunc("/menu/show/", ShowMenu)
	http.ListenAndServe(":8080", nil)
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Split(r.URL.Path[1:], "/")[1])
	var menu []MenuItem
	if id == 0 {
		menu = Menu
	} else {
		menuItem, _ := FindMenuItem(id)
		menu = []MenuItem{menuItem}
	}
	var menuJson []byte
	menuJson, _ = json.Marshal(menu)
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
	recipes := GetRecipe(0) // Get all recipes
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
		price += GetIngredient(item)[0].Price
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

func GetRecipe(id int) []Recipe {
	var recipe []Recipe
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
