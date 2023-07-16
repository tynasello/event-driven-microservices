#!/bin/bash

docker build -t payment-service:latest .
docker tag payment-service:latest tynas/payment-service:latest
docker push tynas/payment-service:latest
