FROM golang:1.23-alpine as base
WORKDIR /src
COPY . /src
RUN CGO_ENABLED=0 go build -o cfpurge

FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=base /src/cfpurge /app
ENTRYPOINT ["/app/cfpurge"]
