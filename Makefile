.DEFAULT_GOAL := dev

.PHONY: dev
dev:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop app && docker-compose rm -f app

.PHONY: build
build:
	docker-compose build --no-cache --build-arg DOCKERFILE=Dockerfile.dev $(s) && docker-compose up -d $(s) --force-recreate

.PHONY: rebuild
rebuild:
	make stop
	make build