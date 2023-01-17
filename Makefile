clean:
	docker container prune
	docker image prune
	docker compose -f ./deployments/docker-compose.yml build --no-cache

deploy: clean
	docker compose -f ./deployments/docker-compose.yml up -d

logs: db-logs op-logs im-logs pp-logs 
	docker compose -f ./deployments/docker-compose.yml logs order_processor

db-logs: 
	docker compose -f ./deployments/docker-compose.yml logs db
op-logs: 
	docker compose -f ./deployments/docker-compose.yml logs order_processor
im-logs: 
	docker compose -f ./deployments/docker-compose.yml logs inventory_manager
pp-logs: 
	docker compose -f ./deployments/docker-compose.yml logs payment_processor
stop:
	docker compose -f ./deployments/docker-compose.yml down