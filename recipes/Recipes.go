package main

import (
	"context"
	"errors"
	pb "github.com/calvinmclean/sandwiches/sandwiches"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// Recipe represents the name and collection of ingredients for a sandwich
type Recipe struct {
	Name     string  `json:"name" yaml:"name"`
	ID       int32   `json:"id" yaml:"id"`
	Bread    int32   `json:"bread" yaml:"bread"`
	Meats    []int32 `json:"meats" yaml:"meats"`
	Cheeses  []int32 `json:"cheeses" yaml:"cheeses"`
	Toppings []int32 `json:"toppings" yaml:"toppings"`
}

var allRecipes []Recipe

func main() {
	allRecipes = append(allRecipes, Recipe{"BLT", 1, 1, []int32{2}, []int32{}, []int32{3, 4, 5}})
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
	return &pb.Recipe{
		Name:     recipe.Name,
		Id:       recipe.ID,
		Bread:    recipe.Bread,
		Meats:    recipe.Meats,
		Cheeses:  recipe.Cheeses,
		Toppings: recipe.Toppings,
	}, nil
}

// FindRecipe returns a Recipe from allRecipes based on ID
func FindRecipe(id int32) (Recipe, error) {
	for _, recipe := range allRecipes {
		if recipe.ID == id {
			return recipe, nil
		}
	}
	var fake Recipe
	return fake, errors.New("No such recipe")
}
