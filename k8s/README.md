# Running on Minikube

1. First start your Minikube VM and wait for setup to finish:
  ```shell
  minikube start -p sandwiches
  minikube profile sandwiches
  ```

2. Make sure to activate Minikube's Docker environment
  ```shell
  eval $(minikube docker-env)
  ```

3. We want these images to be accessible to Kubernetes through the Minikube VM which can be done following these instructions (sourced from [here](https://blog.hasura.io/sharing-a-local-registry-for-minikube-37c7240d0615/)):
  ```shell
  wget https://gist.githubusercontent.com/coco98/b750b3debc6d517308596c248daf3bb1/raw/6efc11eb8c2dce167ba0a5e557833cc4ff38fa7c/kube-registry.yaml
  kubectl create -f kube-registry.yaml
  kubectl port-forward --namespace kube-system $(kubectl get po -n kube-system | grep kube-registry-v0 | \awk '{print $1;}') 5000:5000
  ```

4. Now build and push all images
  ```shell
  docker build -t localhost:5000/sandwiches-recipes:v1 ./recipes/
  docker build -t localhost:5000/sandwiches-ingredients:v1 ./ingredients/
  docker build -t localhost:5000/sandwiches-menu:v1 ./menu/
  docker build -t localhost:5000/sandwiches-clerk:v1 ./clerk/
  docker build -t localhost:5000/sandwiches-nginx:v1 ./nginx/

  docker push localhost:5000/sandwiches-recipes:v1
  docker push localhost:5000/sandwiches-ingredients:v1
  docker push localhost:5000/sandwiches-menu:v1
  docker push localhost:5000/sandwiches-clerk:v1
  docker push localhost:5000/sandwiches-nginx:v1
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
  minikube service nginx
  ```


### Testing rolling updates and rollback

1. Start by entering an order at the `/clerk/order` endpoint by filling out the form in your browser

2. Create and push new image to roll out
  ```shell
  docker build -t localhost:5000/sandwiches-clerk:v2 ./clerk_v2/
  docker push localhost:5000/sandwiches-clerk:v2
  ```

3. Roll out the new image and check its status
  ```shell
  kubectl set image deployment.extensions/clerk clerk=localhost:5000/sandwiches-clerk:v2
  kubectl rollout status deployment.extensions/clerk
  ```

4. Repeat step one and notice that more details are provided about your sandwich

5. Roll back
  ```shell
  kubectl rollout undo deployment.extensions/clerk
  ```
