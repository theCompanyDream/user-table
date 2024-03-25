# Specifies a parent image
FROM golang:1.22.0-bookworm

# Creates an app directory to hold your app’s source code
WORKDIR /opt/app

# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build -o userService

# Specifies the executable command that runs when the container starts
CMD [ “./userService” ]