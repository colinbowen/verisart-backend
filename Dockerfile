# Start from golang v1.8 base image
FROM golang:1.8
# Add Maintainer Info
LABEL maintainer="Colin Bowen <colinbowen2@gmail.com>"
# Set the Current Working Directory inside the container
#WORKDIR $GOPATH/app
WORKDIR /go/src/app
# Copy everything from the current directory to the PWD(Present Working Directory) 
# inside the container
COPY . .
# Download all the dependencies
RUN go get -d -v ./...
# Install the package
RUN go install -v ./...
# Exposes port 8080
EXPOSE 8000
# Run the executable
CMD ["go", "run", "main.go", "models.go"]

