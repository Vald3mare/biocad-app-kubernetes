FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o biocad-app .

FROM scratch

COPY --from=builder /app/biocad-app .

EXPOSE 32777

CMD ["/biocad-app"]
