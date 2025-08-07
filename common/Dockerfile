FROM golang:1.2.4.5 as builder

WORKDIR /app
COPY . .

RUN go build -o common ./common

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/common /app/app
ENTRYPOINT ["/app/app"]
