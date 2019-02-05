FROM golang:1.11.0-stretch 

COPY . $GOPATH/src/MultiStageDocker/SingleStage
WORKDIR $GOPATH/src/MultiStageDocker/SingleStage

ENV GO111MODULE=on

RUN ls

RUN CGO_ENABLED=0 GOOS=linux go build -o SingleStage 

CMD ["./SingleStage"]
