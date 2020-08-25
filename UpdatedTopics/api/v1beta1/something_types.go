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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SomethingSpec defines the desired state of KafkaBundle
type SomethingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Kafka          *Kafka          `json:"kafka"`
	Zookeeper      *Zookeeper      `json:"zookeeper"`
	EntityOperator *EntityOperator `json:"entityOperator"`
	// Foo is an example field of Kafka. Edit Kafka_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// SomethingStatus defines the observed state of KafkaBundle
type SomethingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// Kafka - dass
type Kafka struct {
	Version   string     `json:"version"`
	Replicas  int        `json:"replicas"`
	Listeners *Listeners `json:"listeners"`
	Config    *Config    `json:"config"`
	Storage   *Storage   `json:"storage"`
}

type Listeners struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Plain    Plain  `json:"plain"`
	Tls      Tls    `json:"tls"`
}

type Plain struct {
	Authentiation Authentiation `json:"authentiation"`
}

type Tls struct {
	Authentiation Authentiation `json:"authentiation"`
}

type Authentiation struct {
	Type string `json:"type"`
}

type Config struct {
	OffsetsTopicReplicationFactor        int    `json:"offsets.topic.replication.factor"`
	TransactionStateLogReplicationFactor int    `json:"transaction.state.log.replication.factor"`
	TransactionStateLogMinIsr            int    `json:"transaction.state.log.min.isr"`
	LogMessageFormatVersion              string `json:"log.message.format.version"`
}

type Zookeeper struct {
	Replicas int     `json:"replicas"`
	Storage  Storage `json:"storage"`
}

type Storage struct {
	Type string `json:"type"`
}

type EntityOperator struct {
	TopicOperator TopicOperator `json:"topicOperator"`
	UserOperator  UserOperator  `json:"userOperator"`
}

type TopicOperator struct {
	ReconciliationIntervalSeconds int `json:"reconciliationIntervalSeconds"`
}
type UserOperator struct {
	ReconciliationIntervalSeconds int `json:"reconciliationIntervalSeconds"`
}

// // SomethingStatus defines the observed state of Something
// type SomethingStatus struct {
// 	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
// 	// Important: Run "make" to regenerate code after modifying this file
// }

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Something is the Schema for the somethings API
type Something struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SomethingSpec   `json:"spec,omitempty"`
	Status SomethingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SomethingList contains a list of Something
type SomethingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Something `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Something{}, &SomethingList{})
}
