package main

type Recipe struct {
	Name     string `json:"name" yaml:"name"`
	Id       int    `json:"id" yaml:"id"`
	Bread    int    `json:"bread" yaml:"bread"`
	Meats    []int  `json:"meats" yaml:"meats"`
	Cheeses  []int  `json:"cheeses" yaml:"cheeses"`
	Toppings []int  `json:"toppings" yaml:"toppings"`
}

type MenuItem struct {
	Id    int `json:"id" yaml:"id"`
	Price float64
	Name  string
}

type Ingredient struct {
	Name  string  `json:"name" yaml:"name"`
	Price float64 `json:"price" yaml:"price"`
	Type  string  `json:"type" yaml:"type"`
	Id    int     `json:"id" yaml:"id"`
}
