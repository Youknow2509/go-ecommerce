#!/bin/bash

# This script is used to run a Redis cluster in Docker containers.
docker compose -f ./environment/redis-cluster/docker-compose.yml -p redis-cluster-go-ecommerce up -d