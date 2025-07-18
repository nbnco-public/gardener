// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package controllerregistration

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	operatorv1alpha1 "github.com/gardener/gardener/pkg/apis/operator/v1alpha1"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/utils/gardener/operator"
	"github.com/gardener/gardener/pkg/utils/managedresources"
)

// Interface contains functions to handle the registration of extensions for shoot clusters.
type Interface interface {
	// Reconcile creates or updates the ControllerRegistration and ControllerDeployment for the given extension.
	Reconcile(context.Context, logr.Logger, *operatorv1alpha1.Extension) error
	// Delete deletes the ControllerRegistration and ControllerDeployment for the given extension.
	Delete(context.Context, logr.Logger, *operatorv1alpha1.Extension) error
}

type registration struct {
	runtimeClient client.Client
	recorder      record.EventRecorder

	gardenNamespace string
}

// Reconcile creates or updates the ControllerRegistration and ControllerDeployment for the given extension.
// If the extension doesn't define an extension deployment, the registration is deleted.
func (r *registration) Reconcile(ctx context.Context, log logr.Logger, extension *operatorv1alpha1.Extension) error {
	if extension.Spec.Deployment == nil ||
		extension.Spec.Deployment.ExtensionDeployment == nil ||
		extension.Spec.Deployment.ExtensionDeployment.Helm == nil {
		if err := r.Delete(ctx, log, extension); err != nil {
			return err
		}
		r.recorder.Event(extension, corev1.EventTypeNormal, "Deletion", "ControllerRegistration and ControllerDeployment deleted successfully")

		return nil
	}

	log.Info("Deploying ControllerRegistration and ControllerDeployment")
	if err := r.createOrUpdateControllerRegistration(ctx, extension); err != nil {
		return fmt.Errorf("failed to reconcile ControllerRegistration: %w", err)
	}
	r.recorder.Event(extension, corev1.EventTypeNormal, "Reconciliation", "ControllerRegistration and ControllerDeployment applied successfully")

	return nil
}

func (r *registration) createOrUpdateControllerRegistration(ctx context.Context, extension *operatorv1alpha1.Extension) error {
	var (
		registry = managedresources.NewRegistry(kubernetes.GardenScheme, kubernetes.GardenCodec, kubernetes.GardenSerializer)

		controllerRegistration, controllerDeployment = operator.ControllerRegistrationForExtension(extension)
	)

	objs := []client.Object{controllerRegistration, controllerDeployment}
	if pullSecretRef := GetExtensionPullSecretRef(extension); pullSecretRef != nil {
		secret, err := r.createPullSecretCopy(ctx, extension.Name, pullSecretRef)
		if err != nil {
			return fmt.Errorf("failed to get pull secret: %w", err)
		}
		objs = append(objs, secret)
		controllerDeployment.Helm.OCIRepository.PullSecretRef.Name = secret.Name
	}
	data, err := registry.AddAllAndSerialize(objs...)
	if err != nil {
		return err
	}

	return managedresources.CreateForShoot(ctx, r.runtimeClient, r.gardenNamespace, managedResourceName(extension), managedresources.LabelValueOperator, false, data)
}

// Delete deletes the ControllerRegistration and ControllerDeployment for the given extension.
func (r *registration) Delete(ctx context.Context, log logr.Logger, extension *operatorv1alpha1.Extension) error {
	mrName := managedResourceName(extension)

	log.Info("Deleting extension registration ManagedResource", "managedResource", client.ObjectKey{Name: mrName, Namespace: r.gardenNamespace})
	if err := managedresources.DeleteForShoot(ctx, r.runtimeClient, r.gardenNamespace, mrName); err != nil {
		return fmt.Errorf("failed deleting ManagedResource: %w", err)
	}

	return managedresources.WaitUntilDeleted(ctx, r.runtimeClient, r.gardenNamespace, mrName)
}

func (r *registration) createPullSecretCopy(ctx context.Context, extensionName string, pullSecretRef *corev1.LocalObjectReference) (*corev1.Secret, error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pullSecretRef.Name,
			Namespace: r.gardenNamespace,
		},
	}
	if err := r.runtimeClient.Get(ctx, client.ObjectKeyFromObject(secret), secret); err != nil {
		return nil, fmt.Errorf("failed to get pull secret: %w", err)
	}

	secretCopy := secret.DeepCopy()
	secretCopy.ObjectMeta = metav1.ObjectMeta{
		Name:      fmt.Sprintf("%s-%s", extensionName, pullSecretRef.Name),
		Namespace: v1beta1constants.GardenNamespace,
		Labels:    secretCopy.Labels,
	}

	return secretCopy, nil
}

func managedResourceName(extension *operatorv1alpha1.Extension) string {
	return fmt.Sprintf("extension-registration-%s", extension.Name)
}

// New creates a new handler for ControllerRegistrations.
func New(runtimeClient client.Client, recorder record.EventRecorder, gardenNamespace string) Interface {
	return &registration{
		runtimeClient: runtimeClient,
		recorder:      recorder,

		gardenNamespace: gardenNamespace,
	}
}

// GetExtensionPullSecretRef returns the pull secret reference for the extension's Helm chart.
func GetExtensionPullSecretRef(extension *operatorv1alpha1.Extension) *corev1.LocalObjectReference {
	if extension.Spec.Deployment == nil ||
		extension.Spec.Deployment.ExtensionDeployment == nil ||
		extension.Spec.Deployment.ExtensionDeployment.Helm == nil ||
		extension.Spec.Deployment.ExtensionDeployment.Helm.OCIRepository == nil {
		return nil
	}
	return extension.Spec.Deployment.ExtensionDeployment.Helm.OCIRepository.PullSecretRef
}
