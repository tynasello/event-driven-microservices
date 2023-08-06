#!/bin/bash

kubectl apply -f ./k8s/pre-deployment.yaml
kubectl apply -f ./k8s/message-broker-deployment.yaml
sleep 10
kubectl apply -f ./k8s/order-service-deployment.yaml
kubectl apply -f ./k8s/user-service-deployment.yaml
kubectl apply -f ./k8s/inventory-service-deployment.yaml
kubectl apply -f ./k8s/shipping-service-deployment.yaml
kubectl apply -f ./k8s/payment-service-deployment.yaml
kubectl apply -f ./k8s/management-service-deployment.yaml
