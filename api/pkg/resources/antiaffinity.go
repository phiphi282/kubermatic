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
	var primaryWeight int32 = 100
	if prioritizedClusterName != "" {
		primaryWeight = 5
		delete(labels, "cluster")
	}

	affinity := []corev1.WeightedPodAffinityTerm{
		{
			Weight: primaryWeight,
			PodAffinityTerm: corev1.PodAffinityTerm{
				LabelSelector: &metav1.LabelSelector{
					MatchLabels: labels,
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
