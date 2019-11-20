package middleware

import (
	"context"
	"github.com/kubermatic/kubermatic/api/pkg/keycloak"
	"net/http"

	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"

	"github.com/go-kit/kit/endpoint"
	kubermaticapiv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	k8cerrors "github.com/kubermatic/kubermatic/api/pkg/util/errors"
)

const (
	// KeycloakFacadeContextKey key under which the current keycloak.Facade is kept in the ctx
	KeycloakFacadeContextKey = "keycloak-facade"
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

// Keycloak is a middleware that injects the Keycloak client facade into the ctx
func Keycloak(keycloakFacade keycloak.Facade) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			ctx = context.WithValue(ctx, KeycloakFacadeContextKey, keycloakFacade)
			return next(ctx, request)
		}
	}
}
