FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o goethe

RUN CGO_ENABLED=0 GOOS=linux go build -o goethe

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/goethe .

COPY /public ./public

EXPOSE 8081

CMD ["./goethe"]
