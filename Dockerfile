# build stage
FROM golang:1.21.1-alpine as build

ARG CI_JOB_TOKEN
RUN apk add alpine-sdk

RUN mkdir /app
WORKDIR /app
COPY ./ .

RUN go clean --modcache
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates
RUN apk add --no-cache openssh


RUN GOOS=linux GOARCH=amd64 go build -a -tags musl -a -installsuffix cgo -o main cmd/server/main.go

FROM golang:alpine
COPY --from=build /app/main /

EXPOSE 80
ENTRYPOINT ["/main"]