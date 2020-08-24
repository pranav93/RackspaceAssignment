FROM golang:1.14
# All these steps will be cached
ENV GIN_MODE=release
ENV PORT=3000
RUN mkdir /app
WORKDIR /app
COPY . . 

RUN apt-get update
RUN apt-get install ca-certificates git

RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
EXPOSE 3000
ENTRYPOINT ["/go/bin/app"]