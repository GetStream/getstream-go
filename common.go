package getstream

import (
	"context"
	"net/http"
	"net/url"

	"github.com/getstream/go-client/models"
	"github.com/getstream/go-client"
)


type CommonClient struct {
	client *Client
}

func NewCommonClient(client *Client) *CommonClient {
	return &CommonClient{
		client: NewClient(client),
	}
}


func (c *CommonClient) GetApp(ctx context.Context) (GetApplicationResponse, error) {
	var result GetApplicationResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/app", nil, nil, &result,)
	return result, err
}


func (c *CommonClient) UpdateApp(ctx context.Context,
	UpdateAppRequest UpdateAppRequest
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "PATCH", "/api/v2/app", nil,UpdateAppRequest, &result,)
	return result, err
}


func (c *CommonClient) ListBlockLists(ctx context.Context) (ListBlockListResponse, error) {
	var result ListBlockListResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/blocklists", nil, nil, &result,)
	return result, err
}


func (c *CommonClient) CreateBlockList(ctx context.Context,
	CreateBlockListRequest CreateBlockListRequest
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/blocklists", nil,CreateBlockListRequest, &result,)
	return result, err
}


func (c *CommonClient) DeleteBlockList(ctx context.Context,
	Name string
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/blocklists/{name}", nil, nil, &result, Name,)
	return result, err
}


func (c *CommonClient) GetBlockList(ctx context.Context,
	Name string
) (GetBlockListResponse, error) {
	var result GetBlockListResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/blocklists/{name}", nil, nil, &result, Name,)
	return result, err
}


func (c *CommonClient) UpdateBlockList(ctx context.Context,
	Name string
	, UpdateBlockListRequest UpdateBlockListRequest
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "PUT", "/api/v2/blocklists/{name}", nil,UpdateBlockListRequest, &result, Name,)
	return result, err
}


func (c *CommonClient) CheckPush(ctx context.Context,
	CheckPushRequest CheckPushRequest
) (CheckPushResponse, error) {
	var result CheckPushResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/check_push", nil,CheckPushRequest, &result,)
	return result, err
}


func (c *CommonClient) CheckSNS(ctx context.Context,
	CheckSNSRequest CheckSNSRequest
) (CheckSNSResponse, error) {
	var result CheckSNSResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/check_sns", nil,CheckSNSRequest, &result,)
	return result, err
}


func (c *CommonClient) CheckSQS(ctx context.Context,
	CheckSQSRequest CheckSQSRequest
) (CheckSQSResponse, error) {
	var result CheckSQSResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/check_sqs", nil,CheckSQSRequest, &result,)
	return result, err
}


func (c *CommonClient) DeleteDevice(ctx context.Context,
	Id string
	, UserId *string
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/devices",url.Values{"Id": []string { Id }, "UserId": []string { UserId } }, nil, &result,)
	return result, err
}


func (c *CommonClient) ListDevices(ctx context.Context,
	UserId *string
) (ListDevicesResponse, error) {
	var result ListDevicesResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/devices",url.Values{"UserId": []string { UserId } }, nil, &result,)
	return result, err
}


func (c *CommonClient) CreateDevice(ctx context.Context,
	CreateDeviceRequest CreateDeviceRequest
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/devices", nil,CreateDeviceRequest, &result,)
	return result, err
}


func (c *CommonClient) ExportUsers(ctx context.Context,
	ExportUsersRequest ExportUsersRequest
) (ExportUsersResponse, error) {
	var result ExportUsersResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/export/users", nil,ExportUsersRequest, &result,)
	return result, err
}


func (c *CommonClient) ListExternalStorage(ctx context.Context) (ListExternalStorageResponse, error) {
	var result ListExternalStorageResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/external_storage", nil, nil, &result,)
	return result, err
}


func (c *CommonClient) CreateExternalStorage(ctx context.Context,
	CreateExternalStorageRequest CreateExternalStorageRequest
) (CreateExternalStorageResponse, error) {
	var result CreateExternalStorageResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/external_storage", nil,CreateExternalStorageRequest, &result,)
	return result, err
}


func (c *CommonClient) DeleteExternalStorage(ctx context.Context,
	Name string
) (DeleteExternalStorageResponse, error) {
	var result DeleteExternalStorageResponse
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/external_storage/{name}", nil, nil, &result, Name,)
	return result, err
}


