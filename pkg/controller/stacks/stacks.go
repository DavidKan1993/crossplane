/*
Copyright 2019 The Crossplane Authors.

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

package stacks

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	"github.com/crossplane/crossplane/pkg/controller/stacks/install"
	"github.com/crossplane/crossplane/pkg/controller/stacks/persona"
	"github.com/crossplane/crossplane/pkg/controller/stacks/stack"
)

// Setup Crossplane Stacks controllers.
func Setup(mgr ctrl.Manager, l logging.Logger, hostControllerNamespace, tsControllerImage string, restrictCore bool) error {
	if err := install.SetupStackInstall(mgr, l, hostControllerNamespace, tsControllerImage); err != nil {
		return err
	}

	if err := install.SetupClusterStackInstall(mgr, l, hostControllerNamespace, tsControllerImage); err != nil {
		return err
	}

	if err := persona.Setup(mgr, l); err != nil {
		return nil
	}
	if err := stack.Setup(mgr, l, hostControllerNamespace, restrictCore); err != nil {
		return err
	}

	return nil
}
