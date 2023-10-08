# Build the Go Binary.
FROM golang:1.21.1 as auth_service-api
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container.
COPY . /service

# Build the service binary.
WORKDIR /service/cmd/auth
RUN go build -ldflags "-X main.build=${BUILD_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.18
ARG BUILD_DATE
ARG BUILD_REF
RUN addgroup -g 1000 -S auth && \
    adduser -u 1000 -h /service -G auth -S auth
COPY --from=auth_service-api --chown=auth:auth /service/cmd/auth /service/.
WORKDIR /service
USER auth
CMD ["./auth"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="auth_service" \
      org.opencontainers.image.source="https://github.com/Feedbackify/auth_service" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Feedbackify"