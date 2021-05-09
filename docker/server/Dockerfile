FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build ./product_server/cmd/product_server/server.go
CMD ["./server"]