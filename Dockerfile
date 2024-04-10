FROM golang:bullseye

ARG VERSION="1.0.2"

COPY main.go /
COPY go.mod /
COPY go.sum /

RUN echo "building push server" \
 	&& cd / && go build main.go

COPY entry.sh /


ENTRYPOINT ["/entry.sh"]
RUN ["chmod", "+x", "/entry.sh"]
