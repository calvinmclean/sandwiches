package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "github.com/calvinmclean/sandwiches/sandwiches"
)

func main() {
	ing := getIngredient(int32(1))
	rec := getRecipe(int32(1))

	log.Printf("Ingredient 1 name: %s", ing.Name)
	log.Printf("Recipe 1 name: %s", rec.Name)
}

func getConnection(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func getIngredient(id int32) pb.Ingredient {
	conn := getConnection("localhost:50051")
	defer conn.Close()
	c := pb.NewIngredientsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ing, err := c.GetIngredient(ctx, &pb.IngredientRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *ing
}

func getRecipe(id int32) pb.Recipe {
	conn := getConnection("localhost:50052")
	defer conn.Close()
	c := pb.NewRecipesClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rec, err := c.GetRecipe(ctx, &pb.RecipeRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return *rec
}
