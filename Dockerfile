FROM golang:alpine
LABEL authors="yaucp"

FROM golang:alpine


# Install git and ca-certificates (needed to be able to call HTTPS)
RUN apk --update add ca-certificates git make
RUN apk --update add chromium
RUN which chromium


# Move to working directory /app
WORKDIR /app


# Copy the code into the container
COPY . .


# Download dependencies using go mod
RUN go mod download

# Build the application's binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main ./cmd/main.go

RUN chmod 777 /usr/bin/chromium

CMD ./main
EXPOSE $PORT