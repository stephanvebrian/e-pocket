.PHONY: up-dev down-dev 

up-dev:
	docker-compose -f docker-compose.dev.yaml up -d

down-dev:
	docker-compose -f docker-compose.dev.yaml down