#!/bin/bash

mvn clean install -DskipTests
docker build -t order-service:latest .
docker tag order-service:latest tynas/order-service:latest
docker push tynas/order-service:latest
