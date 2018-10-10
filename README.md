# Knative Caching Resources

[![GoDoc](https://godoc.org/github.com/knative/caching?status.svg)](https://godoc.org/github.com/knative/caching)
[![Go Report Card](https://goreportcard.com/badge/knative/caching)](https://goreportcard.com/report/knative/caching)

Knative `caching` defines resources that can be used to express a desire to
cache things.  These are **just** API definitions for caching custom resources
plugins, it does not include an implementation of this API.

The API definitions here enable loose coupling via [duck
typing](https://docs.google.com/document/d/16j8C91jML4fQRQPhnHihNJUJDcbvW0RM1YAX2REHgyY/edit)
in Knative APIs. This loose coupling allows Knative APIs to work correctly with
arbitrary resources that conform to a partial interface signature.

## Example

The `Image` duck type captures the intent have a container image cached. There
are multiple Knative APIs that express this intent: `BuildTemplate` and
`ClusterBuildTemplate`.

```go
cachingClient, err := cachingclientset.NewForConfig(cfg)
if err != nil {
    // ...
}

cachingInformerFactory := cachinginformers.NewSharedInformerFactory(cachingClient, time.Second*30)
imageInformer := cachingInformerFactory.Caching().V1alpha1().Images()

imageInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
    FilterFunc: controller.Filter(v1alpha1.SchemeGroupVersion.WithKind("BuildTemplate")),
    Handler: cache.ResourceEventHandlerFuncs{
        AddFunc:    impl.EnqueueControllerOf,
        UpdateFunc: controller.PassNew(impl.EnqueueControllerOf),
    },
})
```

To learn more about Knative, please visit our
[Knative docs](https://github.com/knative/docs) repository.

If you are interested in contributing, see [CONTRIBUTING.md](./CONTRIBUTING.md)
and [DEVELOPMENT.md](./DEVELOPMENT.md).
