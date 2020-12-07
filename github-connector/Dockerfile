FROM golang:1.14-alpine3.12 as builder

WORKDIR /go/src/github.com/dekiel/github-slack-connectors/github-connector

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/github-connector
RUN mkdir /app && mv ./main /app/main

FROM alpine:3.12
WORKDIR /app

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /app/main /app

CMD ["./main"]