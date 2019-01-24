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

# Inform

echo ""
echo "*****"
echo "Be carreful to have deploy demo9 before this demo"
echo "*****"

# Apply RBAC

istioctl create -f rbac.yaml
