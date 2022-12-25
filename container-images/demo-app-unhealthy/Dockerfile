FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o /go/bin/app -ldflags "-s -w" main.go

FROM alpine:3 as final

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]
