package provider

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	oslimits "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/limits"
	osimages "github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	"github.com/kubermatic/kubermatic/api/pkg/provider/cloud/openstack"
)

func OpenstackImageEndpoint(providers provider.CloudRegistry) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(OpenstackReq)
		if !ok {
			return nil, fmt.Errorf("incorrect type of request, expected = OpenstackReq, got = %T", request)
		}

		return getOpenstackImages(providers, req.Username, req.Password, req.Tenant, req.Domain, req.DatacenterName)
	}
}

func OpenstackQuotaLimitEndpoint(providers provider.CloudRegistry) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(OpenstackReq)
		if !ok {
			return nil, fmt.Errorf("incorrect type of request, expected = OpenstackReq, got = %T", request)
		}

		return getOpenstackQuotaLimits(providers, req.Username, req.Password, req.Tenant, req.Domain, req.DatacenterName)
	}
}

func OpenstackImageNoCredentialsEndpoint(projectProvider provider.ProjectProvider, providers provider.CloudRegistry) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OpenstackNoCredentialsReq)
		cluster, err := getClusterForOpenstack(ctx, projectProvider, req.ProjectID, req.ClusterID)
		if err != nil {
			return nil, err
		}

		openstackSpec := cluster.Spec.Cloud.Openstack
		datacenterName := cluster.Spec.Cloud.DatacenterName
		return getOpenstackImages(providers, openstackSpec.Username, openstackSpec.Password, openstackSpec.Tenant, openstackSpec.Domain, datacenterName)
	}
}

func OpenstackQuotaLimitNoCredentialsEndpoint(projectProvider provider.ProjectProvider, providers provider.CloudRegistry) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OpenstackNoCredentialsReq)
		cluster, err := getClusterForOpenstack(ctx, projectProvider, req.ProjectID, req.ClusterID)
		if err != nil {
			return nil, err
		}

		openstackSpec := cluster.Spec.Cloud.Openstack
		datacenterName := cluster.Spec.Cloud.DatacenterName
		return getOpenstackQuotaLimits(providers, openstackSpec.Username, openstackSpec.Password, openstackSpec.Tenant, openstackSpec.Domain, datacenterName)
	}
}

func getOpenstackImages(providers provider.CloudRegistry, username, password, tenant, domain, datacenterName string) ([]osimages.Image, error) {
	osProviderInterface, ok := providers[provider.OpenstackCloudProvider]
	if !ok {
		return nil, fmt.Errorf("unable to get %s provider", provider.OpenstackCloudProvider)
	}

	osProvider, ok := osProviderInterface.(*openstack.Provider)
	if !ok {
		return nil, fmt.Errorf("unable to cast osProviderInterface to *openstack.Provider")
	}

	return osProvider.GetImages(kubermaticv1.CloudSpec{
		DatacenterName: datacenterName,
		Openstack: &kubermaticv1.OpenstackCloudSpec{
			Username: username,
			Tenant:   tenant,
			Password: password,
			Domain:   domain,
		},
	})
}

func getOpenstackQuotaLimits(providers provider.CloudRegistry, username, password, tenant, domain, datacenterName string) (*oslimits.Limits, error) {
	osProviderInterface, ok := providers[provider.OpenstackCloudProvider]
	if !ok {
		return nil, fmt.Errorf("unable to get %s provider", provider.OpenstackCloudProvider)
	}

	osProvider, ok := osProviderInterface.(*openstack.Provider)
	if !ok {
		return nil, fmt.Errorf("unable to cast osProviderInterface to *openstack.Provider")
	}

	return osProvider.GetQuotaLimits(kubermaticv1.CloudSpec{
		DatacenterName: datacenterName,
		Openstack: &kubermaticv1.OpenstackCloudSpec{
			Username: username,
			Tenant:   tenant,
			Password: password,
			Domain:   domain,
		},
	})
}
