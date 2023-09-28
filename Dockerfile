FROM golang:1.20 as BUILD

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /main

FROM gcr.io/distroless/base

WORKDIR /

COPY --from=build /main /main

COPY . /

EXPOSE 8080

USER root

ENTRYPOINT ["/main"]