func (c *CommonClient) UpdateExternalStorage(ctx context.Context,
	Name string
	, UpdateExternalStorageRequest UpdateExternalStorageRequest
) (UpdateExternalStorageResponse, error) {
	var result UpdateExternalStorageResponse
	err := MakeRequest(c.client, ctx, "PUT", "/api/v2/external_storage/{name}", nil,UpdateExternalStorageRequest, &result, Name,)
	return result, err
}


func (c *CommonClient) CheckExternalStorage(ctx context.Context,
	Name string
) (CheckExternalStorageResponse, error) {
	var result CheckExternalStorageResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/external_storage/{name}/check", nil, nil, &result, Name,)
	return result, err
}


func (c *CommonClient) CreateGuest(ctx context.Context,
	CreateGuestRequest CreateGuestRequest
) (CreateGuestResponse, error) {
	var result CreateGuestResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/guest", nil,CreateGuestRequest, &result,)
	return result, err
}


func (c *CommonClient) CreateImportURL(ctx context.Context,
	CreateImportURLRequest CreateImportURLRequest
) (CreateImportURLResponse, error) {
	var result CreateImportURLResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/import_urls", nil,CreateImportURLRequest, &result,)
	return result, err
}


func (c *CommonClient) ListImports(ctx context.Context) (ListImportsResponse, error) {
	var result ListImportsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/imports", nil, nil, &result,)
	return result, err
}


func (c *CommonClient) CreateImport(ctx context.Context,
	CreateImportRequest CreateImportRequest
) (CreateImportResponse, error) {
	var result CreateImportResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/imports", nil,CreateImportRequest, &result,)
	return result, err
}


func (c *CommonClient) GetImport(ctx context.Context,
	Id string
) (GetImportResponse, error) {
	var result GetImportResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/imports/{id}", nil, nil, &result, Id,)
	return result, err
}


func (c *CommonClient) Unban(ctx context.Context,
	TargetUserId string
	, ChannelCid *string
	, CreatedBy *string
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/moderation/ban",url.Values{"TargetUserId": []string { TargetUserId }, "ChannelCid": []string { ChannelCid }, "CreatedBy": []string { CreatedBy } }, nil, &result,)
	return result, err
}


func (c *CommonClient) Ban(ctx context.Context,
	BanRequest BanRequest
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/moderation/ban", nil,BanRequest, &result,)
	return result, err
}


func (c *CommonClient) Flag(ctx context.Context,
	FlagRequest FlagRequest
) (FlagResponse, error) {
	var result FlagResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/moderation/flag", nil,FlagRequest, &result,)
	return result, err
}


func (c *CommonClient) MuteUser(ctx context.Context,
	MuteUserRequest MuteUserRequest
) (MuteUserResponse, error) {
	var result MuteUserResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/moderation/mute", nil,MuteUserRequest, &result,)
	return result, err
}


func (c *CommonClient) UnmuteUser(ctx context.Context,
	UnmuteUserRequest UnmuteUserRequest
) (UnmuteResponse, error) {
	var result UnmuteResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/moderation/unmute", nil,UnmuteUserRequest, &result,)
	return result, err
}


func (c *CommonClient) GetOG(ctx context.Context,
	Url string
) (GetOGResponse, error) {
	var result GetOGResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/og",url.Values{"Url": []string { Url } }, nil, &result,)
	return result, err
}


func (c *CommonClient) ListPermissions(ctx context.Context) (ListPermissionsResponse, error) {
	var result ListPermissionsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/permissions", nil, nil, &result,)
	return result, err
}


func (c *CommonClient) GetPermission(ctx context.Context,
	Id string
) (GetCustomPermissionResponse, error) {
	var result GetCustomPermissionResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/permissions/{id}", nil, nil, &result, Id,)
	return result, err
}


func (c *CommonClient) ListPushProviders(ctx context.Context) (ListPushProvidersResponse, error) {
	var result ListPushProvidersResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/push_providers", nil, nil, &result,)
	return result, err
}


func (c *CommonClient) UpsertPushProvider(ctx context.Context,
	UpsertPushProviderRequest UpsertPushProviderRequest
) (UpsertPushProviderResponse, error) {
	var result UpsertPushProviderResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/push_providers", nil,UpsertPushProviderRequest, &result,)
	return result, err
}


func (c *CommonClient) DeletePushProvider(ctx context.Context,
	Type string
	, Name string
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/push_providers/{type}/{name}", nil, nil, &result, Type, Name,)
	return result, err
}


