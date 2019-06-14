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

var IngredientsFromFile []Ingredient

func main() {
	IngredientsFromFile = GetIngredientsFromFile("ingredients.yaml")
	http.HandleFunc("/ingredients/", GetIngredient)
	http.ListenAndServe(":8080", nil)
}

func GetIngredient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Split(r.URL.Path[1:], "/")[1])
	var ingredients []Ingredient
	if id == 0 {
		ingredients = IngredientsFromFile
	} else {
		ingredient, _ := FindIngredient(id)
		ingredients = []Ingredient{ingredient}
	}
	var ingredientJson []byte
	ingredientJson, _ = json.Marshal(ingredients)
	w.Write(ingredientJson)
}

func GetIngredientsFromFile(fname string) []Ingredient {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + fname)
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var ingredients []Ingredient

	ext := strings.Split(fname, ".")[1]
	if ext == "yml" || ext == "yaml" {
		yaml.Unmarshal(byteValue, &ingredients)
	} else if ext == "json" {
		json.Unmarshal(byteValue, &ingredients)
	} else {
		fmt.Println("Cannot parse file without JSON or YAML extension")
		os.Exit(2)
	}
	return ingredients
}

func FindIngredient(id int) (Ingredient, error) {
	if IngredientsFromFile == nil {
		IngredientsFromFile = GetIngredientsFromFile("ingredients.yaml")
	}

	for _, ingredient := range IngredientsFromFile {
		if ingredient.Id == id {
			return ingredient, nil
		}
	}
	var fake Ingredient
	return fake, errors.New("No such ingredient")
}
