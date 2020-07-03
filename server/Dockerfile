FROM golang:alpine AS builder
ENV GOOS=linux GOARCH=amd64
WORKDIR /blog
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build

FROM alpine
WORKDIR /app
EXPOSE 5000
COPY --from=builder /app/server .
ENTRYPOINT [ "/app/server" ]