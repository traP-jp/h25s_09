FROM golang:1.24-alpine AS base
EXPOSE 8080
RUN go install github.com/air-verse/air@latest
WORKDIR /src
COPY . .

FROM base AS build
RUN go mod download
ENTRYPOINT ["air", "-c", ".air.toml"]
