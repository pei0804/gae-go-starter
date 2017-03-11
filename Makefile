ENV:=development
DB_CONFIG:=app/dbconfig.yml
DB_CONTAINER_NAME:=gaegostarter_db
DBNAME:=gaegostarter_db

setup:
	which glide || go get -v github.com/Masterminds/glide
	glide install

dev:
	which sql-migrate || go get -v github.com/rubenv/sql-migrate/...
	which scaneo || go get -v github.com/variadico/scaneo
	which glide || go get -v github.com/Masterminds/glide
	which direnv || go gore -v github.com/zimbatm/direnv
	direnv allow
	glide install

test:
	go test -v $(shell glide novendor)

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/up:
	sql-migrate up -env=$(ENV) -config=$(DB_CONFIG)

migrate/down:
	sql-migrate down -env=$(ENV) -config=$(DB_CONFIG)

migrate/status:
	sql-migrate status -env=$(ENV) -config=$(DB_CONFIG)

migrate/dry:
	sql-migrate up -dryrun -env=$(ENV) -config=$(DB_CONFIG)

docker/build: Dockerfile docker-compose.yml
	docker-compose build

docker/start:
	docker-compose up -d

docker/stop:
	docker-compose down

docker/logs:
	docker-compose logs

docker/clean:
	docker-compose rm

gen:
	cd model && go generate

deploy:
	goapp deploy -application for-handson ./app

rollback:
	appcfg.py rollback ./app -A for-handson

local:
	goapp serve ./app