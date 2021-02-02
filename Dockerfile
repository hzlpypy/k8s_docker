FROM golang:1.15.5-alpine3.12
COPY . /test_docker
WORKDIR /test_docker

ENV GOPROXY=https://goproxy.io,direct
ENV GO111MODULE=on
ENV GOMOD=/test_docker/go.mod

RUN go mod download && go build -o /main .

FROM alpine:3.12
COPY --from=0 /main /main
CMD ["/main"]