func (c *CommonClient) GetRateLimits(ctx context.Context,
	ServerSide *bool
	, Android *bool
	, Ios *bool
	, Web *bool
	, Endpoints *string
) (GetRateLimitsResponse, error) {
	var result GetRateLimitsResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/rate_limits",url.Values{"ServerSide": []string { ServerSide }, "Android": []string { Android }, "Ios": []string { Ios }, "Web": []string { Web }, "Endpoints": []string { Endpoints } }, nil, &result,)
	return result, err
}


func (c *CommonClient) ListRoles(ctx context.Context) (ListRolesResponse, error) {
	var result ListRolesResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/roles", nil, nil, &result,)
	return result, err
}


func (c *CommonClient) CreateRole(ctx context.Context,
	CreateRoleRequest CreateRoleRequest
) (CreateRoleResponse, error) {
	var result CreateRoleResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/roles", nil,CreateRoleRequest, &result,)
	return result, err
}


func (c *CommonClient) DeleteRole(ctx context.Context,
	Name string
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "DELETE", "/api/v2/roles/{name}", nil, nil, &result, Name,)
	return result, err
}


func (c *CommonClient) GetTask(ctx context.Context,
	Id string
) (GetTaskResponse, error) {
	var result GetTaskResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/tasks/{id}", nil, nil, &result, Id,)
	return result, err
}


func (c *CommonClient) QueryUsers(ctx context.Context,
	Payload *QueryUsersPayload
) (QueryUsersResponse, error) {
	var result QueryUsersResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/users",url.Values{"Payload": []string { Payload } }, nil, &result,)
	return result, err
}


func (c *CommonClient) UpdateUsersPartial(ctx context.Context,
	UpdateUsersPartialRequest UpdateUsersPartialRequest
) (UpdateUsersResponse, error) {
	var result UpdateUsersResponse
	err := MakeRequest(c.client, ctx, "PATCH", "/api/v2/users", nil,UpdateUsersPartialRequest, &result,)
	return result, err
}


func (c *CommonClient) UpdateUsers(ctx context.Context,
	UpdateUsersRequest UpdateUsersRequest
) (UpdateUsersResponse, error) {
	var result UpdateUsersResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/users", nil,UpdateUsersRequest, &result,)
	return result, err
}


func (c *CommonClient) DeactivateUsers(ctx context.Context,
	DeactivateUsersRequest DeactivateUsersRequest
) (DeactivateUsersResponse, error) {
	var result DeactivateUsersResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/users/deactivate", nil,DeactivateUsersRequest, &result,)
	return result, err
}


func (c *CommonClient) DeleteUsers(ctx context.Context,
	DeleteUsersRequest DeleteUsersRequest
) (DeleteUsersResponse, error) {
	var result DeleteUsersResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/users/delete", nil,DeleteUsersRequest, &result,)
	return result, err
}


func (c *CommonClient) ReactivateUsers(ctx context.Context,
	ReactivateUsersRequest ReactivateUsersRequest
) (ReactivateUsersResponse, error) {
	var result ReactivateUsersResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/users/reactivate", nil,ReactivateUsersRequest, &result,)
	return result, err
}


func (c *CommonClient) RestoreUsers(ctx context.Context,
	RestoreUsersRequest RestoreUsersRequest
) (Response, error) {
	var result Response
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/users/restore", nil,RestoreUsersRequest, &result,)
	return result, err
}


func (c *CommonClient) DeactivateUser(ctx context.Context,
	UserId string
	, DeactivateUserRequest DeactivateUserRequest
) (DeactivateUserResponse, error) {
	var result DeactivateUserResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/users/{user_id}/deactivate", nil,DeactivateUserRequest, &result, UserId,)
	return result, err
}


func (c *CommonClient) ExportUser(ctx context.Context,
	UserId string
) (ExportUserResponse, error) {
	var result ExportUserResponse
	err := MakeRequest(c.client, ctx, "GET", "/api/v2/users/{user_id}/export", nil, nil, &result, UserId,)
	return result, err
}


func (c *CommonClient) ReactivateUser(ctx context.Context,
	UserId string
	, ReactivateUserRequest ReactivateUserRequest
) (ReactivateUserResponse, error) {
	var result ReactivateUserResponse
	err := MakeRequest(c.client, ctx, "POST", "/api/v2/users/{user_id}/reactivate", nil,ReactivateUserRequest, &result, UserId,)
	return result, err
}










