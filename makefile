VERSION := $(shell cat VERSION)
NEW_VERSION := $(shell date -u +%yw%W.%w.%H)
LDFLAGS := -ldflags "-X github.com/Southclaws/ScavengeSurviveCore/server.version=$(VERSION)"
-include .env


# -
# Local
# -

fast:
	go build $(LDFLAGS) -o ssc

static:
	CGO_ENABLED=0 GOOS=linux go build -a $(LDFLAGS) -o ssc .

next:
	echo $(NEW_VERSION) > VERSION

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
	-docker stop timescaledb
	-docker rm timescaledb
	-docker stop express
	-docker rm express
	-docker stop pgadmin
	-docker rm pgadmin
	docker run \
		--name mongodb \
		--publish 27017:27017 \
		--detach \
		mongo
	docker run \
		--name timescaledb \
		--publish 5432:5432 \
		--detach \
		-e POSTGRES_USER=default \
		-e POSTGRES_PASSWORD=default \
		timescale/timescaledb:0.10.0-pg9.6
	sleep 5
	docker run \
		--name express \
		--publish 8081:8081 \
		--link mongodb:mongo \
		--detach \
		mongo-express
	docker run \
		--name pgadmin \
		--publish 8082:8082 \
		-e "PGADMIN_DEFAULT_EMAIL=u@d.co" \
		-e "PGADMIN_DEFAULT_PASSWORD=password" \
		--detach \
		dpage/pgadmin4


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
