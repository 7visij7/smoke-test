FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY post-deployment-smoke-tests/ /app
RUN go build -o smoke-tests

############################
FROM scratch
COPY --from=builder /app/smoke-tests /app/smoke-tests
ENTRYPOINT ["/app/smoke-tests"]