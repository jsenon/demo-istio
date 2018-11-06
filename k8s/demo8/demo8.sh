#/bin/zsh

istioctl delete -f ../demo7/virtualservice_rate.yml
istioctl create -f external.yml