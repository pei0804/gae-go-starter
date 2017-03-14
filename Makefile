ENV:=development
DB_CONFIG:=app/dbconfig.yml
DB_CONTAINER_NAME:=gaegostarter_db
DBNAME:=gaegostarter_db
PROJECT:=for-handson

setup:
	which glide || go get -v github.com/Masterminds/glide
	glide install

dev:
	which glide || go get -v github.com/Masterminds/glide
	which direnv || go get -v github.com/zimbatm/direnv
	direnv allow
	glide install

test:
	go test -v $(shell glide novendor)

db/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

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

deploy:
	goapp deploy -application $(PROJECT) ./app

rollback:
	appcfg.py rollback ./app -A $(PROJECT)

local:
	goapp serve ./app