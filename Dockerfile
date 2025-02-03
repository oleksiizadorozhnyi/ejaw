FROM golang:1.19
WORKDIR /app
COPY . .
RUN go mod init shop && go mod tidy && go build -o server
CMD ["./server"]