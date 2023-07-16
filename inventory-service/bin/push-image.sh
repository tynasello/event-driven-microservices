#!/bin/bash

docker build -t inventory-service:latest .
docker tag inventory-service:latest tynas/inventory-service:latest
docker push tynas/inventory-service:latest

