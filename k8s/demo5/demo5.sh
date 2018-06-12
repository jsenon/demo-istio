#/bin/zsh

istioctl delete -f ../demo4/virtualservice-delay.yml
kubectl scale deployment demo-istio-v2 --replicas=2
istioctl create -f virtualservice-mirror.yml


