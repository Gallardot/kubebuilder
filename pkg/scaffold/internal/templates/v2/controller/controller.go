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

package controller

import (
	"path/filepath"
	"strings"

	"sigs.k8s.io/kubebuilder/pkg/model/file"
	"sigs.k8s.io/kubebuilder/pkg/model/resource"
)

var _ file.Template = &Controller{}

// Controller scaffolds a Controller for a Resource
type Controller struct {
	file.Input

	// Resource is the Resource to make the Controller for
	Resource *resource.Resource
}

// GetInput implements input.Template
func (f *Controller) GetInput() (file.Input, error) {
	if f.Path == "" {
		if f.MultiGroup {
			f.Path = filepath.Join("controllers", f.Resource.Group,
				strings.ToLower(f.Resource.Kind)+"_controller.go")
		} else {
			f.Path = filepath.Join("controllers",
				strings.ToLower(f.Resource.Kind)+"_controller.go")
		}
	}
	f.TemplateBody = controllerTemplate

	f.Input.IfExistsAction = file.Error
	return f.Input, nil
}

// nolint:lll
const controllerTemplate = `{{ .Boilerplate }}

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	{{ .Resource.ImportAlias }} "{{ .Resource.Package }}"
)

// {{ .Resource.Kind }}Reconciler reconciles a {{ .Resource.Kind }} object
type {{ .Resource.Kind }}Reconciler struct {
	client.Client
	Log logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups={{ .Resource.Domain }},resources={{ .Resource.Plural }},verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups={{ .Resource.Domain }},resources={{ .Resource.Plural }}/status,verbs=get;update;patch

func (r *{{ .Resource.Kind }}Reconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("{{ .Resource.Kind | lower }}", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *{{ .Resource.Kind }}Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&{{ .Resource.ImportAlias }}.{{ .Resource.Kind }}{}).
		Complete(r)
}
`
