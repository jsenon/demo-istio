#/bin/zsh

istioctl delete -f ../demo2/virtualservice-play.yml
istioctl delete -f ../demo3/virtualservice-canary.yml

istioctl create -f virtualservice-delay.yml

kubectl apply -f ../demo3/deployment_v002.yml
kubectl apply -f ../demo3/service_v2.yml

kubectl apply -f ../demo4/circuit-breaker.yml

istioctl create -f outlier.yml




