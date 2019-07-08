package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	pb "sandwiches/protobuf"
)

const (
	port = ":50053"
)

type server struct{}

var allMenu []pb.MenuItem

var recipesClient pb.RecipesClient
var ingredientsClient pb.IngredientsClient

func main() {
	Start()
}

// Start ...
func Start() {
	conn := getConnection("ingredients:50051")
	ingredientsClient = pb.NewIngredientsClient(conn)

	conn = getConnection("recipes:50052")
	recipesClient = pb.NewRecipesClient(conn)

	allMenu = BuildMenu()

	fmt.Println(allMenu)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMenuServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getConnection(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func (s *server) GetMenuItem(ctx context.Context, in *pb.MenuItemRequest) (*pb.MenuItem, error) {
	log.Printf("Received: %d", in.Id)
	menuItem, _ := FindMenuItem(in.Id)
	return &pb.MenuItem{
		Name:  menuItem.Name,
		Price: menuItem.Price,
		Id:    menuItem.Id,
	}, nil
}

func (s *server) GetMenuItems(ctx context.Context, _ *pb.Empty) (*pb.MultipleMenuItem, error) {
	log.Printf("Received request for all MenuItems")
	var result pb.MultipleMenuItem
	for _, menuItem := range allMenu {
		men, _ := FindMenuItem(menuItem.Id)
		result.MenuItems = append(result.MenuItems, &men)
	}
	return &result, nil
}

func (s *server) UpdateMenu(ctx context.Context, _ *pb.Empty) (empty *pb.Empty, err error) {
	log.Printf("Received request to update menu")
	allMenu = BuildMenu()
	return
}

// BuildMenu gathers info from Recipes and Ingredients to create a list of
// MenuItems with calculated prices
func BuildMenu() (menu []pb.MenuItem) {
	recipes := getRecipes()

	// Add prices and names to the MenuItems
	for _, recipe := range recipes.Recipes {
		menu = append(menu, pb.MenuItem{
			Id:    recipe.Id,
			Price: calcRecipePrice(recipe),
			Name:  recipe.Name,
		})
	}
	return
}

func calcRecipePrice(recipe *pb.Recipe) (price float64) {
	price += calcPriceFromIngredients([]int32{recipe.Bread})
	price += calcPriceFromIngredients(recipe.Meats)
	price += calcPriceFromIngredients(recipe.Cheeses)
	price += calcPriceFromIngredients(recipe.Toppings)
	return
}

func calcPriceFromIngredients(items []int32) (price float64) {
	for _, item := range items {
		price += getIngredient(item).Price
	}
	return
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

// FindMenuItem returns a MenuItem from allMenu based on ID
func FindMenuItem(id int32) (menuItem pb.MenuItem, err error) {
	for _, menuItem = range allMenu {
		if menuItem.Id == id {
			return
		}
	}
	err = fmt.Errorf("No such MenuItem %d", id)
	return
}
