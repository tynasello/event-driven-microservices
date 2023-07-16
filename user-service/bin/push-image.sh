#!/bin/bash

docker build -t user-service:latest .
docker tag user-service:latest tynas/user-service:latest
docker push tynas/user-service:latest

