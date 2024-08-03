# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS base
WORKDIR /src
COPY go.mod go.sum .
# RUN go mod download
RUN --mount=type=cache,target=/go/pkg/mod/ \
  go mod download -x
COPY . .

# Step 2
FROM base AS build-client
# RUN go build -o /bin/client ./cmd/client
RUN --mount=type=cache,target=/go/pkg/mod/ \
  go build -o /bin/client ./cmds/client

# Step 3
FROM base AS build-server
# RUN go build -o /bin/server ./cmd/server
RUN --mount=type=cache,target=/go/pkg/mod/ \
  go build -o /bin/server ./cmds/server

# Step 4 A
FROM scratch AS client
COPY --from=build-client /bin/client /bin/
ENTRYPOINT [ "/bin/client" ]

# Step 4 B
FROM scratch AS server
COPY --from=build-server /bin/server /bin/
ENTRYPOINT [ "/bin/server" ]