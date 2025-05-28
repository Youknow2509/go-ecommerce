
# run all containers
# run base
docker compose -f ./environment/docker-compose.yml -p go_ecommerce up  -d
# run redis cluster
docker compose -f ./environment/redis-cluster/docker-compose.yml -p go_ecommerce up -d
# run monitor
docker compose -f ./environment/docker-compose-monitoring.yml -p go_ecommerce up -d
# run elk
docker compose -f ./environment/docker-compose-elk.yml -p go_ecommerce up -d