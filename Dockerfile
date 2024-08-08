# This file is licensed under the Creative Commons Attribution-NonCommercial 4.0 International License.
# You may obtain a copy of the license at https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt

FROM golang:1.22.5 AS build

WORKDIR /app

COPY go.mod go.sum ./

COPY api/ ./api/
COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bff ./cmd/recotis-landing-page-bff

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/bff .

RUN chmod +x bff

EXPOSE 8080

ENTRYPOINT ["./bff"]
