package main

import (
	"sandwiches/client"
	"sandwiches/ingredients"
	"sandwiches/recipes"
)

func main() {
	go ingredients.Start()
	go recipes.Start()
	client.Start()
}
