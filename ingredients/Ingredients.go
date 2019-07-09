package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	pb "sandwiches/protobuf"
	"time"
)

const (
	port = ":50051"
)

type server struct{}

var allIngredients []pb.Ingredient

func main() {
	Start()
}

// Start ...
func Start() {
	allIngredients = getIngredientsFromFile("ingredients.json")
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
	return &ingredient, nil
}

func (s *server) GetIngredients(ctx context.Context, _ *pb.Empty) (*pb.MultipleIngredient, error) {
	log.Printf("Received request for all Ingredients")
	var result pb.MultipleIngredient
	for _, ingredient := range allIngredients {
		ing, _ := FindIngredient(ingredient.Id)
		result.Ingredients = append(result.Ingredients, &ing)
	}
	return &result, nil
}

func (s *server) AddIngredient(ctx context.Context, newIngredient *pb.NewIngredient) (*pb.Ingredient, error) {
	log.Printf("Received request to add new Ingredient")
	ingredient := pb.Ingredient{
		Name:  newIngredient.Name,
		Price: newIngredient.Price,
		Type:  newIngredient.Type,
		Id:    int32(len(allIngredients) + 1),
	}
	allIngredients = append(allIngredients, ingredient)
	// updateMenu()
	return &ingredient, nil
}

func updateMenu() {
	conn, err := grpc.Dial("menu:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	menuClient := pb.NewMenuClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = menuClient.UpdateMenu(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

func getIngredientsFromFile(fname string) (ingredients []pb.Ingredient) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + fname)
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &ingredients)
	return
}

// FindIngredient returns an Ingredient from allIngredients based on ID
func FindIngredient(id int32) (ingredient pb.Ingredient, err error) {
	for _, ingredient = range allIngredients {
		if ingredient.Id == id {
			return
		}
	}
	err = fmt.Errorf("No such ingredient %d", id)
	return
}
