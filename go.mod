module sigs.k8s.io/kind

// NOTE: This is the go language version, NOT the compiler version.
//
// This controls the *minimum* required go version and therefore available Go
// language features.
//
// See ./.go-version for the go compiler version used when building binaries
//
// https://go.dev/doc/modules/gomod-ref#go
go 1.22.0

toolchain go1.22.5

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/alessio/shellescape v1.4.2
	github.com/evanphx/json-patch/v5 v5.6.0
	github.com/google/safetext v0.0.0-20220905092116-b49f7bc46da2
	github.com/mattn/go-isatty v0.0.20
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	gopkg.in/yaml.v3 v3.0.1
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/vishvananda/netns v0.0.0-20191106174202-0a2b9b5464df // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apimachinery v0.30.3 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/vishvananda/netlink v1.1.0
	golang.org/x/sys v0.18.0 // indirect
	k8s.io/api v0.30.3
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8
)
