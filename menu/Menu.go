package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"text/template"

	pb "sandwiches/sandwiches"
)

const (
	port = ":50053"
)

type menuItemPlus struct {
	MenuItem MenuItem
	Bread    pb.Ingredient
	Meats    []pb.Ingredient
	Cheeses  []pb.Ingredient
	Toppings []pb.Ingredient
}

type menuInfo struct {
	Menu        []menuItemPlus
	Ingredients []pb.Ingredient
}

// Recipe represents the name and collection of ingredients for a sandwich
type Recipe struct {
	Name     string  `json:"name" yaml:"name"`
	ID       int32   `json:"id" yaml:"id"`
	Bread    int32   `json:"bread" yaml:"bread"`
	Meats    []int32 `json:"meats" yaml:"meats"`
	Cheeses  []int32 `json:"cheeses" yaml:"cheeses"`
	Toppings []int32 `json:"toppings" yaml:"toppings"`
}

// MenuItem represents an entry in a Menu with an item's name and price
type MenuItem struct {
	ID    int32 `json:"id" yaml:"id"`
	Price float64
	Name  string
}

// Ingredient represents a sandwich ingredient with its name, price, and type
type Ingredient struct {
	Name  string  `json:"name" yaml:"name"`
	Price float64 `json:"price" yaml:"price"`
	Type  string  `json:"type" yaml:"type"`
	ID    int32   `json:"id" yaml:"id"`
}

type server struct{}

var allMenu []MenuItem

func main() {
	allMenu = BuildMenu()
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

func (s *server) GetMenuItem(ctx context.Context, in *pb.MenuItemRequest) (*pb.MenuItem, error) {
	log.Printf("Received: %d", in.Id)
	menuItem, _ := FindMenuItem(in.Id)
	return &pb.MenuItem{
		Name:  menuItem.Name,
		Price: menuItem.Price,
		Id:    menuItem.ID,
	}, nil
}

// ShowMenu responds to GET requests and builds new menuItemPlus
// in order to populate an HTML template with each MenuItem, its price, and a
// list of Ingredients by type
func ShowMenu(w http.ResponseWriter, r *http.Request) {
	ingredients, _ := getIngredients()

	var menuWithIngredients []menuItemPlus
	for _, menuItem := range allMenu {
		recipe := getRecipe(menuItem.ID)

		var meats []pb.Ingredient
		for _, id := range recipe.Meats {
			meats = append(meats, getIngredient(id))
		}
		var cheeses []pb.Ingredient
		for _, id := range recipe.Cheeses {
			cheeses = append(cheeses, getIngredient(id))
		}
		var toppings []pb.Ingredient
		for _, id := range recipe.Toppings {
			toppings = append(toppings, getIngredient(id))
		}
		menuWithIngredients = append(menuWithIngredients,
			menuItemPlus{menuItem, getIngredient(recipe.Bread), meats, cheeses, toppings})
	}

	info := menuInfo{menuWithIngredients, ingredients}
	tmpl, _ := template.ParseFiles("menu.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, info)
}

// BuildMenu gathers info from Recipes and Ingredients to create a list of
// MenuItems with calculated prices
func BuildMenu() (menu []MenuItem) {
	recipes := getRecipes()

	// Add prices and names to the MenuItems
	for _, recipe := range recipes {
		menu = append(menu, MenuItem{
			ID:    recipe.ID,
			Price: calcRecipePrice(recipe),
			Name:  recipe.Name,
		})
	}
	return
}

func calcRecipePrice(recipe Recipe) (price float64) {
	price += calcPriceFromIngredients([]int{recipe.Bread})
	price += calcPriceFromIngredients(recipe.Meats)
	price += calcPriceFromIngredients(recipe.Cheeses)
	price += calcPriceFromIngredients(recipe.Toppings)
	return
}

func calcPriceFromIngredients(items []int) (price float64) {
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
func FindMenuItem(id int32) (menuItem MenuItem, err error) {
	for _, menuItem = range allMenu {
		if menuItem.ID == id {
			return
		}
	}
	err = fmt.Errorf("No such MenuItem %d", id)
	return
}
