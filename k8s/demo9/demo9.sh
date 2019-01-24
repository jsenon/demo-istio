#/bin/zsh

# Cleaning All

kubectl delete -f ../demo1/
kubectl delete -f ../demo2/
kubectl delete -f ../demo3/
kubectl delete -f ../demo4/
kubectl delete -f ../demo5/
kubectl delete -f ../demo6/
kubectl delete -f ../demo7/
kubectl delete -f ../demo8/

# Create v1

kubectl apply -f service.yml
kubectl apply -f deployment_v001.yml

# Create v2

kubectl apply -f service_v2.yml
kubectl apply -f deployment_v002.yml

# Create Virtual Service

istioctl create -f virtualservice.yml

# Inform

echo ""
echo "*****"
echo "You can now configure Auth0, and adapt jwt-policy.yaml file"
echo "After that applied it with istioctl create -f jwt-policy.yaml"
echo "*****"
