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
func (c *CommonClient) GetApp(ctx context.Context) (*GetApplicationResponse, error) {
	var result GetApplicationResponse
	err := MakeRequest[any, GetApplicationResponse, any](c.client, ctx, "GET", "/api/v2/app", nil, nil, &result, nil)
	return &result, err
}

// This Method updates one or more application settings
func (c *CommonClient) UpdateApp(ctx context.Context, request *UpdateAppRequest) (*Response, error) {
	var result Response
	err := MakeRequest[UpdateAppRequest, Response, any](c.client, ctx, "PATCH", "/api/v2/app", nil, request, &result, nil)
	return &result, err
}

// Returns all available block lists
func (c *CommonClient) ListBlockLists(ctx context.Context) (*ListBlockListResponse, error) {
	var result ListBlockListResponse
	err := MakeRequest[any, ListBlockListResponse, any](c.client, ctx, "GET", "/api/v2/blocklists", nil, nil, &result, nil)
	return &result, err
}

// Creates a new application blocklist, once created the blocklist can be used by any channel type
func (c *CommonClient) CreateBlockList(ctx context.Context, request *CreateBlockListRequest) (*Response, error) {
	var result Response
	err := MakeRequest[CreateBlockListRequest, Response, any](c.client, ctx, "POST", "/api/v2/blocklists", nil, request, &result, nil)
	return &result, err
}

// Deletes previously created application blocklist
func (c *CommonClient) DeleteBlockList(ctx context.Context, name string) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/blocklists/{name}", nil, nil, &result, pathParams)
	return &result, err
}

// Returns block list by given name
func (c *CommonClient) GetBlockList(ctx context.Context, name string) (*GetBlockListResponse, error) {
	var result GetBlockListResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, GetBlockListResponse, any](c.client, ctx, "GET", "/api/v2/blocklists/{name}", nil, nil, &result, pathParams)
	return &result, err
}

// Updates contents of the block list
func (c *CommonClient) UpdateBlockList(ctx context.Context, name string, request *UpdateBlockListRequest) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[UpdateBlockListRequest, Response, any](c.client, ctx, "PUT", "/api/v2/blocklists/{name}", nil, request, &result, pathParams)
	return &result, err
}

// Sends a test message via push, this is a test endpoint to verify your push settings
func (c *CommonClient) CheckPush(ctx context.Context, request *CheckPushRequest) (*CheckPushResponse, error) {
	var result CheckPushResponse
	err := MakeRequest[CheckPushRequest, CheckPushResponse, any](c.client, ctx, "POST", "/api/v2/check_push", nil, request, &result, nil)
	return &result, err
}

// Validates Amazon SNS configuration
func (c *CommonClient) CheckSNS(ctx context.Context, request *CheckSNSRequest) (*CheckSNSResponse, error) {
	var result CheckSNSResponse
	err := MakeRequest[CheckSNSRequest, CheckSNSResponse, any](c.client, ctx, "POST", "/api/v2/check_sns", nil, request, &result, nil)
	return &result, err
}

// Validates Amazon SQS credentials
func (c *CommonClient) CheckSQS(ctx context.Context, request *CheckSQSRequest) (*CheckSQSResponse, error) {
	var result CheckSQSResponse
	err := MakeRequest[CheckSQSRequest, CheckSQSResponse, any](c.client, ctx, "POST", "/api/v2/check_sqs", nil, request, &result, nil)
	return &result, err
}

// Deletes one device
func (c *CommonClient) DeleteDevice(ctx context.Context, id string, userId *string) (*Response, error) {
	var result Response
	queryParams := map[string]interface{}{
		"id":      id,
		"user_id": userId,
	}
	err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/devices", queryParams, nil, &result, nil)
	return &result, err
}

// Returns all available devices
func (c *CommonClient) ListDevices(ctx context.Context, userId *string) (*ListDevicesResponse, error) {
	var result ListDevicesResponse
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, ListDevicesResponse](c.client, ctx, "GET", "/api/v2/devices", queryParams, nil, &result, nil)
	return &result, err
}

// Adds a new device to a user, if the same device already exists the call will have no effect
func (c *CommonClient) CreateDevice(ctx context.Context, request *CreateDeviceRequest) (*Response, error) {
	var result Response
	err := MakeRequest[CreateDeviceRequest, Response, any](c.client, ctx, "POST", "/api/v2/devices", nil, request, &result, nil)
	return &result, err
}

// Exports user profile, reactions and messages for list of given users
func (c *CommonClient) ExportUsers(ctx context.Context, request *ExportUsersRequest) (*ExportUsersResponse, error) {
	var result ExportUsersResponse
	err := MakeRequest[ExportUsersRequest, ExportUsersResponse, any](c.client, ctx, "POST", "/api/v2/export/users", nil, request, &result, nil)
	return &result, err
}

