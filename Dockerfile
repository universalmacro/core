FROM golang:1.21.6 as build-stage
WORKDIR /app
ARG GIT_TOKEN
ARG GIT_NAME
COPY go.mod go.sum ./
RUN go env -w GOPRIVATE=github.com/universalmacro/*
RUN git config --global url."https://${GIT_NAME}:${GIT_TOKEN}@github.com".insteadOf "https://github.com"
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /docker-gs-ping /docker-gs-ping
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/docker-gs-ping"]