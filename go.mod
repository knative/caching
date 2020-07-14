module knative.dev/caching

go 1.14

require (
	github.com/c2h5oh/datasize v0.0.0-20200112174442-28bbd4740fee // indirect
	github.com/dgryski/go-lttb v0.0.0-20180810165845-318fcdf10a77 // indirect
	github.com/google/go-cmp v0.5.0
	github.com/google/licenseclassifier v0.0.0-20200708223521-3d09a0ea2f39
	github.com/influxdata/tdigest v0.0.1 // indirect
	github.com/miekg/dns v1.1.29 // indirect
	github.com/tsenart/go-tsz v0.0.0-20180814235614-0bd30b3df1c3 // indirect
	github.com/tsenart/vegeta v12.7.1-0.20190725001342-b5f4fca92137+incompatible
	k8s.io/api v0.17.6
	k8s.io/apimachinery v0.18.5
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.18.0
	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29
	knative.dev/pkg v0.0.0-20200714070918-ac02cac99b88
	knative.dev/test-infra v0.0.0-20200713220518-5a4c4cad5372
)

replace (
	k8s.io/api => k8s.io/api v0.17.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/apiserver => k8s.io/apiserver v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
	k8s.io/code-generator => k8s.io/code-generator v0.17.6
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29
)
