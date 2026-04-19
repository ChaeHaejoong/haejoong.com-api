FROM golang:1.26.2-alpine AS base
WORKDIR /app
RUN apk add --no-cache tzdata
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM base AS dev
CMD ["air"]

FROM base AS builder
RUN go build -o server ./cmd/api

FROM alpine:latest AS prod
WORKDIR /app
RUN apk add --no-cache tzdata
COPY --from=builder /app/server .
EXPOSE 8080
CMD [ "./server" ]