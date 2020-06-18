FROM golang

RUN mkdir /opt/app

WORKDIR /opt/app

# Copy the local package files to the container's workspace.
ADD go-calc /opt/app

RUN go get github.com/gorilla/mux

RUN go build .
EXPOSE 8080
ENTRYPOINT  go run main.go
