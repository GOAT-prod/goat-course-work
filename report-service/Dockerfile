FROM golang:1.23 as build
WORKDIR /build
COPY . .

RUN mkdir out && \
    mkdir out/database && \
    mv database/migrations/ out/database/migrations/ && \
    mv .config/ out/

RUN apt-get update && \
    apt-get install -y build-essential pkg-config libssl-dev

RUN go build -mod vendor -o out/app

FROM alpine
EXPOSE 8080

RUN apk update && apk add --no-cache tzdata bash musl libc6-compat

WORKDIR /app

COPY --from=build /build/out ./
RUN chmod +x ./app

ENTRYPOINT ./app