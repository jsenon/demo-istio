#/bin/zsh

kubectl apply -f service_200.yml
kubectl apply -f deployment_200.yml
kubectl apply -f service_503.yml
kubectl apply -f deployment_503.yml
kubectl apply -f service_200_stdspan.yml
kubectl apply -f deployment_200_stdspan.yml
istioctl create -f virtualservice.yml


