package getstream

import (
	"context"
)

// This Method returns the application settings
func (c *Client) GetApp(ctx context.Context) (*StreamResponse[GetApplicationResponse], error) {
	var result GetApplicationResponse
	res, err := MakeRequest[any, GetApplicationResponse](c, ctx, "GET", "/api/v2/app", nil, nil, &result, nil)
	return res, err
}

// This Method updates one or more application settings
func (c *Client) UpdateApp(ctx context.Context, request *UpdateAppRequest) (*StreamResponse[Response], error) {
	var result Response
	res, err := MakeRequest[UpdateAppRequest, Response](c, ctx, "PATCH", "/api/v2/app", nil, request, &result, nil)
	return res, err
}

// Returns all available block lists
func (c *Client) ListBlockLists(ctx context.Context) (*StreamResponse[ListBlockListResponse], error) {
	var result ListBlockListResponse
	res, err := MakeRequest[any, ListBlockListResponse](c, ctx, "GET", "/api/v2/blocklists", nil, nil, &result, nil)
	return res, err
}

// Creates a new application blocklist, once created the blocklist can be used by any channel type
func (c *Client) CreateBlockList(ctx context.Context, request *CreateBlockListRequest) (*StreamResponse[Response], error) {
	var result Response
	res, err := MakeRequest[CreateBlockListRequest, Response](c, ctx, "POST", "/api/v2/blocklists", nil, request, &result, nil)
	return res, err
}

// Deletes previously created application blocklist
func (c *Client) DeleteBlockList(ctx context.Context, name string) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, Response](c, ctx, "DELETE", "/api/v2/blocklists/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Returns block list by given name
func (c *Client) GetBlockList(ctx context.Context, name string) (*StreamResponse[GetBlockListResponse], error) {
	var result GetBlockListResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, GetBlockListResponse](c, ctx, "GET", "/api/v2/blocklists/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Updates contents of the block list
func (c *Client) UpdateBlockList(ctx context.Context, name string, request *UpdateBlockListRequest) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[UpdateBlockListRequest, Response](c, ctx, "PUT", "/api/v2/blocklists/{name}", nil, request, &result, pathParams)
	return res, err
}

// Sends a test message via push, this is a test endpoint to verify your push settings
func (c *Client) CheckPush(ctx context.Context, request *CheckPushRequest) (*StreamResponse[CheckPushResponse], error) {
	var result CheckPushResponse
	res, err := MakeRequest[CheckPushRequest, CheckPushResponse](c, ctx, "POST", "/api/v2/check_push", nil, request, &result, nil)
	return res, err
}

// Validates Amazon SNS configuration
func (c *Client) CheckSNS(ctx context.Context, request *CheckSNSRequest) (*StreamResponse[CheckSNSResponse], error) {
	var result CheckSNSResponse
	res, err := MakeRequest[CheckSNSRequest, CheckSNSResponse](c, ctx, "POST", "/api/v2/check_sns", nil, request, &result, nil)
	return res, err
}

// Validates Amazon SQS credentials
func (c *Client) CheckSQS(ctx context.Context, request *CheckSQSRequest) (*StreamResponse[CheckSQSResponse], error) {
	var result CheckSQSResponse
	res, err := MakeRequest[CheckSQSRequest, CheckSQSResponse](c, ctx, "POST", "/api/v2/check_sqs", nil, request, &result, nil)
	return res, err
}

// Deletes one device
func (c *Client) DeleteDevice(ctx context.Context, request *DeleteDeviceRequest) (*StreamResponse[Response], error) {
	var result Response
	params := extractQueryParams(request)
	res, err := MakeRequest[any, Response](c, ctx, "DELETE", "/api/v2/devices", params, nil, &result, nil)
	return res, err
}

// Returns all available devices
func (c *Client) ListDevices(ctx context.Context, request *ListDevicesRequest) (*StreamResponse[ListDevicesResponse], error) {
	var result ListDevicesResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[any, ListDevicesResponse](c, ctx, "GET", "/api/v2/devices", params, nil, &result, nil)
	return res, err
}

// Adds a new device to a user, if the same device already exists the call will have no effect
func (c *Client) CreateDevice(ctx context.Context, request *CreateDeviceRequest) (*StreamResponse[Response], error) {
	var result Response
	res, err := MakeRequest[CreateDeviceRequest, Response](c, ctx, "POST", "/api/v2/devices", nil, request, &result, nil)
	return res, err
}

// Exports user profile, reactions and messages for list of given users
func (c *Client) ExportUsers(ctx context.Context, request *ExportUsersRequest) (*StreamResponse[ExportUsersResponse], error) {
	var result ExportUsersResponse
	res, err := MakeRequest[ExportUsersRequest, ExportUsersResponse](c, ctx, "POST", "/api/v2/export/users", nil, request, &result, nil)
	return res, err
}

// Lists external storage
func (c *Client) ListExternalStorage(ctx context.Context) (*StreamResponse[ListExternalStorageResponse], error) {
	var result ListExternalStorageResponse
	res, err := MakeRequest[any, ListExternalStorageResponse](c, ctx, "GET", "/api/v2/external_storage", nil, nil, &result, nil)
	return res, err
}

// Creates new external storage
func (c *Client) CreateExternalStorage(ctx context.Context, request *CreateExternalStorageRequest) (*StreamResponse[CreateExternalStorageResponse], error) {
	var result CreateExternalStorageResponse
	res, err := MakeRequest[CreateExternalStorageRequest, CreateExternalStorageResponse](c, ctx, "POST", "/api/v2/external_storage", nil, request, &result, nil)
	return res, err
}

// Deletes external storage
func (c *Client) DeleteExternalStorage(ctx context.Context, name string) (*StreamResponse[DeleteExternalStorageResponse], error) {
	var result DeleteExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, DeleteExternalStorageResponse](c, ctx, "DELETE", "/api/v2/external_storage/{name}", nil, nil, &result, pathParams)
	return res, err
}

