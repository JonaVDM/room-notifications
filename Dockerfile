FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build main.go

FROM alpine
COPY --from=builder /app/main /etc/periodic/daily/app
CMD ["crond", "-f"]
