FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o server ./cmd
CMD ["./server"]
