FROM golang:1.22-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/pcs-api

FROM alpine:latest


WORKDIR /app
COPY --from=build ./app/config.env .
COPY --from=build /app/pcs-api .

CMD ["./pcs-api"]