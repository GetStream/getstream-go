package getstream

import (
	"context"
)

type CommonClient struct {
	client *Client
}

func NewCommonClient(client *Client) *CommonClient {
	return &CommonClient{
		client: client,
	}
}

// This Method returns the application settings
func (c *CommonClient) GetApp(ctx context.Context) (*StreamResponse[GetApplicationResponse], error) {
	var result GetApplicationResponse
	res, err := MakeRequest[any, GetApplicationResponse, any](c.client, ctx, "GET", "/api/v2/app", nil, nil, &result, nil)
	return res, err
}

// This Method updates one or more application settings
func (c *CommonClient) UpdateApp(ctx context.Context, request *UpdateAppRequest) (*StreamResponse[Response], error) {
	var result Response
	res, err := MakeRequest[UpdateAppRequest, Response, any](c.client, ctx, "PATCH", "/api/v2/app", nil, request, &result, nil)
	return res, err
}

// Returns all available block lists
func (c *CommonClient) ListBlockLists(ctx context.Context) (*StreamResponse[ListBlockListResponse], error) {
	var result ListBlockListResponse
	res, err := MakeRequest[any, ListBlockListResponse, any](c.client, ctx, "GET", "/api/v2/blocklists", nil, nil, &result, nil)
	return res, err
}

// Creates a new application blocklist, once created the blocklist can be used by any channel type
func (c *CommonClient) CreateBlockList(ctx context.Context, request *CreateBlockListRequest) (*StreamResponse[Response], error) {
	var result Response
	res, err := MakeRequest[CreateBlockListRequest, Response, any](c.client, ctx, "POST", "/api/v2/blocklists", nil, request, &result, nil)
	return res, err
}

// Deletes previously created application blocklist
func (c *CommonClient) DeleteBlockList(ctx context.Context, name string) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/blocklists/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Returns block list by given name
func (c *CommonClient) GetBlockList(ctx context.Context, name string) (*StreamResponse[GetBlockListResponse], error) {
	var result GetBlockListResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, GetBlockListResponse, any](c.client, ctx, "GET", "/api/v2/blocklists/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Updates contents of the block list
func (c *CommonClient) UpdateBlockList(ctx context.Context, name string, request *UpdateBlockListRequest) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[UpdateBlockListRequest, Response, any](c.client, ctx, "PUT", "/api/v2/blocklists/{name}", nil, request, &result, pathParams)
	return res, err
}

// Sends a test message via push, this is a test endpoint to verify your push settings
func (c *CommonClient) CheckPush(ctx context.Context, request *CheckPushRequest) (*StreamResponse[CheckPushResponse], error) {
	var result CheckPushResponse
	res, err := MakeRequest[CheckPushRequest, CheckPushResponse, any](c.client, ctx, "POST", "/api/v2/check_push", nil, request, &result, nil)
	return res, err
}

// Validates Amazon SNS configuration
func (c *CommonClient) CheckSNS(ctx context.Context, request *CheckSNSRequest) (*StreamResponse[CheckSNSResponse], error) {
	var result CheckSNSResponse
	res, err := MakeRequest[CheckSNSRequest, CheckSNSResponse, any](c.client, ctx, "POST", "/api/v2/check_sns", nil, request, &result, nil)
	return res, err
}

// Validates Amazon SQS credentials
func (c *CommonClient) CheckSQS(ctx context.Context, request *CheckSQSRequest) (*StreamResponse[CheckSQSResponse], error) {
	var result CheckSQSResponse
	res, err := MakeRequest[CheckSQSRequest, CheckSQSResponse, any](c.client, ctx, "POST", "/api/v2/check_sqs", nil, request, &result, nil)
	return res, err
}

// Deletes one device
func (c *CommonClient) DeleteDevice(ctx context.Context, queryParams *DeleteDeviceParams) (*StreamResponse[Response], error) {
	var result Response
	params, err := ToMap(queryParams)
	if err != nil {
		return nil, err
	}
	res, err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/devices", params, nil, &result, nil)
	return res, err
}

// Returns all available devices
func (c *CommonClient) ListDevices(ctx context.Context, queryParams *ListDevicesParams) (*StreamResponse[ListDevicesResponse], error) {
	var result ListDevicesResponse
	params, err := ToMap(queryParams)
	if err != nil {
		return nil, err
	}
	res, err := MakeRequest[any, ListDevicesResponse](c.client, ctx, "GET", "/api/v2/devices", params, nil, &result, nil)
	return res, err
}

