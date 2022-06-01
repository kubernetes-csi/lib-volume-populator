FROM alpine
LABEL maintainers="Kubernetes Authors"
LABEL description="csi hostpath transfer"
ARG binary=./bin/csi-hostpath-transfer

COPY ${binary} csi-hostpath-transfer
ENTRYPOINT ["/csi-hostpath-transfer"]
