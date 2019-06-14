package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var RecipesFromFile []Recipe

func main() {
	RecipesFromFile = GetRecipesFromFile("recipes.yaml")
	http.HandleFunc("/recipes/", GetRecipe)
	http.ListenAndServe(":8080", nil)
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Split(r.URL.Path[1:], "/")[1])
	var recipes []Recipe
	if id == 0 {
		recipes = RecipesFromFile
	} else {
		recipe, _ := FindRecipe(id)
		recipes = []Recipe{recipe}
	}
	var recipeJson []byte
	recipeJson, _ = json.Marshal(recipes)
	w.Write(recipeJson)
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

	ext := strings.Split(fname, ".")[1]
	if ext == "yml" || ext == "yaml" {
		yaml.Unmarshal(byteValue, &recipes)
	} else if ext == "json" {
		json.Unmarshal(byteValue, &recipes)
	} else {
		fmt.Println("Cannot parse file without JSON or YAML extension")
		os.Exit(2)
	}
	return recipes
}

func FindRecipe(id int) (Recipe, error) {
	if RecipesFromFile == nil {
		RecipesFromFile = GetRecipesFromFile("recipes.yaml")
	}

	for _, recipe := range RecipesFromFile {
		if recipe.Id == id {
			return recipe, nil
		}
	}
	var fake Recipe
	return fake, errors.New("No such recipe")
}
