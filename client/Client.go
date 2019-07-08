package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "sandwiches/sandwiches"
)

var recipesClient pb.RecipesClient
var ingredientsClient pb.IngredientsClient
var menuClient pb.MenuClient

func main() {
	conn := getConnection("localhost:50051")
	ingredientsClient = pb.NewIngredientsClient(conn)

	conn = getConnection("localhost:50052")
	recipesClient = pb.NewRecipesClient(conn)

	conn = getConnection("localhost:50053")
	menuClient = pb.NewMenuClient(conn)
	defer conn.Close()

	ing := getSingleIngredient(int32(1))
	rec := getSingleRecipe(int32(1))
	allIngredients := getAllIngredients()
	allRecipes := getAllRecipes()
	// menu := getMenuItems()

	log.Printf("Ingredient 1 name: %s", ing.Name)
	log.Printf("All Ingredients: %v", allIngredients.Ingredients)

	log.Printf("Recipe 1 name: %s", rec.Name)
	log.Printf("    Meats:    %v", rec.Meats)
	log.Printf("    Cheeses:  %v", rec.Cheeses)
	log.Printf("    Toppings: %v", rec.Toppings)
	log.Printf("All Recipes: %v", allRecipes.Recipes)

	// log.Printf("Menu: %v", menu)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ingredients", apiGetAllIngredients).Methods("GET")
	// router.HandleFunc("/ingredients", AddIngredient).Methods("POST")
	router.HandleFunc("/ingredients/{id}", apiGetSingleIngredient).Methods("GET")
	// router.HandleFunc("/ingredients/{id}", DeleteIngredient).Methods("DELETE")

	router.HandleFunc("/recipes", apiGetAllRecipes).Methods("GET")
	// router.HandleFunc("/recipes", AddRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", apiGetSingleRecipe).Methods("GET")
	// router.HandleFunc("/recipes/{id}", DeleteRecipe).Methods("DELETE")

	// router.HandleFunc("/menu", apiGetAllMenuItems).Methods("GET")
	// router.HandleFunc("/menu/show", apiShowMenu).Methods("GET")
	// router.HandleFunc("/menu/{id}", apiGetSingleMenuItem).Methods("GET")
	// router.HandleFunc("/menu", UpdateMenu).Methods("POST")

	http.ListenAndServe(":8080", router)
}

// Start gRPC functions

func getConnection(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func getSingleIngredient(id int32) pb.Ingredient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ing, err := ingredientsClient.GetIngredient(ctx, &pb.IngredientRequest{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *ing
}

func getAllIngredients() pb.MultipleIngredient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ing, err := ingredientsClient.GetIngredients(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *ing
}

func getSingleRecipe(id int32) pb.Recipe {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rec, err := recipesClient.GetRecipe(ctx, &pb.RecipeRequest{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *rec
}

func getAllRecipes() pb.MultipleRecipe {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rec, err := recipesClient.GetRecipes(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *rec
}

func getMenuItem(id int32) pb.MenuItem {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	menu, err := menuClient.GetMenuItem(ctx, &pb.MenuItemRequest{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *menu
}

func getMenuItems() pb.MultipleMenuItem {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	menu, err := menuClient.GetMenuItems(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *menu
}

// End gRPC functions

// Start REST API functions

func apiGetSingleIngredient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ingredientJSON := getIngredientJSON(id)
	w.Header().Set("Content-Type", "application/json")
	w.Write(ingredientJSON)
}

func apiGetAllIngredients(w http.ResponseWriter, r *http.Request) {
	var ingredientsJSON []byte
	ingredientsJSON, _ = json.Marshal(getAllIngredients())
	w.Header().Set("Content-Type", "application/json")
	w.Write(ingredientsJSON)
}

func getIngredientJSON(id int) (result []byte) {
	ingredient := getSingleIngredient(int32(id))
	result, _ = json.Marshal(ingredient)
	return
}

func apiGetSingleRecipe(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	recipeJSON := getRecipeJSON(id)
	w.Header().Set("Content-Type", "application/json")
	w.Write(recipeJSON)
}

func apiGetAllRecipes(w http.ResponseWriter, r *http.Request) {
	var recipesJSON []byte
	recipesJSON, _ = json.Marshal(getAllRecipes())
	w.Header().Set("Content-Type", "application/json")
	w.Write(recipesJSON)
}

func getRecipeJSON(id int) (result []byte) {
	recipe := getSingleRecipe(int32(id))
	result, _ = json.Marshal(recipe)
	return
}
