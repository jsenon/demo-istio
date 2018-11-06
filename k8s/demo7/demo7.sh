#/bin/zsh

istioctl delete -f ../demo5/virtualservice-mirror.yml
istioctl delete -f ../demo4/outlier.yml
kubectl scale deployment demo-istio-v2 --replicas=2
istioctl create -f virtualservice_rate.yml
istioctl create -f quotas.yml
