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

package kubernetes_test

import (
	"testing"

	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	"github.com/kubermatic/kubermatic/api/pkg/provider/kubernetes"
	"k8s.io/apimachinery/pkg/util/diff"

	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	restclient "k8s.io/client-go/rest"
)

func TestCreateCluster(t *testing.T) {
	// test data
	testcases := []struct {
		name                      string
		workerName                string
		existingKubermaticObjects []runtime.Object
		project                   *kubermaticv1.Project
		userInfo                  *provider.UserInfo
		spec                      *kubermaticv1.ClusterSpec
		clusterType               string
		expectedCluster           *kubermaticv1.Cluster
		expectedError             string
		shareKubeconfig           bool
	}{
		{
			name:            "scenario 1, create kubernetes cluster",
			shareKubeconfig: false,
			workerName:      "test-kubernetes",
			userInfo:        &provider.UserInfo{Email: "john@acme.com", Group: "owners-abcd"},
			project:         genDefaultProject(),
			spec:            genClusterSpec("test-k8s"),
			clusterType:     "kubernetes",
			existingKubermaticObjects: []runtime.Object{
				createAuthenitactedUser(),
				genDefaultProject(),
			},
			expectedCluster: genCluster("test-k8s", "kubernetes", "my-first-project-ID", "test-kubernetes", "john@acme.com"),
		},
		{
			name:            "scenario 2, create OpenShift cluster",
			shareKubeconfig: false,
			workerName:      "test-openshift",
			userInfo:        &provider.UserInfo{Email: "john@acme.com", Group: "owners-abcd"},
			project:         genDefaultProject(),
			spec:            genClusterSpec("test-openshift"),
			clusterType:     "openshift",
			existingKubermaticObjects: []runtime.Object{
				createAuthenitactedUser(),
				genDefaultProject(),
			},
			expectedCluster: genCluster("test-openshift", "openshift", "my-first-project-ID", "test-openshift", "john@acme.com"),
		},
		{
			name:            "scenario 3, create kubernetes cluster when share kubeconfig is enabled and OIDC is set",
			shareKubeconfig: true,
			workerName:      "test-kubernetes",
			userInfo:        &provider.UserInfo{Email: "john@acme.com", Group: "owners-abcd"},
			project:         genDefaultProject(),
			spec: func() *kubermaticv1.ClusterSpec {
				spec := genClusterSpec("test-k8s")
				spec.OIDC = kubermaticv1.OIDCSettings{
					IssuerURL: "http://test",
					ClientID:  "test",
				}
				return spec
			}(),
			clusterType: "kubernetes",
			existingKubermaticObjects: []runtime.Object{
				createAuthenitactedUser(),
				genDefaultProject(),
			},
			expectedError: "can not set OIDC for the cluster when share config feature is enabled",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			impersonationClient, _, _, err := createFakeKubermaticClients(tc.existingKubermaticObjects)
			if err != nil {
				t.Fatalf("unable to create fake clients, err = %v", err)
			}

			// act
			target := kubernetes.NewClusterProvider(&restclient.Config{}, impersonationClient.CreateFakeImpersonatedClientSet, nil, tc.workerName, nil, nil, nil, tc.shareKubeconfig)
			partialCluster := &kubermaticv1.Cluster{}
			partialCluster.Spec = *tc.spec
			if tc.clusterType == "openshift" {
				partialCluster.Annotations = map[string]string{
					"kubermatic.io/openshift": "true",
				}
			}
			if tc.expectedCluster != nil {
				partialCluster.Finalizers = tc.expectedCluster.Finalizers
			}

			cluster, err := target.New(tc.project, tc.userInfo, partialCluster)
			if len(tc.expectedError) > 0 {
				if err == nil {
					t.Fatalf("expected error: %s", tc.expectedError)
				}
				if tc.expectedError != err.Error() {
					t.Fatalf("expected error: %s got %v", tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}

				// override autogenerated field
				cluster.Name = tc.expectedCluster.Name
				cluster.Status.NamespaceName = tc.expectedCluster.Status.NamespaceName

				if !equality.Semantic.DeepEqual(cluster, tc.expectedCluster) {
					t.Fatalf("%v", diff.ObjectDiff(tc.expectedCluster, cluster))
				}
			}

		})
	}
}
