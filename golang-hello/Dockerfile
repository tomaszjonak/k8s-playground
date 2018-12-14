FROM golang:1.11.3-alpine3.8 AS build
WORKDIR /go/src
COPY main.go .
RUN CGO_ENABLED=0 go build -installsuffix cgo -a -o app .

FROM scratch
COPY --from=build /go/src/app .
ENTRYPOINT [ "./app" ]
