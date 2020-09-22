.PHONY: build run

CONTAINER_TAG := twitter-random-irasutoya

build:
	docker build -t $(CONTAINER_TAG) .

run:
	docker run --rm --env-file .env $(CONTAINER_TAG)
