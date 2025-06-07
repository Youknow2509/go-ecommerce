#! /bin/bash

# Run all containers for application
echo "Starting all containers..."


echo "Running metrics container..."
docker compose -f ./docker-compose/docker-compose-metrics.yml -p java-ddd-001 up -d