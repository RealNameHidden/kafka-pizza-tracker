FROM golang:1.24.5 as builder

WORKDIR /app
COPY . .

RUN go build -o . /app/app

FROM gcr.io/distroless/base
COPY --from=builder /app/order-service /app/order-service 
ENTRYPOINT ["/app/app "]
