FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build main.go

FROM alpine
COPY --from=builder /app/main /app
CMD ["/app"]
