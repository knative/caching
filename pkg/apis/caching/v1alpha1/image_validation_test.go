/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/pkg/apis"
)

func TestImageValidation(t *testing.T) {
	tests := []struct {
		name string
		r    *Image
		want *apis.FieldError
	}{{
		name: "valid",
		r: &Image{
			Spec: ImageSpec{
				Image: "busybox",
			},
		},
		want: nil,
	}, {
		name: "missing image",
		r:    &Image{},
		want: apis.ErrMissingField("spec.image"),
	}, {
		name: "nested spec error",
		r: &Image{
			Spec: ImageSpec{
				Image:            "busybox",
				ImagePullSecrets: []corev1.LocalObjectReference{{}},
			},
		},
		want: apis.ErrMissingField("spec.imagePullSecrets[0].name"),
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.r.Validate(context.Background())
			if (test.want == nil) != (got == nil) {
				t.Fatalf("Validate() = %v, wanted %v", got, test.want)
			}
			if want, got := test.want.Error(), got.Error(); want != got {
				t.Fatalf("Validate() = %v, wanted %v", got, want)
			}
		})
	}
}
