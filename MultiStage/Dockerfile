#first stage - builder
FROM golang:1.11.0-stretch as builder

COPY . /MultiStage
WORKDIR /MultiStage

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o MultiStage 


#second stage 
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /MultiStage .

CMD ["./SingleStage"]