// Adds a new device to a user, if the same device already exists the call will have no effect
func (c *CommonClient) CreateDevice(ctx context.Context, request *CreateDeviceRequest) (*StreamResponse[Response], error) {
	var result Response
	res, err := MakeRequest[CreateDeviceRequest, Response, any](c.client, ctx, "POST", "/api/v2/devices", nil, request, &result, nil)
	return res, err
}

// Exports user profile, reactions and messages for list of given users
func (c *CommonClient) ExportUsers(ctx context.Context, request *ExportUsersRequest) (*StreamResponse[ExportUsersResponse], error) {
	var result ExportUsersResponse
	res, err := MakeRequest[ExportUsersRequest, ExportUsersResponse, any](c.client, ctx, "POST", "/api/v2/export/users", nil, request, &result, nil)
	return res, err
}

// Lists external storage
func (c *CommonClient) ListExternalStorage(ctx context.Context) (*StreamResponse[ListExternalStorageResponse], error) {
	var result ListExternalStorageResponse
	res, err := MakeRequest[any, ListExternalStorageResponse, any](c.client, ctx, "GET", "/api/v2/external_storage", nil, nil, &result, nil)
	return res, err
}

// Creates new external storage
func (c *CommonClient) CreateExternalStorage(ctx context.Context, request *CreateExternalStorageRequest) (*StreamResponse[CreateExternalStorageResponse], error) {
	var result CreateExternalStorageResponse
	res, err := MakeRequest[CreateExternalStorageRequest, CreateExternalStorageResponse, any](c.client, ctx, "POST", "/api/v2/external_storage", nil, request, &result, nil)
	return res, err
}

// Deletes external storage
func (c *CommonClient) DeleteExternalStorage(ctx context.Context, name string) (*StreamResponse[DeleteExternalStorageResponse], error) {
	var result DeleteExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, DeleteExternalStorageResponse, any](c.client, ctx, "DELETE", "/api/v2/external_storage/{name}", nil, nil, &result, pathParams)
	return res, err
}

func (c *CommonClient) UpdateExternalStorage(ctx context.Context, name string, request *UpdateExternalStorageRequest) (*StreamResponse[UpdateExternalStorageResponse], error) {
	var result UpdateExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[UpdateExternalStorageRequest, UpdateExternalStorageResponse, any](c.client, ctx, "PUT", "/api/v2/external_storage/{name}", nil, request, &result, pathParams)
	return res, err
}

func (c *CommonClient) CheckExternalStorage(ctx context.Context, name string) (*StreamResponse[CheckExternalStorageResponse], error) {
	var result CheckExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, CheckExternalStorageResponse, any](c.client, ctx, "GET", "/api/v2/external_storage/{name}/check", nil, nil, &result, pathParams)
	return res, err
}

func (c *CommonClient) CreateGuest(ctx context.Context, request *CreateGuestRequest) (*StreamResponse[CreateGuestResponse], error) {
	var result CreateGuestResponse
	res, err := MakeRequest[CreateGuestRequest, CreateGuestResponse, any](c.client, ctx, "POST", "/api/v2/guest", nil, request, &result, nil)
	return res, err
}

// Creates a new import URL
func (c *CommonClient) CreateImportURL(ctx context.Context, request *CreateImportURLRequest) (*StreamResponse[CreateImportURLResponse], error) {
	var result CreateImportURLResponse
	res, err := MakeRequest[CreateImportURLRequest, CreateImportURLResponse, any](c.client, ctx, "POST", "/api/v2/import_urls", nil, request, &result, nil)
	return res, err
}

// Gets an import
func (c *CommonClient) ListImports(ctx context.Context) (*StreamResponse[ListImportsResponse], error) {
	var result ListImportsResponse
	res, err := MakeRequest[any, ListImportsResponse, any](c.client, ctx, "GET", "/api/v2/imports", nil, nil, &result, nil)
	return res, err
}

