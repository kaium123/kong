ARG GO_VERSION=1.25

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /src
COPY ./ ./


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix 'static' -o /api_gaurd/app main.go

FROM alpine AS final

WORKDIR /api_gaurd

# copy the binary into /app/app
COPY --from=builder /api_gaurd/app /api_gaurd/app

# config.yml will be mounted into /api_gaurd/config.yml
ENTRYPOINT ["/api_gaurd/app"]
