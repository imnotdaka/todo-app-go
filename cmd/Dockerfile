
FROM golang:1.21.6 AS builder

WORKDIR /build
COPY go.mod /build
COPY go.sum /build
RUN go mod download
COPY ./cmd /build/cmd
COPY ./internal /build/internal
RUN go build -o todo-app ./cmd/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/todo-app /app/todo-app
EXPOSE 8080
CMD ["./todo-app"]