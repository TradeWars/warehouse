VERSION := $(shell date -u +%yw%W.%w.%H)
LDFLAGS := -ldflags "-X github.com/Southclaws/ScavengeSurviveCore/server.version=$(VERSION)"
-include .env


# -
# Local
# -

fast:
	go build $(LDFLAGS) -o ssc

static:
	CGO_ENABLED=0 GOOS=linux go build -a $(LDFLAGS) -o ssc .

release: build push
	# re-tag this commit
	-git tag -d $(VERSION)
	git tag $(VERSION)
	# build release binaries with current version tag
	GITHUB_TOKEN=$(GITHUB_TOKEN) goreleaser --rm-dist

# -
# Testing
# -

test:
	TESTING=1 go test -v -race ./server

databases:
	-docker stop mongodb
	-docker rm mongodb
	-docker stop postgres
	-docker rm postgres
	-docker stop express
	-docker rm express
	docker run \
		--name mongodb \
		--publish 27017:27017 \
		--detach \
		mongo
	docker run \
		--name postgres \
		--publish 5432:5432 \
		--detach \
		postgres
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
	docker push southclaws/ssc:$(VERSION)

run:
	docker run \
		--name ssc \
		--env-file .env \
		southclaws/ssc:$(VERSION)
