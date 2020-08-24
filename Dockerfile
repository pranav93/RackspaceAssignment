FROM golang:1.14
# All these steps will be cached
ENV GO111MODULE=on
RUN chmod o+r /etc/resolv.conf
RUN cat /etc/resolv.conf
RUN mkdir /app
WORKDIR /app
COPY . . 
# <- COPY go.mod and go.sum files to the workspace
# COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN apt-get update
RUN apt-get install ca-certificates git
RUN ls

RUN go mod download
# COPY the source code as the last step
# COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
# FROM scratch
# <- Second step to build minimal image
# COPY --from=build-env /go/bin/app /go/bin/app
EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]