FROM golang:1.20.4 AS build
WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOPROXY=direct
RUN go build -v -o app .

# Install SQLite
FROM alpine:latest as db
RUN apk --no-cache add sqlite

FROM scratch
COPY --from=db ./example.db .
COPY --from=build /go/src/app/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]
