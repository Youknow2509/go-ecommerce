#!/bin/bash

# This script is used to stop a Redis cluster in Docker containers.
docker compose -f ./environment/redis-cluster/docker-compose.yml -p redis-cluster-go-ecommerce down
# Remove the Redis data directory
rm -rf ./environment/redis-cluster/data