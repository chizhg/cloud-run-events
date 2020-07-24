/*
Copyright 2020 Google LLC.

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

package v1beta1

import (
	"time"

	gcpauthtesthelper "github.com/google/knative-gcp/pkg/apis/configs/gcpauth/testhelper"
	"github.com/google/knative-gcp/pkg/reconciler/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/google/knative-gcp/pkg/apis/events/v1beta1"
)

// CloudPubSubSourceOption enables further configuration of a CloudPubSubSource.
type CloudPubSubSourceOption func(*v1beta1.CloudPubSubSource)

// NewCloudPubSubSource creates a CloudPubSubSource with CloudPubSubSourceOptions
func NewCloudPubSubSource(name, namespace string, so ...CloudPubSubSourceOption) *v1beta1.CloudPubSubSource {
	ps := &v1beta1.CloudPubSubSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			UID:       "test-pubsub-uid",
		},
	}
	for _, opt := range so {
		opt(ps)
	}
	return ps
}

func WithCloudPubSubSourceSink(gvk metav1.GroupVersionKind, name string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Spec.Sink = duckv1.Destination{
			Ref: &duckv1.KReference{
				APIVersion: testing.ApiVersion(gvk),
				Kind:       gvk.Kind,
				Name:       name,
			},
		}
	}
}

func WithCloudPubSubSourceServiceAccount(kServiceAccount string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Spec.ServiceAccountName = kServiceAccount
	}
}

func WithCloudPubSubSourceDeletionTimestamp(s *v1beta1.CloudPubSubSource) {
	t := metav1.NewTime(time.Unix(1e9, 0))
	s.ObjectMeta.SetDeletionTimestamp(&t)
}

func WithCloudPubSubSourceProject(project string) CloudPubSubSourceOption {
	return func(s *v1beta1.CloudPubSubSource) {
		s.Spec.Project = project
	}
}

func WithCloudPubSubSourceTopic(topicID string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Spec.Topic = topicID
	}
}

// WithInitCloudPubSubSourceConditions initializes the CloudPubSubSource's conditions.
func WithInitCloudPubSubSourceConditions(ps *v1beta1.CloudPubSubSource) {
	ps.Status.InitializeConditions()
}

// WithCloudPubSubSourceServiceAccountName will give status.ServiceAccountName a k8s service account name, which is related on Workload Identity's Google service account.
func WithCloudPubSubSourceServiceAccountName(name string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Status.ServiceAccountName = name
	}
}

func WithCloudPubSubSourceWorkloadIdentityFailed(reason, message string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Status.MarkWorkloadIdentityFailed(ps.ConditionSet(), reason, message)
	}
}

// WithCloudPubSubSourcePullSubscriptionFailed marks the condition that the
// status of PullSubscription is False
func WithCloudPubSubSourcePullSubscriptionFailed(reason, message string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Status.MarkPullSubscriptionFailed(ps.ConditionSet(), reason, message)
	}
}

// WithCloudPubSubSourcePullSubscriptionUnknown marks the condition that the
// topic is Unknown
func WithCloudPubSubSourcePullSubscriptionUnknown(reason, message string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Status.MarkPullSubscriptionUnknown(ps.ConditionSet(), reason, message)
	}
}

// WithCloudPubSubSourcePullSubscriptionReady marks the condition that the
// topic is not ready
func WithCloudPubSubSourcePullSubscriptionReady(ps *v1beta1.CloudPubSubSource) {
	ps.Status.MarkPullSubscriptionReady(ps.ConditionSet())
}

// WithCloudPubSubSourceSinkURI sets the status for sink URI
func WithCloudPubSubSourceSinkURI(url *apis.URL) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Status.SinkURI = url
	}
}

func WithCloudPubSubSourceSubscriptionID(subscriptionID string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Status.SubscriptionID = subscriptionID
	}
}

func WithCloudPubSubSourceFinalizers(finalizers ...string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Finalizers = finalizers
	}
}

func WithCloudPubSubSourceStatusObservedGeneration(generation int64) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.Status.Status.ObservedGeneration = generation
	}
}

func WithCloudPubSubSourceObjectMetaGeneration(generation int64) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.ObjectMeta.Generation = generation
	}
}

func WithCloudPubSubSourceAnnotations(Annotations map[string]string) CloudPubSubSourceOption {
	return func(ps *v1beta1.CloudPubSubSource) {
		ps.ObjectMeta.Annotations = Annotations
	}
}

func WithCloudPubSubSourceSetDefaults(ps *v1beta1.CloudPubSubSource) {
	ps.SetDefaults(gcpauthtesthelper.ContextWithDefaults())
}