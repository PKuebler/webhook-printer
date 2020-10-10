# Start by building the application.
FROM golang:1.14 as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN CGO_ENABLED=0 go build -o /go/bin/webhook-printer

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /go/bin/webhook-printer /
CMD ["/webhook-printer"]