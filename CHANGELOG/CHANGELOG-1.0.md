# Release notes for v1.0.0

## Changes by Kind

### Feature

- Hello-populator: add --namespace flag to the hello-populator example to deploy in another namespace. ([#17](https://github.com/kubernetes-csi/lib-volume-populator/pull/17), [@mkimuram](https://github.com/mkimuram))
- The library now collects metrics internally and can expose them through an HTTP listener if the importing application desires. The metrics include a histogram of populator operation durations (and their results) as well as a gauge of the number of populator operations in flight. ([#18](https://github.com/kubernetes-csi/lib-volume-populator/pull/18), [@bswartz](https://github.com/bswartz))
- The example "hello" populator is updated for compatibility with v1.0 of volume-data-source-validator and its v1beta1 API. ([#20](https://github.com/kubernetes-csi/lib-volume-populator/pull/20), [@bswartz](https://github.com/bswartz))
 
### Uncategorized

- Kubernetes client dependencies are updated to v1.23.4 in this release. ([#13](https://github.com/kubernetes-csi/lib-volume-populator/pull/13), [@humblec](https://github.com/humblec))

## Dependencies

### Added
- cloud.google.com/go/firestore: v1.1.0
- github.com/Azure/go-ansiterm: [d185dfc](https://github.com/Azure/go-ansiterm/tree/d185dfc)
- github.com/OneOfOne/xxhash: [v1.2.2](https://github.com/OneOfOne/xxhash/tree/v1.2.2)
- github.com/alecthomas/template: [fb15b89](https://github.com/alecthomas/template/tree/fb15b89)
- github.com/alecthomas/units: [f65c72e](https://github.com/alecthomas/units/tree/f65c72e)
- github.com/antihax/optional: [v1.0.0](https://github.com/antihax/optional/tree/v1.0.0)
- github.com/armon/circbuf: [bbbad09](https://github.com/armon/circbuf/tree/bbbad09)
- github.com/armon/go-metrics: [f0300d1](https://github.com/armon/go-metrics/tree/f0300d1)
- github.com/armon/go-radix: [7fddfc3](https://github.com/armon/go-radix/tree/7fddfc3)
- github.com/benbjohnson/clock: [v1.1.0](https://github.com/benbjohnson/clock/tree/v1.1.0)
- github.com/beorn7/perks: [v1.0.1](https://github.com/beorn7/perks/tree/v1.0.1)
- github.com/bgentry/speakeasy: [v0.1.0](https://github.com/bgentry/speakeasy/tree/v0.1.0)
- github.com/bketelsen/crypt: [v0.0.4](https://github.com/bketelsen/crypt/tree/v0.0.4)
- github.com/blang/semver: [v3.5.1+incompatible](https://github.com/blang/semver/tree/v3.5.1)
- github.com/cespare/xxhash/v2: [v2.1.1](https://github.com/cespare/xxhash/v2/tree/v2.1.1)
- github.com/cespare/xxhash: [v1.1.0](https://github.com/cespare/xxhash/tree/v1.1.0)
- github.com/cncf/xds/go: [fbca930](https://github.com/cncf/xds/go/tree/fbca930)
- github.com/coreos/go-semver: [v0.3.0](https://github.com/coreos/go-semver/tree/v0.3.0)
- github.com/coreos/go-systemd/v22: [v22.3.2](https://github.com/coreos/go-systemd/v22/tree/v22.3.2)
- github.com/cpuguy83/go-md2man/v2: [v2.0.0](https://github.com/cpuguy83/go-md2man/v2/tree/v2.0.0)
- github.com/fatih/color: [v1.7.0](https://github.com/fatih/color/tree/v1.7.0)
- github.com/felixge/httpsnoop: [v1.0.1](https://github.com/felixge/httpsnoop/tree/v1.0.1)
- github.com/go-kit/kit: [v0.9.0](https://github.com/go-kit/kit/tree/v0.9.0)
- github.com/go-kit/log: [v0.1.0](https://github.com/go-kit/log/tree/v0.1.0)
- github.com/go-logfmt/logfmt: [v0.5.0](https://github.com/go-logfmt/logfmt/tree/v0.5.0)
- github.com/go-logr/zapr: [v1.2.0](https://github.com/go-logr/zapr/tree/v1.2.0)
- github.com/go-stack/stack: [v1.8.0](https://github.com/go-stack/stack/tree/v1.8.0)
- github.com/godbus/dbus/v5: [v5.0.4](https://github.com/godbus/dbus/v5/tree/v5.0.4)
- github.com/gopherjs/gopherjs: [0766667](https://github.com/gopherjs/gopherjs/tree/0766667)
- github.com/grpc-ecosystem/grpc-gateway: [v1.16.0](https://github.com/grpc-ecosystem/grpc-gateway/tree/v1.16.0)
- github.com/hashicorp/consul/api: [v1.1.0](https://github.com/hashicorp/consul/api/tree/v1.1.0)
- github.com/hashicorp/consul/sdk: [v0.1.1](https://github.com/hashicorp/consul/sdk/tree/v0.1.1)
- github.com/hashicorp/errwrap: [v1.0.0](https://github.com/hashicorp/errwrap/tree/v1.0.0)
- github.com/hashicorp/go-cleanhttp: [v0.5.1](https://github.com/hashicorp/go-cleanhttp/tree/v0.5.1)
- github.com/hashicorp/go-immutable-radix: [v1.0.0](https://github.com/hashicorp/go-immutable-radix/tree/v1.0.0)
- github.com/hashicorp/go-msgpack: [v0.5.3](https://github.com/hashicorp/go-msgpack/tree/v0.5.3)
- github.com/hashicorp/go-multierror: [v1.0.0](https://github.com/hashicorp/go-multierror/tree/v1.0.0)
- github.com/hashicorp/go-rootcerts: [v1.0.0](https://github.com/hashicorp/go-rootcerts/tree/v1.0.0)
- github.com/hashicorp/go-sockaddr: [v1.0.0](https://github.com/hashicorp/go-sockaddr/tree/v1.0.0)
- github.com/hashicorp/go-syslog: [v1.0.0](https://github.com/hashicorp/go-syslog/tree/v1.0.0)
- github.com/hashicorp/go-uuid: [v1.0.1](https://github.com/hashicorp/go-uuid/tree/v1.0.1)
- github.com/hashicorp/go.net: [v0.0.1](https://github.com/hashicorp/go.net/tree/v0.0.1)
- github.com/hashicorp/hcl: [v1.0.0](https://github.com/hashicorp/hcl/tree/v1.0.0)
- github.com/hashicorp/logutils: [v1.0.0](https://github.com/hashicorp/logutils/tree/v1.0.0)
- github.com/hashicorp/mdns: [v1.0.0](https://github.com/hashicorp/mdns/tree/v1.0.0)
- github.com/hashicorp/memberlist: [v0.1.3](https://github.com/hashicorp/memberlist/tree/v0.1.3)
- github.com/hashicorp/serf: [v0.8.2](https://github.com/hashicorp/serf/tree/v0.8.2)
- github.com/inconshreveable/mousetrap: [v1.0.0](https://github.com/inconshreveable/mousetrap/tree/v1.0.0)
- github.com/jpillora/backoff: [v1.0.0](https://github.com/jpillora/backoff/tree/v1.0.0)
- github.com/jtolds/gls: [v4.20.0+incompatible](https://github.com/jtolds/gls/tree/v4.20.0)
- github.com/julienschmidt/httprouter: [v1.3.0](https://github.com/julienschmidt/httprouter/tree/v1.3.0)
- github.com/konsorten/go-windows-terminal-sequences: [v1.0.3](https://github.com/konsorten/go-windows-terminal-sequences/tree/v1.0.3)
- github.com/kr/fs: [v0.1.0](https://github.com/kr/fs/tree/v0.1.0)
- github.com/kr/logfmt: [b84e30a](https://github.com/kr/logfmt/tree/b84e30a)
- github.com/magiconair/properties: [v1.8.5](https://github.com/magiconair/properties/tree/v1.8.5)
- github.com/mattn/go-colorable: [v0.0.9](https://github.com/mattn/go-colorable/tree/v0.0.9)
- github.com/mattn/go-isatty: [v0.0.3](https://github.com/mattn/go-isatty/tree/v0.0.3)
- github.com/matttproud/golang_protobuf_extensions: [c182aff](https://github.com/matttproud/golang_protobuf_extensions/tree/c182aff)
- github.com/miekg/dns: [v1.0.14](https://github.com/miekg/dns/tree/v1.0.14)
- github.com/mitchellh/cli: [v1.0.0](https://github.com/mitchellh/cli/tree/v1.0.0)
- github.com/mitchellh/go-homedir: [v1.0.0](https://github.com/mitchellh/go-homedir/tree/v1.0.0)
- github.com/mitchellh/go-testing-interface: [v1.0.0](https://github.com/mitchellh/go-testing-interface/tree/v1.0.0)
- github.com/mitchellh/gox: [v0.4.0](https://github.com/mitchellh/gox/tree/v0.4.0)
- github.com/mitchellh/iochan: [v1.0.0](https://github.com/mitchellh/iochan/tree/v1.0.0)
- github.com/moby/term: [9d4ed18](https://github.com/moby/term/tree/9d4ed18)
- github.com/mwitkow/go-conntrack: [2f06839](https://github.com/mwitkow/go-conntrack/tree/2f06839)
- github.com/pascaldekloe/goe: [57f6aae](https://github.com/pascaldekloe/goe/tree/57f6aae)
- github.com/pelletier/go-toml: [v1.9.3](https://github.com/pelletier/go-toml/tree/v1.9.3)
- github.com/pkg/sftp: [v1.10.1](https://github.com/pkg/sftp/tree/v1.10.1)
- github.com/posener/complete: [v1.1.1](https://github.com/posener/complete/tree/v1.1.1)
- github.com/prometheus/client_golang: [v1.11.0](https://github.com/prometheus/client_golang/tree/v1.11.0)
- github.com/prometheus/common: [v0.28.0](https://github.com/prometheus/common/tree/v0.28.0)
- github.com/prometheus/procfs: [v0.6.0](https://github.com/prometheus/procfs/tree/v0.6.0)
- github.com/rogpeppe/fastuuid: [v1.2.0](https://github.com/rogpeppe/fastuuid/tree/v1.2.0)
- github.com/russross/blackfriday/v2: [v2.0.1](https://github.com/russross/blackfriday/v2/tree/v2.0.1)
- github.com/ryanuber/columnize: [9b3edd6](https://github.com/ryanuber/columnize/tree/9b3edd6)
- github.com/sean-/seed: [e2103e2](https://github.com/sean-/seed/tree/e2103e2)
- github.com/shurcooL/sanitized_anchor_name: [v1.0.0](https://github.com/shurcooL/sanitized_anchor_name/tree/v1.0.0)
- github.com/sirupsen/logrus: [v1.6.0](https://github.com/sirupsen/logrus/tree/v1.6.0)
- github.com/smartystreets/assertions: [b2de0cb](https://github.com/smartystreets/assertions/tree/b2de0cb)
- github.com/smartystreets/goconvey: [v1.6.4](https://github.com/smartystreets/goconvey/tree/v1.6.4)
- github.com/spaolacci/murmur3: [f09979e](https://github.com/spaolacci/murmur3/tree/f09979e)
- github.com/spf13/cast: [v1.3.1](https://github.com/spf13/cast/tree/v1.3.1)
- github.com/spf13/cobra: [v1.2.1](https://github.com/spf13/cobra/tree/v1.2.1)
- github.com/spf13/jwalterweatherman: [v1.1.0](https://github.com/spf13/jwalterweatherman/tree/v1.1.0)
- github.com/spf13/viper: [v1.8.1](https://github.com/spf13/viper/tree/v1.8.1)
- github.com/subosito/gotenv: [v1.2.0](https://github.com/subosito/gotenv/tree/v1.2.0)
- go.etcd.io/etcd/api/v3: v3.5.0
- go.etcd.io/etcd/client/pkg/v3: v3.5.0
- go.etcd.io/etcd/client/v2: v2.305.0
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.20.0
- go.opentelemetry.io/contrib: v0.20.0
- go.opentelemetry.io/otel/exporters/otlp: v0.20.0
- go.opentelemetry.io/otel/metric: v0.20.0
- go.opentelemetry.io/otel/oteltest: v0.20.0
- go.opentelemetry.io/otel/sdk/export/metric: v0.20.0
- go.opentelemetry.io/otel/sdk/metric: v0.20.0
- go.opentelemetry.io/otel/sdk: v0.20.0
- go.opentelemetry.io/otel/trace: v0.20.0
- go.opentelemetry.io/otel: v0.20.0
- go.opentelemetry.io/proto/otlp: v0.7.0
- go.uber.org/atomic: v1.7.0
- go.uber.org/goleak: v1.1.10
- go.uber.org/multierr: v1.6.0
- go.uber.org/zap: v1.19.0
- gopkg.in/alecthomas/kingpin.v2: v2.2.6
- gopkg.in/ini.v1: v1.62.0
- gotest.tools/v3: v3.0.3
- k8s.io/component-base: v0.23.5

### Changed
- github.com/creack/pty: [v1.1.9 → v1.1.11](https://github.com/creack/pty/compare/v1.1.9...v1.1.11)
- github.com/envoyproxy/go-control-plane: [fd9021f → 63b5d3c](https://github.com/envoyproxy/go-control-plane/compare/fd9021f...63b5d3c)
- github.com/mitchellh/mapstructure: [v1.1.2 → v1.4.1](https://github.com/mitchellh/mapstructure/compare/v1.1.2...v1.4.1)
- github.com/prometheus/client_model: [14fe0d1 → v0.2.0](https://github.com/prometheus/client_model/compare/14fe0d1...v0.2.0)
- github.com/spf13/afero: [v1.2.2 → v1.6.0](https://github.com/spf13/afero/compare/v1.2.2...v1.6.0)
- github.com/stretchr/objx: [v0.1.0 → v0.1.1](https://github.com/stretchr/objx/compare/v0.1.0...v0.1.1)
- github.com/yuin/goldmark: [v1.3.5 → v1.4.0](https://github.com/yuin/goldmark/compare/v1.3.5...v1.4.0)
- golang.org/x/lint: 83fdc39 → 6edffad
- golang.org/x/tools: v0.1.5 → d4cc65f
- google.golang.org/api: v0.43.0 → v0.44.0
- google.golang.org/genproto: 6c239bb → fe13028
- google.golang.org/grpc: v1.36.1 → v1.40.0
- k8s.io/api: v0.23.4 → v0.23.5
- k8s.io/apimachinery: v0.23.4 → v0.23.5
- k8s.io/client-go: v0.23.4 → v0.23.5

### Removed
_Nothing has changed._
