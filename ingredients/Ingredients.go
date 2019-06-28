package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var allIngredients []Ingredient

func main() {
	allIngredients = GetIngredientsFromFile("ingredients.yaml")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ingredients", GetAllIngredients).Methods("GET")
	router.HandleFunc("/ingredients", AddIngredient).Methods("POST")
	router.HandleFunc("/ingredients/{id}", GetSingleIngredient).Methods("GET")
	router.HandleFunc("/ingredients/{id}", DeleteIngredient).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

// GetSingleIngredient responds to GET requests and returns an Ingredient by ID
func GetSingleIngredient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ingredientJSON := GetIngredientJSON(id)
	w.Header().Set("Content-Type", "application/json")
	w.Write(ingredientJSON)
}

// GetIngredientJSON is used to convert Ingredient to JSON
func GetIngredientJSON(id int) (result []byte) {
	ingredient, _ := FindIngredient(id)
	result, _ = json.Marshal(ingredient)
	return
}

// GetAllIngredients reponds to GET requests and returns all Ingredients
func GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	var ingredientsJSON []byte
	ingredientsJSON, _ = json.Marshal(allIngredients)
	w.Header().Set("Content-Type", "application/json")
	w.Write(ingredientsJSON)
}

// AddIngredient responds to POST requests to create a new Ingredient from the
// JSON, add it to allIngredients, and write to file
func AddIngredient(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var ingredient Ingredient
	json.Unmarshal(reqBody, &ingredient)
	ingredient.ID = len(allIngredients) + 1
	allIngredients = append(allIngredients, ingredient)
	WriteIngredientsToFile(allIngredients, "ingredients.yaml")
	UpdateMenu()

	w.Header().Set("Content-Type", "application/json")
	w.Write(GetIngredientJSON(ingredient.ID))
}

// DeleteIngredient responds to DELETE requests to delete an Ingredient based on
// its ID and writes this change to file
func DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i, ingredient := range allIngredients {
		if ingredient.ID == id {
			allIngredients = append(allIngredients[:i], allIngredients[i+1:]...)
		}
	}
	WriteIngredientsToFile(allIngredients, "ingredients.yaml")
	UpdateMenu()
}

// UpdateMenu is called after adding or deleting an Ingredient and sends an
// empty POST request to the Menu microservice telling it to update
func UpdateMenu() {
	var httpClient = &http.Client{}
	r, _ := httpClient.PostForm("http://menu:8080/menu", url.Values{})
	defer r.Body.Close()
}

// WriteIngredientsToFile is used to write allIngredients to a YAML file
func WriteIngredientsToFile(ingredients []Ingredient, fname string) {
	var ingredientsData []byte
	ingredientsData, _ = yaml.Marshal(ingredients)
	ioutil.WriteFile(fname, ingredientsData, 0644)
}

// GetIngredientsFromFile is used to read a YAML file into allIngredients
func GetIngredientsFromFile(fname string) (ingredients []Ingredient) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + fname)
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	yaml.Unmarshal(byteValue, &ingredients)
	return
}

// FindIngredient returns an Ingredient from allIngredients based on ID
func FindIngredient(id int) (ingredient Ingredient, err error) {
	for _, ingredient = range allIngredients {
		if ingredient.ID == id {
			return
		}
	}
	err = fmt.Errorf("No such ingredient %d", id)
	return
}
