#first stage - builder
FROM golang:1.11.0-stretch as builder

COPY . /SingleStage
WORKDIR /SingleStage

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o SingleStage 



#second stage 
FROM alpine:latest

RUN apk add --no-cache tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /root/

COPY --from=builder /SingleStage .

CMD ["./SingleStage"]