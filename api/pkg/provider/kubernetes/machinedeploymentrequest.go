package kubernetes

import (
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	"github.com/kubermatic/machine-controller/pkg/apis/cluster/v1alpha1"

	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MachineDeploymentRequestProvider struct that holds required components of the MachineDeploymentRequestProvider implementation
type MachineDeploymentRequestProvider struct {
	// createSeedImpersonatedClient is used as a ground for impersonation
	// whenever a connection to Seed API server is required
	createSeedImpersonatedClient kubermaticImpersonationClient
}

// NewMachineDeploymentRequestProvider returns a new MachineDeploymentRequest provider that respects RBAC policies
// it uses createSeedImpersonatedClient to create a connection that uses user impersonation
func NewMachineDeploymentRequestProvider(
	createSeedImpersonatedClient kubermaticImpersonationClient) *MachineDeploymentRequestProvider {
	return &MachineDeploymentRequestProvider{
		createSeedImpersonatedClient: createSeedImpersonatedClient,
	}
}

// New creates a new MachineDeploymentRequest for the given cluster
func (p *MachineDeploymentRequestProvider) New(userInfo *provider.UserInfo, cluster *kubermaticv1.Cluster, mdrName string, mdSpec *v1alpha1.MachineDeploymentSpec) (*kubermaticv1.MachineDeploymentRequest, error) {
	seedImpersonatedClient, err := createKubermaticImpersonationClientWrapperFromUserInfo(userInfo, p.createSeedImpersonatedClient)
	if err != nil {
		return nil, err
	}

	gv := kubermaticv1.SchemeGroupVersion

	mdr := &kubermaticv1.MachineDeploymentRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:            mdrName,
			Namespace:       cluster.Status.NamespaceName,
			OwnerReferences: []metav1.OwnerReference{*metav1.NewControllerRef(cluster, gv.WithKind("Cluster"))},
			Labels:          map[string]string{},
		},
		Spec: kubermaticv1.MachineDeploymentRequestSpec{
			MdSpec: *mdSpec,
		},
	}

	mdr, err = seedImpersonatedClient.MachineDeploymentRequests(cluster.Status.NamespaceName).Create(mdr)
	if err != nil {
		return nil, err
	}

	return mdr, nil
}

// Get returns the given MachineDeploymentRequest
func (p *MachineDeploymentRequestProvider) Get(userInfo *provider.UserInfo, cluster *kubermaticv1.Cluster, mdrName string) (*kubermaticv1.MachineDeploymentRequest, error) {
	seedImpersonatedClient, err := createKubermaticImpersonationClientWrapperFromUserInfo(userInfo, p.createSeedImpersonatedClient)
	if err != nil {
		return nil, err
	}

	mdr, err := seedImpersonatedClient.MachineDeploymentRequests(cluster.Status.NamespaceName).Get(mdrName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return mdr, nil
}

// List returns all MachineDeploymentRequests in the given cluster
func (p *MachineDeploymentRequestProvider) List(userInfo *provider.UserInfo, cluster *kubermaticv1.Cluster) ([]*kubermaticv1.MachineDeploymentRequest, error) {
	seedImpersonatedClient, err := createKubermaticImpersonationClientWrapperFromUserInfo(userInfo, p.createSeedImpersonatedClient)
	if err != nil {
		return nil, err
	}

	mdrList, err := seedImpersonatedClient.MachineDeploymentRequests(cluster.Status.NamespaceName).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	result := []*kubermaticv1.MachineDeploymentRequest{}
	for _, mdr := range mdrList.Items {
		result = append(result, mdr.DeepCopy())
	}

	return result, nil
}

// Update updates an MachineDeploymentRequest
func (p *MachineDeploymentRequestProvider) Update(userInfo *provider.UserInfo, cluster *kubermaticv1.Cluster, newMdRequest *kubermaticv1.MachineDeploymentRequest) (*kubermaticv1.MachineDeploymentRequest, error) {
	seedImpersonatedClient, err := createKubermaticImpersonationClientWrapperFromUserInfo(userInfo, p.createSeedImpersonatedClient)
	if err != nil {
		return nil, err
	}

	newMdRequest, err = seedImpersonatedClient.MachineDeploymentRequests(cluster.Status.NamespaceName).Update(newMdRequest)
	if err != nil {
		return nil, err
	}

	return newMdRequest, nil
}

// Delete deletes the given MachineDeploymentRequest
func (p *MachineDeploymentRequestProvider) Delete(userInfo *provider.UserInfo, cluster *kubermaticv1.Cluster, mdrName string) error {
	seedImpersonatedClient, err := createKubermaticImpersonationClientWrapperFromUserInfo(userInfo, p.createSeedImpersonatedClient)
	if err != nil {
		return err
	}

	return seedImpersonatedClient.MachineDeploymentRequests(cluster.Status.NamespaceName).Delete(mdrName, &metav1.DeleteOptions{})
}

func MachineDeploymentRequestProviderFactory(seedKubeconfigGetter provider.SeedKubeconfigGetter) provider.MachineDeploymentRequestProviderGetter {
	return func(seed *kubermaticv1.Seed) (provider.MachineDeploymentRequestProvider, error) {
		cfg, err := seedKubeconfigGetter(seed)
		if err != nil {
			return nil, err
		}
		defaultImpersonationClientForSeed := NewKubermaticImpersonationClient(cfg)

		return NewMachineDeploymentRequestProvider(defaultImpersonationClientForSeed.CreateImpersonatedKubermaticClientSet), nil
	}
}
