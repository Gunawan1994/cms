FROM golang:1.23.1 AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/app
COPY . .

RUN rm -r go.mod
RUN rm -r go.sum

RUN go mod init payment
RUN go mod tidy

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app

FROM alpine:latest
RUN apk add tzdata
ENV TZ Asia/Jakarta
COPY --from=builder /go/bin/app .
ENTRYPOINT ["./app"]
