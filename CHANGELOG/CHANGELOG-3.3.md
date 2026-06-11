## Changes by Kind

### Feature

- Update Kubernetes dependencies to v1.36.0 ([#257](https://github.com/kubernetes-csi/lib-volume-populator/pull/257), [@sunnylovestiramisu](https://github.com/sunnylovestiramisu))

## Dependencies

### Added
- github.com/cenkalti/backoff/v5: [v5.0.3](https://github.com/cenkalti/backoff/v5/tree/v5.0.3)
- k8s.io/streaming: v0.36.1

### Changed
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.26.3 → v2.27.7](https://github.com/grpc-ecosystem/grpc-gateway/v2/compare/v2.26.3...v2.27.7)
- github.com/mailru/easyjson: [v0.9.0 → v0.9.1](https://github.com/mailru/easyjson/compare/v0.9.0...v0.9.1)
- github.com/moby/spdystream: [v0.5.0 → v0.5.1](https://github.com/moby/spdystream/compare/v0.5.0...v0.5.1)
- github.com/onsi/ginkgo/v2: [v2.27.2 → v2.28.0](https://github.com/onsi/ginkgo/v2/compare/v2.27.2...v2.28.0)
- github.com/onsi/gomega: [v1.38.2 → v1.39.1](https://github.com/onsi/gomega/compare/v1.38.2...v1.39.1)
- github.com/spf13/cobra: [v1.10.0 → v1.10.2](https://github.com/spf13/cobra/compare/v1.10.0...v1.10.2)
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.61.0 → v0.65.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.34.0 → v1.40.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.34.0 → v1.40.0
- go.opentelemetry.io/otel/metric: v1.39.0 → v1.41.0
- go.opentelemetry.io/otel/sdk: v1.36.0 → v1.40.0
- go.opentelemetry.io/otel/trace: v1.39.0 → v1.41.0
- go.opentelemetry.io/otel: v1.39.0 → v1.41.0
- go.opentelemetry.io/proto/otlp: v1.5.0 → v1.9.0
- go.uber.org/zap: v1.27.0 → v1.27.1
- golang.org/x/tools/go/expect: v0.1.1-deprecated → v0.1.0-deprecated
- google.golang.org/genproto/googleapis/api: a0af3ef → 8636f87
- google.golang.org/genproto/googleapis/rpc: ef028d9 → 8636f87
- google.golang.org/grpc: v1.75.1 → v1.79.3
- google.golang.org/protobuf: v1.36.11 → f2248ac
- k8s.io/api: v0.35.0 → v0.36.0
- k8s.io/apimachinery: v0.35.0 → v0.36.0
- k8s.io/client-go: v0.35.0 → v0.36.0
- k8s.io/component-base: v0.35.0 → v0.36.0
- k8s.io/component-helpers: v0.35.0 → v0.36.0
- k8s.io/gengo/v2: c297c0c → 85fd79d
- k8s.io/klog/v2: v2.130.1 → v2.140.0
- k8s.io/kube-openapi: 4e65d59 → 43fb72c
- k8s.io/utils: 914a6e7 → b8788ab
- sigs.k8s.io/gateway-api: v1.4.1 → v1.5.1
- sigs.k8s.io/structured-merge-diff/v6: v6.3.1 → v6.3.2

### Removed
- github.com/Masterminds/goutils: [v1.1.1](https://github.com/Masterminds/goutils/tree/v1.1.1)
- github.com/Masterminds/semver/v3: [v3.4.0](https://github.com/Masterminds/semver/v3/tree/v3.4.0)
- github.com/Masterminds/semver: [v1.5.0](https://github.com/Masterminds/semver/tree/v1.5.0)
- github.com/Masterminds/sprig: [v2.22.0+incompatible](https://github.com/Masterminds/sprig/tree/v2.22.0)
- github.com/armon/go-socks5: [e753329](https://github.com/armon/go-socks5/tree/e753329)
- github.com/cenkalti/backoff/v4: [v4.3.0](https://github.com/cenkalti/backoff/v4/tree/v4.3.0)
- github.com/elastic/crd-ref-docs: [v0.2.0](https://github.com/elastic/crd-ref-docs/tree/v0.2.0)
- github.com/evanphx/json-patch/v5: [v5.9.11](https://github.com/evanphx/json-patch/v5/tree/v5.9.11)
- github.com/fatih/color: [v1.18.0](https://github.com/fatih/color/tree/v1.18.0)
- github.com/go-task/slim-sprig/v3: [v3.0.0](https://github.com/go-task/slim-sprig/v3/tree/v3.0.0)
- github.com/gobuffalo/flect: [v1.0.3](https://github.com/gobuffalo/flect/tree/v1.0.3)
- github.com/goccy/go-yaml: [v1.18.0](https://github.com/goccy/go-yaml/tree/v1.18.0)
- github.com/gogo/protobuf: [v1.3.2](https://github.com/gogo/protobuf/tree/v1.3.2)
- github.com/google/pprof: [27863c8](https://github.com/google/pprof/tree/27863c8)
- github.com/gregjones/httpcache: [901d907](https://github.com/gregjones/httpcache/tree/901d907)
- github.com/huandu/xstrings: [v1.3.3](https://github.com/huandu/xstrings/tree/v1.3.3)
- github.com/imdario/mergo: [v0.3.11](https://github.com/imdario/mergo/tree/v0.3.11)
- github.com/mattn/go-colorable: [v0.1.13](https://github.com/mattn/go-colorable/tree/v0.1.13)
- github.com/mattn/go-isatty: [v0.0.20](https://github.com/mattn/go-isatty/tree/v0.0.20)
- github.com/miekg/dns: [v1.1.68](https://github.com/miekg/dns/tree/v1.1.68)
- github.com/mitchellh/copystructure: [v1.2.0](https://github.com/mitchellh/copystructure/tree/v1.2.0)
- github.com/mitchellh/reflectwalk: [v1.0.2](https://github.com/mitchellh/reflectwalk/tree/v1.0.2)
- google.golang.org/grpc/cmd/protoc-gen-go-grpc: v1.5.1
- gopkg.in/yaml.v2: v2.4.0
- k8s.io/apiextensions-apiserver: v0.34.1
- k8s.io/code-generator: v0.34.1
- sigs.k8s.io/controller-runtime: v0.22.1
- sigs.k8s.io/controller-tools: v0.19.0
