FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
WORKDIR /app 
COPY --from=build /app/main .
ENV AUTH_SERVICE_PORT=${AUTH_SERVICE_PORT}
ENV POSTGRES_URI=${POSTGRES_URI}
CMD [ "./main"]