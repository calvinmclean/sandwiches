# Sandwich Shop Microservices


### How to use

Start with `docker-compose up`

Then, you can visit http://localhost/clerk/order to create a sandwich by giving a sandwich ID and listing extra ingredient IDs like: `1 2 3`. This is not very fancy but it works for now.

Ordering from the Clerk is the only interactive part of the program at the moment. Other links can be explored to get JSON responses:

| Link                           | Purpose                                         |
| :----------------------------- | :---------------------------------------------- |
| http://localhost/menu/show     | show the menu (this one is plain text not JSON) |
| http://localhost/menu/         | get the menu JSON                               |
| http://localhost/menu/1        | get the first item on the menu                  |
| http://localhost/ingredients   | get all ingredients                             |
| http://localhost/ingredients/1 | get the first ingredient                        |
| http://localhost/recipes       | get all recipes                                 |
| http://localhost/recipes/1     | get the first recipe                            |


## Microservices

This program is composed of 4 different services and an Nginx proxy. The microservices are:
  1. Clerk: responsible for handling purchases
    - This is an effective microservice because all it does is receive information from users and then requests additional information from other services in order to return information to the user
    - This can be scaled to increase simultaneous order processing
  2. Ingredients: holds a list of ingredients and their prices
    - This is an effective microservice because it only handles one thing: ingredients. It exists to provide information about ingredients to other services
  3. Menu: holds a list of recipes and their prices to serve to users
    - This is an effective microservice because it just holds a list of menu items that can be requested by the Clerk. It relies on the Ingredients and Recipes microservices to provide up-to-date and accurate information when creating the menu
    - Since all values are from other services, if the price of an ingredient changes or a new recipe is added, the menu will be able to easily update its information
  4. recipes: holds a list of recipes and their required ingredients
    - This is an effective microservice because it simply serves information about recipes to the Menu service


## Event-driven aspects of this program/service

| Event                      | Response                                                                                    | Implemented? |
| :------------------------- | :------------------------------------------------------------------------------------------ | :----------- |
| Ingredient price change    | trigger an event in the Menu service to update the prices of all items with that ingredient | No           |
| New recipe is added        | trigger an event in the Menu service to add the new recipe to the menu                      | No           |
| Customer visits store page | clerk requests menu from Menu service                                                       | No           |
| Customer submits order     | clerk consults Menu and Ingredients to calculate price                                      | Yes          |

**Note**: Some of these event-response ideas are not implemented because they require more aspects of a RESTful API that I haven't programmed yet. Also I am currently using YAML files instead of a database to hold information about recipes and ingredients which also makes it more difficult to change/update recipes and ingredients.

**Note**: Currently the `models.go` file is required by each service so I just copy it to each directory. If this was a more robust service, the models would be published on Github and imported by each service.
