FROM golang:1.24 as build
WORKDIR /build
COPY . .

RUN mkdir out && \
    mv .config/ out/

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o out/app

FROM alpine
EXPOSE 8080

RUN apk update && apk add --no-cache tzdata

WORKDIR /app

COPY --from=build /build/out ./
RUN chmod +x ./app

ENTRYPOINT ./app