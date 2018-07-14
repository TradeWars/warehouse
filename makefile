VERSION := $(shell cat VERSION)
NEW_VERSION := $(shell date -u +%yw%W.%w.%H)
LDFLAGS := -ldflags "-X github.com/TradeWars/warehouse/server.version=$(VERSION)"
-include .env


# -
# Local
# -

fast:
	go build $(LDFLAGS) -o warehouse

static:
	CGO_ENABLED=0 GOOS=linux go build -a $(LDFLAGS) -o warehouse .

next:
	echo $(NEW_VERSION) > VERSION

release: build push
	# re-tag this commit
	-git tag -d $(VERSION)
	git tag $(VERSION)
	# build release binaries with current version tag
	GITHUB_TOKEN=$(GITHUB_TOKEN) goreleaser --rm-dist

local:
	WAREHOUSE_TEMPORARY=false \
	WAREHOUSE_BIND=localhost:7788 \
	WAREHOUSE_AUTH=cunning_fox \
	WAREHOUSE_MONGO_HOST=localhost \
	WAREHOUSE_MONGO_PORT=27017 \
	WAREHOUSE_MONGO_NAME=warehouse \
	WAREHOUSE_MONGO_USER=warehouse \
	WAREHOUSE_MONGO_PASS=warehouse \
	DEBUG=1 \
	./warehouse

# -
# Testing
# -

test:
	TESTING=1 go test -v -race ./server

test-pawn:
	sampctl p build
	sampctl p run

test-all: test test-pawn

databases:
	-docker stop mongodb
	-docker rm mongodb
	-docker stop express
	-docker rm express
	docker run \
		--name mongodb \
		--publish 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME=warehouse \
		-e MONGO_INITDB_ROOT_PASSWORD=warehouse \
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
	docker build -t southclaws/tw-warehouse:$(VERSION) .

push:
	docker push southclaws/tw-warehouse:$(VERSION)

run:
	docker run \
		--name warehouse \
		--env-file .env \
		southclaws/tw-warehouse:$(VERSION)
