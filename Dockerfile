FROM golang:1.21-alpine3.17 AS build
WORKDIR /
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /main
# RUN ls -ltra /
# EXPOSE 8080
# ENTRYPOINT ["/main"]

FROM alpine:3.18
COPY --from=build /main /main
ENTRYPOINT ["/main"]