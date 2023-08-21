/*
Copyright 2023.

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

package v1

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var sumlog = logf.Log.WithName("sum-resource")

func (r *Sum) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-kalkulator-kalkulator-io-v1-sum,mutating=false,failurePolicy=fail,sideEffects=None,groups=kalkulator.kalkulator.io,resources=sums,verbs=create;update,versions=v1,name=vsum.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Sum{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Sum) ValidateCreate() (admission.Warnings, error) {
	klog.Infof("In validate create of %s", r.Name)
	if r.Spec.NumberOne < 0 || r.Spec.NumberTwo < 0 {
		return nil, fmt.Errorf("The input values Number One: %d Number Two: %d cannot be negative ", r.Spec.NumberOne, r.Spec.NumberTwo)
	}
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Sum) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	sumlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Sum) ValidateDelete() (admission.Warnings, error) {
	sumlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
