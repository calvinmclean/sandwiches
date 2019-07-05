package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "sandwiches/sandwiches"
)

const (
	port = ":50051"
)

type server struct{}

// Ingredient represents a sandwich ingredient with its name, price, and type
type Ingredient struct {
	Name  string  `json:"name" yaml:"name"`
	Price float64 `json:"price" yaml:"price"`
	Type  string  `json:"type" yaml:"type"`
	ID    int32   `json:"id" yaml:"id"`
}

var allIngredients []Ingredient

func main() {
	allIngredients = append(allIngredients, Ingredient{"Cheddar", 1.50, "cheese", 1})
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterIngredientsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetIngredient(ctx context.Context, in *pb.IngredientRequest) (*pb.Ingredient, error) {
	log.Printf("Received: %d", in.Id)
	ingredient, _ := FindIngredient(in.Id)
	return &pb.Ingredient{
		Name:  ingredient.Name,
		Price: ingredient.Price,
		Type:  ingredient.Type,
		Id:    ingredient.ID,
	}, nil
}

func (s *server) GetIngredients(ctx context.Context, _ *pb.Empty) (*pb.MultipleIngredient, error) {
	log.Printf("Received request for all Ingredients")
	var result pb.MultipleIngredient
	for _, ingredient := range allIngredients {
		result.Ingredients = append(result.Ingredients, &pb.Ingredient{
			Name:  ingredient.Name,
			Price: ingredient.Price,
			Type:  ingredient.Type,
			Id:    ingredient.ID,
		})
	}
	return &result, nil
}

// FindIngredient returns an Ingredient from allIngredients based on ID
func FindIngredient(id int32) (Ingredient, error) {
	for _, ingredient := range allIngredients {
		if ingredient.ID == id {
			return ingredient, nil
		}
	}
	var fake Ingredient
	return fake, errors.New("No such ingredient")
}
