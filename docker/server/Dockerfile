FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
RUN go build -o ./product_server/cmd/product_server product_server/cmd/product_server/server.go
CMD [".product_server/cmd/product_server"]