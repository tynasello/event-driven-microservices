#!/bin/bash

kubectl apply -f ./k8s/pre-deployment.yaml
kubectl apply -f ./k8s/message-broker-deployment.yaml
kubectl apply -f ./k8s/order-service-deployment.yaml
kubectl apply -f ./k8s/user-service-deployment.yaml
