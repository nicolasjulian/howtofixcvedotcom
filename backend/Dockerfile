FROM golang:1.21-alpine AS build

WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /out/backend_howtofixcvedotcom .

FROM alpine:3
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true
RUN mkdir /apps
RUN mkdir /apps/configs

COPY --from=build /out/backend_howtofixcvedotcom /apps/

WORKDIR /apps

ENTRYPOINT ["./backend_howtofixcvedotcom"]
