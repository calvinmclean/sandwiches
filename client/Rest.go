package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	pb "sandwiches/protobuf"
)

type menuItemPlus struct {
	MenuItem *pb.MenuItem
	Bread    pb.Ingredient
	Meats    []pb.Ingredient
	Cheeses  []pb.Ingredient
	Toppings []pb.Ingredient
}

type menuInfo struct {
	Menu        []menuItemPlus
	Ingredients []*pb.Ingredient
}

func apiGetSingleIngredient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ingredientJSON := getIngredientJSON(int32(id))
	w.Header().Set("Content-Type", "application/json")
	w.Write(ingredientJSON)
}

func apiGetAllIngredients(w http.ResponseWriter, r *http.Request) {
	var ingredientsJSON []byte
	ingredientsJSON, _ = json.Marshal(getAllIngredients())
	w.Header().Set("Content-Type", "application/json")
	w.Write(ingredientsJSON)
}

func getIngredientJSON(id int32) (result []byte) {
	ingredient := getSingleIngredient(id)
	result, _ = json.Marshal(ingredient)
	return
}

func apiAddIngredient(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var ingredient pb.Ingredient
	json.Unmarshal(reqBody, &ingredient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	newIngredient, err := ingredientsClient.AddIngredient(ctx, &pb.NewIngredient{
		Name:  ingredient.Name,
		Price: ingredient.Price,
		Type:  ingredient.Type,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(getIngredientJSON(newIngredient.Id))
}

func apiGetSingleRecipe(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	recipeJSON := getRecipeJSON(int32(id))
	w.Header().Set("Content-Type", "application/json")
	w.Write(recipeJSON)
}

func apiGetAllRecipes(w http.ResponseWriter, r *http.Request) {
	var recipesJSON []byte
	recipesJSON, _ = json.Marshal(getAllRecipes())
	w.Header().Set("Content-Type", "application/json")
	w.Write(recipesJSON)
}

func getRecipeJSON(id int32) (result []byte) {
	recipe := getSingleRecipe(int32(id))
	result, _ = json.Marshal(recipe)
	return
}

func apiAddRecipe(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var recipe pb.NewRecipe
	json.Unmarshal(reqBody, &recipe)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	newRecipe, err := recipesClient.AddRecipe(ctx, &recipe)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(getRecipeJSON(newRecipe.Id))
}

// ShowMenu responds to GET requests and builds new menuItemPlus
// in order to populate an HTML template with each MenuItem, its price, and a
// list of Ingredients by type
func apiShowMenu(w http.ResponseWriter, r *http.Request) {
	ingredients := getAllIngredients()
	menu := getMenuItems()

	var menuWithIngredients []menuItemPlus
	for _, menuItem := range menu.MenuItems {
		recipe := getSingleRecipe(menuItem.Id)

		var meats []pb.Ingredient
		for _, id := range recipe.Meats {
			meats = append(meats, getSingleIngredient(id))
		}
		var cheeses []pb.Ingredient
		for _, id := range recipe.Cheeses {
			cheeses = append(cheeses, getSingleIngredient(id))
		}
		var toppings []pb.Ingredient
		for _, id := range recipe.Toppings {
			toppings = append(toppings, getSingleIngredient(id))
		}
		menuWithIngredients = append(menuWithIngredients,
			menuItemPlus{menuItem, getSingleIngredient(recipe.Bread), meats, cheeses, toppings})
	}

	info := menuInfo{menuWithIngredients, ingredients.Ingredients}
	tmpl, _ := template.ParseFiles("menu.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, info)
}
