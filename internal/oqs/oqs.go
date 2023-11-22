package oqs

// import (
// 	"context"

// 	quantummanageriov1 "github.com/ShubhamTatvamasi/quantum-manager/api/v1"
// 	corev1 "k8s.io/api/core/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	ctrl "sigs.k8s.io/controller-runtime"
// 	"sigs.k8s.io/controller-runtime/pkg/client"
// )

// func GenerateSecret(r client.Client, ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

// 	// GenerateRandomNumber(r, 64)
// 	keyrequest := &quantummanageriov1.KeyRequest{}
// 	err := r.Get(ctx, req.NamespacedName, keyrequest)
// 	if err != nil {
// 		return ctrl.Result{}, client.IgnoreNotFound(err)
// 	}

// 	secret := &corev1.Secret{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      keyrequest.Name,
// 			Namespace: keyrequest.Namespace,
// 		},
// 		StringData: map[string]string{
// 			"some-key":  "myusername", //keyrequest.Spec.SomeKey,
// 			"other-key": "mypassword", //keyrequest.Spec.OtherKey,
// 			// Add other key-value pairs as needed
// 		},
// 	}

// 	err = CreateOrUpdateSecret(r, ctx, secret)
// 	if err != nil {
// 		return ctrl.Result{}, err
// 	}

// 	return ctrl.Result{}, nil
// }
