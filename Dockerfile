# Use the official Golang image with version 1.22 on Alpine Linux for the build stage
FROM golang:1.22-alpine AS buildstage

# Set the working directory inside the container to /LaptopLounge
WORKDIR /LaptopLounge

# Copy all files from the current directory on the host machine to /LaptopLounge in the container
COPY . /LaptopLounge/

# Download Go module dependencies specified in the go.mod file
RUN go mod download

# Install ca-certificates package for enabling HTTPS support
RUN apk --no-cache add ca-certificates

# Build the Go application and output the binary executable
RUN go build -o binary ./cmd/main.go

# Use a minimal Alpine Linux image as the base for the final stage
FROM alpine:3.18

# Set the working directory inside the final container to /LaptopLounge
WORKDIR /LaptopLounge

# Copy the compiled binary from the build stage to the final container
COPY --from=buildstage /LaptopLounge/binary .

# Copy the .env file from the build stage to the final container
COPY --from=buildstage /LaptopLounge/.env .env

# Copy the SSL certificates from the build stage to the final container
COPY --from=buildstage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the template directory from the build stage to the final container
COPY --from=buildstage /LaptopLounge/template /LaptopLounge/template/

# Set the command to run the compiled binary when the container starts
CMD ["./binary"]
