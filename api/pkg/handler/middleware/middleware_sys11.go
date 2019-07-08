package middleware

import (
	"context"
	"net/http"

	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"

	"github.com/go-kit/kit/endpoint"
	kubermaticapiv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	k8cerrors "github.com/kubermatic/kubermatic/api/pkg/util/errors"
)

const (
	SanitizeKubeConfigKey contextKey = "sanitize-kubeconfig"
)

func PrivilegedUserGroupVerifier(userProjectMapper provider.ProjectMemberMapper, privilegedUserGroups map[string]bool) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var projectID string
			prjIDGetter, ok := request.(common.ProjectIDGetter)
			if ok {
				projectID = prjIDGetter.GetProjectID()
			}

			user, ok := ctx.Value(UserCRContextKey).(*kubermaticapiv1.User)
			if !ok {
				return nil, k8cerrors.New(http.StatusInternalServerError, "unable to get authenticated user object")
			}

			group, err := userProjectMapper.MapUserToGroup(user.Spec.Email, projectID)
			if err != nil {
				return nil, err
			}

			// group is in format '${group}-${projectId}'
			group = group[:len(group)-len(projectID)-1]

			if trusted, ok := privilegedUserGroups[group]; ok && trusted {
				return next(context.WithValue(ctx, UserCRContextKey, user), request)
			}

			return nil, k8cerrors.New(http.StatusForbidden, "you don't have permission to access this resource")
		}
	}
}

// Will add the SanitizeKubeConfig Key to the current context, to ensure, that the admin kubeconfig is sanitized before
// handing out. The key is used in cluster.GetAdminKubeconfigEndpoint()
func SanitizeKubeconfig() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// TODO check  for parameter (e.g. ?sanitized&token='...' vs. ?token='...')
			return next(context.WithValue(ctx, SanitizeKubeConfigKey, true), request)
		}
	}
}
