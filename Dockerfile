FROM scratch as base
EXPOSE 8090

FROM golang:1.13 as build
WORKDIR $GOPATH/src/spyglass
COPY go.mod .
RUN go get -d -v ./...
RUN go install -v ./...
COPY spyglass.go .
RUN go build -o /build/spyglass

FROM base as final
WORKDIR /app
COPY --from=build /build/spyglass .

ENTRYPOINT ./spyglass
