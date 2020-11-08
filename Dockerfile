FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/todo-now/todo-now
COPY . .

#fetch dependencies

#using go get 
RUN go get -d -v

#Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" -o /go/bin/todo-now .

#####
# Build minimal img with the previous build
#####
FROM scratch

COPY --from=builder /go/bin/todo-now /usr/bin/todo-now

ENTRYPOINT [ "/usr/bin/todo-now" ]
