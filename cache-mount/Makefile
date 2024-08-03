#------------------------------------------------------------------------------
# go code
#------------------------------------------------------------------------------
.PHONY: lint
lint:
	go fmt ./...
	go vet ./...

.PHONY: build
build:
	go build -o server ./cmds/server
	go build -o client ./cmds/client

.PHONY: build
clean:
	rm -rf client server

# default version must be v5.0.13
update-pkg:
	go get github.com/go-chi/chi/v5@v5.0.14

revert-pkg:
	go get github.com/go-chi/chi/v5@v5.0.13

#------------------------------------------------------------------------------
# Docker
#------------------------------------------------------------------------------
.PHONY: check
check:
	docker build -f ./Dockerfile . --check

.PHONY: build-image
build-image:
	docker build --progress=plain . 2> log-all.log

.PHONY: build-image-server
build-image-server:
	docker build --target=server --progress=plain . 2> log-server.log

.PHONY: clean-build-cache
clean-build-cache:
	docker builder prune -af
