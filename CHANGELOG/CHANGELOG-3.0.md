# Release notes for v3.0.0

# Changelog since v2.0.0

## Changes by Kind

### Feature

- Add MutatorConfig to Volume Populator Controller to allow specify mutator functions to modify the Kubernetes resources used for volume population. Add the PvcPrimeMutator function to allow modify PVC prime before object creation ([#176](https://github.com/kubernetes-csi/lib-volume-populator/pull/176), [@dannawang0221](https://github.com/dannawang0221))
- Add Workqueue to the VolumePopulatorConfig. Allows users to run controller with customized workqueue. ([#202](https://github.com/kubernetes-csi/lib-volume-populator/pull/202), [@dannawang0221](https://github.com/dannawang0221))
- Add support for provider populate functions ([#133](https://github.com/kubernetes-csi/lib-volume-populator/pull/133), [@dannawang0221](https://github.com/dannawang0221))
- Allow users to delete their PVC during volume population and clean up temporary resources create by volume populators ([#190](https://github.com/kubernetes-csi/lib-volume-populator/pull/190), [@dannawang0221](https://github.com/dannawang0221))
- The library now has a provider metric manager. Similarly to the "ProviderFunctionConfig" which holds provider specific functions for "PopulateFn", "PopulateCleanupFn", and "PopulateCompleteFn", "ProviderMetricManager" holds a provider specific metric handling function called "HandleMetricFn". ([#173](https://github.com/kubernetes-csi/lib-volume-populator/pull/173), [@uriel-guzman](https://github.com/uriel-guzman))
- Update kubernetes dependencies to v1.33.0 ([#212](https://github.com/kubernetes-csi/lib-volume-populator/pull/212), [@sunnylovestiramisu](https://github.com/sunnylovestiramisu))

### Bug or Regression

- BugFix: Fix API group nil pointer for PVCs with dataSource of another PVC ([#82](https://github.com/kubernetes-csi/lib-volume-populator/pull/82), [@akalenyu](https://github.com/akalenyu))
- PersistentVolumeClaim will now get a propagated accessModes. ([#65](https://github.com/kubernetes-csi/lib-volume-populator/pull/65), [@liranr23](https://github.com/liranr23))

### Other (Cleanup or Flake)

- Remove PV from the `PopulatorParams` API ([#195](https://github.com/kubernetes-csi/lib-volume-populator/pull/195), [@dannawang0221](https://github.com/dannawang0221))

### Uncategorized

- Add an optional stop channel that overrides the controller's stop channel for downstream starting/stopping of the controller. ([#164](https://github.com/kubernetes-csi/lib-volume-populator/pull/164), [@uriel-guzman](https://github.com/uriel-guzman))
- Add support for cross-namespace data sources alpha feature ([#117](https://github.com/kubernetes-csi/lib-volume-populator/pull/117), [@ttakahashi21](https://github.com/ttakahashi21))
- Update kubernetes dependencies to v1.31.0 ([#192](https://github.com/kubernetes-csi/lib-volume-populator/pull/192), [@sunnylovestiramisu](https://github.com/sunnylovestiramisu))
- Update kubernetes dependencies to v1.32.0 ([#194](https://github.com/kubernetes-csi/lib-volume-populator/pull/194), [@sunnylovestiramisu](https://github.com/sunnylovestiramisu))

## Dependencies

### Added
- github.com/ahmetb/gen-crd-api-reference-docs: [v0.3.0](https://github.com/ahmetb/gen-crd-api-reference-docs/tree/v0.3.0)
- github.com/alecthomas/kingpin/v2: [v2.4.0](https://github.com/alecthomas/kingpin/v2/tree/v2.4.0)
- github.com/evanphx/json-patch/v5: [v5.9.0](https://github.com/evanphx/json-patch/v5/tree/v5.9.0)
- github.com/fatih/color: [v1.17.0](https://github.com/fatih/color/tree/v1.17.0)
- github.com/fxamacker/cbor/v2: [v2.7.0](https://github.com/fxamacker/cbor/v2/tree/v2.7.0)
- github.com/go-task/slim-sprig/v3: [v3.0.0](https://github.com/go-task/slim-sprig/v3/tree/v3.0.0)
- github.com/gobuffalo/flect: [v1.0.2](https://github.com/gobuffalo/flect/tree/v1.0.2)
- github.com/google/gnostic-models: [v0.6.9](https://github.com/google/gnostic-models/tree/v0.6.9)
- github.com/google/pprof: [d1b30fe](https://github.com/google/pprof/tree/d1b30fe)
- github.com/gorilla/websocket: [e064f32](https://github.com/gorilla/websocket/tree/e064f32)
- github.com/klauspost/compress: [v1.18.0](https://github.com/klauspost/compress/tree/v1.18.0)
- github.com/kylelemons/godebug: [v1.1.0](https://github.com/kylelemons/godebug/tree/v1.1.0)
- github.com/mattn/go-colorable: [v0.1.13](https://github.com/mattn/go-colorable/tree/v0.1.13)
- github.com/mattn/go-isatty: [v0.0.20](https://github.com/mattn/go-isatty/tree/v0.0.20)
- github.com/miekg/dns: [v1.1.62](https://github.com/miekg/dns/tree/v1.1.62)
- github.com/rogpeppe/go-internal: [v1.13.1](https://github.com/rogpeppe/go-internal/tree/v1.13.1)
- github.com/russross/blackfriday/v2: [v2.1.0](https://github.com/russross/blackfriday/v2/tree/v2.1.0)
- github.com/x448/float16: [v0.8.4](https://github.com/x448/float16/tree/v0.8.4)
- github.com/xhit/go-str2duration/v2: [v2.1.0](https://github.com/xhit/go-str2duration/v2/tree/v2.1.0)
- go.opentelemetry.io/auto/sdk: v1.1.0
- google.golang.org/genproto/googleapis/api: e6fa225
- google.golang.org/genproto/googleapis/rpc: e6fa225
- google.golang.org/grpc/cmd/protoc-gen-go-grpc: v1.5.1
- gopkg.in/evanphx/json-patch.v4: v4.12.0
- k8s.io/apiextensions-apiserver: v0.31.1
- k8s.io/code-generator: v0.31.1
- k8s.io/gengo/v2: a7b603a
- k8s.io/klog: v0.2.0
- sigs.k8s.io/controller-runtime: v0.18.0
- sigs.k8s.io/controller-tools: v0.16.3
- sigs.k8s.io/gateway-api: v1.2.1
- sigs.k8s.io/randfill: v1.0.0

### Changed
- cloud.google.com/go/compute/metadata: v0.2.0 → v0.3.0
- github.com/Azure/go-ansiterm: [d185dfc → 306776e](https://github.com/Azure/go-ansiterm/compare/d185dfc...306776e)
- github.com/NYTimes/gziphandler: [56545f4 → v1.1.1](https://github.com/NYTimes/gziphandler/compare/56545f4...v1.1.1)
- github.com/alecthomas/units: [f65c72e → b94a6e3](https://github.com/alecthomas/units/compare/f65c72e...b94a6e3)
- github.com/cenkalti/backoff/v4: [v4.1.3 → v4.3.0](https://github.com/cenkalti/backoff/v4/compare/v4.1.3...v4.3.0)
- github.com/cespare/xxhash/v2: [v2.1.2 → v2.3.0](https://github.com/cespare/xxhash/v2/compare/v2.1.2...v2.3.0)
- github.com/davecgh/go-spew: [v1.1.1 → d8f796a](https://github.com/davecgh/go-spew/compare/v1.1.1...d8f796a)
- github.com/emicklei/go-restful/v3: [v3.9.0 → v3.12.0](https://github.com/emicklei/go-restful/v3/compare/v3.9.0...v3.12.0)
- github.com/evanphx/json-patch: [v4.12.0+incompatible → v5.7.0+incompatible](https://github.com/evanphx/json-patch/compare/v4.12.0...v5.7.0)
- github.com/felixge/httpsnoop: [v1.0.3 → v1.0.4](https://github.com/felixge/httpsnoop/compare/v1.0.3...v1.0.4)
- github.com/go-logr/logr: [v1.2.3 → v1.4.2](https://github.com/go-logr/logr/compare/v1.2.3...v1.4.2)
- github.com/go-logr/zapr: [v1.2.3 → v1.3.0](https://github.com/go-logr/zapr/compare/v1.2.3...v1.3.0)
- github.com/go-openapi/jsonpointer: [v0.19.5 → v0.21.0](https://github.com/go-openapi/jsonpointer/compare/v0.19.5...v0.21.0)
- github.com/go-openapi/jsonreference: [v0.20.0 → v0.21.0](https://github.com/go-openapi/jsonreference/compare/v0.20.0...v0.21.0)
- github.com/go-openapi/swag: [v0.22.3 → v0.23.0](https://github.com/go-openapi/swag/compare/v0.22.3...v0.23.0)
- github.com/golang/protobuf: [v1.5.2 → v1.5.4](https://github.com/golang/protobuf/compare/v1.5.2...v1.5.4)
- github.com/google/btree: [v1.0.1 → v1.1.3](https://github.com/google/btree/compare/v1.0.1...v1.1.3)
- github.com/google/go-cmp: [v0.5.9 → v0.7.0](https://github.com/google/go-cmp/compare/v0.5.9...v0.7.0)
- github.com/google/uuid: [v1.1.2 → v1.6.0](https://github.com/google/uuid/compare/v1.1.2...v1.6.0)
- github.com/gregjones/httpcache: [9cad4c3 → 901d907](https://github.com/gregjones/httpcache/compare/9cad4c3...901d907)
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.7.0 → v2.24.0](https://github.com/grpc-ecosystem/grpc-gateway/v2/compare/v2.7.0...v2.24.0)
- github.com/imdario/mergo: [v0.3.13 → v0.3.16](https://github.com/imdario/mergo/compare/v0.3.13...v0.3.16)
- github.com/inconshreveable/mousetrap: [v1.0.1 → v1.1.0](https://github.com/inconshreveable/mousetrap/compare/v1.0.1...v1.1.0)
- github.com/kr/pretty: [v0.2.0 → v0.3.1](https://github.com/kr/pretty/compare/v0.2.0...v0.3.1)
- github.com/moby/spdystream: [v0.2.0 → v0.5.0](https://github.com/moby/spdystream/compare/v0.2.0...v0.5.0)
- github.com/moby/term: [39b0c02 → v0.5.0](https://github.com/moby/term/compare/39b0c02...v0.5.0)
- github.com/onsi/ginkgo/v2: [v2.4.0 → v2.21.0](https://github.com/onsi/ginkgo/v2/compare/v2.4.0...v2.21.0)
- github.com/onsi/gomega: [v1.23.0 → v1.35.1](https://github.com/onsi/gomega/compare/v1.23.0...v1.35.1)
- github.com/pmezard/go-difflib: [v1.0.0 → 5d4384e](https://github.com/pmezard/go-difflib/compare/v1.0.0...5d4384e)
- github.com/prometheus/client_golang: [v1.14.0 → v1.22.0](https://github.com/prometheus/client_golang/compare/v1.14.0...v1.22.0)
- github.com/prometheus/client_model: [v0.3.0 → v0.6.1](https://github.com/prometheus/client_model/compare/v0.3.0...v0.6.1)
- github.com/prometheus/common: [v0.39.0 → v0.62.0](https://github.com/prometheus/common/compare/v0.39.0...v0.62.0)
- github.com/prometheus/procfs: [v0.8.0 → v0.15.1](https://github.com/prometheus/procfs/compare/v0.8.0...v0.15.1)
- github.com/spf13/cobra: [v1.6.0 → v1.8.1](https://github.com/spf13/cobra/compare/v1.6.0...v1.8.1)
- github.com/stretchr/objx: [v0.1.0 → v0.5.2](https://github.com/stretchr/objx/compare/v0.1.0...v0.5.2)
- github.com/stretchr/testify: [v1.8.0 → v1.10.0](https://github.com/stretchr/testify/compare/v1.8.0...v1.10.0)
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.35.0 → v0.58.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.10.0 → v1.33.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.10.0 → v1.33.0
- go.opentelemetry.io/otel/metric: v0.31.0 → v1.33.0
- go.opentelemetry.io/otel/sdk: v1.10.0 → v1.33.0
- go.opentelemetry.io/otel/trace: v1.10.0 → v1.33.0
- go.opentelemetry.io/otel: v1.10.0 → v1.33.0
- go.opentelemetry.io/proto/otlp: v0.19.0 → v1.4.0
- go.uber.org/goleak: v1.2.0 → v1.3.0
- go.uber.org/multierr: v1.6.0 → v1.11.0
- go.uber.org/zap: v1.19.0 → v1.27.0
- golang.org/x/crypto: 75b2880 → v0.36.0
- golang.org/x/exp: 509febe → fe59bbe
- golang.org/x/mod: 86c51ed → v0.20.0
- golang.org/x/net: v0.4.0 → v0.38.0
- golang.org/x/oauth2: v0.3.0 → v0.27.0
- golang.org/x/sync: 0de741c → v0.12.0
- golang.org/x/sys: v0.3.0 → v0.31.0
- golang.org/x/term: v0.3.0 → v0.30.0
- golang.org/x/text: v0.5.0 → v0.23.0
- golang.org/x/time: v0.1.0 → v0.9.0
- golang.org/x/tools: v0.1.12 → v0.26.0
- google.golang.org/grpc: v1.49.0 → v1.68.1
- google.golang.org/protobuf: v1.28.1 → v1.36.5
- k8s.io/api: v0.26.0 → v0.33.0
- k8s.io/apimachinery: v0.26.0 → v0.33.0
- k8s.io/client-go: v0.26.0 → v0.33.0
- k8s.io/component-base: v0.26.0 → v0.33.0
- k8s.io/component-helpers: v0.26.0 → v0.33.0
- k8s.io/gengo: 485abfe → 9cce18d
- k8s.io/klog/v2: v2.80.1 → v2.130.1
- k8s.io/kube-openapi: 172d655 → c8a335a
- k8s.io/utils: 1a15be2 → 3ea5e8c
- sigs.k8s.io/json: f223a00 → 9aa6b5e
- sigs.k8s.io/structured-merge-diff/v4: v4.2.3 → v4.6.0
- sigs.k8s.io/yaml: v1.3.0 → v1.4.0

### Removed
- cloud.google.com/go: v0.34.0
- github.com/BurntSushi/toml: [v0.3.1](https://github.com/BurntSushi/toml/tree/v0.3.1)
- github.com/OneOfOne/xxhash: [v1.2.2](https://github.com/OneOfOne/xxhash/tree/v1.2.2)
- github.com/PuerkitoBio/purell: [v1.1.1](https://github.com/PuerkitoBio/purell/tree/v1.1.1)
- github.com/PuerkitoBio/urlesc: [de5bf2a](https://github.com/PuerkitoBio/urlesc/tree/de5bf2a)
- github.com/alecthomas/template: [fb15b89](https://github.com/alecthomas/template/tree/fb15b89)
- github.com/antihax/optional: [v1.0.0](https://github.com/antihax/optional/tree/v1.0.0)
- github.com/asaskevich/govalidator: [f61b66f](https://github.com/asaskevich/govalidator/tree/f61b66f)
- github.com/buger/jsonparser: [v1.1.1](https://github.com/buger/jsonparser/tree/v1.1.1)
- github.com/census-instrumentation/opencensus-proto: [v0.2.1](https://github.com/census-instrumentation/opencensus-proto/tree/v0.2.1)
- github.com/cespare/xxhash: [v1.1.0](https://github.com/cespare/xxhash/tree/v1.1.0)
- github.com/client9/misspell: [v0.3.4](https://github.com/client9/misspell/tree/v0.3.4)
- github.com/cncf/udpa/go: [5459f2c](https://github.com/cncf/udpa/go/tree/5459f2c)
- github.com/cncf/xds/go: [fbca930](https://github.com/cncf/xds/go/tree/fbca930)
- github.com/docopt/docopt-go: [ee0de3b](https://github.com/docopt/docopt-go/tree/ee0de3b)
- github.com/elazarl/goproxy: [947c36d](https://github.com/elazarl/goproxy/tree/947c36d)
- github.com/envoyproxy/go-control-plane: [63b5d3c](https://github.com/envoyproxy/go-control-plane/tree/63b5d3c)
- github.com/envoyproxy/protoc-gen-validate: [v0.1.0](https://github.com/envoyproxy/protoc-gen-validate/tree/v0.1.0)
- github.com/flowstack/go-jsonschema: [v0.1.1](https://github.com/flowstack/go-jsonschema/tree/v0.1.1)
- github.com/ghodss/yaml: [v1.0.0](https://github.com/ghodss/yaml/tree/v1.0.0)
- github.com/go-kit/log: [v0.2.1](https://github.com/go-kit/log/tree/v0.2.1)
- github.com/go-logfmt/logfmt: [v0.5.1](https://github.com/go-logfmt/logfmt/tree/v0.5.1)
- github.com/golang/glog: [23def4e](https://github.com/golang/glog/tree/23def4e)
- github.com/golang/groupcache: [41bb18b](https://github.com/golang/groupcache/tree/41bb18b)
- github.com/golang/mock: [v1.1.1](https://github.com/golang/mock/tree/v1.1.1)
- github.com/google/gnostic: [v0.6.9](https://github.com/google/gnostic/tree/v0.6.9)
- github.com/grpc-ecosystem/grpc-gateway: [v1.16.0](https://github.com/grpc-ecosystem/grpc-gateway/tree/v1.16.0)
- github.com/kr/pty: [v1.1.1](https://github.com/kr/pty/tree/v1.1.1)
- github.com/matttproud/golang_protobuf_extensions: [v1.0.4](https://github.com/matttproud/golang_protobuf_extensions/tree/v1.0.4)
- github.com/mitchellh/mapstructure: [v1.1.2](https://github.com/mitchellh/mapstructure/tree/v1.1.2)
- github.com/niemeyer/pretty: [a10e7ca](https://github.com/niemeyer/pretty/tree/a10e7ca)
- github.com/rogpeppe/fastuuid: [v1.2.0](https://github.com/rogpeppe/fastuuid/tree/v1.2.0)
- github.com/spaolacci/murmur3: [f09979e](https://github.com/spaolacci/murmur3/tree/f09979e)
- github.com/stoewer/go-strcase: [v1.2.0](https://github.com/stoewer/go-strcase/tree/v1.2.0)
- github.com/xeipuuv/gojsonpointer: [4e3ac27](https://github.com/xeipuuv/gojsonpointer/tree/4e3ac27)
- github.com/xeipuuv/gojsonreference: [bd5ef7b](https://github.com/xeipuuv/gojsonreference/tree/bd5ef7b)
- github.com/xeipuuv/gojsonschema: [v1.2.0](https://github.com/xeipuuv/gojsonschema/tree/v1.2.0)
- go.opentelemetry.io/otel/exporters/otlp/internal/retry: v1.10.0
- go.uber.org/atomic: v1.7.0
- golang.org/x/lint: d0100b6
- google.golang.org/appengine: v1.6.7
- google.golang.org/genproto: c8bf987
- gopkg.in/alecthomas/kingpin.v2: v2.2.6
- gotest.tools/v3: v3.0.3
- honnef.co/go/tools: ea95bdf
