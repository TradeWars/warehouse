VERSION := $(shell date +%Y-%m-%dT%H-%M-%S)
LDFLAGS := -ldflags "-X main.version=$(VERSION)"
-include .env


# -
# Local
# -

fast:
	go build $(LDFLAGS) -o ssc

static:
	CGO_ENABLED=0 GOOS=linux go build -a $(LDFLAGS) -o ssc .


# -
# Testing
# -

test:
	TESTING=1 go test -v -race ./server

mongodb:
	-docker stop mongodb
	-docker rm mongodb
	-docker stop express
	-docker rm express
	docker run \
		--name mongodb \
		--publish 27017:27017 \
		--detach \
		mongo
	sleep 5
	docker run \
		--name express \
		--publish 8081:8081 \
		--link mongodb:mongo \
		--detach \
		mongo-express


# -
# Docker
# -

build: static
	docker build -t southclaws/ssc:$(VERSION) .

push:
	docker push -t southclaws/ssc:$(VERSION)

run:
	docker run \
		--name ssc \
		--env-file .env \
		southclaws/ssc:$(VERSION)
