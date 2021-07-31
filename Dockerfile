# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.16-alpine

# Instruct Docker to use this directory as the default destination for all subsequent commands
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Move Go source files
COPY *.go ./

# Compile Go package and dependencies and keep static application binary output
RUN go build -o /go-margarita

# Expose docker port
EXPOSE 8080

# Tell Docker what command to execute when our image is used to start a container
CMD [ "/go-margarita" ]