#! /bin/bash

# Run all containers for application
echo "Starting all containers..."

echo "Running base container..."
docker compose -f ./docker-compose/docker-compose.yml -p go-ecommerce up -d

echo "Running elk container..."
docker compose -f ./docker-compose/docker-compose-elk.yml -p go-ecommerce up -d

echo "Running kafka container..."
docker compose -f ./docker-compose/docker-compose-kafka.yml -p go-ecommerce up -d

echo "Running redis cluster container..."
docker compose -f ./cluster-redis/docker-compose.yml -p go-ecommerce up -d
