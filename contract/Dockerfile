FROM golang:1.23
LABEL authors="madosov"

WORKDIR /contract

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 go build -o service-contract cmd/*

ENTRYPOINT ["./service-contract"]
