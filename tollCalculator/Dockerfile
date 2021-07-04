FROM golang:alpine

ARG PORT=4000
ENV SERVER_PORT=$PORT
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE $SERVER_PORT

CMD ["/dist/main"]