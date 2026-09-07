# renovate: datasource=docker depName=europe-docker.pkg.dev/mailerlitehub/base-images/go
ARG BASE_IMAGE=europe-docker.pkg.dev/mailerlitehub/base-images/go:v1.0.10-1.26
FROM ${BASE_IMAGE} AS build

WORKDIR /src
COPY . /src

RUN CGO_ENABLED=0 go build -o cfpurge

# renovate: datasource=docker depName=europe-docker.pkg.dev/mailerlitehub/base-images/base
FROM europe-docker.pkg.dev/mailerlitehub/base-images/base:v1.0.44

WORKDIR /mlr
COPY --from=build /src/cfpurge /mlr/

USER 24128

ENTRYPOINT ["./cfpurge"]
