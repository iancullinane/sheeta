# Docker container with nothing
FROM scratch

# Copy the binary into the container
ADD bin/sheeta /sheeta

# The service listens on these ports
EXPOSE 8080

# On run, execute the binary
ENTRYPOINT ["/sheeta"]
