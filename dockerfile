##
## Build
##

FROM golang:1.16-buster AS build

ENV APPDIR /usr/src/app
ENV GO111MODULE="on" \
    GOOS=linux

RUN mkdir -p ${APPDIR}
WORKDIR ${APPDIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -a -o /build

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /build /build

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/build"]
