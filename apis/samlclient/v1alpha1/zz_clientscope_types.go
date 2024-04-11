// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClientScopeInitParameters struct {

	// When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
	ConsentScreenText *string `json:"consentScreenText,omitempty" tf:"consent_screen_text,omitempty"`

	// The description of this client scope in the GUI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specify order of the client scope in GUI (such as in Consent page) as integer.
	GuiOrder *float64 `json:"guiOrder,omitempty" tf:"gui_order,omitempty"`

	// The display name of this client scope in the GUI.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm this client scope belongs to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type ClientScopeObservation struct {

	// When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
	ConsentScreenText *string `json:"consentScreenText,omitempty" tf:"consent_screen_text,omitempty"`

	// The description of this client scope in the GUI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specify order of the client scope in GUI (such as in Consent page) as integer.
	GuiOrder *float64 `json:"guiOrder,omitempty" tf:"gui_order,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The display name of this client scope in the GUI.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm this client scope belongs to.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type ClientScopeParameters struct {

	// When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
	// +kubebuilder:validation:Optional
	ConsentScreenText *string `json:"consentScreenText,omitempty" tf:"consent_screen_text,omitempty"`

	// The description of this client scope in the GUI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specify order of the client scope in GUI (such as in Consent page) as integer.
	// +kubebuilder:validation:Optional
	GuiOrder *float64 `json:"guiOrder,omitempty" tf:"gui_order,omitempty"`

	// The display name of this client scope in the GUI.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm this client scope belongs to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// ClientScopeSpec defines the desired state of ClientScope
type ClientScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClientScopeParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClientScopeInitParameters `json:"initProvider,omitempty"`
}

// ClientScopeStatus defines the observed state of ClientScope.
type ClientScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClientScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClientScope is the Schema for the ClientScopes API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type ClientScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ClientScopeSpec   `json:"spec"`
	Status ClientScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientScopeList contains a list of ClientScopes
type ClientScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClientScope `json:"items"`
}

// Repository type metadata.
var (
	ClientScope_Kind             = "ClientScope"
	ClientScope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClientScope_Kind}.String()
	ClientScope_KindAPIVersion   = ClientScope_Kind + "." + CRDGroupVersion.String()
	ClientScope_GroupVersionKind = CRDGroupVersion.WithKind(ClientScope_Kind)
)

func init() {
	SchemeBuilder.Register(&ClientScope{}, &ClientScopeList{})
}
