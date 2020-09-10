/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package resources

import (
	nodelabelerapi "github.com/kubermatic/kubermatic/api/pkg/controller/user-cluster-controller-manager/node-labeler/api"
	"github.com/kubermatic/kubermatic/api/pkg/resources"
	"github.com/kubermatic/kubermatic/api/pkg/resources/reconciling"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	AgentDaemonSetName = "flatcar-linux-update-agent"
)

var (
	daemonSetMaxUnavailable = intstr.FromInt(1)
	hostPathType            = corev1.HostPathUnset
)

func AgentDaemonSetCreator(getRegistry GetImageRegistry) reconciling.NamedDaemonSetCreatorGetter {
	return func() (string, reconciling.DaemonSetCreator) {
		return AgentDaemonSetName, func(ds *appsv1.DaemonSet) (*appsv1.DaemonSet, error) {
			ds.Spec.UpdateStrategy = appsv1.DaemonSetUpdateStrategy{
				Type: appsv1.RollingUpdateDaemonSetStrategyType,
				RollingUpdate: &appsv1.RollingUpdateDaemonSet{
					MaxUnavailable: &daemonSetMaxUnavailable,
				},
			}

			labels := map[string]string{"app": AgentDaemonSetName}
			ds.Spec.Selector = &metav1.LabelSelector{MatchLabels: labels}
			ds.Spec.Template.ObjectMeta.Labels = labels

			// The agent should only run on Flatcar nodes
			ds.Spec.Template.Spec.NodeSelector = map[string]string{nodelabelerapi.DistributionLabelKey: nodelabelerapi.FlatcarLabelValue}

			ds.Spec.Template.Spec.ServiceAccountName = AgentServiceAccountName

			ds.Spec.Template.Spec.Containers = []corev1.Container{
				{
					Name:    "update-agent",
					Image:   getRegistry(resources.RegistryQuay) + "/kinvolk/flatcar-linux-update-operator:v0.7.3",
					Command: []string{"/bin/update-agent"},
					Env: []corev1.EnvVar{
						{
							Name: "UPDATE_AGENT_NODE",
							ValueFrom: &corev1.EnvVarSource{
								FieldRef: &corev1.ObjectFieldSelector{
									APIVersion: "v1",
									FieldPath:  "spec.nodeName",
								},
							},
						},
						{
							Name: "POD_NAMESPACE",
							ValueFrom: &corev1.EnvVarSource{
								FieldRef: &corev1.ObjectFieldSelector{
									APIVersion: "v1",
									FieldPath:  "metadata.namespace",
								},
							},
						},
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "var-run-dbus",
							MountPath: "/var/run/dbus",
						},
						{
							Name:      "etc-flatcar",
							MountPath: "/etc/flatcar",
						},
						{
							Name:      "usr-share-flatcar",
							MountPath: "/usr/share/flatcar",
						},
						{
							Name:      "etc-os-release",
							MountPath: "/etc/os-release",
						},
					},
				},
			}

			ds.Spec.Template.Spec.Tolerations = []corev1.Toleration{
				{
					Key:      "node-role.kubernetes.io/master",
					Operator: "Exists",
					Effect:   "NoSchedule",
				},
			}

			ds.Spec.Template.Spec.Volumes = []corev1.Volume{
				{
					Name: "var-run-dbus",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/var/run/dbus",
							Type: &hostPathType,
						},
					},
				},
				{
					Name: "etc-flatcar",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/etc/flatcar",
							Type: &hostPathType,
						},
					},
				},
				{
					Name: "usr-share-flatcar",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/usr/share/flatcar",
							Type: &hostPathType,
						},
					},
				},
				{
					Name: "etc-os-release",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/etc/os-release",
							Type: &hostPathType,
						},
					},
				},
			}

			return ds, nil
		}
	}
}
