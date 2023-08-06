#!/bin/bash

CLI_COMMAND="$1" docker-compose -f docker-compose.local.yml up --build
