/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and secret-generator-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	"github.com/sap/component-operator-runtime/pkg/component"
	componentoperatorruntimetypes "github.com/sap/component-operator-runtime/pkg/types"
)

// SecretGeneratorSpec defines the desired state of SecretGenerator.
type SecretGeneratorSpec struct {
	component.Spec `json:",inline"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1
	ReplicaCount int `json:"replicaCount,omitempty"`
	// +optional
	Image                          component.ImageSpec `json:"image"`
	component.KubernetesProperties `json:",inline"`
	ObjectSelector                 *metav1.LabelSelector `json:"objectSelector,omitempty"`
	NamespaceSelector              *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
	LogLevel                       int                   `json:"logLevel,omitempty"`
}

// SecretGeneratorStatus defines the observed state of SecretGenerator.
type SecretGeneratorStatus struct {
	component.Status `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +genclient

// SecretGenerator is the Schema for the secretgenerators API.
type SecretGenerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SecretGeneratorSpec `json:"spec,omitempty"`
	// +kubebuilder:default={"observedGeneration":-1}
	Status SecretGeneratorStatus `json:"status,omitempty"`
}

var _ component.Component = &SecretGenerator{}

// +kubebuilder:object:root=true

// SecretGeneratorList contains a list of SecretGenerator.
type SecretGeneratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretGenerator `json:"items"`
}

func (s *SecretGeneratorSpec) ToUnstructured() map[string]any {
	result, err := runtime.DefaultUnstructuredConverter.ToUnstructured(s)
	if err != nil {
		panic(err)
	}
	return result
}

func (c *SecretGenerator) GetDeploymentNamespace() string {
	if c.Spec.Namespace != "" {
		return c.Spec.Namespace
	}
	return c.Namespace
}

func (c *SecretGenerator) GetDeploymentName() string {
	if c.Spec.Name != "" {
		return c.Spec.Name
	}
	return c.Name
}

func (c *SecretGenerator) GetSpec() componentoperatorruntimetypes.Unstructurable {
	return &c.Spec
}

func (c *SecretGenerator) GetStatus() *component.Status {
	return &c.Status.Status
}

func init() {
	SchemeBuilder.Register(&SecretGenerator{}, &SecretGeneratorList{})
}
