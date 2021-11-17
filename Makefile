CMD=default

echo:
	@echo $(CMD)

##############################
# make docker environmental
##############################
up:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose down

down-rmi:
	docker-compose down --rmi all
ps:
	docker-compose ps

dev:
	sh ./scripts/container.sh

##############################
# make frontend production in nginx container
##############################
frontend-install:
	docker-compose exec nginx ash -c 'cd /var/www/frontend && yarn install'

frontend-build:
	docker-compose exec nginx ash -c 'cd /var/www/frontend && yarn build'

##############################
# backend
##############################
container-dev:
	docker-compose exec app go run src/main.go $(CMD)

container-build:
	docker-compose exec app go build -o dist/app src/main.go

container-tidy:
	docker-compose exec app go mod tidy

container-test:
	docker-compose exec app go test -v ./src/...

##############################
# backend go
##############################

back-serve:
	cd app/backend && ./dist/app $(CMD)

back-dev:
	cd app/backend && go run src/main.go $(CMD)

main:
	cd app/backend && go build -o dist/app src/main.go

tidy:
	cd app/backend && go mod tidy

module-list:
	cd app/backend && go list -m all

test:
	cd app/backend && go test -v ./src/...

module:
ifeq ($(CMD),default)
	@echo invalid parameter!
else
	@echo make new module: $(CMD)
	@cd app/backend && mkdir src/$(CMD) && cd src/$(CMD) && go mod init $(CMD) && touch $(CMD).go
endif

controller:
ifeq ($(CMD),default)
	@echo invalid parameter!
else
	@echo make new controller: $(CMD)Controller
	@cd app/backend && mkdir src/controllers/$(CMD)Controller && cd src/controllers/${CMD}Controller && go mod init ${CMD}Controller && touch ${CMD}Controller.go
endif

repository:
ifeq ($(CMD),default)
	@echo invalid parameter!
else
	@echo make new repository: $(CMD)Repository
	@cd app/backend && mkdir src/repositories/$(CMD)Repository && cd src/repositories/$(CMD)Repository && go mod init $(CMD)Repository && touch $(CMD)Repository.go
endif

service:
ifeq ($(CMD),default)
	@echo invalid parameter!
else
	@echo make new service: $(CMD)
	@cd app/backend && mkdir src/services/$(CMD)Service && cd src/services/$(CMD)Service && go mod init $(CMD)Service && touch $(CMD)Service.go
endif

##############################
# web server(nginx)
##############################
nginx-t:
	docker-compose exec nginx ash -c 'nginx -t'

nginx-reload:
	docker-compose exec nginx ash -c 'nginx -s reload'

nginx-stop:
	docker-compose exec nginx ash -c 'nginx -s stop'


##############################
# db container(mysql)
##############################
mysql:
	docker-compose exec db bash -c 'mysql -u $$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE'

db:
	sh ./scripts/db.sh

##############################
# circle ci
##############################
circleci:
	cd app/backend && circleci build

ci:
	circleci build

##############################
# mock-server docker container
##############################
mock-up:
	docker-compose -f ./docker-compose.mock.yml up -d

mock-down:
	docker-compose -f ./docker-compose.mock.yml down

mock-ps:
	docker-compose -f ./docker-compose.mock.yml ps

##############################
# swagger docker container
##############################
swagger-up:
	docker-compose -f ./docker-compose.swagger.yml up -d

swagger-down:
	docker-compose -f ./docker-compose.swagger.yml down

swagger-ps:
	docker-compose -f ./docker-compose.swagger.yml ps

##############################
# swagger codegen mock-server
##############################
codegen-mock:
	rm -rf api/node-mock/* && \
	swagger-codegen generate -i api/api.yaml -l nodejs-server -o api/node-mock && \
	sed -i -e "s/serverPort = 8080/serverPort = 3200/g" api/node-mock/index.js && \
	cd api/node-mock && npm run prestart

codegen-changeport:
	sed -i -e "s/serverPort = 8080/serverPort = 3200/g" api/node-mock/index.js

codegen-prestart:
	cd api/node-mock && npm run prestart

codegen-start:
	cd api/node-mock && npm run start

##############################
# integration
##############################
prod:
	sh ./scripts/db.sh && sh ./scripts/container.sh

##############################
# etc
##############################
request:
	curl localhost:8080/public
