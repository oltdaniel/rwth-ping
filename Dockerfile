# build the application in a full environment
FROM golang:latest AS builder

RUN go version

WORKDIR /go/src/github.com/oltdaniel/rwth-ping

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app

CMD ["/app"]

EXPOSE 4001

# just a small alpine step to load the ca certificates
FROM alpine:3 as extras

RUN apk --no-cache add tzdata zip ca-certificates

WORKDIR /usr/share/zoneinfo
RUN zip -q -r -0 /zoneinfo.zip .

# Create a small scratch image as the final one
FROM scratch

COPY --from=builder /app .
COPY --from=extras /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV ZONEINFO /zoneinfo.zip
COPY --from=extras /zoneinfo.zip /

EXPOSE 4001

CMD ["/app"]