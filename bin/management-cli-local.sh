#!/bin/bash

POD_NAME_SUBSTRING="management-service"
EXECUTABLE_NAME="target/release/management-service"

if [ $# -eq 0 ]; then
  echo "Usage: $0 [cli_arguments]"
  exit 1
fi

POD_NAME=$(kubectl get pods -n edms -o=name | grep "$POD_NAME_SUBSTRING")

if [ -z "$POD_NAME" ]; then
  echo "Error: Pod with the specified substring not found."
  exit 1
fi

kubectl exec -it "$POD_NAME" -n edms -- "$EXECUTABLE_NAME" "$@"
