package resources

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HostnameAntiAffinity returns a simple Affinity rule to prevent* scheduling of same kind pods on the same node. If
// prioritizedClusterName is not empty an extra affinity will be added to distribute all pods matching the labels except
// "cluster".
// *if scheduling is not possible with this rule, it will be ignored.
func HostnameAntiAffinity(labels map[string]string, prioritizedClusterName string) *corev1.Affinity {
	var reducedWeight int32 = 100
	reducedLabels := make(map[string]string)
	for k, v := range labels {
		reducedLabels[k] = v
	}
	if prioritizedClusterName != "" {
		reducedWeight = 5
		delete(reducedLabels, "cluster")
	}

	affinity := []corev1.WeightedPodAffinityTerm{
		{
			Weight: reducedWeight,
			PodAffinityTerm: corev1.PodAffinityTerm{
				LabelSelector: &metav1.LabelSelector{
					MatchLabels: reducedLabels,
				},
				TopologyKey: TopologyKeyHostname,
			},
		},
	}

	if prioritizedClusterName != "" {
		labels["cluster"] = prioritizedClusterName
		affinity = append(affinity, []corev1.WeightedPodAffinityTerm{
			{
				Weight: 100,
				PodAffinityTerm: corev1.PodAffinityTerm{
					LabelSelector: &metav1.LabelSelector{
						MatchLabels: labels,
					},
					TopologyKey: TopologyKeyHostname,
				},
			},
		}...)
	}

	return &corev1.Affinity{
		PodAntiAffinity: &corev1.PodAntiAffinity{
			PreferredDuringSchedulingIgnoredDuringExecution: affinity,
		},
	}
}
