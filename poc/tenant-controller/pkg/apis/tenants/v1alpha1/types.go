// Copyright 2017 The Kubernetes Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Tenant is the resource represents a group of users belonging to the same tenant.
// A Tenant is a grouping concept of resources belong to a group of users (the tenant).
// Under a tenant, one or more namespaces are created.
// The OwerReferences in namespace resource will point to this Tenant resource, so
// once the Tenant resource is deleted, the namespaces will be garbage collected.
// Beyond this, the following labels are proposed to be associated with namespaces:
//     tenants.k8s.io/tenant=<name of Tenant resource>
type Tenant struct {
	metav1.TypeMeta `json:",inline"`
	// ObjectMeta is standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec is the details of the tenant.
	// +optional
	Spec TenantSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	// Status is the status of the tenant.
	// +optional
	Status TenantStatus `json:"status" protobuf:"bytes,3,opt,name=status"`
}

// TenantSpec defines the spec of a tenant resource.
// The Tenant controller will use the list of Namespaces here as the source of truth
// to reconciliate the actual namespaces belong to the tenant.
// Updating the namespace list here will trigger the reconciliation of namespaces.
type TenantSpec struct {
	// Admins are the identities with admin privilege in namespaces.
	// +optional
	Admins []rbacv1.Subject `json:"admins"`

	// Namespaces are the namespaces created for the tenant.
	// +optional
	Namespaces []TenantNamespace `json:"namespaces"`
}

// TenantStatus defines the status of a tenant resource.
type TenantStatus struct {
	// Phase indicates if the tenant is Pending, Creating, Active or Terminating.
	// +optional
	Phase TenantPhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=TenantPhase"`

	// Message provides human-readable information of current status.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`

	// Reason is a brief CamelCase string describing the status.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
}

// TenantNamespace defines the namespaces belonging to this tenant.
type TenantNamespace struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

// TenantPhase defines the phase of tenant status.
type TenantPhase string

// Known tenant phases.
const (
	// TenantPending means the tenant is going to be created, but not happening yet.
	// This is set right after the tenant is created.
	TenantPending TenantPhase = "Pending"
	// TenantCreating means tenant is being created.
	TenantCreating TenantPhase = "Creating"
	// TenantActive means tenant is ready and being used.
	TenantActive TenantPhase = "Active"
	// TenantTerminating means tenant is being removed.
	TenantTerminating TenantPhase = "Terminating"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	// ObjectMeta is standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items are list of Tenant objects.
	Items []Tenant `json:"items" protobuf:"bytes,2,rep,name=items`
}
