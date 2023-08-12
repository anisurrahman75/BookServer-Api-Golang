#----------------------GO----------------------
appName ?= serverBinary
RELEASE_NAME ?= bookserver-api-mysql
build-binary:
	@echo building Binary.....
	go build -o  ${appName} .
git-push:
	git add .; git commit -m "From MakeFile"; git push
run-binary: build-binary
	./${appName} startServer -a=false
#------------------------Helm---------------------
helm-install:
	@echo helm-chart installing.....
	helm install ${RELEASE_NAME} helmChart
helm-uninstall:
	@echo helm-chart Un-installing.....
	helm uninstall ${RELEASE_NAME}
#------------------------Docker--------------------
DOCKER_USER:=anisurrahman75
COMPOSE_FILE_PATH := -f docker-compose.yaml
IMAGE := anisurrahman75/bookserver-api-mysql
TAG:=latest
compose-build:
	$(info Make: docker compose building.....)
	docker-compose $(COMPOSE_FILE_PATH) build --no-cache
	@make -s clean
compose-start:
	$(info Make: docker compose starting.....)
	docker-compose $(COMPOSE_FILE_PATH) up
compose-stop:
	$(info Make: Stopping containers.....)
	docker-compose $(COMPOSE_FILE_PATH) stop
compose-down:
	$(info Make: Stopping containers.....)
	docker-compose $(COMPOSE_FILE_PATH) down
compose-restart:
	$(info Make: Restarting  containers......)
	@make -s stop
	@make -s start
docker-push:
	$(info Make: Pushing  image.....)
	@docker push $(IMAGE):$(TAG)
docker-pull:
	$(info Make: Pulling  image.....)
	@docker pull $(IMAGE):$(TAG)
docker-clean:
	@docker system prune --volumes --force
docker-system-clean:
	@docker system prune -a
docker-login:
	$(info Make: Login to Docker Hub.)
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)