package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "sandwiches/sandwiches"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

var allRecipes []pb.Recipe

func main() {
	allRecipes = append(allRecipes, pb.Recipe{
		Name:     "BLT",
		Id:       1,
		Bread:    1,
		Meats:    []int32{2},
		Cheeses:  []int32{},
		Toppings: []int32{3, 4, 5},
	})
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRecipesServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetRecipe(ctx context.Context, in *pb.RecipeRequest) (*pb.Recipe, error) {
	log.Printf("Received: %d", in.Id)
	recipe, _ := FindRecipe(in.Id)
	return &recipe, nil
}

func (s *server) GetRecipes(ctx context.Context, _ *pb.Empty) (*pb.MultipleRecipe, error) {
	log.Printf("Received request for all Recipes")
	var result pb.MultipleRecipe
	for _, recipe := range allRecipes {
		rec, _ := FindRecipe(recipe.Id)
		result.Recipes = append(result.Recipes, &rec)
	}
	return &result, nil
}

func (s *server) AddRecipe(ctx context.Context, newRecipe *pb.NewRecipe) (*pb.Recipe, error) {
	log.Printf("Received request to add new Recipe")
	recipe := pb.Recipe{
		Name:     newRecipe.Name,
		Id:       int32(len(allRecipes) + 1),
		Bread:    newRecipe.Bread,
		Cheeses:  newRecipe.Cheeses,
		Meats:    newRecipe.Meats,
		Toppings: newRecipe.Toppings,
	}
	allRecipes = append(allRecipes, recipe)
	return &recipe, nil
}

// FindRecipe returns a Recipe from allRecipes based on ID
func FindRecipe(id int32) (recipe pb.Recipe, err error) {
	for _, recipe = range allRecipes {
		if recipe.Id == id {
			return
		}
	}
	err = fmt.Errorf("No such recipe %d", id)
	return
}
