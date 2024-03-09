FROM golang:1.22 AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o blog-api .

FROM scratch

COPY --from=build /build/blog-api .
