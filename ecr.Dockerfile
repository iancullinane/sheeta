FROM public.ecr.aws/lambda/go:1

# Copy the binary into the container
ADD bin/main /main

# The service listens on these ports
EXPOSE 8080

# On run, execute the binary
ENTRYPOINT ["/main"]
