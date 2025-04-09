FROM golang:alpine

RUN mkdir -p /app
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /app/main .
EXPOSE 8080
CMD ["/app/main"]