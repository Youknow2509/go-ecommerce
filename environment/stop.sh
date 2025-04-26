
# Stop all containers
docker compose -f ./environment/docker-compose.yml -p go_ecommerce down
docker compose -f ./environment/redis-cluster/docker-compose.yml -p redis-cluster-go-ecommerce down
docker compose -f ./environment/docker-compose-monitoring.yml -p go_ecommerce down
# Stop all containers and remove volumes