// Creates a new import
func (c *CommonClient) CreateImport(ctx context.Context, request *CreateImportRequest) (*StreamResponse[CreateImportResponse], error) {
	var result CreateImportResponse
	res, err := MakeRequest[CreateImportRequest, CreateImportResponse, any](c.client, ctx, "POST", "/api/v2/imports", nil, request, &result, nil)
	return res, err
}

// Gets an import
func (c *CommonClient) GetImport(ctx context.Context, id string) (*StreamResponse[GetImportResponse], error) {
	var result GetImportResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[any, GetImportResponse, any](c.client, ctx, "GET", "/api/v2/imports/{id}", nil, nil, &result, pathParams)
	return res, err
}

// Removes previously applied ban
//
// Sends events:
// - user.unbanned
//
// Required permissions:
// - BanChannelMember
// - BanUser
func (c *CommonClient) Unban(ctx context.Context, queryParams *UnbanParams) (*StreamResponse[Response], error) {
	var result Response
	params, err := ToMap(queryParams)
	if err != nil {
		return nil, err
	}
	res, err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/moderation/ban", params, nil, &result, nil)
	return res, err
}

// Restricts user activity either in specific channel or globally
//
// Sends events:
// - user.banned
//
// Required permissions:
// - BanChannelMember
// - BanUser
func (c *CommonClient) Ban(ctx context.Context, request *BanRequest) (*StreamResponse[Response], error) {
	var result Response
	res, err := MakeRequest[BanRequest, Response, any](c.client, ctx, "POST", "/api/v2/moderation/ban", nil, request, &result, nil)
	return res, err
}

// Reports message or user for review by moderators
//
// Sends events:
// - message.flagged
// - user.flagged
//
// Required permissions:
// - FlagMessage
// - FlagUser
func (c *CommonClient) Flag(ctx context.Context, request *FlagRequest) (*StreamResponse[FlagResponse], error) {
	var result FlagResponse
	res, err := MakeRequest[FlagRequest, FlagResponse, any](c.client, ctx, "POST", "/api/v2/moderation/flag", nil, request, &result, nil)
	return res, err
}

// Mutes one or several users
//
// Sends events:
// - user.muted
//
// Required permissions:
// - MuteUser
func (c *CommonClient) MuteUser(ctx context.Context, request *MuteUserRequest) (*StreamResponse[MuteUserResponse], error) {
	var result MuteUserResponse
	res, err := MakeRequest[MuteUserRequest, MuteUserResponse, any](c.client, ctx, "POST", "/api/v2/moderation/mute", nil, request, &result, nil)
	return res, err
}

// Unmutes previously muted user
//
// Sends events:
// - user.unmuted
//
// Required permissions:
// - MuteUser
func (c *CommonClient) UnmuteUser(ctx context.Context, request *UnmuteUserRequest) (*StreamResponse[UnmuteResponse], error) {
	var result UnmuteResponse
	res, err := MakeRequest[UnmuteUserRequest, UnmuteResponse, any](c.client, ctx, "POST", "/api/v2/moderation/unmute", nil, request, &result, nil)
	return res, err
}

// Get an OpenGraph attachment for a link
func (c *CommonClient) GetOG(ctx context.Context, queryParams *GetOGParams) (*StreamResponse[GetOGResponse], error) {
	var result GetOGResponse
	params, err := ToMap(queryParams)
	if err != nil {
		return nil, err
	}
	res, err := MakeRequest[any, GetOGResponse](c.client, ctx, "GET", "/api/v2/og", params, nil, &result, nil)
	return res, err
}

// Lists all available permissions
func (c *CommonClient) ListPermissions(ctx context.Context) (*StreamResponse[ListPermissionsResponse], error) {
	var result ListPermissionsResponse
	res, err := MakeRequest[any, ListPermissionsResponse, any](c.client, ctx, "GET", "/api/v2/permissions", nil, nil, &result, nil)
	return res, err
}

// Gets custom permission
func (c *CommonClient) GetPermission(ctx context.Context, id string) (*StreamResponse[GetCustomPermissionResponse], error) {
	var result GetCustomPermissionResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[any, GetCustomPermissionResponse, any](c.client, ctx, "GET", "/api/v2/permissions/{id}", nil, nil, &result, pathParams)
	return res, err
}

