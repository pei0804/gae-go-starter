DBNAME:=motivating_menu
ENV:=development

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
	sql-migrate up -env=$(ENV)

migrate/down:
	sql-migrate down -env=$(ENV)

migrate/status:
	sql-migrate status -env=$(ENV)

migrate/dry:
	sql-migrate up -dryrun -env=$(ENV)

gen:
	cd model && go generate

deploy:
	goapp deploy -application hoge ./app

rollback:
	appcfg.py rollback ./app -A hoge

local:
	goapp serve ./app