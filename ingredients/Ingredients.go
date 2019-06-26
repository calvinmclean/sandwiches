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

var Ingredients []Ingredient

func main() {
	Ingredients = GetIngredientsFromFile("ingredients.yaml")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ingredients", GetAllIngredients).Methods("GET")
	router.HandleFunc("/ingredients", AddIngredient).Methods("POST")
	router.HandleFunc("/ingredients/{id}", GetSingleIngredient).Methods("GET")
	router.HandleFunc("/ingredients/{id}", DeleteIngredient).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

func GetSingleIngredient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ingredientJson := GetIngredientJson(id)
	w.Header().Set("Content-Type", "application/json")
	w.Write(ingredientJson)
}

func GetIngredientJson(id int) []byte {
	ingredient, _ := FindIngredient(id)
	ingredientJson, _ := json.Marshal(ingredient)
	return ingredientJson
}

func GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	var ingredientsJson []byte
	ingredientsJson, _ = json.Marshal(Ingredients)
	w.Header().Set("Content-Type", "application/json")
	w.Write(ingredientsJson)
}

func AddIngredient(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var ingredient Ingredient
	json.Unmarshal(reqBody, &ingredient)
	ingredient.Id = len(Ingredients) + 1
	Ingredients = append(Ingredients, ingredient)
	WriteIngredientsToFile(Ingredients, "ingredients.yaml")
	UpdateMenu()

	w.Header().Set("Content-Type", "application/json")
	w.Write(GetIngredientJson(ingredient.Id))
}

func UpdateMenu() {
	var httpClient = &http.Client{}
	r, _ := httpClient.PostForm("http://menu:8080/menu", url.Values{})
	defer r.Body.Close()
}

func DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i, ingredient := range Ingredients {
		if ingredient.Id == id {
			Ingredients = append(Ingredients[:i], Ingredients[i+1:]...)
		}
	}
	WriteIngredientsToFile(Ingredients, "ingredients.yaml")
}

func WriteIngredientsToFile(ingredients []Ingredient, fname string) {
	var ingredientsData []byte
	ingredientsData, _ = yaml.Marshal(ingredients)
	ioutil.WriteFile(fname, ingredientsData, 0644)
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

	yaml.Unmarshal(byteValue, &ingredients)
	return ingredients
}

func FindIngredient(id int) (Ingredient, error) {
	if Ingredients == nil {
		Ingredients = GetIngredientsFromFile("ingredients.yaml")
	}

	for _, ingredient := range Ingredients {
		if ingredient.Id == id {
			return ingredient, nil
		}
	}
	var fake Ingredient
	return fake, errors.New("No such ingredient")
}
