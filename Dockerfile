# This Dockerfile bundles the following things:
# - API binary
# - UI code / assets

# Note these things:
# - app is exposed on port 26000 (cannot be changed)

FROM golang:1.14 AS build_api
WORKDIR /go/src/github.com/ishkanan/tienlen
RUN go get -u github.com/alecthomas/gometalinter && \
    gometalinter -i
COPY api api
RUN cd api && \
    GOOS=linux go build -o /tmp/tienlen-server main.go && \
    go test ./... && \
    gometalinter --deadline=10m --config=linter.json ./...

FROM node:14.5.0 AS build_ui
COPY ui/ .
RUN NODE_ENV= npm ci && \
    npm run build

FROM ubuntu:focal AS final
WORKDIR /root/
RUN apt update && \
    DEBIAN_FRONTEND=noninteractive apt install -y tzdata
COPY --from=build_api /tmp/tienlen-server .
COPY --from=build_ui dist/ ui/
CMD ["/root/tienlen-server", "-addr", "0.0.0.0:26000", "-ui", "ui"]
EXPOSE 26000