func (c *Client) UpdateExternalStorage(ctx context.Context, name string, request *UpdateExternalStorageRequest) (*StreamResponse[UpdateExternalStorageResponse], error) {
	var result UpdateExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[UpdateExternalStorageRequest, UpdateExternalStorageResponse](c, ctx, "PUT", "/api/v2/external_storage/{name}", nil, request, &result, pathParams)
	return res, err
}

func (c *Client) CheckExternalStorage(ctx context.Context, name string) (*StreamResponse[CheckExternalStorageResponse], error) {
	var result CheckExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, CheckExternalStorageResponse](c, ctx, "GET", "/api/v2/external_storage/{name}/check", nil, nil, &result, pathParams)
	return res, err
}

func (c *Client) CreateGuest(ctx context.Context, request *CreateGuestRequest) (*StreamResponse[CreateGuestResponse], error) {
	var result CreateGuestResponse
	res, err := MakeRequest[CreateGuestRequest, CreateGuestResponse](c, ctx, "POST", "/api/v2/guest", nil, request, &result, nil)
	return res, err
}

// Creates a new import URL
func (c *Client) CreateImportURL(ctx context.Context, request *CreateImportURLRequest) (*StreamResponse[CreateImportURLResponse], error) {
	var result CreateImportURLResponse
	res, err := MakeRequest[CreateImportURLRequest, CreateImportURLResponse](c, ctx, "POST", "/api/v2/import_urls", nil, request, &result, nil)
	return res, err
}

// Gets an import
func (c *Client) ListImports(ctx context.Context) (*StreamResponse[ListImportsResponse], error) {
	var result ListImportsResponse
	res, err := MakeRequest[any, ListImportsResponse](c, ctx, "GET", "/api/v2/imports", nil, nil, &result, nil)
	return res, err
}

// Creates a new import
func (c *Client) CreateImport(ctx context.Context, request *CreateImportRequest) (*StreamResponse[CreateImportResponse], error) {
	var result CreateImportResponse
	res, err := MakeRequest[CreateImportRequest, CreateImportResponse](c, ctx, "POST", "/api/v2/imports", nil, request, &result, nil)
	return res, err
}

// Gets an import
func (c *Client) GetImport(ctx context.Context, id string) (*StreamResponse[GetImportResponse], error) {
	var result GetImportResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[any, GetImportResponse](c, ctx, "GET", "/api/v2/imports/{id}", nil, nil, &result, pathParams)
	return res, err
}

// Get an OpenGraph attachment for a link
func (c *Client) GetOG(ctx context.Context, request *GetOGRequest) (*StreamResponse[GetOGResponse], error) {
	var result GetOGResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetOGResponse](c, ctx, "GET", "/api/v2/og", params, nil, &result, nil)
	return res, err
}

