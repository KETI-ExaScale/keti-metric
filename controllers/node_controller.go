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

package controllers

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ketimetricv1beta1 "github.com/KETI-ExaScale/keti-metric/api/v1beta1"
)

// NodeReconciler reconciles a Node object
type NodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=keti-metric.keti-exascale.keti-metric,resources=nodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=keti-metric.keti-exascale.keti-metric,resources=nodes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=keti-metric.keti-exascale.keti-metric,resources=nodes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Node object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *NodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	nodeMetricList := &ketimetricv1beta1.NodeList{}
	err := r.List(ctx, nodeMetricList)
	if err != nil {
		klog.Errorln(err)
	}
	minTime := time.Now()
	minObj := &ketimetricv1beta1.Node{}
	if len(nodeMetricList.Items) > 30 {
		for _, metric := range nodeMetricList.Items {
			if !minTime.Before(metric.CreationTimestamp.Time) {
				minTime = metric.CreationTimestamp.Time
				minObj = &metric
			}
		}
		err = r.Delete(ctx, minObj, client.GracePeriodSeconds(1))
		if err != nil {
			klog.Errorln(err)
		}
	}
	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ketimetricv1beta1.Node{}).
		Complete(r)
}
