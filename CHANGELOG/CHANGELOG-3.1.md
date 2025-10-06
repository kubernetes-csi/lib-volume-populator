# Release notes for v3.1.0

# Changelog since v3.0.0

## Changes by Kind

### Feature

- Add the ability to adjust the PV/PVC/Pod/StorageClass shared informer options with a SharedInformerOptions array on VolumePopulatorConfig. ([#219](https://github.com/kubernetes-csi/lib-volume-populator/pull/219), [@pwschuurman](https://github.com/pwschuurman))
- Adds detailed structured logging to the volume populator controller, improving observability and making it easier to debug and monitor volume population status. ([#222](https://github.com/kubernetes-csi/lib-volume-populator/pull/222), [@Abrahami2](https://github.com/Abrahami2))
- Update dependencies for Kubernetes 1.34 ([#231](https://github.com/kubernetes-csi/lib-volume-populator/pull/231), [@sunnylovestiramisu](https://github.com/sunnylovestiramisu))
- Update gateway-api to v1.4.0 ([#236](https://github.com/kubernetes-csi/lib-volume-populator/pull/236), [@sunnylovestiramisu](https://github.com/sunnylovestiramisu))
- Update to gateway-api to v1.4.0-rc.2 ([#234](https://github.com/kubernetes-csi/lib-volume-populator/pull/234), [@sunnylovestiramisu](https://github.com/sunnylovestiramisu))

### Bug or Regression

- Fixed the module name so that it can import the latest major version. You can now import the latest module with 'import "github.com/kubernetes-csi/lib-volume-populator/v3"' ([#225](https://github.com/kubernetes-csi/lib-volume-populator/pull/225), [@everpeace](https://github.com/everpeace))

### Other (Cleanup or Flake)

- The reasonPopulateOperationStartSuccess event gets removed from syncPVC. You can add it in the PopulateFn. This change will reduce the overall number of reported events.
  - EventRecorder gets added to the VolumePopulatorConfig. This allows you to use customized  EventRecorder ([#217](https://github.com/kubernetes-csi/lib-volume-populator/pull/217), [@dannawang0221](https://github.com/dannawang0221))


## Dependencies

### Added
- github.com/Masterminds/goutils: [v1.1.1](https://github.com/Masterminds/goutils/tree/v1.1.1)
- github.com/Masterminds/semver: [v1.5.0](https://github.com/Masterminds/semver/tree/v1.5.0)
- github.com/Masterminds/sprig: [v2.22.0+incompatible](https://github.com/Masterminds/sprig/tree/v2.22.0)
- github.com/elastic/crd-ref-docs: [v0.2.0](https://github.com/elastic/crd-ref-docs/tree/v0.2.0)
- github.com/goccy/go-yaml: [v1.18.0](https://github.com/goccy/go-yaml/tree/v1.18.0)
- github.com/huandu/xstrings: [v1.3.3](https://github.com/huandu/xstrings/tree/v1.3.3)
- github.com/mitchellh/copystructure: [v1.2.0](https://github.com/mitchellh/copystructure/tree/v1.2.0)
- github.com/mitchellh/reflectwalk: [v1.0.2](https://github.com/mitchellh/reflectwalk/tree/v1.0.2)
- go.yaml.in/yaml/v2: v2.4.2
- go.yaml.in/yaml/v3: v3.0.4
- golang.org/x/tools/go/expect: v0.1.1-deprecated
- sigs.k8s.io/structured-merge-diff/v6: v6.3.0

### Changed
- github.com/emicklei/go-restful/v3: [v3.12.0 → v3.13.0](https://github.com/emicklei/go-restful/v3/compare/v3.12.0...v3.13.0)
- github.com/evanphx/json-patch/v5: [v5.9.0 → v5.9.11](https://github.com/evanphx/json-patch/v5/compare/v5.9.0...v5.9.11)
- github.com/fatih/color: [v1.17.0 → v1.18.0](https://github.com/fatih/color/compare/v1.17.0...v1.18.0)
- github.com/fxamacker/cbor/v2: [v2.7.0 → v2.9.0](https://github.com/fxamacker/cbor/v2/compare/v2.7.0...v2.9.0)
- github.com/go-logr/logr: [v1.4.2 → v1.4.3](https://github.com/go-logr/logr/compare/v1.4.2...v1.4.3)
- github.com/go-openapi/jsonpointer: [v0.21.0 → v0.21.2](https://github.com/go-openapi/jsonpointer/compare/v0.21.0...v0.21.2)
- github.com/go-openapi/swag: [v0.23.0 → v0.23.1](https://github.com/go-openapi/swag/compare/v0.23.0...v0.23.1)
- github.com/gobuffalo/flect: [v1.0.2 → v1.0.3](https://github.com/gobuffalo/flect/compare/v1.0.2...v1.0.3)
- github.com/golang/protobuf: [v1.5.4 → v1.5.0](https://github.com/golang/protobuf/compare/v1.5.4...v1.5.0)
- github.com/google/gnostic-models: [v0.6.9 → v0.7.0](https://github.com/google/gnostic-models/compare/v0.6.9...v0.7.0)
- github.com/google/gofuzz: [v1.2.0 → v1.0.0](https://github.com/google/gofuzz/compare/v1.2.0...v1.0.0)
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.24.0 → v2.26.3](https://github.com/grpc-ecosystem/grpc-gateway/v2/compare/v2.24.0...v2.26.3)
- github.com/imdario/mergo: [v0.3.16 → v0.3.11](https://github.com/imdario/mergo/compare/v0.3.16...v0.3.11)
- github.com/mailru/easyjson: [v0.7.7 → v0.9.0](https://github.com/mailru/easyjson/compare/v0.7.7...v0.9.0)
- github.com/miekg/dns: [v1.1.62 → v1.1.68](https://github.com/miekg/dns/compare/v1.1.62...v1.1.68)
- github.com/modern-go/reflect2: [v1.0.2 → 35a7c28](https://github.com/modern-go/reflect2/compare/v1.0.2...35a7c28)
- github.com/prometheus/client_golang: [v1.22.0 → v1.23.0](https://github.com/prometheus/client_golang/compare/v1.22.0...v1.23.0)
- github.com/prometheus/client_model: [v0.6.1 → v0.6.2](https://github.com/prometheus/client_model/compare/v0.6.1...v0.6.2)
- github.com/prometheus/common: [v0.62.0 → v0.65.0](https://github.com/prometheus/common/compare/v0.62.0...v0.65.0)
- github.com/prometheus/procfs: [v0.15.1 → v0.17.0](https://github.com/prometheus/procfs/compare/v0.15.1...v0.17.0)
- github.com/spf13/cobra: [v1.8.1 → v1.9.1](https://github.com/spf13/cobra/compare/v1.8.1...v1.9.1)
- github.com/spf13/pflag: [v1.0.5 → v1.0.7](https://github.com/spf13/pflag/compare/v1.0.5...v1.0.7)
- github.com/stretchr/testify: [v1.10.0 → v1.11.0](https://github.com/stretchr/testify/compare/v1.10.0...v1.11.0)
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.33.0 → v1.34.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.33.0 → v1.34.0
- go.opentelemetry.io/otel/metric: v1.33.0 → v1.35.0
- go.opentelemetry.io/otel/sdk: v1.33.0 → v1.34.0
- go.opentelemetry.io/otel/trace: v1.33.0 → v1.35.0
- go.opentelemetry.io/otel: v1.33.0 → v1.35.0
- go.opentelemetry.io/proto/otlp: v1.4.0 → v1.5.0
- golang.org/x/crypto: v0.36.0 → v0.41.0
- golang.org/x/mod: v0.20.0 → v0.27.0
- golang.org/x/net: v0.38.0 → v0.43.0
- golang.org/x/oauth2: v0.27.0 → v0.30.0
- golang.org/x/sync: v0.12.0 → v0.16.0
- golang.org/x/sys: v0.31.0 → v0.35.0
- golang.org/x/term: v0.30.0 → v0.34.0
- golang.org/x/text: v0.23.0 → v0.28.0
- golang.org/x/time: v0.9.0 → v0.12.0
- golang.org/x/tools: v0.26.0 → v0.36.0
- google.golang.org/genproto/googleapis/api: e6fa225 → a0af3ef
- google.golang.org/genproto/googleapis/rpc: e6fa225 → ef028d9
- google.golang.org/grpc: v1.68.1 → v1.75.1
- google.golang.org/protobuf: v1.36.5 → v1.36.8
- gopkg.in/evanphx/json-patch.v4: v4.12.0 → v4.13.0
- k8s.io/api: v0.33.0 → v0.34.1
- k8s.io/apiextensions-apiserver: v0.31.1 → v0.34.1
- k8s.io/apimachinery: v0.33.0 → v0.34.1
- k8s.io/client-go: v0.33.0 → v0.34.1
- k8s.io/code-generator: v0.31.1 → v0.34.1
- k8s.io/component-base: v0.33.0 → v0.34.0
- k8s.io/component-helpers: v0.33.0 → v0.34.0
- k8s.io/gengo/v2: a7b603a → c297c0c
- k8s.io/kube-openapi: c8a335a → d7b6acb
- k8s.io/utils: 3ea5e8c → 0af2bda
- sigs.k8s.io/controller-runtime: v0.18.0 → v0.22.1
- sigs.k8s.io/controller-tools: v0.16.3 → v0.19.0
- sigs.k8s.io/gateway-api: v1.2.1 → v1.4.0
- sigs.k8s.io/json: 9aa6b5e → 2d32026
- sigs.k8s.io/yaml: v1.4.0 → v1.6.0

### Removed
- github.com/ahmetb/gen-crd-api-reference-docs: [v0.3.0](https://github.com/ahmetb/gen-crd-api-reference-docs/tree/v0.3.0)
- github.com/evanphx/json-patch: [v5.7.0+incompatible](https://github.com/evanphx/json-patch/tree/v5.7.0)
- github.com/russross/blackfriday/v2: [v2.1.0](https://github.com/russross/blackfriday/v2/tree/v2.1.0)
- golang.org/x/exp: fe59bbe
- k8s.io/gengo: 9cce18d
- k8s.io/klog: v0.2.0
- sigs.k8s.io/structured-merge-diff/v4: v4.6.0
