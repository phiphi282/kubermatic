package kubermatic

import (
	"fmt"
	"strings"

	operatorv1alpha1 "github.com/kubermatic/kubermatic/api/pkg/crd/operator/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

const (
	// nameLabel is the recommended name for an identifying label.
	nameLabel = "app.kubernetes.io/name"

	// versionLabel is the recommended name for a version label.
	versionLabel = "app.kubernetes.io/version"
)

func i32ptr(i int32) *int32 {
	return &i
}

// mergeServicePort searches the src port inside the dst slice and if it
// finds it, updates its properties. If the src port is not found yet,
// it will be appended verbatim to the dst list. If dst is nil, a new
// slice will be created.
func mergeServicePort(dst []corev1.ServicePort, src corev1.ServicePort) []corev1.ServicePort {
	for idx, port := range dst {
		if port.Name == src.Name {
			dst[idx].Port = src.Port
			dst[idx].TargetPort = src.TargetPort
			dst[idx].Protocol = src.Protocol

			return dst
		}
	}

	if dst == nil {
		dst = make([]corev1.ServicePort, 0)
	}

	return append(dst, src)
}

func featureGates(cfg *operatorv1alpha1.KubermaticConfiguration) string {
	features := make([]string, 0)
	for _, feature := range cfg.Spec.EnabledFeatureGates() {
		features = append(features, fmt.Sprintf("%s=true", feature))
	}

	return strings.Join(features, ",")
}

func dockerImage(img operatorv1alpha1.DockerImage) string {
	if img.Tag == "" {
		return img.Repository
	}

	return fmt.Sprintf("%s:%s", img.Repository, img.Tag)
}
