appName ?= serverBinary
build-binnary:
	@echo building Binary.....
	go build -o  ${appName} .
git-push:
	git add .; git commit -m "commit_message"; git push
clean:
	@echo Cleaning all created file
	rm -rf ${appName}
