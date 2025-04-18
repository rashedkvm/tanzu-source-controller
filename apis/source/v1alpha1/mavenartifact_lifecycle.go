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
	"reconciler.io/runtime/apis"
)

var (
	MavenArtifactLabelKey = GroupVersion.Group + "/maven-artifact"
)

const (
	MavenArtifactConditionReady             = apis.ConditionReady
	MavenArtifactConditionArtifactResolved  = "ArtifactVersionResolved"
	MavenArtifactConditionArtifactAvailable = "ArtifactAvailable"
)

var MavenArtifactCondSet = apis.NewLivingConditionSet(
	MavenArtifactConditionArtifactResolved,
	MavenArtifactConditionArtifactAvailable,
)

func (c *MavenArtifact) ManageConditions() apis.ConditionManager {
	return c.GetConditionSet().Manage(c.GetConditionsAccessor())
}

func (c *MavenArtifact) GetConditionsAccessor() apis.ConditionsAccessor {
	return &c.Status
}

func (c *MavenArtifact) GetConditionSet() apis.ConditionSet {
	return MavenArtifactCondSet
}

func (s *MavenArtifactStatus) InitializeConditions() {
	MavenArtifactCondSet.Manage(s).InitializeConditions()
}
