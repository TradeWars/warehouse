# -
# `VERSION` contains the current version - aka the last tagged version of the
# repository. It's stored in a file named `VERSION` and should only ever be
# updated by running `make next`.
#
# `NEXT_VERSION` contains the next version number which is generated using the
# date and contains: the two-digit year, the week number, day of week and hour.
#
# `next` updates the `VERSION` file with the `NEXT_VERSION` value which
# signifies a new version has been reached and will be released.
# `release` applies the `VERSION` to the repository by tagging the current
# repository state with the `VERSION` value and then pushes the repo to the
# git server. It's important that your `.git/config` file contains the necessary
# `push` fields under the `origin` remote (or whatever remote you use).
# -

VERSION := $(shell cat VERSION)
NEXT_VERSION := $(shell date -u +%yw%W.%w.%H)

next:
	echo $(NEXT_VERSION) > VERSION

release:
	# re-tag this commit
	-git tag -d $(VERSION)
	git tag $(VERSION)
	# note: this requires that the configuration contains:
	# [remote "origin"]
	#     url = ...
	#     fetch = +refs/heads/*:refs/remotes/origin/*
	#     push = +refs/heads/*
	#     push = +refs/tags/*
	# in order to force tags to push alongside everything else.
	git push


# -
# Local Build
# -

LDFLAGS := -ldflags "-X github.com/TradeWars/warehouse/server.version=$(VERSION)"
-include .env

fast:
	go build $(LDFLAGS) -o warehouse

static:
	CGO_ENABLED=0 GOOS=linux go build -a $(LDFLAGS) -o warehouse .

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
