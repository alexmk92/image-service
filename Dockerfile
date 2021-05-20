FROM golang:1.16 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# Now we've built our app, let's port to a lightweight container
# using alpine (the builder image is heavy, not great for prod)
FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]

