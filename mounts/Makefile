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
	docker build -f ./Dockerfile_bind . --check

# cache mounts
.PHONY: build-image
build-image:
	docker build -f ./Dockerfile --progress=plain . 2> log-all.log

.PHONY: build-image-server
build-image-server:
	docker build -f ./Dockerfile --target=server --progress=plain . 2> log-server.log

# bind mounts
.PHONY: build-image2
build-image2:
	docker build -f ./Dockerfile_bind --progress=plain . 2> bind-log-all.log

.PHONY: build-image-server2
build-image-server2:
	docker build -f ./Dockerfile_bind --target=server --progress=plain . 2> bind-log-server.log


.PHONY: clean-build-cache
clean-build-cache:
	docker builder prune -af

#------------------------------------------------------------------------------
# utility
#------------------------------------------------------------------------------
.PHONY: clean
clean:
	rm -rf client server *.log
