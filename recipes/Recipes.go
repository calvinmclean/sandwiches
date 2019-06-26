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

var Recipes []Recipe

func main() {
	Recipes = GetRecipesFromFile("recipes.yaml")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recipes", GetAllRecipes).Methods("GET")
	router.HandleFunc("/recipes", AddRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", GetSingleRecipe).Methods("GET")
	router.HandleFunc("/recipes/{id}", DeleteRecipe).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

func GetSingleRecipe(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	recipeJson := GetRecipeJson(id)
	w.Header().Set("Content-Type", "application/json")
	w.Write(recipeJson)
}

func GetRecipeJson(id int) []byte {
	recipe, _ := FindRecipe(id)
	recipeJson, _ := json.Marshal(recipe)
	return recipeJson
}

func GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	var recipeJson []byte
	recipeJson, _ = json.Marshal(Recipes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(recipeJson)
}

func AddRecipe(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var recipe Recipe
	json.Unmarshal(reqBody, &recipe)
	recipe.Id = len(Recipes) + 1
	Recipes = append(Recipes, recipe)
	WriteRecipesToFile(Recipes, "recipes.yaml")
	UpdateMenu()

	w.Header().Set("Content-Type", "application/json")
	w.Write(GetRecipeJson(recipe.Id))
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i, recipe := range Recipes {
		if recipe.Id == id {
			Recipes = append(Recipes[:i], Recipes[i+1:]...)
		}
	}
	WriteRecipesToFile(Recipes, "recipes.yaml")
	UpdateMenu()
}

func UpdateMenu() {
	var httpClient = &http.Client{}
	r, _ := httpClient.PostForm("http://menu:8080/menu", url.Values{})
	defer r.Body.Close()
}

func WriteRecipesToFile(recipes []Recipe, fname string) {
	var recipesData []byte
	recipesData, _ = yaml.Marshal(recipes)
	ioutil.WriteFile(fname, recipesData, 0644)
}

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

func FindRecipe(id int) (Recipe, error) {
	if Recipes == nil {
		Recipes = GetRecipesFromFile("recipes.yaml")
	}

	for _, recipe := range Recipes {
		if recipe.Id == id {
			return recipe, nil
		}
	}
	var fake Recipe
	return fake, errors.New("No such recipe")
}
