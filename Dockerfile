FROM golang:1.17-alpine as build

WORKDIR /app

COPY . .

RUN go build -o semver main.go

FROM alpine:3.15 as dist

COPY --from=build /app/semver /usr/local/bin/semver

CMD [ "semver" ]

