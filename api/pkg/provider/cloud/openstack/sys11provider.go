package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	goopenstack "github.com/gophercloud/gophercloud/openstack"
	oslimits "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/limits"
	osimages "github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
)

func (os *Provider) GetImages(cloud kubermaticv1.CloudSpec) ([]osimages.Image, error) {
	serviceClient, err := os.getComputeClient(cloud)

	if err != nil {
		return nil, fmt.Errorf("couldn't get auth client: %v", err)
	}

	images, err := getAllImages(serviceClient, osimages.ListOpts{})

	if err != nil {
		return nil, fmt.Errorf("couldn't get images: %v", err)
	}

	return images, nil
}

// GetSubnets list all available subnet ids fot a given CloudSpec
func (os *Provider) GetQuotaLimits(cloud kubermaticv1.CloudSpec) (*oslimits.Limits, error) {
	serviceClient, err := os.getComputeClient(cloud)
	if err != nil {
		return nil, fmt.Errorf("couldn't get auth client: %v", err)
	}

	limits, err := getLimits(serviceClient, oslimits.GetOpts{})
	if err != nil {
		return nil, err
	}

	return limits, nil
}

func getLimits(netClient *gophercloud.ServiceClient, opts oslimits.GetOpts) (*oslimits.Limits, error) {
	limits, err := oslimits.Get(netClient, opts).Extract()
	if err != nil {
		return nil, err
	}

	return limits, nil
}

func getAllImages(netClient *gophercloud.ServiceClient, opts osimages.ListOpts) ([]osimages.Image, error) {
	allPages, err := osimages.ListDetail(netClient, opts).AllPages()
	if err != nil {
		return nil, err
	}

	allImages, err := osimages.ExtractImages(allPages)
	if err != nil {
		return nil, err
	}

	return allImages, nil
}

func (os *Provider) getComputeClient(cloud kubermaticv1.CloudSpec) (*gophercloud.ServiceClient, error) {
	authClient, err := os.getAuthClient(cloud)
	if err != nil {
		return nil, err
	}

	dc, found := os.dcs[cloud.DatacenterName]
	if !found || dc.Spec.Openstack == nil {
		return nil, fmt.Errorf("invalid datacenter %q", cloud.DatacenterName)
	}

	return goopenstack.NewComputeV2(authClient, gophercloud.EndpointOpts{Region: dc.Spec.Openstack.Region})
}
