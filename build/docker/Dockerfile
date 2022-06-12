ARG GO_VER
ARG ALPINE_VER
# FROM $DOCKER_REGISTRY/golang:${GO_VER}-alpine${ALPINE_VER} as builder
# FROM $DOCKER_REGISTRY/golang:1.16.5-alpine${ALPINE_VER} as builder
FROM golang:1.16.5-alpine${ALPINE_VER} as builder

# ARG GO_MAIN_PATH
ARG VERSION
WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -installsuffix cgo -o app -ldflags "-X 'main.version=${VERSION}'" ./app
# FROM $DOCKER_REGISTRY/alpine:${ALPINE_VER}
FROM alpine:${ALPINE_VER}

WORKDIR /root/
RUN touch checkfile
COPY --from=builder /src/app .
ADD /configs /root/configs
CMD ["./app"]
