module knative.dev/caching

go 1.13

require (
	github.com/c2h5oh/datasize v0.0.0-20200112174442-28bbd4740fee // indirect
	github.com/dgryski/go-lttb v0.0.0-20180810165845-318fcdf10a77 // indirect
	github.com/google/go-cmp v0.4.0
	github.com/google/licenseclassifier v0.0.0-20190103191631-c2a262e3078a
	github.com/influxdata/tdigest v0.0.1 // indirect
	github.com/miekg/dns v1.1.29 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/tsenart/go-tsz v0.0.0-20180814235614-0bd30b3df1c3 // indirect
	github.com/tsenart/vegeta v12.7.1-0.20190725001342-b5f4fca92137+incompatible
	go.uber.org/zap v1.10.0 // indirect
	k8s.io/api v0.16.4
	k8s.io/apimachinery v0.16.5-beta.1
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.18.0
	k8s.io/kube-openapi v0.0.0-20190816220812-743ec37842bf
	knative.dev/pkg v0.0.0-20200501005942-d980c0865972
	knative.dev/test-infra v0.0.0-20200430225942-f7c1fafc1cde
)

replace (
	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/apiserver => k8s.io/apiserver v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
	k8s.io/code-generator => k8s.io/code-generator v0.16.4
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190816220812-743ec37842bf
)
