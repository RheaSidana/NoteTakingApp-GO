# Use a base image with Go installed
FROM golang:1.18-bullseye

# Set the working directory inside the container
WORKDIR /app

# Add everything to the working directory
COPY go.mod go.sum ./

# Install all the dependencies inside the container 
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

# Build the application inside the container 
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/build/noteApp .

# Expose a port 
EXPOSE 3033

# Define the command to run your application
CMD ["/app/build/noteApp"]

# To build the dockerfile: "docker build -t test-go-docker:latest . "