FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o /go/bin/app -ldflags "-s -w" main.go

FROM alpine:3 as final

RUN apk update && \
    apk add --update-cache curl && \
    rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]
