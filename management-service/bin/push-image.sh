#!/bin/bash

docker build --no-cache -t management-service:latest .
docker tag management-service:latest tynas/management-service:latest
docker push tynas/management-service:latest

