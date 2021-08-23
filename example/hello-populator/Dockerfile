FROM gcr.io/distroless/static:latest
LABEL maintainers="Kubernetes Authors"
LABEL description="hello populator"
ARG binary=./bin/hello-populator

COPY ${binary} hello-populator
ENTRYPOINT ["/hello-populator"]
