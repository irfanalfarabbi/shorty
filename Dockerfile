FROM alpine:3.6

# Copy binary
COPY shorty /usr/local/bin

# Expose your service port
EXPOSE 8080

# Set your entrypoint here just to start the webserver
ENTRYPOINT ["shorty"]
