
FROM golang:1.11-alpine

ENV GO111MODULE=on

WORKDIR /src

ADD ./ ./

RUN ls -la
RUN go build main.go

EXPOSE 3000

ENTRYPOINT ["./main"]