FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -o syncservice .
CMD ["./syncservice"]
