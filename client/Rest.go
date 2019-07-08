package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "sandwiches/protobuf"
)

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
