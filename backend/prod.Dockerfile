FROM golang:1.24-alpine AS base
EXPOSE 8080
WORKDIR /src
COPY . .

FROM base AS build
WORKDIR /src
RUN ["go", "mod", "tidy"]
RUN ["go", "build", "-o", "./bin/app"]

FROM alpine:latest
WORKDIR /
COPY --from=build /src/bin/app .
ENTRYPOINT ["./app"]
