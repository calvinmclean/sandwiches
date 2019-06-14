# Running on Minikube

1. First start your Minikube VM and wait for setup to finish:
  ```shell
  minikube start -p sandwiches
  ```

2. Make sure to activate Minikube's Docker environment
  ```shell
  eval $(minikube -p sandwiches docker-env)
  ```

3. We want these images to be accessible to Kubernetes through the Minikube VM which can be done following these instructions (sourced from [here](https://blog.hasura.io/sharing-a-local-registry-for-minikube-37c7240d0615/)):
  ```shell
  wget https://gist.githubusercontent.com/coco98/b750b3debc6d517308596c248daf3bb1/raw/6efc11eb8c2dce167ba0a5e557833cc4ff38fa7c/kube-registry.yaml
  kubectl create -f kube-registry.yaml
  kubectl port-forward --namespace kube-system $(kubectl get po -n kube-system | grep kube-registry-v0 | \awk '{print $1;}') 5000:5000
  ```

4. Now build and push all images
  ```hell
  docker build -t localhost:5000/sandwiches-recipes:latest ./recipes/
  docker build -t localhost:5000/sandwiches-ingredients:latest ./ingredients/
  docker build -t localhost:5000/sandwiches-menu:latest ./menu/
  docker build -t localhost:5000/sandwiches-clerk:latest ./clerk/
  docker build -t localhost:5000/sandwiches-nginx:latest ./nginx/

  docker push localhost:5000/sandwiches-recipes:latest
  docker push localhost:5000/sandwiches-ingredients:latest
  docker push localhost:5000/sandwiches-menu:latest
  docker push localhost:5000/sandwiches-clerk:latest
  docker push localhost:5000/sandwiches-nginx:latest
  ```

5. Now start applying the Kubernetes stuff:
  ```shell
  kubectl apply -f k8s/recipes.yaml
  kubectl apply -f k8s/ingredients.yaml
  # Make sure recipes and ingredients are up and running before continuing
  kubectl apply -f k8s/menu.yaml
  kubectl apply -f k8s/clerk.yaml
  kubectl apply -f k8s/nginx.yaml
  ```

6. Finally, use Minikube to access the application
  ```shell
  minikube -p sandwiches service nginx
  ```
