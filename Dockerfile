FROM golang:alpine AS build

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w -extldflags '-static'" -o main ./cmd/app/main.go

FROM alpine

WORKDIR /build/app/

COPY --from=build /build/ .

RUN apk add --no-cache tzdata

ADD https://github.com/pressly/goose/releases/download/v3.5.3/goose_linux_x86_64 /bin/goose

RUN chmod +x /bin/goose

CMD /build/app/main

#goose -dir migrations postgres "user=root password=rpass dbname=ps_db host=localhost port=5432 sslmode=disable" up