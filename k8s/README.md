# Running on Minikube

1. Start your Minikube VM and wait for setup to finish
  ```shell
  minikube start -p sandwiches
  minikube profile sandwiches
  ```

2. Start the application services
  ```shell
  kubectl apply -Rf k8s/
  kubectl get svc,pods,deploy
  ```

3. Use Minikube to access the application
  ```shell
  minikube service nginx
  ```


### Testing rolling updates and rollback

1. Start by entering an order at the `/clerk/order` endpoint by filling out the form in your browser

2. Roll out the new image and check its status
  ```shell
  kubectl set image deployment.extensions/clerk clerk=calvinmclean/sandwiches-clerk:v2
  kubectl rollout status deployment.extensions/clerk
  ```

3. Repeat step one and notice that more details are provided about your sandwich

4. Roll back
  ```shell
  kubectl rollout undo deployment.extensions/clerk
  ```