// Lists external storage
func (c *CommonClient) ListExternalStorage(ctx context.Context) (*ListExternalStorageResponse, error) {
	var result ListExternalStorageResponse
	err := MakeRequest[any, ListExternalStorageResponse, any](c.client, ctx, "GET", "/api/v2/external_storage", nil, nil, &result, nil)
	return &result, err
}

// Creates new external storage
func (c *CommonClient) CreateExternalStorage(ctx context.Context, request *CreateExternalStorageRequest) (*CreateExternalStorageResponse, error) {
	var result CreateExternalStorageResponse
	err := MakeRequest[CreateExternalStorageRequest, CreateExternalStorageResponse, any](c.client, ctx, "POST", "/api/v2/external_storage", nil, request, &result, nil)
	return &result, err
}

// Deletes external storage
func (c *CommonClient) DeleteExternalStorage(ctx context.Context, name string) (*DeleteExternalStorageResponse, error) {
	var result DeleteExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, DeleteExternalStorageResponse, any](c.client, ctx, "DELETE", "/api/v2/external_storage/{name}", nil, nil, &result, pathParams)
	return &result, err
}

func (c *CommonClient) UpdateExternalStorage(ctx context.Context, name string, request *UpdateExternalStorageRequest) (*UpdateExternalStorageResponse, error) {
	var result UpdateExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[UpdateExternalStorageRequest, UpdateExternalStorageResponse, any](c.client, ctx, "PUT", "/api/v2/external_storage/{name}", nil, request, &result, pathParams)
	return &result, err
}

func (c *CommonClient) CheckExternalStorage(ctx context.Context, name string) (*CheckExternalStorageResponse, error) {
	var result CheckExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, CheckExternalStorageResponse, any](c.client, ctx, "GET", "/api/v2/external_storage/{name}/check", nil, nil, &result, pathParams)
	return &result, err
}

func (c *CommonClient) CreateGuest(ctx context.Context, request *CreateGuestRequest) (*CreateGuestResponse, error) {
	var result CreateGuestResponse
	err := MakeRequest[CreateGuestRequest, CreateGuestResponse, any](c.client, ctx, "POST", "/api/v2/guest", nil, request, &result, nil)
	return &result, err
}

// Creates a new import URL
func (c *CommonClient) CreateImportURL(ctx context.Context, request *CreateImportURLRequest) (*CreateImportURLResponse, error) {
	var result CreateImportURLResponse
	err := MakeRequest[CreateImportURLRequest, CreateImportURLResponse, any](c.client, ctx, "POST", "/api/v2/import_urls", nil, request, &result, nil)
	return &result, err
}

// Gets an import
func (c *CommonClient) ListImports(ctx context.Context) (*ListImportsResponse, error) {
	var result ListImportsResponse
	err := MakeRequest[any, ListImportsResponse, any](c.client, ctx, "GET", "/api/v2/imports", nil, nil, &result, nil)
	return &result, err
}

// Creates a new import
func (c *CommonClient) CreateImport(ctx context.Context, request *CreateImportRequest) (*CreateImportResponse, error) {
	var result CreateImportResponse
	err := MakeRequest[CreateImportRequest, CreateImportResponse, any](c.client, ctx, "POST", "/api/v2/imports", nil, request, &result, nil)
	return &result, err
}

// Gets an import
func (c *CommonClient) GetImport(ctx context.Context, id string) (*GetImportResponse, error) {
	var result GetImportResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[any, GetImportResponse, any](c.client, ctx, "GET", "/api/v2/imports/{id}", nil, nil, &result, pathParams)
	return &result, err
}

// Removes previously applied ban
//
// Sends events:
// - user.unbanned
//
// Required permissions:
// - BanChannelMember
// - BanUser
func (c *CommonClient) Unban(ctx context.Context, targetUserId string, channelCid *string, createdBy *string) (*Response, error) {
	var result Response
	queryParams := map[string]interface{}{
		"target_user_id": targetUserId,
		"created_by":     createdBy,
	}
	if channelCid != nil {
		queryParams["channel_cid"] = channelCid
	}
	err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/moderation/ban", queryParams, nil, &result, nil)
	return &result, err
}

