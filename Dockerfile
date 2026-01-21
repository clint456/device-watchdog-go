ARG BASE=golang:1.25-alpine3.22
FROM ${BASE} AS builder

ARG ADD_BUILD_TAGS=""
ARG MAKE="make -e ADD_BUILD_TAGS=$ADD_BUILD_TAGS build"
ENV GOPROXY=https://goproxy.cn,direct
RUN apk add --update --no-cache make git openssh

# set the working directory
WORKDIR /device-demo-go

COPY go.mod vendor* ./
RUN [ ! -d "vendor" ] && go mod download all || echo "skipping..."

COPY . .
RUN ${MAKE}

FROM alpine:3.22


RUN apk add --update --no-cache dumb-init
# Ensure using latest versions of all installed packages to avoid any recent CVEs
RUN apk --no-cache upgrade

COPY --from=builder /device-demo-go/cmd/device-demo /
COPY --from=builder /device-demo-go/cmd/res /res

RUN chmod -R 755 /res

EXPOSE 59911

ENTRYPOINT ["/device-demo"]
CMD ["-cp=keeper.http://edgex-core-keeper:59890", "--registry"]