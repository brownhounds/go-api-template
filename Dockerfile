FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /api

COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/api

FROM golang:alpine

RUN apk update && apk --no-cache add curl

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/api /go/bin/api
COPY --from=builder /api/schema.yml /go/schema.yml
COPY --from=builder /api/.certs /go/.certs
COPY --from=builder /api/.swagger /go/.swagger

HEALTHCHECK --interval=5s --timeout=10s --retries=5 \
  CMD curl --fail --cacert /go/.certs/server.pem https://127.0.0.1:8080/health || exit 1

ENTRYPOINT ["/go/bin/api"]