// Restricts user activity either in specific channel or globally
//
// Sends events:
// - user.banned
//
// Required permissions:
// - BanChannelMember
// - BanUser
func (c *CommonClient) Ban(ctx context.Context, request *BanRequest) (*Response, error) {
	var result Response
	err := MakeRequest[BanRequest, Response, any](c.client, ctx, "POST", "/api/v2/moderation/ban", nil, request, &result, nil)
	return &result, err
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
func (c *CommonClient) Flag(ctx context.Context, request *FlagRequest) (*FlagResponse, error) {
	var result FlagResponse
	err := MakeRequest[FlagRequest, FlagResponse, any](c.client, ctx, "POST", "/api/v2/moderation/flag", nil, request, &result, nil)
	return &result, err
}

// Mutes one or several users
//
// Sends events:
// - user.muted
//
// Required permissions:
// - MuteUser
func (c *CommonClient) MuteUser(ctx context.Context, request *MuteUserRequest) (*MuteUserResponse, error) {
	var result MuteUserResponse
	err := MakeRequest[MuteUserRequest, MuteUserResponse, any](c.client, ctx, "POST", "/api/v2/moderation/mute", nil, request, &result, nil)
	return &result, err
}

// Unmutes previously muted user
//
// Sends events:
// - user.unmuted
//
// Required permissions:
// - MuteUser
func (c *CommonClient) UnmuteUser(ctx context.Context, request *UnmuteUserRequest) (*UnmuteResponse, error) {
	var result UnmuteResponse
	err := MakeRequest[UnmuteUserRequest, UnmuteResponse, any](c.client, ctx, "POST", "/api/v2/moderation/unmute", nil, request, &result, nil)
	return &result, err
}

// Get an OpenGraph attachment for a link
func (c *CommonClient) GetOG(ctx context.Context, url string) (*GetOGResponse, error) {
	var result GetOGResponse
	queryParams := map[string]interface{}{
		"url": url,
	}
	err := MakeRequest[any, GetOGResponse](c.client, ctx, "GET", "/api/v2/og", queryParams, nil, &result, nil)
	return &result, err
}

// Lists all available permissions
func (c *CommonClient) ListPermissions(ctx context.Context) (*ListPermissionsResponse, error) {
	var result ListPermissionsResponse
	err := MakeRequest[any, ListPermissionsResponse, any](c.client, ctx, "GET", "/api/v2/permissions", nil, nil, &result, nil)
	return &result, err
}

// Gets custom permission
func (c *CommonClient) GetPermission(ctx context.Context, id string) (*GetCustomPermissionResponse, error) {
	var result GetCustomPermissionResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[any, GetCustomPermissionResponse, any](c.client, ctx, "GET", "/api/v2/permissions/{id}", nil, nil, &result, pathParams)
	return &result, err
}

// List details of all push providers.
func (c *CommonClient) ListPushProviders(ctx context.Context) (*ListPushProvidersResponse, error) {
	var result ListPushProvidersResponse
	err := MakeRequest[any, ListPushProvidersResponse, any](c.client, ctx, "GET", "/api/v2/push_providers", nil, nil, &result, nil)
	return &result, err
}

// Upsert a push provider for v2 with multi bundle/package support
func (c *CommonClient) UpsertPushProvider(ctx context.Context, request *UpsertPushProviderRequest) (*UpsertPushProviderResponse, error) {
	var result UpsertPushProviderResponse
	err := MakeRequest[UpsertPushProviderRequest, UpsertPushProviderResponse, any](c.client, ctx, "POST", "/api/v2/push_providers", nil, request, &result, nil)
	return &result, err
}

// Delete a push provider from v2 with multi bundle/package support. v1 isn't supported in this endpoint
func (c *CommonClient) DeletePushProvider(ctx context.Context, _type string, name string) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"type": _type,
		"name": name,
	}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/push_providers/{type}/{name}", nil, nil, &result, pathParams)
	return &result, err
}

// Get rate limits usage and quotas
func (c *CommonClient) GetRateLimits(ctx context.Context, serverSide *bool, android *bool, ios *bool, web *bool, endpoints *string) (*GetRateLimitsResponse, error) {
	var result GetRateLimitsResponse
	queryParams := map[string]interface{}{
		"server_side": serverSide,
		"android":     android,
		"ios":         ios,
		"web":         web,
		"endpoints":   endpoints,
	}
	err := MakeRequest[any, GetRateLimitsResponse](c.client, ctx, "GET", "/api/v2/rate_limits", queryParams, nil, &result, nil)
	return &result, err
}

// Lists all available roles
func (c *CommonClient) ListRoles(ctx context.Context) (*ListRolesResponse, error) {
	var result ListRolesResponse
	err := MakeRequest[any, ListRolesResponse, any](c.client, ctx, "GET", "/api/v2/roles", nil, nil, &result, nil)
	return &result, err
}

// Creates custom role
func (c *CommonClient) CreateRole(ctx context.Context, request *CreateRoleRequest) (*CreateRoleResponse, error) {
	var result CreateRoleResponse
	err := MakeRequest[CreateRoleRequest, CreateRoleResponse, any](c.client, ctx, "POST", "/api/v2/roles", nil, request, &result, nil)
	return &result, err
}