// List details of all push providers.
func (c *CommonClient) ListPushProviders(ctx context.Context) (*StreamResponse[ListPushProvidersResponse], error) {
	var result ListPushProvidersResponse
	res, err := MakeRequest[any, ListPushProvidersResponse, any](c.client, ctx, "GET", "/api/v2/push_providers", nil, nil, &result, nil)
	return res, err
}

// Upsert a push provider for v2 with multi bundle/package support
func (c *CommonClient) UpsertPushProvider(ctx context.Context, request *UpsertPushProviderRequest) (*StreamResponse[UpsertPushProviderResponse], error) {
	var result UpsertPushProviderResponse
	res, err := MakeRequest[UpsertPushProviderRequest, UpsertPushProviderResponse, any](c.client, ctx, "POST", "/api/v2/push_providers", nil, request, &result, nil)
	return res, err
}

// Delete a push provider from v2 with multi bundle/package support. v1 isn't supported in this endpoint
func (c *CommonClient) DeletePushProvider(ctx context.Context, _type string, name string) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"type": _type,
		"name": name,
	}
	res, err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/push_providers/{type}/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Get rate limits usage and quotas
func (c *CommonClient) GetRateLimits(ctx context.Context, queryParams *GetRateLimitsParams) (*StreamResponse[GetRateLimitsResponse], error) {
	var result GetRateLimitsResponse
	params, err := ToMap(queryParams)
	if err != nil {
		return nil, err
	}
	res, err := MakeRequest[any, GetRateLimitsResponse](c.client, ctx, "GET", "/api/v2/rate_limits", params, nil, &result, nil)
	return res, err
}

// Lists all available roles
func (c *CommonClient) ListRoles(ctx context.Context) (*StreamResponse[ListRolesResponse], error) {
	var result ListRolesResponse
	res, err := MakeRequest[any, ListRolesResponse, any](c.client, ctx, "GET", "/api/v2/roles", nil, nil, &result, nil)
	return res, err
}

// Creates custom role
func (c *CommonClient) CreateRole(ctx context.Context, request *CreateRoleRequest) (*StreamResponse[CreateRoleResponse], error) {
	var result CreateRoleResponse
	res, err := MakeRequest[CreateRoleRequest, CreateRoleResponse, any](c.client, ctx, "POST", "/api/v2/roles", nil, request, &result, nil)
	return res, err
}

// Deletes custom role
func (c *CommonClient) DeleteRole(ctx context.Context, name string) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/roles/{name}", nil, nil, &result, pathParams)
	return res, err
}

// Gets status of a task
func (c *CommonClient) GetTask(ctx context.Context, id string) (*StreamResponse[GetTaskResponse], error) {
	var result GetTaskResponse
	pathParams := map[string]string{
		"id": id,
	}
	res, err := MakeRequest[any, GetTaskResponse, any](c.client, ctx, "GET", "/api/v2/tasks/{id}", nil, nil, &result, pathParams)
	return res, err
}

// Find and filter users
//
// Required permissions:
// - SearchUser
func (c *CommonClient) QueryUsers(ctx context.Context, queryParams *QueryUsersParams) (*StreamResponse[QueryUsersResponse], error) {
	var result QueryUsersResponse
	params, err := ToMap(queryParams)
	if err != nil {
		return nil, err
	}
	res, err := MakeRequest[any, QueryUsersResponse](c.client, ctx, "GET", "/api/v2/users", params, nil, &result, nil)
	return res, err
}

// Updates certain fields of the user
//
// Sends events:
// - user.presence.changed
// - user.updated
func (c *CommonClient) UpdateUsersPartial(ctx context.Context, request *UpdateUsersPartialRequest) (*StreamResponse[UpdateUsersResponse], error) {
	var result UpdateUsersResponse
	res, err := MakeRequest[UpdateUsersPartialRequest, UpdateUsersResponse, any](c.client, ctx, "PATCH", "/api/v2/users", nil, request, &result, nil)
	return res, err
}

// Update or create users in bulk
//
// Sends events:
// - user.updated
func (c *CommonClient) UpdateUsers(ctx context.Context, request *UpdateUsersRequest) (*StreamResponse[UpdateUsersResponse], error) {
	var result UpdateUsersResponse
	res, err := MakeRequest[UpdateUsersRequest, UpdateUsersResponse, any](c.client, ctx, "POST", "/api/v2/users", nil, request, &result, nil)
	return res, err
}

