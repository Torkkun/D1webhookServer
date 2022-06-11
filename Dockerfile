FROM golang:latest as builder
WORKDIR /workdir
COPY . .
# set env
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# build
RUN go build -o webhookser

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /workdir/webhookser /webhookser

ENTRYPOINT ["/webhookser"]