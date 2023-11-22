package oqs

import (
	"context"
	"fmt"

	quantummanageriov1 "github.com/ShubhamTatvamasi/quantum-manager/api/v1"
	oqsrand "github.com/open-quantum-safe/liboqs-go/oqs/rand"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GenerateRandomNumber(r client.Client, ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	randomNumber := oqsrand.RandomBytes(32)

	// delete this in future
	fmt.Println(randomNumber)

	keyrequest := &quantummanageriov1.KeyRequest{}
	err := r.Get(ctx, req.NamespacedName, keyrequest)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyrequest.Name,
			Namespace: keyrequest.Namespace,
		},
		StringData: map[string]string{
			"random-number": string(randomNumber),
		},
	}

	err = CreateOrUpdateSecret(r, ctx, secret)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}
