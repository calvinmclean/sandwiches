package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
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

	ing := getIngredient(int32(1))
	rec := getRecipe(int32(1))
	allIngredients := getIngredients()
	allRecipes := getRecipes()
	// menu := getMenuItems()

	log.Printf("Ingredient 1 name: %s", ing.Name)
	log.Printf("All Ingredients: %v", allIngredients.Ingredients)

	log.Printf("Recipe 1 name: %s", rec.Name)
	log.Printf("    Meats:    %v", rec.Meats)
	log.Printf("    Cheeses:  %v", rec.Cheeses)
	log.Printf("    Toppings: %v", rec.Toppings)
	log.Printf("All Recipes: %v", allRecipes.Recipes)

	// log.Printf("Menu: %v", menu)
}

func getConnection(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func getIngredient(id int32) pb.Ingredient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ing, err := ingredientsClient.GetIngredient(ctx, &pb.IngredientRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *ing
}

func getIngredients() pb.MultipleIngredient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ing, err := ingredientsClient.GetIngredients(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *ing
}

func getRecipe(id int32) pb.Recipe {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rec, err := recipesClient.GetRecipe(ctx, &pb.RecipeRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *rec
}

func getRecipes() pb.MultipleRecipe {
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
	menu, err := menuClient.GetMenuItem(ctx, &pb.MenuItemRequest{Id: 1})
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
