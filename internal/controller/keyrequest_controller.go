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

package controller

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	quantummanageriov1 "github.com/ShubhamTatvamasi/quantum-manager/api/v1"
	oqsrand "github.com/open-quantum-safe/liboqs-go/oqs/rand"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KeyRequestReconciler reconciles a KeyRequest object
type KeyRequestReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=quantum-manager.io,resources=keyrequests,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=quantum-manager.io,resources=keyrequests/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=quantum-manager.io,resources=keyrequests/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KeyRequest object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *KeyRequestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// TODO(user): your logic here
	log := log.FromContext(ctx)

	keyrequest := &quantummanageriov1.KeyRequest{}
	err := r.Get(ctx, req.NamespacedName, keyrequest)
	if err != nil {
		log.Error(err, "Failed to get KeyRequest")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if err = r.CreateSecret(keyrequest, ctx); err != nil {
		log.Error(err, "Failed to Create or Update Secret")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeyRequestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&quantummanageriov1.KeyRequest{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}

func (r *KeyRequestReconciler) CreateSecret(keyrequest *quantummanageriov1.KeyRequest, ctx context.Context) error {
	log := log.FromContext(ctx)

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyrequest.Name,
			Namespace: keyrequest.Namespace,
		},
	}

	// existingSecret := &corev1.Secret{}
	err := r.Get(ctx, client.ObjectKey{Namespace: secret.Namespace, Name: secret.Name}, secret)
	if err != nil && client.IgnoreNotFound(err) != nil {
		return err
	}

	// If Secret doesn't exist, create itß
	if err != nil {

		randomNumber := oqsrand.RandomBytes(keyrequest.Spec.Bytes)

		// delete this in future
		fmt.Println(randomNumber)

		newSecret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      keyrequest.Name,
				Namespace: keyrequest.Namespace,
			},
			StringData: map[string]string{
				"random-number": string(randomNumber),
			},
		}

		if err := ctrl.SetControllerReference(keyrequest, newSecret, r.Scheme); err != nil {
			log.Error(err, "Failed to Set Controller Reference")
			return err
		}

		err = r.Create(ctx, newSecret)
		if err != nil {
			return err
		}
		log.Info("Created Secret")
	}

	return nil
}
