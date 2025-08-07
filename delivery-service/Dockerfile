FROM golang:1.24.5 as builder

WORKDIR /app
COPY . .

RUN go build -o delivery-service ./delivery-service

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/delivery-service /app/app
ENTRYPOINT ["/app/app"]
