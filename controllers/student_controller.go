/*
Copyright 2021.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	esdevopscomv1 "github.com/TangSmith31/school-controller-example/api/v1"
)

// StudentReconciler reconciles a Student object
type StudentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=esdevops.com,resources=students,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=esdevops.com,resources=students/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=esdevops.com,resources=students/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Student object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile

var SchoolCRName = types.NamespacedName{
	Namespace: "default",
	Name:      "beijing-university",
}

func (r *StudentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// Get All Events: OnAdd/OnUpdate/OnDelete

	// Step-1: Get Student CR
	var StudentCR = esdevopscomv1.Student{}
	if r.Get(ctx, req.NamespacedName, &StudentCR) != nil {
		klog.Errorf("Error: Get Student %s %s", req.NamespacedName)
		return ctrl.Result{}, nil
	}

	// Step-2: Get School CR
	var SchoolCR = esdevopscomv1.School{}
	if r.Get(ctx, SchoolCRName, &SchoolCR) != nil {
		klog.Errorf("Error: Get School %s", SchoolCRName)
		return ctrl.Result{}, nil
	}

	// Step-3: Check Student Card
	if !StudentCR.Status.StudentCard {
		// New Student is coming
		// Step-3-1: Provide Student Card
		patch := client.MergeFrom(StudentCR.DeepCopy())
		StudentCR.Status.StudentCard = true

		error := r.Status().Patch(ctx, &StudentCR, patch)
		if error != nil {
			klog.Errorf("%s", error)
			return ctrl.Result{}, nil
		}

		// Step-3-2: Add Student Number into School
		SchoolDeepCopy := SchoolCR.DeepCopy()
		SchoolDeepCopy.Spec.StudentNumber++
		if r.Patch(ctx, SchoolDeepCopy, client.MergeFrom(&SchoolCR)) != nil {
			klog.Errorf("Error: Patch School Student Number")
			return ctrl.Result{}, nil
		}
		klog.Infof(
			"Student %s had been registered, School Student Number is [%d] Now",
			StudentCR.Spec.Name,
			SchoolDeepCopy.Spec.StudentNumber,
		)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *StudentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&esdevopscomv1.Student{}).
		Complete(r)
}
