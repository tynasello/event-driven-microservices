#!/bin/bash

docker build -t shipping-service:latest .
docker tag shipping-service:latest tynas/shipping-service:latest
docker push tynas/shipping-service:latest

