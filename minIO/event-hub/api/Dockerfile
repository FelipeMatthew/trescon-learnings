# build environment
FROM golang:1.22 AS build

WORKDIR /go/src/ms

ADD . /go/src/ms

RUN go mod tidy
RUN go build -o eventhub

# production environment
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=build /go/src/ms/eventhub .

COPY .env .env

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["./eventhub"]