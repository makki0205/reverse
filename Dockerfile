FROM golang:1.11-alpine
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /src
RUN apk add --no-cache --virtual git
ADD ./ ./

RUN echo $GOPATH
RUN ls -la
RUN go build main.go
RUN rm -rf /go 
RUN apk del --purge git
RUN rm -rf /usr/local/

EXPOSE 3000

ENTRYPOINT ["./main"]
