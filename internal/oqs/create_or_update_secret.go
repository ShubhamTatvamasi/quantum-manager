package oqs

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// other imports
)

func CreateOrUpdateSecret(r client.Client, ctx context.Context, secret *corev1.Secret) error {
	// Function implementation...

	existingSecret := &corev1.Secret{}
	err := r.Get(ctx, client.ObjectKey{Namespace: secret.Namespace, Name: secret.Name}, existingSecret)
	if err != nil && client.IgnoreNotFound(err) != nil {
		return err
	}

	if err != nil { // Secret doesn't exist, create it
		err = r.Create(ctx, secret)
		if err != nil {
			return err
		}
		fmt.Printf("Created Secret: %s\n", secret.Name)
	} else { // Secret exists, update it
		existingSecret.StringData = secret.StringData
		err = r.Update(ctx, existingSecret)
		if err != nil {
			return err
		}
		fmt.Printf("Updated Secret: %s\n", secret.Name)
	}

	return nil

}
