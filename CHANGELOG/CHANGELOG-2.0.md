# Release notes for v2.0.0

# Changelog since v1.2.0

## Changes by Kind

### Feature

- Action Needed: Emit events when a populator pod starts / finishes or on Pod/PVC creation error. The populator now needs RBACs to `create` `events`! ([#57](https://github.com/kubernetes-csi/lib-volume-populator/pull/57), [@jsafrane](https://github.com/jsafrane))

### Bug or Regression

- Populating of in-tree volumes has been disabled. It was an oversight, we want to limit volume populators only to CSI. ([#52](https://github.com/kubernetes-csi/lib-volume-populator/pull/52), [@jsafrane](https://github.com/jsafrane))

## Dependencies

### Added
- cloud.google.com/go/compute/metadata: v0.2.0
- github.com/OneOfOne/xxhash: [v1.2.2](https://github.com/OneOfOne/xxhash/tree/v1.2.2)
- github.com/antihax/optional: [v1.0.0](https://github.com/antihax/optional/tree/v1.0.0)
- github.com/buger/jsonparser: [v1.1.1](https://github.com/buger/jsonparser/tree/v1.1.1)
- github.com/cenkalti/backoff/v4: [v4.1.3](https://github.com/cenkalti/backoff/v4/tree/v4.1.3)
- github.com/cespare/xxhash: [v1.1.0](https://github.com/cespare/xxhash/tree/v1.1.0)
- github.com/cncf/xds/go: [fbca930](https://github.com/cncf/xds/go/tree/fbca930)
- github.com/flowstack/go-jsonschema: [v0.1.1](https://github.com/flowstack/go-jsonschema/tree/v0.1.1)
- github.com/go-logr/stdr: [v1.2.2](https://github.com/go-logr/stdr/tree/v1.2.2)
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.7.0](https://github.com/grpc-ecosystem/grpc-gateway/v2/tree/v2.7.0)
- github.com/rogpeppe/fastuuid: [v1.2.0](https://github.com/rogpeppe/fastuuid/tree/v1.2.0)
- github.com/spaolacci/murmur3: [f09979e](https://github.com/spaolacci/murmur3/tree/f09979e)
- github.com/xeipuuv/gojsonpointer: [4e3ac27](https://github.com/xeipuuv/gojsonpointer/tree/4e3ac27)
- github.com/xeipuuv/gojsonreference: [bd5ef7b](https://github.com/xeipuuv/gojsonreference/tree/bd5ef7b)
- github.com/xeipuuv/gojsonschema: [v1.2.0](https://github.com/xeipuuv/gojsonschema/tree/v1.2.0)
- go.opentelemetry.io/otel/exporters/otlp/internal/retry: v1.10.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.10.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.10.0
- go.uber.org/goleak: v1.2.0
- k8s.io/component-helpers: v0.26.0

### Changed
- cloud.google.com/go: v0.97.0 → v0.34.0
- github.com/cncf/udpa/go: [269d4d4 → 5459f2c](https://github.com/cncf/udpa/go/compare/269d4d4...5459f2c)
- github.com/emicklei/go-restful/v3: [v3.8.0 → v3.9.0](https://github.com/emicklei/go-restful/v3/compare/v3.8.0...v3.9.0)
- github.com/envoyproxy/go-control-plane: [v0.9.4 → 63b5d3c](https://github.com/envoyproxy/go-control-plane/compare/v0.9.4...63b5d3c)
- github.com/felixge/httpsnoop: [v1.0.1 → v1.0.3](https://github.com/felixge/httpsnoop/compare/v1.0.1...v1.0.3)
- github.com/go-kit/log: [v0.1.0 → v0.2.1](https://github.com/go-kit/log/compare/v0.1.0...v0.2.1)
- github.com/go-logfmt/logfmt: [v0.5.0 → v0.5.1](https://github.com/go-logfmt/logfmt/compare/v0.5.0...v0.5.1)
- github.com/go-openapi/jsonreference: [v0.19.5 → v0.20.0](https://github.com/go-openapi/jsonreference/compare/v0.19.5...v0.20.0)
- github.com/go-openapi/swag: [v0.19.14 → v0.22.3](https://github.com/go-openapi/swag/compare/v0.19.14...v0.22.3)
- github.com/golang/mock: [v1.4.4 → v1.1.1](https://github.com/golang/mock/compare/v1.4.4...v1.1.1)
- github.com/google/gnostic: [v0.5.7-v3refs → v0.6.9](https://github.com/google/gnostic/compare/v0.5.7-v3refs...v0.6.9)
- github.com/google/go-cmp: [v0.5.6 → v0.5.9](https://github.com/google/go-cmp/compare/v0.5.6...v0.5.9)
- github.com/google/gofuzz: [v1.1.0 → v1.2.0](https://github.com/google/gofuzz/compare/v1.1.0...v1.2.0)
- github.com/imdario/mergo: [v0.3.6 → v0.3.13](https://github.com/imdario/mergo/compare/v0.3.6...v0.3.13)
- github.com/inconshreveable/mousetrap: [v1.0.0 → v1.0.1](https://github.com/inconshreveable/mousetrap/compare/v1.0.0...v1.0.1)
- github.com/mailru/easyjson: [v0.7.6 → v0.7.7](https://github.com/mailru/easyjson/compare/v0.7.6...v0.7.7)
- github.com/matttproud/golang_protobuf_extensions: [c182aff → v1.0.4](https://github.com/matttproud/golang_protobuf_extensions/compare/c182aff...v1.0.4)
- github.com/moby/term: [3f7ff69 → 39b0c02](https://github.com/moby/term/compare/3f7ff69...39b0c02)
- github.com/onsi/ginkgo/v2: [v2.1.4 → v2.4.0](https://github.com/onsi/ginkgo/v2/compare/v2.1.4...v2.4.0)
- github.com/onsi/gomega: [v1.19.0 → v1.23.0](https://github.com/onsi/gomega/compare/v1.19.0...v1.23.0)
- github.com/prometheus/client_golang: [v1.12.1 → v1.14.0](https://github.com/prometheus/client_golang/compare/v1.12.1...v1.14.0)
- github.com/prometheus/client_model: [v0.2.0 → v0.3.0](https://github.com/prometheus/client_model/compare/v0.2.0...v0.3.0)
- github.com/prometheus/common: [v0.32.1 → v0.39.0](https://github.com/prometheus/common/compare/v0.32.1...v0.39.0)
- github.com/prometheus/procfs: [v0.7.3 → v0.8.0](https://github.com/prometheus/procfs/compare/v0.7.3...v0.8.0)
- github.com/spf13/cobra: [v1.4.0 → v1.6.0](https://github.com/spf13/cobra/compare/v1.4.0...v1.6.0)
- github.com/stretchr/objx: [v0.1.1 → v0.1.0](https://github.com/stretchr/objx/compare/v0.1.1...v0.1.0)
- github.com/stretchr/testify: [v1.7.0 → v1.8.0](https://github.com/stretchr/testify/compare/v1.7.0...v1.8.0)
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.20.0 → v0.35.0
- go.opentelemetry.io/otel/metric: v0.20.0 → v0.31.0
- go.opentelemetry.io/otel/sdk: v0.20.0 → v1.10.0
- go.opentelemetry.io/otel/trace: v0.20.0 → v1.10.0
- go.opentelemetry.io/otel: v0.20.0 → v1.10.0
- go.opentelemetry.io/proto/otlp: v0.7.0 → v0.19.0
- golang.org/x/crypto: 3147a52 → 75b2880
- golang.org/x/exp: 6cc2880 → 509febe
- golang.org/x/lint: 6edffad → d0100b6
- golang.org/x/mod: 9b9b3d8 → 86c51ed
- golang.org/x/net: a158d28 → v0.4.0
- golang.org/x/oauth2: d3ed0bb → v0.3.0
- golang.org/x/sync: 09787c9 → 0de741c
- golang.org/x/sys: 8c9f86f → v0.3.0
- golang.org/x/term: 03fcf44 → v0.3.0
- golang.org/x/text: v0.3.7 → v0.5.0
- golang.org/x/time: 90d013b → v0.1.0
- google.golang.org/grpc: v1.47.0 → v1.49.0
- google.golang.org/protobuf: v1.28.0 → v1.28.1
- gopkg.in/check.v1: 8fa4692 → 10cb982
- honnef.co/go/tools: v0.0.1-2020.1.4 → ea95bdf
- k8s.io/api: v0.25.0 → v0.26.0
- k8s.io/apimachinery: v0.25.0 → v0.26.0
- k8s.io/client-go: v0.25.0 → v0.26.0
- k8s.io/component-base: v0.25.0 → v0.26.0
- k8s.io/klog/v2: v2.70.1 → v2.80.1
- k8s.io/kube-openapi: 67bda5d → 172d655
- k8s.io/utils: ee6ede2 → 1a15be2
- sigs.k8s.io/yaml: v1.2.0 → v1.3.0

### Removed
- cloud.google.com/go/bigquery: v1.8.0
- cloud.google.com/go/datastore: v1.1.0
- cloud.google.com/go/pubsub: v1.3.1
- cloud.google.com/go/storage: v1.10.0
- dmitri.shuralyov.com/gpu/mtl: 666a987
- github.com/Azure/go-autorest/autorest/adal: [v0.9.20](https://github.com/Azure/go-autorest/autorest/adal/tree/v0.9.20)
- github.com/Azure/go-autorest/autorest/date: [v0.3.0](https://github.com/Azure/go-autorest/autorest/date/tree/v0.3.0)
- github.com/Azure/go-autorest/autorest: [v0.11.27](https://github.com/Azure/go-autorest/autorest/tree/v0.11.27)
- github.com/Azure/go-autorest/logger: [v0.2.1](https://github.com/Azure/go-autorest/logger/tree/v0.2.1)
- github.com/Azure/go-autorest/tracing: [v0.6.0](https://github.com/Azure/go-autorest/tracing/tree/v0.6.0)
- github.com/Azure/go-autorest: [v14.2.0+incompatible](https://github.com/Azure/go-autorest/tree/v14.2.0)
- github.com/BurntSushi/xgb: [27f1227](https://github.com/BurntSushi/xgb/tree/27f1227)
- github.com/chzyer/logex: [v1.1.10](https://github.com/chzyer/logex/tree/v1.1.10)
- github.com/chzyer/readline: [2972be2](https://github.com/chzyer/readline/tree/2972be2)
- github.com/chzyer/test: [a1ea475](https://github.com/chzyer/test/tree/a1ea475)
- github.com/creack/pty: [v1.1.9](https://github.com/creack/pty/tree/v1.1.9)
- github.com/getkin/kin-openapi: [v0.76.0](https://github.com/getkin/kin-openapi/tree/v0.76.0)
- github.com/go-gl/glfw/v3.3/glfw: [6f7a984](https://github.com/go-gl/glfw/v3.3/glfw/tree/6f7a984)
- github.com/go-gl/glfw: [e6da0ac](https://github.com/go-gl/glfw/tree/e6da0ac)
- github.com/go-kit/kit: [v0.9.0](https://github.com/go-kit/kit/tree/v0.9.0)
- github.com/go-stack/stack: [v1.8.0](https://github.com/go-stack/stack/tree/v1.8.0)
- github.com/golang-jwt/jwt/v4: [v4.2.0](https://github.com/golang-jwt/jwt/v4/tree/v4.2.0)
- github.com/google/martian/v3: [v3.0.0](https://github.com/google/martian/v3/tree/v3.0.0)
- github.com/google/martian: [v2.1.0+incompatible](https://github.com/google/martian/tree/v2.1.0)
- github.com/google/pprof: [1a94d86](https://github.com/google/pprof/tree/1a94d86)
- github.com/google/renameio: [v0.1.0](https://github.com/google/renameio/tree/v0.1.0)
- github.com/googleapis/gax-go/v2: [v2.0.5](https://github.com/googleapis/gax-go/v2/tree/v2.0.5)
- github.com/hashicorp/golang-lru: [v0.5.1](https://github.com/hashicorp/golang-lru/tree/v0.5.1)
- github.com/ianlancetaylor/demangle: [5e5cf60](https://github.com/ianlancetaylor/demangle/tree/5e5cf60)
- github.com/jstemmer/go-junit-report: [v0.9.1](https://github.com/jstemmer/go-junit-report/tree/v0.9.1)
- github.com/konsorten/go-windows-terminal-sequences: [v1.0.3](https://github.com/konsorten/go-windows-terminal-sequences/tree/v1.0.3)
- github.com/kr/logfmt: [b84e30a](https://github.com/kr/logfmt/tree/b84e30a)
- github.com/rogpeppe/go-internal: [v1.3.0](https://github.com/rogpeppe/go-internal/tree/v1.3.0)
- github.com/sirupsen/logrus: [v1.6.0](https://github.com/sirupsen/logrus/tree/v1.6.0)
- github.com/spf13/afero: [v1.2.2](https://github.com/spf13/afero/tree/v1.2.2)
- go.opencensus.io: v0.22.4
- go.opentelemetry.io/contrib: v0.20.0
- go.opentelemetry.io/otel/exporters/otlp: v0.20.0
- go.opentelemetry.io/otel/sdk/export/metric: v0.20.0
- go.opentelemetry.io/otel/sdk/metric: v0.20.0
- golang.org/x/image: cff245a
- golang.org/x/mobile: d2bd2a2
- google.golang.org/api: v0.30.0
- gopkg.in/errgo.v2: v2.1.0
- rsc.io/binaryregexp: v0.2.0
- rsc.io/quote/v3: v3.1.0
- rsc.io/sampler: v1.3.0
