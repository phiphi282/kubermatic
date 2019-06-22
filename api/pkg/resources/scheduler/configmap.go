package scheduler

import (
	"bytes"
	"fmt"
	"github.com/kubermatic/kubermatic/api/pkg/resources"
	"github.com/kubermatic/kubermatic/api/pkg/resources/reconciling"
	corev1 "k8s.io/api/core/v1"
	"text/template"
)

const (
	schedulerPolicyTpl = `{
    "apiVersion": "v1",
    "kind": "Policy",
    "predicates": [
{{- if .Cluster.Spec.Cloud.Azure }}
        {
            "name": "MaxAzureDiskVolumeCount"
        },
{{- end }}
        {
            "name": "MatchInterPodAffinity"
        },
        {
            "name": "GeneralPredicates"
        },
        {
            "name": "CheckVolumeBinding"
        },
        {
            "name": "CheckNodeUnschedulable"
        },
{{- if .Cluster.Spec.Cloud.AWS }}
        {
            "name": "MaxEBSVolumeCount"
        },
{{- end }}
{{- if false }}
        {
            "name": "MaxGCEPDVolumeCount"
        },
{{- end }}
        {
            "name": "NoDiskConflict"
        },
        {
            "name": "NoVolumeZoneConflict"
        },
        {
            "name": "MaxCSIVolumeCountPred"
        },
{{- if and .Cluster.Spec.Cloud.Openstack (ge .Cluster.Spec.Version.Semver.Minor 14) }}
        {
            "name": "MaxCinderVolumeCount"
        },
{{- end }}
        {
            "name": "PodToleratesNodeTaints"
        }
    ],
    "priorities": [
        {
            "name": "SelectorSpreadPriority",
            "weight": 1
        },
        {
            "name": "InterPodAffinityPriority",
            "weight": 1
        },
        {
            "name": "LeastRequestedPriority",
            "weight": 1
        },
        {
            "name": "BalancedResourceAllocation",
            "weight": 1
        },
        {
            "name": "ImageLocalityPriority",
            "weight": 1
        },
        {
            "name": "NodePreferAvoidPodsPriority",
            "weight": 10000
        },
        {
            "name": "NodeAffinityPriority",
            "weight": 1
        },
        {
            "name": "TaintTolerationPriority",
            "weight": 1
        }
    ]
}`
)

// PolicyConfigMapCreator returns a function to create the ConfigMap containing the scheduler configuration
func PolicyConfigMapCreator(data *resources.TemplateData) reconciling.NamedConfigMapCreatorGetter {
	return func() (string, reconciling.ConfigMapCreator) {
		return resources.SchedulerPolicyConfigMapName, func(cm *corev1.ConfigMap) (*corev1.ConfigMap, error) {
			if cm.Data == nil {
				cm.Data = map[string]string{}
			}

			policyConfig, err := PolicyConfig(data)
			if err != nil {
				return nil, fmt.Errorf("failed to create scheduler policy: %v", err)
			}

			cm.Labels = resources.BaseAppLabel(name, nil)
			cm.Data[resources.SchedulerPolicyFileName] = policyConfig

			return cm, nil
		}
	}
}

// PolicyConfig returns the scheduler configuration for the supplied data
func PolicyConfig(data *resources.TemplateData) (policyConfig string, err error) {
	tpl, err := template.New("scheduler-policy").Parse(schedulerPolicyTpl)
	if err != nil {
		return "", fmt.Errorf("failed to parse the scheduler policy template: %v", err)
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data); err != nil {
		return "", fmt.Errorf("failed to execute scheduler policy template: %v", err)
	}

	return buf.String(), nil
}
