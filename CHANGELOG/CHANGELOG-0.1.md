# Release notes for v0.1.0

## Changes by Kind

### Feature

- Add library to perform the function of volume population.
- Add example "hello" volume populator that demonstrates how to use the library.

### Uncategorized

- Kubernetes v1.22 or later is required, and the AnyVolumeDataSource feature gate must be enabled ([#3](https://github.com/kubernetes-csi/lib-volume-populator/pull/3), [@bswartz](https://github.com/bswartz))

## Dependencies

### Added
- github.com/Azure/go-autorest: [v14.2.0+incompatible](https://github.com/Azure/go-autorest/tree/v14.2.0)
- github.com/asaskevich/govalidator: [f61b66f](https://github.com/asaskevich/govalidator/tree/f61b66f)
- github.com/creack/pty: [v1.1.9](https://github.com/creack/pty/tree/v1.1.9)
- github.com/form3tech-oss/jwt-go: [v3.2.3+incompatible](https://github.com/form3tech-oss/jwt-go/tree/v3.2.3)
- github.com/go-gl/glfw: [e6da0ac](https://github.com/go-gl/glfw/tree/e6da0ac)
- github.com/gorilla/websocket: [v1.4.2](https://github.com/gorilla/websocket/tree/v1.4.2)
- github.com/mitchellh/mapstructure: [v1.1.2](https://github.com/mitchellh/mapstructure/tree/v1.1.2)
- github.com/moby/spdystream: [v0.2.0](https://github.com/moby/spdystream/tree/v0.2.0)
- github.com/niemeyer/pretty: [a10e7ca](https://github.com/niemeyer/pretty/tree/a10e7ca)
- github.com/nxadm/tail: [v1.4.4](https://github.com/nxadm/tail/tree/v1.4.4)
- github.com/stoewer/go-strcase: [v1.2.0](https://github.com/stoewer/go-strcase/tree/v1.2.0)
- github.com/yuin/goldmark: [v1.2.1](https://github.com/yuin/goldmark/tree/v1.2.1)
- golang.org/x/term: 6a3ed07
- gopkg.in/yaml.v3: 496545a
- rsc.io/quote/v3: v3.1.0
- rsc.io/sampler: v1.3.0

### Changed
- cloud.google.com/go/bigquery: v1.0.1 → v1.4.0
- cloud.google.com/go/datastore: v1.0.0 → v1.1.0
- cloud.google.com/go/pubsub: v1.0.1 → v1.2.0
- cloud.google.com/go/storage: v1.0.0 → v1.6.0
- cloud.google.com/go: v0.51.0 → v0.54.0
- github.com/Azure/go-autorest/autorest/adal: [v0.8.2 → v0.9.13](https://github.com/Azure/go-autorest/autorest/adal/compare/v0.8.2...v0.9.13)
- github.com/Azure/go-autorest/autorest/date: [v0.2.0 → v0.3.0](https://github.com/Azure/go-autorest/autorest/date/compare/v0.2.0...v0.3.0)
- github.com/Azure/go-autorest/autorest/mocks: [v0.3.0 → v0.4.1](https://github.com/Azure/go-autorest/autorest/mocks/compare/v0.3.0...v0.4.1)
- github.com/Azure/go-autorest/autorest: [v0.9.6 → v0.11.18](https://github.com/Azure/go-autorest/autorest/compare/v0.9.6...v0.11.18)
- github.com/Azure/go-autorest/logger: [v0.1.0 → v0.2.1](https://github.com/Azure/go-autorest/logger/compare/v0.1.0...v0.2.1)
- github.com/Azure/go-autorest/tracing: [v0.5.0 → v0.6.0](https://github.com/Azure/go-autorest/tracing/compare/v0.5.0...v0.6.0)
- github.com/PuerkitoBio/purell: [v1.0.0 → v1.1.1](https://github.com/PuerkitoBio/purell/compare/v1.0.0...v1.1.1)
- github.com/PuerkitoBio/urlesc: [5bd2802 → de5bf2a](https://github.com/PuerkitoBio/urlesc/compare/5bd2802...de5bf2a)
- github.com/evanphx/json-patch: [v4.9.0+incompatible → v4.11.0+incompatible](https://github.com/evanphx/json-patch/compare/v4.9.0...v4.11.0)
- github.com/go-gl/glfw/v3.3/glfw: [12ad95a → 6f7a984](https://github.com/go-gl/glfw/v3.3/glfw/compare/12ad95a...6f7a984)
- github.com/go-openapi/jsonpointer: [46af16f → v0.19.3](https://github.com/go-openapi/jsonpointer/compare/46af16f...v0.19.3)
- github.com/go-openapi/jsonreference: [13c6e35 → v0.19.3](https://github.com/go-openapi/jsonreference/compare/13c6e35...v0.19.3)
- github.com/go-openapi/swag: [1d0bd11 → v0.19.5](https://github.com/go-openapi/swag/compare/1d0bd11...v0.19.5)
- github.com/gogo/protobuf: [v1.3.1 → v1.3.2](https://github.com/gogo/protobuf/compare/v1.3.1...v1.3.2)
- github.com/golang/groupcache: [215e871 → 41bb18b](https://github.com/golang/groupcache/compare/215e871...41bb18b)
- github.com/golang/mock: [v1.3.1 → v1.4.1](https://github.com/golang/mock/compare/v1.3.1...v1.4.1)
- github.com/golang/protobuf: [v1.4.2 → v1.5.2](https://github.com/golang/protobuf/compare/v1.4.2...v1.5.2)
- github.com/google/btree: [v1.0.0 → v1.0.1](https://github.com/google/btree/compare/v1.0.0...v1.0.1)
- github.com/google/go-cmp: [v0.4.0 → v0.5.5](https://github.com/google/go-cmp/compare/v0.4.0...v0.5.5)
- github.com/google/pprof: [d4f498a → 1ebb73c](https://github.com/google/pprof/compare/d4f498a...1ebb73c)
- github.com/google/uuid: [v1.1.1 → v1.1.2](https://github.com/google/uuid/compare/v1.1.1...v1.1.2)
- github.com/googleapis/gnostic: [v0.4.1 → v0.5.5](https://github.com/googleapis/gnostic/compare/v0.4.1...v0.5.5)
- github.com/json-iterator/go: [v1.1.10 → v1.1.11](https://github.com/json-iterator/go/compare/v1.1.10...v1.1.11)
- github.com/kisielk/errcheck: [v1.2.0 → v1.5.0](https://github.com/kisielk/errcheck/compare/v1.2.0...v1.5.0)
- github.com/kr/text: [v0.1.0 → v0.2.0](https://github.com/kr/text/compare/v0.1.0...v0.2.0)
- github.com/mailru/easyjson: [d5b7844 → b2ccc51](https://github.com/mailru/easyjson/compare/d5b7844...b2ccc51)
- github.com/onsi/ginkgo: [v1.11.0 → v1.14.0](https://github.com/onsi/ginkgo/compare/v1.11.0...v1.14.0)
- github.com/onsi/gomega: [v1.7.0 → v1.10.1](https://github.com/onsi/gomega/compare/v1.7.0...v1.10.1)
- github.com/stretchr/testify: [v1.4.0 → v1.7.0](https://github.com/stretchr/testify/compare/v1.4.0...v1.7.0)
- go.opencensus.io: v0.22.2 → v0.22.3
- golang.org/x/crypto: 75b2880 → 5ea612d
- golang.org/x/exp: da58074 → 6cc2880
- golang.org/x/lint: fdd1cda → 738671d
- golang.org/x/mod: c90efee → v0.3.0
- golang.org/x/net: 69a7880 → 37e1c6a
- golang.org/x/oauth2: 858c2ad → bf48bf1
- golang.org/x/sync: cd5d95a → 67f06af
- golang.org/x/sys: 5cba982 → 59db8d7
- golang.org/x/text: v0.3.3 → v0.3.6
- golang.org/x/time: 555d28b → 1f47c86
- golang.org/x/tools: 7b8e75d → 113979e
- golang.org/x/xerrors: 9bdfabe → 5ec99f8
- google.golang.org/api: v0.15.0 → v0.20.0
- google.golang.org/genproto: cb27e3a → 1ed22bb
- google.golang.org/grpc: v1.27.0 → v1.27.1
- google.golang.org/protobuf: v1.24.0 → v1.26.0
- gopkg.in/check.v1: 41f04d3 → 8fa4692
- gopkg.in/yaml.v2: v2.2.8 → v2.4.0
- honnef.co/go/tools: v0.0.1-2019.2.3 → v0.0.1-2020.1.3
- k8s.io/api: v0.19.9 → v0.22.0
- k8s.io/apimachinery: v0.19.9 → v0.22.0
- k8s.io/client-go: v0.19.9 → v0.22.0
- k8s.io/klog/v2: v2.8.0 → v2.9.0
- k8s.io/kube-openapi: 6aeccd4 → 9528897
- k8s.io/utils: d5654de → 4b05e18
- sigs.k8s.io/structured-merge-diff/v4: v4.0.1 → v4.1.2

### Removed
- github.com/dgrijalva/jwt-go: [v3.2.0+incompatible](https://github.com/dgrijalva/jwt-go/tree/v3.2.0)
- github.com/docker/spdystream: [449fdfc](https://github.com/docker/spdystream/tree/449fdfc)
- github.com/ghodss/yaml: [73d445a](https://github.com/ghodss/yaml/tree/73d445a)
- github.com/go-openapi/spec: [6aced65](https://github.com/go-openapi/spec/tree/6aced65)
