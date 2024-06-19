# Build stage
FROM golang:1.19 as builder
WORKDIR /app
COPY . /app
RUN go mod download && go mod verify
RUN cd /app && go build -o goapp

# Final stage
# Use the compiled binary from the first
FROM alpine
WORKDIR /app
COPY --from=builder /app/goapp /app/
ENTRYPOINT ./goapp
