FROM golang:1.19-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary internal/cmd/app/*.go

EXPOSE 8080

ENTRYPOINT ["/app/binary"]