// Lists all available permissions
func (c *Client) ListPermissions(ctx context.Context) (*StreamResponse[ListPermissionsResponse], error) {
	var result ListPermissionsResponse
	res, err := MakeRequest[any, ListPermissionsResponse](c, ctx, "GET", "/api/v2/permissions", nil, nil, &result, nil)
	return res, err
}

// Gets custom permission
func (c *Client) GetPermission(ctx context.Context, id string) (*StreamResponse[GetCustomPermissionResponse], error) {
	var result GetCustomPermissionResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[any, GetCustomPermissionResponse](c, ctx, "GET", "/api/v2/permissions/{id}", nil, nil, &result, pathParams)
	return res, err
}

// List details of all push providers.
func (c *Client) ListPushProviders(ctx context.Context) (*StreamResponse[ListPushProvidersResponse], error) {
	var result ListPushProvidersResponse
	res, err := MakeRequest[any, ListPushProvidersResponse](c, ctx, "GET", "/api/v2/push_providers", nil, nil, &result, nil)
	return res, err
}

// Upsert a push provider for v2 with multi bundle/package support
func (c *Client) UpsertPushProvider(ctx context.Context, request *UpsertPushProviderRequest) (*StreamResponse[UpsertPushProviderResponse], error) {
	var result UpsertPushProviderResponse
	res, err := MakeRequest[UpsertPushProviderRequest, UpsertPushProviderResponse](c, ctx, "POST", "/api/v2/push_providers", nil, request, &result, nil)
	return res, err
}

// Delete a push provider from v2 with multi bundle/package support. v1 isn't supported in this endpoint
func (c *Client) DeletePushProvider(ctx context.Context, _type string, name string) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"type": _type,
		"name": name,
	}
	res, err := MakeRequest[any, Response](c, ctx, "DELETE", "/api/v2/push_providers/{type}/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Get rate limits usage and quotas
func (c *Client) GetRateLimits(ctx context.Context, request *GetRateLimitsRequest) (*StreamResponse[GetRateLimitsResponse], error) {
	var result GetRateLimitsResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetRateLimitsResponse](c, ctx, "GET", "/api/v2/rate_limits", params, nil, &result, nil)
	return res, err
}

// Lists all available roles
func (c *Client) ListRoles(ctx context.Context) (*StreamResponse[ListRolesResponse], error) {
	var result ListRolesResponse
	res, err := MakeRequest[any, ListRolesResponse](c, ctx, "GET", "/api/v2/roles", nil, nil, &result, nil)
	return res, err
}

// Creates custom role
func (c *Client) CreateRole(ctx context.Context, request *CreateRoleRequest) (*StreamResponse[CreateRoleResponse], error) {
	var result CreateRoleResponse
	res, err := MakeRequest[CreateRoleRequest, CreateRoleResponse](c, ctx, "POST", "/api/v2/roles", nil, request, &result, nil)
	return res, err
}

// Deletes custom role
func (c *Client) DeleteRole(ctx context.Context, name string) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, Response](c, ctx, "DELETE", "/api/v2/roles/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Gets status of a task
func (c *Client) GetTask(ctx context.Context, id string) (*StreamResponse[GetTaskResponse], error) {
	var result GetTaskResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[any, GetTaskResponse](c, ctx, "GET", "/api/v2/tasks/{id}", nil, nil, &result, pathParams)
	return res, err
}

// Find and filter users
//
// Required permissions:
// - SearchUser
func (c *Client) QueryUsers(ctx context.Context, request *QueryUsersRequest) (*StreamResponse[QueryUsersResponse], error) {
	var result QueryUsersResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[any, QueryUsersResponse](c, ctx, "GET", "/api/v2/users", params, nil, &result, nil)
	return res, err
}

// Updates certain fields of the user
//
// Sends events:
// - user.updated
// - user.presence.changed
func (c *Client) UpdateUsersPartial(ctx context.Context, request *UpdateUsersPartialRequest) (*StreamResponse[UpdateUsersResponse], error) {
	var result UpdateUsersResponse
	res, err := MakeRequest[UpdateUsersPartialRequest, UpdateUsersResponse](c, ctx, "PATCH", "/api/v2/users", nil, request, &result, nil)
	return res, err
}

// Update or create users in bulk
//
// Sends events:
// - user.updated
func (c *Client) UpdateUsers(ctx context.Context, request *UpdateUsersRequest) (*StreamResponse[UpdateUsersResponse], error) {
	var result UpdateUsersResponse
	res, err := MakeRequest[UpdateUsersRequest, UpdateUsersResponse](c, ctx, "POST", "/api/v2/users", nil, request, &result, nil)
	return res, err
}

// Get list of blocked Users
func (c *Client) GetBlockedUsers(ctx context.Context, request *GetBlockedUsersRequest) (*StreamResponse[GetBlockedUsersResponse], error) {
	var result GetBlockedUsersResponse
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetBlockedUsersResponse](c, ctx, "GET", "/api/v2/users/block", params, nil, &result, nil)
	return res, err
}

