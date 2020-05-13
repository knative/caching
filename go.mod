module knative.dev/caching

go 1.13

require (
	github.com/c2h5oh/datasize v0.0.0-20200112174442-28bbd4740fee // indirect
	github.com/dgryski/go-lttb v0.0.0-20180810165845-318fcdf10a77 // indirect
	github.com/google/go-cmp v0.4.0
	github.com/google/licenseclassifier v0.0.0-20200402202327-879cb1424de0
	github.com/influxdata/tdigest v0.0.1 // indirect
	github.com/miekg/dns v1.1.29 // indirect
	github.com/tsenart/go-tsz v0.0.0-20180814235614-0bd30b3df1c3 // indirect
	github.com/tsenart/vegeta v12.7.1-0.20190725001342-b5f4fca92137+incompatible
	k8s.io/api v0.17.3
	k8s.io/apimachinery v0.17.3
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.18.0
	k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a
	knative.dev/pkg v0.0.0-20200512191130-b10849aacf82
	knative.dev/test-infra v0.0.0-20200513011557-d03429a76034
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
