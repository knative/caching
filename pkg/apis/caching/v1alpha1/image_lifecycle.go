/*
Copyright 2021 The Knative Authors

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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"knative.dev/pkg/apis"
)

const (
	// ImageCacheConditionReady has status True when the SampleSource is ready to send events.
	ImageCacheConditionReady = apis.ConditionReady
)

var condSet = apis.NewLivingConditionSet()

// GetGroupVersionKind implements kmeta.OwnerRefable
func (i *Image) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Image")
}

// GetConditionSet retrieves the condition set for this resource. Implements the KRShaped interface.
func (i *Image) GetConditionSet() apis.ConditionSet {
	return condSet
}

// InitializeConditions sets the initial values to the conditions.
func (is *ImageStatus) InitializeConditions() {
	condSet.Manage(is).InitializeConditions()
}

// IsReady looks at the conditions and if the Status has a condition
// ImageConditionReady returns true if ConditionStatus is True
func (is *ImageStatus) IsReady() bool {
	if c := is.GetCondition(ImageCacheConditionReady); c != nil {
		return c.Status == corev1.ConditionTrue
	}
	return false
}

// GetStatus retrieves the status of the Image. Implements the KRShaped interface.
func (t *Image) GetStatus() *duckv1.Status {
	return &t.Status.Status
}
