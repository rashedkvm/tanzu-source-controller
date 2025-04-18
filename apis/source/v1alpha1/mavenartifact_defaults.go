/*
Copyright 2022 VMware, Inc.

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

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// +kubebuilder:webhook:path=/mutate-source-apps-tanzu-vmware-com-v1alpha1-mavenartifact,mutating=true,failurePolicy=fail,sideEffects=none,admissionReviewVersions=v1beta1,groups=source.apps.tanzu.vmware.com,resources=mavenartifacts,verbs=create;update,versions=v1alpha1,name=mavenartifacts.source.apps.tanzu.vmware.com

type MavenArtifactDefaulter struct{}

var _ webhook.CustomDefaulter = &MavenArtifactDefaulter{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (c *MavenArtifactDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	mavenArtifact := obj.(*MavenArtifact)
	return mavenArtifact.Spec.Default()
}

func (s *MavenArtifactSpec) Default() error {
	s.Artifact.Default()
	s.Repository.Default()
	if s.Timeout == nil {
		s.Timeout = s.Interval.DeepCopy()
	}
	return nil
}

func (s *MavenArtifactType) Default() {
	if s.Type == "" {
		s.Type = "jar"
	}
}

func (s *Repository) Default() {
	// nothing to default
}
