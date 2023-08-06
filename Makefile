DOCKER_USERNAME ?= anisurrahman75
APPLICATION_NAME ?= bookserver-golang-mysql
appName ?= serverBinary
RELEASE_NAME ?= bookserver

build-binary:
	@echo building Binary.....
	go build -o  ${appName} .

git-push:
	git add .; git commit -m "From MakeFile"; git push

run-binary: build-binary
	./${appName} startServer -a=false

helm-install:
	@echo helm-chart installing.....
	helm install ${RELEASE_NAME} helmChart

helm-uninstall:
	@echo helm-chart Un-installing.....
	helm uninstall ${RELEASE_NAME}

clean:
	@echo Cleaning up.....
	rm -rf ${appName}

CONTAINERS := book_server mysql
.PHONY: delete-containers
delete_container_%:
	@container_name=$*; \
	if docker ps --format '{{.Names}}' | grep -q "$$container_name"; then \
		docker kill "$$container_name"; \
	fi;
	@container_name=$*; \
	if docker ps -a --format '{{.Names}}' | grep -q "$$container_name"; then \
		docker rm "$$container_name"; \
	fi;
	@container_name=$*; \
	if docker images "$$container_name"; then \
		docker rmi "$$container_name"; \
	fi;

delete-containers: $(addprefix delete_container_, $(CONTAINERS))

run docker-compose:
	sudo systemctl restart docker
	docker-compose up --remove-orphan