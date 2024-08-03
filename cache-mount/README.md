# cache-mount

This example uses [cache mount](https://docs.docker.com/build/guide/mounts/)

## How to check

1. build all by `docker build --progress=plain . 2> log-all.log` then check log
2. clean build cache by `docker builder prune -af`
3. build only server by `docker build --target=server --progress=plain . 2> log-server1.log` then check log
4. build only server by `docker build --target=server --progress=plain . 2> log-server2.log` again then check log
5. upgrade package `chi` by `go get github.com/go-chi/chi/v5@v5.0.14`
6. build only server by `docker build --target=server --progress=plain . 2> log-server3.log` again then check log
