FROM node:20.10-alpine as payments
WORKDIR /app
COPY ./payments/ .
CMD ["npm", "run", "dev"]

FROM node:20.10-alpine as shippings
WORKDIR /app
COPY shippings/ .
CMD ["npm", "run", "dev"]

FROM golang:1.21-alpine as gateway
WORKDIR /gateway
COPY ./gateway/ .
RUN apk add git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /gt cmd/main.go
CMD [ "/gt", "config.yml" ]