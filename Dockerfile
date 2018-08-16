FROM golang:1.10-alpine AS gobuild-base
RUN apk add --no-cache \
	git \
	make

FROM gobuild-base AS base
WORKDIR /go/src/github.com/ehotinger/scratch
COPY . .
RUN make static && mv scratch /usr/bin/scratch

FROM ubuntu:bionic
RUN groupadd -r -g 127001 container && \
    useradd -r -u 127001 -g container container
COPY --from=base /usr/bin/scratch scratch
ADD --chown=container:container scratch /usr/bin/scratch
USER container
ENTRYPOINT [ "scratch" ]
CMD [ ]