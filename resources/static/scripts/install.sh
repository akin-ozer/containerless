#!/usr/bin/env bash

kind create cluster --config config/kind-config.yaml
#dashboard
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml
kubectl create clusterrolebinding default-admin --clusterrole cluster-admin --serviceaccount=default:default
token=$(kubectl get secrets -o jsonpath="{.items[?(@.metadata.annotations['kubernetes\.io/service-account\.name']=='default')].data.token}"|base64 --decode)
echo $token

#istio
#config/istio/bin/istioctl install --set profile=demo

#knative basic
kubectl apply --filename https://github.com/knative/serving/releases/download/v0.14.0/serving-crds.yaml
kubectl apply --filename https://github.com/knative/serving/releases/download/v0.14.0/serving-core.yaml
sleep 120

#ambassador
kubectl create namespace ambassador
kubectl apply --namespace ambassador \
  --filename https://getambassador.io/yaml/ambassador/ambassador-rbac.yaml \
  --filename config/ambassador-service.yaml
kubectl patch clusterrolebinding ambassador -p '{"subjects":[{"kind": "ServiceAccount", "name": "ambassador", "namespace": "ambassador"}]}'
kubectl set env --namespace ambassador  deployments/ambassador AMBASSADOR_KNATIVE_SUPPORT=true
sleep 120
kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress.class":"ambassador.ingress.networking.knative.dev"}}'
kubectl patch configmap/config-domain \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"knative.example.com":""}}'

#hpa
kubectl apply --filename https://github.com/knative/serving/releases/download/v0.14.0/serving-hpa.yaml

#eventing
kubectl apply  --selector knative.dev/crd-install=true \
--filename https://github.com/knative/eventing/releases/download/v0.14.0/eventing.yaml
kubectl apply --filename https://github.com/knative/eventing/releases/download/v0.14.0/eventing.yaml
kubectl apply --filename https://github.com/knative/eventing/releases/download/v0.14.0/in-memory-channel.yaml
sleep 120
#eventing
kubectl apply  --selector knative.dev/crd-install=true \
--filename https://github.com/knative/eventing/releases/download/v0.14.0/eventing.yaml
kubectl apply --filename https://github.com/knative/eventing/releases/download/v0.14.0/eventing.yaml
kubectl apply --filename https://github.com/knative/eventing/releases/download/v0.14.0/in-memory-channel.yaml
kubectl apply --filename config/imc-channel.yaml
kubectl apply --filename config/config-br-defaults.yaml
kubectl apply --filename https://github.com/knative/eventing/releases/download/v0.14.0/channel-broker.yaml

#restart
kubectl delete po --all -n knative-serving
sleep 120

#deploy app
#kubectl apply -f config/service.yaml
#kubectl run curl --image kloiadocker/curl
#nohup kubectl port-forward -n ambassador svc/ambassador 10000:80 &>/dev/null &
#echo "access within cluster:"
#echo "curl -H \"Host: helloworld-java-quarkus.default.knative.example.com\" http://$(kubectl get svc ambassador -n ambassador | awk '{print $3}' | sed -n 2p)"
#echo "-----------"
#echo "access from host: "
#echo "curl -H \"Host: helloworld-java-quarkus.default.knative.example.com\" http://127.0.0.1:10000"
