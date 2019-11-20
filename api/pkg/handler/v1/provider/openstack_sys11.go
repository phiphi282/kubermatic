package provider

import (
	"context"
	"fmt"
	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"
	"net/http"

	v1 "github.com/kubermatic/kubermatic/api/pkg/api/v1"

	"github.com/go-kit/kit/endpoint"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	"github.com/kubermatic/kubermatic/api/pkg/provider/cloud/openstack"
	"github.com/kubermatic/kubermatic/api/pkg/util/errors"
)

func OpenstackImageEndpoint(seedsGetter provider.SeedsGetter, credentialManager common.PresetsManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(OpenstackReq)
		if !ok {
			return nil, fmt.Errorf("incorrect type of request, expected = OpenstackReq, got = %T", request)
		}
		userInfo, ok := ctx.Value(middleware.UserInfoContextKey).(*provider.UserInfo)
		if !ok {
			return nil, errors.New(http.StatusInternalServerError, "can not get user info")
		}

		datacenterName := req.DatacenterName
		_, datacenter, err := provider.DatacenterFromSeedMap(userInfo, seedsGetter, datacenterName)
		if err != nil {
			return nil, fmt.Errorf("error getting dc: %v", err)
		}

		username, password, domain, tenant, tenantID, err := getOpenstackCredentials(userInfo, req.Credential, req.Username, req.Password, req.Domain, req.Tenant, req.TenantID, credentialManager)
		if err != nil {
			return nil, fmt.Errorf("error getting OpenStack credentials: %v", err)
		}

		return openstack.GetImages(username, password, domain, tenant, tenantID, datacenter.Spec.Openstack.AuthURL, datacenter.Spec.Openstack.Region)
	}
}

func OpenstackQuotaLimitEndpoint(seedsGetter provider.SeedsGetter, credentialManager common.PresetsManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(OpenstackReq)
		if !ok {
			return nil, fmt.Errorf("incorrect type of request, expected = OpenstackReq, got = %T", request)
		}
		userInfo, ok := ctx.Value(middleware.UserInfoContextKey).(*provider.UserInfo)
		if !ok {
			return nil, errors.New(http.StatusInternalServerError, "can not get user info")
		}

		datacenterName := req.DatacenterName
		_, datacenter, err := provider.DatacenterFromSeedMap(userInfo, seedsGetter, datacenterName)
		if err != nil {
			return nil, fmt.Errorf("error getting dc: %v", err)
		}

		username, password, domain, tenant, tenantID, err := getOpenstackCredentials(userInfo, req.Credential, req.Username, req.Password, req.Domain, req.Tenant, req.TenantID, credentialManager)
		if err != nil {
			return nil, fmt.Errorf("error getting OpenStack credentials: %v", err)
		}

		return getOpenstackQuotaLimits(username, password, domain, tenant, tenantID, datacenter.Spec.Openstack.AuthURL, datacenter.Spec.Openstack.Region)
	}
}

func OpenstackImageNoCredentialsEndpoint(projectProvider provider.ProjectProvider, seedsGetter provider.SeedsGetter, credentialManager common.PresetsManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OpenstackNoCredentialsReq)
		cluster, err := getClusterForOpenstack(ctx, projectProvider, req.ProjectID, req.ClusterID)
		if err != nil {
			return nil, err
		}

		userInfo, ok := ctx.Value(middleware.UserInfoContextKey).(*provider.UserInfo)
		if !ok {
			return nil, errors.New(http.StatusInternalServerError, "can not get user info")
		}

		openstackSpec := cluster.Spec.Cloud.Openstack
		datacenterName := cluster.Spec.Cloud.DatacenterName
		_, datacenter, err := provider.DatacenterFromSeedMap(userInfo, seedsGetter, datacenterName)
		if err != nil {
			return nil, fmt.Errorf("error getting dc: %v", err)
		}

		return openstack.GetImages(openstackSpec.Username, openstackSpec.Password, openstackSpec.Domain, openstackSpec.Tenant, openstackSpec.TenantID, datacenter.Spec.Openstack.AuthURL, datacenter.Spec.Openstack.Region)
	}
}

func OpenstackQuotaLimitNoCredentialsEndpoint(projectProvider provider.ProjectProvider, seedsGetter provider.SeedsGetter, credentialManager common.PresetsManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OpenstackNoCredentialsReq)
		cluster, err := getClusterForOpenstack(ctx, projectProvider, req.ProjectID, req.ClusterID)
		if err != nil {
			return nil, err
		}

		userInfo, ok := ctx.Value(middleware.UserInfoContextKey).(*provider.UserInfo)
		if !ok {
			return nil, errors.New(http.StatusInternalServerError, "can not get user info")
		}

		openstackSpec := cluster.Spec.Cloud.Openstack
		datacenterName := cluster.Spec.Cloud.DatacenterName
		_, datacenter, err := provider.DatacenterFromSeedMap(userInfo, seedsGetter, datacenterName)
		if err != nil {
			return nil, fmt.Errorf("error getting dc: %v", err)
		}

		return getOpenstackQuotaLimits(openstackSpec.Username, openstackSpec.Password, openstackSpec.Domain, openstackSpec.Tenant, openstackSpec.TenantID, datacenter.Spec.Openstack.AuthURL, datacenter.Spec.Openstack.Region)
	}
}

func getOpenstackQuotaLimits(username, password, domain, tenant, tenantID, authURL, region string) (*v1.Quotas, error) {
	limits, err := openstack.GetQuotaLimits(username, password, domain, tenant, tenantID, authURL, region)
	if err != nil {
		return nil, err
	}

	usedFloatingIPCount, err := openstack.GetUsedFloatingIPCount(username, password, domain, tenant, tenantID, authURL, region)
	if err != nil {
		return nil, err
	}

	floatingIPQuota, err := openstack.GetFloatingIPQuota(username, password, domain, tenant, tenantID, authURL, region)
	if err != nil {
		return nil, err
	}

	apiLimits := &v1.Quotas{
		Limits:              limits,
		UsedFloatingIPCount: usedFloatingIPCount,
		FloatingIPQuota:     floatingIPQuota,
	}

	return apiLimits, nil
}
