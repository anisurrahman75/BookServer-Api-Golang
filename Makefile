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