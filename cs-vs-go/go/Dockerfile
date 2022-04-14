FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /basic-api-calls -ldflags "-s -w"

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /basic-api-calls /basic-api-calls

EXPOSE 8080
ENV GIN_MODE=release
USER $APP_USER

ENTRYPOINT ["/basic-api-calls"]