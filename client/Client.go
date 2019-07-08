package main

import (
	"context"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"

	pb "sandwiches/protobuf"
)

var recipesClient pb.RecipesClient
var ingredientsClient pb.IngredientsClient
var menuClient pb.MenuClient

func main() {
	Start()
}

// Start ...
func Start() {
	conn := getConnection("ingredients:50051")
	ingredientsClient = pb.NewIngredientsClient(conn)

	conn = getConnection("recipes:50052")
	recipesClient = pb.NewRecipesClient(conn)

	conn = getConnection("menu:50053")
	menuClient = pb.NewMenuClient(conn)
	defer conn.Close()

	allIngredients := getAllIngredients()
	allRecipes := getAllRecipes()
	// menu := getMenuItems()

	log.Printf("All Ingredients: %v", allIngredients.Ingredients)
	log.Printf("All Recipes: %v", allRecipes.Recipes)
	// log.Printf("Menu: %v", menu)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ingredients", apiGetAllIngredients).Methods("GET")
	router.HandleFunc("/ingredients", apiAddIngredient).Methods("POST")
	router.HandleFunc("/ingredients/{id}", apiGetSingleIngredient).Methods("GET")
	// router.HandleFunc("/ingredients/{id}", DeleteIngredient).Methods("DELETE")

	router.HandleFunc("/recipes", apiGetAllRecipes).Methods("GET")
	router.HandleFunc("/recipes", apiAddRecipe).Methods("POST")
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
