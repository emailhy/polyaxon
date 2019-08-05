/*
Copyright 2019 Polyaxon, Inc.

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

package utils

import (
	"reflect"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// DefaultBackoffLimit for PlxJobs
	DefaultBackoffLimit = 1
	// DefaultRestartPolicy for PlxJobs
	DefaultRestartPolicy = "Never"
)

// CopyJobFields copies the owned fields from one Job to another
// Returns true if the fields copied from don't match to.
func CopyJobFields(from, to *batchv1.Job) bool {
	requireUpdate := false
	for k, v := range to.Labels {
		if from.Labels[k] != v {
			requireUpdate = true
		}
	}
	to.Labels = from.Labels

	if to.Spec.ActiveDeadlineSeconds != from.Spec.ActiveDeadlineSeconds {
		to.Spec.ActiveDeadlineSeconds = from.Spec.ActiveDeadlineSeconds
		requireUpdate = true
	}

	if to.Spec.BackoffLimit != from.Spec.BackoffLimit {
		to.Spec.BackoffLimit = from.Spec.BackoffLimit
		requireUpdate = true
	}

	if to.Spec.TTLSecondsAfterFinished != from.Spec.TTLSecondsAfterFinished {
		to.Spec.TTLSecondsAfterFinished = from.Spec.TTLSecondsAfterFinished
		requireUpdate = true
	}

	if !reflect.DeepEqual(to.Spec.Template.Spec, from.Spec.Template.Spec) {
		requireUpdate = true
		to.Spec.Template.Spec = from.Spec.Template.Spec
	}

	return requireUpdate
}

// IsPlxJobSucceded return true if job is running
func IsPlxJobSucceded(jc batchv1.JobCondition) bool {
	return jc.Type == batchv1.JobComplete && jc.Status == corev1.ConditionTrue
}

// IsPlxJobFailed return true if job is running
func IsPlxJobFailed(jc batchv1.JobCondition) bool {
	return jc.Type == batchv1.JobFailed && jc.Status == corev1.ConditionTrue
}

// GeneratePlxJob returns a a batch job given a PolyaxonBaseJobSpec
func GeneratePlxJob(
	name string,
	namespace string,
	labels map[string]string,
	backoffLimit *int32,
	podSpec corev1.PodSpec,
) *batchv1.Job {
	jobBackoffLimit := int32(DefaultBackoffLimit)
	if backoffLimit != nil {
		jobBackoffLimit = *backoffLimit
	}

	if podSpec.RestartPolicy == "" {
		podSpec.RestartPolicy = DefaultRestartPolicy
	}

	PlxJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: &jobBackoffLimit,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{}},
				Spec:       podSpec,
			},
		},
	}
	// copy all of the labels to the pod including poddefault related labels
	l := &PlxJob.Spec.Template.ObjectMeta.Labels
	for k, v := range labels {
		(*l)[k] = v
	}

	return PlxJob
}
