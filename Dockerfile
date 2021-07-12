FROM golang:1.16-alpine as base

WORKDIR /src

COPY . .

RUN go build -o cfpurge

FROM gcr.io/distroless/base

COPY --from=base /src/cfpurge /

CMD ["/cfpurge"]
