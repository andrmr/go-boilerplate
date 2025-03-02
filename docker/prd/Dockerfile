FROM golang:1.24-bookworm AS build

ARG BUILD_OUT="/out/"

WORKDIR /code
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=.,target=. \
    make deps && make build BUILD_OUT=${BUILD_OUT}

FROM gcr.io/distroless/static-debian12:nonroot AS runtime

WORKDIR /app
COPY --from=build ${BUILD_OUT} .

USER nonroot:nonroot
CMD ["/golnib"]