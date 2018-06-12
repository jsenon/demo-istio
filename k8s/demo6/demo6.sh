#/bin/zsh

istioctl delete -f ../demo5/virtualservice-mirror.yml

kubectl create ns mtls
kubectl label namespace mtls istio-injection=enabled

kubectl apply -f service_v2.yml
kubectl apply -f deployment_v002.yml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml -n mtls

kubectl create ns legacy

kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml -n legacy
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml -n default

