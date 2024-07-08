FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /api

COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/api

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/api /go/bin/api
COPY --from=builder /api/schema.yml /schema.yml
COPY --from=builder /api/.certs /.certs
COPY --from=builder /api/.swagger /.swagger

ENTRYPOINT ["/go/bin/api"]
