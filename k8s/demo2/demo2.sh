#/bin/zsh

istioctl delete -f ../demo1/virtualservice.yml
kubectl apply -f service.yml
kubectl apply -f deployment_v001.yml

kubectl apply -f service_pong.yml
kubectl apply -f deployment_pong_srv.yml

istioctl create -f virtualservice-play.yml