package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var allRecipes []Recipe

func main() {
	allRecipes = GetRecipesFromFile("recipes.yaml")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recipes", GetAllRecipes).Methods("GET")
	router.HandleFunc("/recipes", AddRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", GetSingleRecipe).Methods("GET")
	router.HandleFunc("/recipes/{id}", DeleteRecipe).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

// GetSingleRecipe responds to GET requests and returns a single Recipe from ID
func GetSingleRecipe(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	recipeJSON := GetRecipeJSON(id)
	w.Header().Set("Content-Type", "application/json")
	w.Write(recipeJSON)
}

// GetRecipeJSON is used to convert Recipe to JSON
func GetRecipeJSON(id int) []byte {
	recipe, _ := FindRecipe(id)
	recipeJSON, _ := json.Marshal(recipe)
	return recipeJSON
}

// GetAllRecipes responds to GET requests and returns all Recipes
func GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	var recipeJSON []byte
	recipeJSON, _ = json.Marshal(allRecipes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(recipeJSON)
}

// AddRecipe responds to POST requests to create a new Recipe from the JSON, add
// it to allRecipes, and write to file
func AddRecipe(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var recipe Recipe
	json.Unmarshal(reqBody, &recipe)
	recipe.ID = len(allRecipes) + 1
	allRecipes = append(allRecipes, recipe)
	WriteRecipesToFile(allRecipes, "recipes.yaml")
	UpdateMenu()

	w.Header().Set("Content-Type", "application/json")
	w.Write(GetRecipeJSON(recipe.ID))
}

// DeleteRecipe responds to DELETE requests to delete a Recipe based on its ID
// and writes this change to file
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i, recipe := range allRecipes {
		if recipe.ID == id {
			allRecipes = append(allRecipes[:i], allRecipes[i+1:]...)
		}
	}
	WriteRecipesToFile(allRecipes, "recipes.yaml")
	UpdateMenu()
}

// UpdateMenu is called after adding or deleting a Recipe and sends an empty
// POST request to the Menu microservice telling it to update
func UpdateMenu() {
	var httpClient = &http.Client{}
	r, _ := httpClient.PostForm("http://menu:8080/menu", url.Values{})
	defer r.Body.Close()
}

// WriteRecipesToFile is used to write allRecipes to a YAML file
func WriteRecipesToFile(recipes []Recipe, fname string) {
	var recipesData []byte
	recipesData, _ = yaml.Marshal(recipes)
	ioutil.WriteFile(fname, recipesData, 0644)
}

// GetRecipesFromFile is used to read a YAML file into allRecipes
func GetRecipesFromFile(fname string) []Recipe {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + fname)
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var recipes []Recipe
	yaml.Unmarshal(byteValue, &recipes)
	return recipes
}

// FindRecipe returns a Recipe from allRecipes based on ID
func FindRecipe(id int) (Recipe, error) {
	for _, recipe := range allRecipes {
		if recipe.ID == id {
			return recipe, nil
		}
	}
	var fake Recipe
	return fake, errors.New("No such recipe")
}
