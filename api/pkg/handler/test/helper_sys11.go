package test

import (
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GenTestMachineDeploymentRequest(cluster *kubermaticv1.Cluster, name, rawProviderSpec string, selector map[string]string) *kubermaticv1.MachineDeploymentRequest {
	md := GenTestMachineDeployment(name, rawProviderSpec, selector)
	return &kubermaticv1.MachineDeploymentRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: cluster.Status.NamespaceName,
		},
		Spec: kubermaticv1.MachineDeploymentRequestSpec{
			Md: *md,
		},
	}
}
