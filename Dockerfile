FROM golang:1.12.5-alpine AS gobuild-base
RUN apk add --no-cache \
	git \
	make

FROM gobuild-base AS base
WORKDIR /go/src/github.com/ehotinger/scratch
COPY . .
RUN make static && mv scratch /usr/bin/scratch

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=base /usr/bin/scratch /usr/bin/scratch
ENTRYPOINT [ "scratch" ]