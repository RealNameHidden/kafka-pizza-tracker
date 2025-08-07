FROM golang:1.24.5 as builder

WORKDIR /app
COPY . .

RUN go build -o /app/order-service ./order-service

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/order-service /app/order-service 
ENTRYPOINT ["/app/order-service "]
