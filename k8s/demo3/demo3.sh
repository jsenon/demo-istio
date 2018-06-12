#/bin/zsh

istioctl delete -f ../demo2/virtualservice-play.yml
istioctl create -f virtualservice-canary.yml

kubectl apply -f service_v2.yml
kubectl apply -f deployment_v002.yml

