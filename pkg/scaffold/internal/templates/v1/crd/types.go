/*
Copyright 2018 The Kubernetes Authors.

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

package crd

import (
	"fmt"
	"path/filepath"
	"strings"

	"sigs.k8s.io/kubebuilder/pkg/model/file"
	"sigs.k8s.io/kubebuilder/pkg/model/resource"
)

var _ file.Template = &Types{}

// Types scaffolds the pkg/apis/group/version/kind_types.go file to define the schema for an API
type Types struct {
	file.Input

	// Resource is the resource to scaffold the types_test.go file for
	Resource *resource.Resource
}

// GetInput implements input.Template
func (f *Types) GetInput() (file.Input, error) {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "apis", f.Resource.GroupPackageName, f.Resource.Version,
			fmt.Sprintf("%s_types.go", strings.ToLower(f.Resource.Kind)))
	}
	f.TemplateBody = typesTemplate
	f.IfExistsAction = file.Error
	return f.Input, nil
}

// Validate validates the values
func (f *Types) Validate() error {
	return f.Resource.Validate()
}

const typesTemplate = `{{ .Boilerplate }}

package {{ .Resource.Version }}

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// {{ .Resource.Kind }}Spec defines the desired state of {{ .Resource.Kind }}
type {{ .Resource.Kind }}Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// {{ .Resource.Kind }}Status defines the observed state of {{ .Resource.Kind }}
type {{ .Resource.Kind }}Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
{{- if not .Resource.Namespaced }}
// +genclient:nonNamespaced
{{- end }}

// {{ .Resource.Kind }} is the Schema for the {{ .Resource.Plural }} API
// +k8s:openapi-gen=true
type {{ .Resource.Kind }} struct {
	metav1.TypeMeta   ` + "`" + `json:",inline"` + "`" + `
	metav1.ObjectMeta ` + "`" + `json:"metadata,omitempty"` + "`" + `

	Spec   {{ .Resource.Kind }}Spec   ` + "`" + `json:"spec,omitempty"` + "`" + `
	Status {{ .Resource.Kind }}Status ` + "`" + `json:"status,omitempty"` + "`" + `
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
{{- if not .Resource.Namespaced }}
// +genclient:nonNamespaced
{{- end }}

// {{ .Resource.Kind }}List contains a list of {{ .Resource.Kind }}
type {{ .Resource.Kind }}List struct {
	metav1.TypeMeta ` + "`" + `json:",inline"` + "`" + `
	metav1.ListMeta ` + "`" + `json:"metadata,omitempty"` + "`" + `
	Items           []{{ .Resource.Kind }} ` + "`" + `json:"items"` + "`" + `
}

func init() {
	SchemeBuilder.Register(&{{ .Resource.Kind }}{}, &{{ .Resource.Kind }}List{})
}
`
