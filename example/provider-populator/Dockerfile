FROM gcr.io/distroless/static:latest
LABEL maintainers="Kubernetes Authors"
LABEL description="provider populator"
ARG binary=./provider-populator

COPY ${binary} provider-populator
ENTRYPOINT ["/provider-populator"]