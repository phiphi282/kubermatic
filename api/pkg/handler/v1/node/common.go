package node

import (
	"fmt"
	apiv1 "github.com/kubermatic/kubermatic/api/pkg/api/v1"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/label"
	machineconversions "github.com/kubermatic/kubermatic/api/pkg/machine"
	clusterv1alpha1 "github.com/kubermatic/machine-controller/pkg/apis/cluster/v1alpha1"
)

func mdSpecToNdSpec(mdSpec *clusterv1alpha1.MachineDeploymentSpec) (*apiv1.NodeDeploymentSpec, error) {
	operatingSystemSpec, err := machineconversions.GetAPIV1OperatingSystemSpec(mdSpec.Template.Spec)
	if err != nil {
		return nil, fmt.Errorf("failed to get operating system spec from machine deployment: %v", err)
	}

	cloudSpec, err := machineconversions.GetAPIV2NodeCloudSpec(mdSpec.Template.Spec)
	if err != nil {
		return nil, fmt.Errorf("failed to get node cloud spec from machine deployment: %v", err)
	}

	taints := make([]apiv1.TaintSpec, len(mdSpec.Template.Spec.Taints))
	for i, taint := range mdSpec.Template.Spec.Taints {
		taints[i] = apiv1.TaintSpec{
			Effect: string(taint.Effect),
			Key:    taint.Key,
			Value:  taint.Value,
		}
	}

	return &apiv1.NodeDeploymentSpec{
		Replicas:    *mdSpec.Replicas,
		Template: apiv1.NodeSpec{
			Labels: label.FilterLabels(label.NodeDeploymentResourceType, mdSpec.Template.Spec.Labels),
			Taints: taints,
			Versions: apiv1.NodeVersionInfo{
				Kubelet: mdSpec.Template.Spec.Versions.Kubelet,
			},
			OperatingSystem: *operatingSystemSpec,
			Cloud:           *cloudSpec,
		},
		Paused: &mdSpec.Paused,
	}, nil
}
