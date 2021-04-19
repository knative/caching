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
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

func TestIsReady(t *testing.T) {
	cases := []struct {
		name       string
		generation int64
		status     ImageStatus
		isReady    bool
	}{{
		name:    "empty status should not be ready",
		status:  ImageStatus{},
		isReady: false,
	}, {
		name:       "Different condition type should not be ready",
		generation: 0,
		status: ImageStatus{
			Status: duckv1.Status{
				ObservedGeneration: 0,
				Conditions: duckv1.Conditions{{
					Type:   "foo",
					Status: corev1.ConditionTrue,
				},
				}},
		},
		isReady: false,
	}, {
		name:       "False condition status should not be ready",
		generation: 0,
		status: ImageStatus{
			Status: duckv1.Status{
				ObservedGeneration: 0,
				Conditions: duckv1.Conditions{{
					Type:   ImageConditionReady,
					Status: corev1.ConditionFalse,
				}},
			},
		},
		isReady: false,
	}, {
		name:       "Unknown condition status should not be ready",
		generation: 0,
		status: ImageStatus{
			Status: duckv1.Status{
				ObservedGeneration: 0,
				Conditions: duckv1.Conditions{{
					Type:   ImageConditionReady,
					Status: corev1.ConditionUnknown,
				}},
			},
		},
		isReady: false,
	}, {
		name:       "Missing condition status should not be ready",
		generation: 0,
		status: ImageStatus{
			Status: duckv1.Status{
				ObservedGeneration: 0,
				Conditions: duckv1.Conditions{{
					Type: ImageConditionReady,
				}},
			},
		},
		isReady: false,
	}, {
		name:       "True condition status should be ready",
		generation: 0,
		status: ImageStatus{
			Status: duckv1.Status{
				ObservedGeneration: 0,
				Conditions: duckv1.Conditions{{
					Type:   ImageConditionReady,
					Status: corev1.ConditionTrue,
				}},
			},
		},
		isReady: true,
	}, {
		name:       "Multiple conditions with ready status should be ready",
		generation: 0,
		status: ImageStatus{
			Status: duckv1.Status{
				ObservedGeneration: 0,
				Conditions: duckv1.Conditions{{
					Type:   "foo",
					Status: corev1.ConditionTrue,
				}, {
					Type:   ImageConditionReady,
					Status: corev1.ConditionTrue,
				}},
			},
		},
		isReady: true,
	}, {
		name:       "Multiple conditions with ready status false should not be ready",
		generation: 0,
		status: ImageStatus{
			Status: duckv1.Status{
				ObservedGeneration: 0,
				Conditions: duckv1.Conditions{{
					Type:   "foo",
					Status: corev1.ConditionTrue,
				}, {
					Type:   ImageConditionReady,
					Status: corev1.ConditionFalse,
				}},
			},
		},
		isReady: false,
	}, {
		name:       "Generation not equal ObservedGeneration should not be ready",
		generation: 1,
		status: ImageStatus{
			Status: duckv1.Status{
				ObservedGeneration: 2,
				Conditions: duckv1.Conditions{{
					Type:   "foo",
					Status: corev1.ConditionTrue,
				}, {
					Type:   ImageConditionReady,
					Status: corev1.ConditionFalse,
				}},
			},
		},
		isReady: false,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := Image{
				ObjectMeta: metav1.ObjectMeta{
					Generation: tc.generation,
				},
				Status: tc.status}
			if e, a := tc.isReady, m.IsReady(); e != a {
				t.Errorf("Ready = %v, want: %v", a, e)
			}
		})
	}
}
