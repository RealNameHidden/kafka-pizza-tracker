FROM golang:1.24.5 as builder

WORKDIR /app
COPY . .

RUN go build -o kitchen-service ./kitchen-service

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/kitchen-service /app/app
ENTRYPOINT ["/app/app"]
