FROM debian:jessie
MAINTAINER Hofstadter, Inc <support@hofstadter.io>

COPY hof /
ENTRYPOINT ["/hof"]

VOLUME ["/app"]
WORKDIR /app
