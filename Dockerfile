FROM golang:1.17-bullseye as build
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./...
RUN go build -o /go/bin/server ./cmd

FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/server /
CMD ["/server"]