DB_NAME ?= postgres
DB_TYPE ?= postgres
DB_USER ?= postgres
DB_PWD ?= passwordnya
IP=localhost

PSQLURL ?= $(DB_TYPE)://$(DB_USER):$(DB_PWD)@$(IP):5432/$(DB_NAME)

NET_NAME ?= postgres_net
CON_NAME ?= postgres_db
POST_VERSION ?= postgres:15.2-alpine

.PHONY : postgresup postgresdown test build

postgresup:
	docker run \
	--name $(CON_NAME) \
	--network $(NET_NAME) \
	-p 5433:5432 \
	-e POSTGRES_PASSWORD=$(DB_PWD) \
	-v $(PWD):/var/lib/postgresql/data \
	-t $(POST_VERSION) \
	-d $(CON_NAME)

postgresdown:
	docker stop go-api-postgres && docker rm go-api-postgres 

psql:
	docker exec -it go-api-postgres psql $(PSQLURL)

test: 
	go test ./test -v

build:
	docker build -t go-rest-api:0.0.1 .

go_app:
	docker run --name go-rest-api \
	-p 8080:8080 \
	--network $(NET_NAME) \
	-d -t go-rest-api:0.0.1

run: go_app postgresup
IMAGE_NAME ?= program_1

delete_container:
	docker rm -f $(CON_NAME)
delete_image:
	docker rmi -f $(IMAGE_NAME)