// Block users
func (c *Client) BlockUsers(ctx context.Context, request *BlockUsersRequest) (*StreamResponse[BlockUsersResponse], error) {
	var result BlockUsersResponse
	res, err := MakeRequest[BlockUsersRequest, BlockUsersResponse](c, ctx, "POST", "/api/v2/users/block", nil, request, &result, nil)
	return res, err
}

// Deactivate users in batches
//
// Sends events:
// - user.deactivated
func (c *Client) DeactivateUsers(ctx context.Context, request *DeactivateUsersRequest) (*StreamResponse[DeactivateUsersResponse], error) {
	var result DeactivateUsersResponse
	res, err := MakeRequest[DeactivateUsersRequest, DeactivateUsersResponse](c, ctx, "POST", "/api/v2/users/deactivate", nil, request, &result, nil)
	return res, err
}

// Deletes users and optionally all their belongings asynchronously.
//
// Sends events:
// - channel.deleted
// - user.deleted
func (c *Client) DeleteUsers(ctx context.Context, request *DeleteUsersRequest) (*StreamResponse[DeleteUsersResponse], error) {
	var result DeleteUsersResponse
	res, err := MakeRequest[DeleteUsersRequest, DeleteUsersResponse](c, ctx, "POST", "/api/v2/users/delete", nil, request, &result, nil)
	return res, err
}

// Reactivate users in batches
//
// Sends events:
// - user.reactivated
func (c *Client) ReactivateUsers(ctx context.Context, request *ReactivateUsersRequest) (*StreamResponse[ReactivateUsersResponse], error) {
	var result ReactivateUsersResponse
	res, err := MakeRequest[ReactivateUsersRequest, ReactivateUsersResponse](c, ctx, "POST", "/api/v2/users/reactivate", nil, request, &result, nil)
	return res, err
}

// Restore soft deleted users
func (c *Client) RestoreUsers(ctx context.Context, request *RestoreUsersRequest) (*StreamResponse[Response], error) {
	var result Response
	res, err := MakeRequest[RestoreUsersRequest, Response](c, ctx, "POST", "/api/v2/users/restore", nil, request, &result, nil)
	return res, err
}

// Unblock users
func (c *Client) UnblockUsers(ctx context.Context, request *UnblockUsersRequest) (*StreamResponse[UnblockUsersResponse], error) {
	var result UnblockUsersResponse
	res, err := MakeRequest[UnblockUsersRequest, UnblockUsersResponse](c, ctx, "POST", "/api/v2/users/unblock", nil, request, &result, nil)
	return res, err
}

// Deactivates user with possibility to activate it back
//
// Sends events:
// - user.deactivated
func (c *Client) DeactivateUser(ctx context.Context, userId string, request *DeactivateUserRequest) (*StreamResponse[DeactivateUserResponse], error) {
	var result DeactivateUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	res, err := MakeRequest[DeactivateUserRequest, DeactivateUserResponse](c, ctx, "POST", "/api/v2/users/{user_id}/deactivate", nil, request, &result, pathParams)
	return res, err
}

// Exports the user's profile, reactions and messages. Raises an error if a user has more than 10k messages or reactions
func (c *Client) ExportUser(ctx context.Context, userId string) (*StreamResponse[ExportUserResponse], error) {
	var result ExportUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	res, err := MakeRequest[any, ExportUserResponse](c, ctx, "GET", "/api/v2/users/{user_id}/export", nil, nil, &result, pathParams)
	return res, err
}

// Activates user who's been deactivated previously
//
// Sends events:
// - user.reactivated
func (c *Client) ReactivateUser(ctx context.Context, userId string, request *ReactivateUserRequest) (*StreamResponse[ReactivateUserResponse], error) {
	var result ReactivateUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	res, err := MakeRequest[ReactivateUserRequest, ReactivateUserResponse](c, ctx, "POST", "/api/v2/users/{user_id}/reactivate", nil, request, &result, pathParams)
	return res, err
}
