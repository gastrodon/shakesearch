FROM golang:alpine as builder

WORKDIR /app
COPY . .

RUN go get -u ./...
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/build .
COPY completeworks.txt .

ENTRYPOINT ./build