// Deletes custom role
func (c *CommonClient) DeleteRole(ctx context.Context, name string) (*Response, error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/roles/{name}", nil, nil, &result, pathParams)
	return &result, err
}

// Gets status of a task
func (c *CommonClient) GetTask(ctx context.Context, id string) (*GetTaskResponse, error) {
	var result GetTaskResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[any, GetTaskResponse, any](c.client, ctx, "GET", "/api/v2/tasks/{id}", nil, nil, &result, pathParams)
	return &result, err
}

// Find and filter users
//
// Required permissions:
// - SearchUser
func (c *CommonClient) QueryUsers(ctx context.Context, payload *QueryUsersPayload) (*QueryUsersResponse, error) {
	var result QueryUsersResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, QueryUsersResponse](c.client, ctx, "GET", "/api/v2/users", queryParams, nil, &result, nil)
	return &result, err
}

// Updates certain fields of the user
//
// Sends events:
// - user.presence.changed
// - user.updated
func (c *CommonClient) UpdateUsersPartial(ctx context.Context, request *UpdateUsersPartialRequest) (*UpdateUsersResponse, error) {
	var result UpdateUsersResponse
	err := MakeRequest[UpdateUsersPartialRequest, UpdateUsersResponse, any](c.client, ctx, "PATCH", "/api/v2/users", nil, request, &result, nil)
	return &result, err
}

// Update or create users in bulk
//
// Sends events:
// - user.updated
func (c *CommonClient) UpdateUsers(ctx context.Context, request *UpdateUsersRequest) (*UpdateUsersResponse, error) {
	var result UpdateUsersResponse
	err := MakeRequest[UpdateUsersRequest, UpdateUsersResponse, any](c.client, ctx, "POST", "/api/v2/users", nil, request, &result, nil)
	return &result, err
}

// Deactivate users in batches
//
// Sends events:
// - user.deactivated
func (c *CommonClient) DeactivateUsers(ctx context.Context, request *DeactivateUsersRequest) (*DeactivateUsersResponse, error) {
	var result DeactivateUsersResponse
	err := MakeRequest[DeactivateUsersRequest, DeactivateUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/deactivate", nil, request, &result, nil)
	return &result, err
}

// Deletes users and optionally all their belongings asynchronously.
//
// Sends events:
// - channel.deleted
// - user.deleted
func (c *CommonClient) DeleteUsers(ctx context.Context, request *DeleteUsersRequest) (*DeleteUsersResponse, error) {
	var result DeleteUsersResponse
	err := MakeRequest[DeleteUsersRequest, DeleteUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/delete", nil, request, &result, nil)
	return &result, err
}

// Reactivate users in batches
//
// Sends events:
// - user.reactivated
func (c *CommonClient) ReactivateUsers(ctx context.Context, request *ReactivateUsersRequest) (*ReactivateUsersResponse, error) {
	var result ReactivateUsersResponse
	err := MakeRequest[ReactivateUsersRequest, ReactivateUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/reactivate", nil, request, &result, nil)
	return &result, err
}

// Restore soft deleted users
func (c *CommonClient) RestoreUsers(ctx context.Context, request *RestoreUsersRequest) (*Response, error) {
	var result Response
	err := MakeRequest[RestoreUsersRequest, Response, any](c.client, ctx, "POST", "/api/v2/users/restore", nil, request, &result, nil)
	return &result, err
}

// Deactivates user with possibility to activate it back
//
// Sends events:
// - user.deactivated
func (c *CommonClient) DeactivateUser(ctx context.Context, userId string, request *DeactivateUserRequest) (*DeactivateUserResponse, error) {
	var result DeactivateUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	err := MakeRequest[DeactivateUserRequest, DeactivateUserResponse, any](c.client, ctx, "POST", "/api/v2/users/{user_id}/deactivate", nil, request, &result, pathParams)
	return &result, err
}

// Exports the user's profile, reactions and messages. Raises an error if a user has more than 10k messages or reactions
func (c *CommonClient) ExportUser(ctx context.Context, userId string) (*ExportUserResponse, error) {
	var result ExportUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	err := MakeRequest[any, ExportUserResponse, any](c.client, ctx, "GET", "/api/v2/users/{user_id}/export", nil, nil, &result, pathParams)
	return &result, err
}

// Activates user who's been deactivated previously
//
// Sends events:
// - user.reactivated
func (c *CommonClient) ReactivateUser(ctx context.Context, userId string, request *ReactivateUserRequest) (*ReactivateUserResponse, error) {
	var result ReactivateUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	err := MakeRequest[ReactivateUserRequest, ReactivateUserResponse, any](c.client, ctx, "POST", "/api/v2/users/{user_id}/reactivate", nil, request, &result, pathParams)
	return &result, err
}