// Get list of blocked Users
func (c *CommonClient) GetBlockedUsers(ctx context.Context, queryParams *GetBlockedUsersParams) (*StreamResponse[GetBlockedUsersResponse], error) {
	var result GetBlockedUsersResponse
	params, err := ToMap(queryParams)
	if err != nil {
		return nil, err
	}
	res, err := MakeRequest[any, GetBlockedUsersResponse](c.client, ctx, "GET", "/api/v2/users/block", params, nil, &result, nil)
	return res, err
}

// Block users
func (c *CommonClient) BlockUsers(ctx context.Context, request *BlockUsersRequest) (*StreamResponse[BlockUsersResponse], error) {
	var result BlockUsersResponse
	res, err := MakeRequest[BlockUsersRequest, BlockUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/block", nil, request, &result, nil)
	return res, err
}

// Deactivate users in batches
//
// Sends events:
// - user.deactivated
func (c *CommonClient) DeactivateUsers(ctx context.Context, request *DeactivateUsersRequest) (*StreamResponse[DeactivateUsersResponse], error) {
	var result DeactivateUsersResponse
	res, err := MakeRequest[DeactivateUsersRequest, DeactivateUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/deactivate", nil, request, &result, nil)
	return res, err
}

// Deletes users and optionally all their belongings asynchronously.
//
// Sends events:
// - channel.deleted
// - user.deleted
func (c *CommonClient) DeleteUsers(ctx context.Context, request *DeleteUsersRequest) (*StreamResponse[DeleteUsersResponse], error) {
	var result DeleteUsersResponse
	res, err := MakeRequest[DeleteUsersRequest, DeleteUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/delete", nil, request, &result, nil)
	return res, err
}

// Reactivate users in batches
//
// Sends events:
// - user.reactivated
func (c *CommonClient) ReactivateUsers(ctx context.Context, request *ReactivateUsersRequest) (*StreamResponse[ReactivateUsersResponse], error) {
	var result ReactivateUsersResponse
	res, err := MakeRequest[ReactivateUsersRequest, ReactivateUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/reactivate", nil, request, &result, nil)
	return res, err
}

// Restore soft deleted users
func (c *CommonClient) RestoreUsers(ctx context.Context, request *RestoreUsersRequest) (*StreamResponse[Response], error) {
	var result Response
	res, err := MakeRequest[RestoreUsersRequest, Response, any](c.client, ctx, "POST", "/api/v2/users/restore", nil, request, &result, nil)
	return res, err
}

// Unblock users
func (c *CommonClient) UnblockUsers(ctx context.Context, request *UnblockUsersRequest) (*StreamResponse[UnblockUsersResponse], error) {
	var result UnblockUsersResponse
	res, err := MakeRequest[UnblockUsersRequest, UnblockUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/unblock", nil, request, &result, nil)
	return res, err
}

// Deactivates user with possibility to activate it back
//
// Sends events:
// - user.deactivated
func (c *CommonClient) DeactivateUser(ctx context.Context, userId string, request *DeactivateUserRequest) (*StreamResponse[DeactivateUserResponse], error) {
	var result DeactivateUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	res, err := MakeRequest[DeactivateUserRequest, DeactivateUserResponse, any](c.client, ctx, "POST", "/api/v2/users/{user_id}/deactivate", nil, request, &result, pathParams)
	return res, err
}

// Exports the user's profile, reactions and messages. Raises an error if a user has more than 10k messages or reactions
func (c *CommonClient) ExportUser(ctx context.Context, userId string) (*StreamResponse[ExportUserResponse], error) {
	var result ExportUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	res, err := MakeRequest[any, ExportUserResponse, any](c.client, ctx, "GET", "/api/v2/users/{user_id}/export", nil, nil, &result, pathParams)
	return res, err
}

// Activates user who's been deactivated previously
//
// Sends events:
// - user.reactivated
func (c *CommonClient) ReactivateUser(ctx context.Context, userId string, request *ReactivateUserRequest) (*StreamResponse[ReactivateUserResponse], error) {
	var result ReactivateUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	res, err := MakeRequest[ReactivateUserRequest, ReactivateUserResponse, any](c.client, ctx, "POST", "/api/v2/users/{user_id}/reactivate", nil, request, &result, pathParams)
	return res, err
}
