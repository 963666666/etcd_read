package v3

import (
	"context"
)

type Auth interface {
	Authenticate(ctx context.Context, name, password string)

	AuthEnable(ctx context.Context)

	AuthDisable(ctx context.Context)

	AuthStatus(ctx context.Context)

	UserAdd(ctx context.Context, name, password string)

	UserGrantRole(ctx context.Context, user, role string)

	UserGet(ctx context.Context, name string)

	UserList(ctx context.Context)

	UserRevokeRole(ctx context.Context, name, role string)

	RoleAdd(ctx context.Context, name string)

	RoleGrantPermission(ctx context.Context, name, key, rangeEnd string)

	RoleGet(ctx context.Context, role string)

	RoleList(ctx context.Context)

	RoleRevokePermission(ctx context.Context, role, key, rangeEnd string)

	RoleDelete(ctx context.Context, role string)
}

type authClient struct {

}

func NewAuth(c *Client) Auth {
	return nil
}
