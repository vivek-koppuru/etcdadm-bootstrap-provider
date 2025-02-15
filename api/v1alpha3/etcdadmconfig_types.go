/*


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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	capbk "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha3"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
const (
	DataSecretAvailableCondition clusterv1.ConditionType = "DataSecretAvailable"
	// CloudConfig make the bootstrap data to be of cloud-config format.
	CloudConfig Format = "cloud-config"
	// Bottlerocket make the bootstrap data to be of bottlerocket format.
	Bottlerocket Format = "bottlerocket"
)

// Format specifies the output format of the bootstrap data
// +kubebuilder:validation:Enum=cloud-config;bottlerocket
type Format string

// EtcdadmConfigSpec defines the desired state of EtcdadmConfig
type EtcdadmConfigSpec struct {
	// Users specifies extra users to add
	// +optional
	Users []capbk.User `json:"users,omitempty"`

	// +optional
	EtcdadmBuiltin bool `json:"etcdadmBuiltin,omitempty"`

	// +optional
	EtcdadmInstallCommands []string `json:"etcdadmInstallCommands,omitempty"`

	// PreEtcdadmCommands specifies extra commands to run before kubeadm runs
	// +optional
	PreEtcdadmCommands []string `json:"preEtcdadmCommands,omitempty"`

	// PostEtcdadmCommands specifies extra commands to run after kubeadm runs
	// +optional
	PostEtcdadmCommands []string `json:"postEtcdadmCommands,omitempty"`

	// Format specifies the output format of the bootstrap data
	// +optional
	Format Format `json:"format,omitempty"`

	// BottlerocketConfig specifies the configuration for the bottlerocket bootstrap data
	// +optional
	BottlerocketConfig *BottlerocketConfig `json:"bottlerocketConfig,omitempty"`

	// CloudInitConfig specifies the configuration for the cloud-init bootstrap data
	// +optional
	CloudInitConfig *CloudInitConfig `json:"cloudInitConfig,omitempty"`

	// Files specifies extra files to be passed to user_data upon creation.
	// +optional
	Files []capbk.File `json:"files,omitempty"`

	// Proxy holds the https and no proxy information
	// This is only used for bottlerocket
	// +optional
	Proxy *ProxyConfiguration `json:"proxy,omitempty"`

	// RegistryMirror holds the image registry mirror information
	// This is only used for bottlerocket
	// +optional
	RegistryMirror *RegistryMirrorConfiguration `json:"registryMirror,omitempty"`
}

type BottlerocketConfig struct {
	// EtcdImage specifies the etcd image to use by etcdadm
	EtcdImage string `json:"etcdImage,omitempty"`

	// BootstrapImage specifies the container image to use for bottlerocket's bootstrapping
	BootstrapImage string `json:"bootstrapImage"`

	// PauseImage specifies the image to use for the pause container
	PauseImage string `json:"pauseImage"`
}

type CloudInitConfig struct {
	// +optional
	Version string `json:"version,omitempty"`

	// EtcdReleaseURL is an optional field to specify where etcdadm can download etcd from
	// +optional
	EtcdReleaseURL string `json:"etcdReleaseURL,omitempty"`

	// InstallDir is an optional field to specify where etcdadm will extract etcd binaries to
	// +optional
	InstallDir string `json:"installDir,omitempty"`
}

// ProxyConfiguration holds the settings for proxying bottlerocket services
type ProxyConfiguration struct {
	// HTTP Proxy
	HTTPProxy string `json:"httpProxy,omitempty"`

	// HTTPS proxy
	HTTPSProxy string `json:"httpsProxy,omitempty"`

	// No proxy, list of ips that should not use proxy
	NoProxy []string `json:"noProxy,omitempty"`
}

// RegistryMirrorConfiguration holds the settings for image registry mirror
type RegistryMirrorConfiguration struct {
	// Endpoint defines the registry mirror endpoint to use for pulling images
	Endpoint string `json:"endpoint,omitempty"`

	// CACert defines the CA cert for the registry mirror
	CACert string `json:"caCert,omitempty"`
}

// EtcdadmConfigStatus defines the observed state of EtcdadmConfig
type EtcdadmConfigStatus struct {
	// Conditions defines current service state of the KubeadmConfig.
	// +optional
	Conditions clusterv1.Conditions `json:"conditions,omitempty"`

	DataSecretName *string `json:"dataSecretName,omitempty"`

	Ready bool `json:"ready,omitempty"`
}

func (c *EtcdadmConfig) GetConditions() clusterv1.Conditions {
	return c.Status.Conditions
}

func (c *EtcdadmConfig) SetConditions(conditions clusterv1.Conditions) {
	c.Status.Conditions = conditions
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// EtcdadmConfig is the Schema for the etcdadmconfigs API
type EtcdadmConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdadmConfigSpec   `json:"spec,omitempty"`
	Status EtcdadmConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EtcdadmConfigList contains a list of EtcdadmConfig
type EtcdadmConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EtcdadmConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EtcdadmConfig{}, &EtcdadmConfigList{})
}
