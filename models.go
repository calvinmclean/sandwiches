package main

// Recipe represents the name and collection of ingredients for a sandwich
type Recipe struct {
	Name     string `json:"name" yaml:"name"`
	ID       int    `json:"id" yaml:"id"`
	Bread    int    `json:"bread" yaml:"bread"`
	Meats    []int  `json:"meats" yaml:"meats"`
	Cheeses  []int  `json:"cheeses" yaml:"cheeses"`
	Toppings []int  `json:"toppings" yaml:"toppings"`
}

// MenuItem represents an entry in a Menu with an item's name and price
type MenuItem struct {
	ID    int `json:"id" yaml:"id"`
	Price float64
	Name  string
}

// Ingredient represents a sandwich ingredient with its name, price, and type
type Ingredient struct {
	Name  string  `json:"name" yaml:"name"`
	Price float64 `json:"price" yaml:"price"`
	Type  string  `json:"type" yaml:"type"`
	ID    int     `json:"id" yaml:"id"`
}
