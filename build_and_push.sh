#!/bin/bash
docker build -t calvinmclean/sandwiches-recipes:v1 ./recipes/
docker build -t calvinmclean/sandwiches-ingredients:v1 ./ingredients/
docker build -t calvinmclean/sandwiches-menu:v1 ./menu/
docker build -t calvinmclean/sandwiches-menu:v2 ./menu_v2/
docker build -t calvinmclean/sandwiches-clerk:v1 ./clerk/
docker build -t calvinmclean/sandwiches-clerk:v2 ./clerk_v2/
docker build -t calvinmclean/sandwiches-nginx:v1 ./nginx/

docker push calvinmclean/sandwiches-recipes:v1
docker push calvinmclean/sandwiches-ingredients:v1
docker push calvinmclean/sandwiches-menu:v1
docker push calvinmclean/sandwiches-menu:v2
docker push calvinmclean/sandwiches-clerk:v1
docker push calvinmclean/sandwiches-clerk:v2
docker push calvinmclean/sandwiches-nginx:v1
