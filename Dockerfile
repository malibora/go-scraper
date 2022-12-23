FROM golang:1.19 as builder

# first (build) stage

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o go-scraper

# final (target) stage

FROM alpine:latest
COPY --from=builder /app/go-scraper /
CMD ["/go-scraper"]