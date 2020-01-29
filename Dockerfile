################################################################################
## Build
################################################################################
ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /user && echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor

WORKDIR /src

COPY go.mod go.sum ./
COPY *.go ./
COPY packages/ ./packages/
COPY resources/ ./resources/
COPY vendor/ ./vendor/

RUN go build -installsuffix 'static' -o /app .


################################################################################
## Final
################################################################################
FROM scratch AS final

LABEL maintainer="William Santos <w@wsantos.net>"

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /app /app

USER nobody:nobody

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["/app"]
