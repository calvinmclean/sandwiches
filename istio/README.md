# Using Istio
Note: these instructions are a modified version if [Istio's demo](https://istio.io/docs/examples/bookinfo/#if-you-are-running-on-kubernetes)

1. Prepare Minikube environment
  ```shell
  minikube start -p istio-sandwiches --memory=16384 --cpus=4 --kubernetes-version=v1.14.2
  minikube profile istio-sandwiches
  ```

2. Install Istio following [these instructions](https://istio.io/docs/setup/kubernetes/install/kubernetes/#prerequisites) (enable "strict mutual TLS")

3. Label the default namespace to allow [Istio automatic sidecar injection](https://istio.io/docs/setup/kubernetes/additional-setup/sidecar-injection/#automatic-sidecar-injection)
  ```shell
  kubectl label namespace default istio-injection=enabled
  ```

4. Deploy the Sandwiches application and make sure it is running
  ```shell
  kubectl apply -f istio/sandwiches.yaml
  kubectl get svc,pods,deploy
  kubectl exec -it $(kubectl get pod -l app=menu -o jsonpath='{.items[0].metadata.name}') -c menu -- curl menu:8080/menu/show/
  # If the Menu prints, this confirms that Menu, Recipes, and Ingredients services are working
  # If the Menu does not print, you should delete the Menu pod so it will restart:
  kubectl delete pod <pod_name>
  ```

5. Define the Ingress gateway for the Sandwiches application
  ```shell
  kubectl apply -f istio/sandwiches-gateway.yaml
  kubectl get gateway
  ```

6. Determine the Ingress host and IP
  ```shell
  kubectl get svc istio-ingressgateway -n istio-system
  export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
  export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].nodePort}')
  export INGRESS_HOST=$(minikube ip)
  export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
  ```

7. Confirm that the app is running by hitting a few endpoints
  ```shell
  curl http://${GATEWAY_URL}/menu/show/
  # Check in browser by echoing and clicking link
  echo http://${GATEWAY_URL}/clerk/order/
  ```

8. Define Destination Rules
  ```shell
  kubectl apply -f istio/destination-rule.yaml
  ```

9. Apply initial Virtual Service (_note_: now I am following [this guide](https://istio.io/docs/tasks/traffic-management/request-routing/))
  ```shell
  kubectl apply -f istio/virtual-service-all-v1.yaml
  kubectl get vs
  ```