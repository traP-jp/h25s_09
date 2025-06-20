FROM golang:1.24-alpine AS base
EXPOSE 8080
WORKDIR /src
COPY . .

FROM base AS build
WORKDIR /src
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o ./bin/app

FROM alpine:latest
WORKDIR /
COPY --from=build /src/bin/app .
ENTRYPOINT ["./app"]
