// Code generated by GetStream internal OpenAPI code generator. DO NOT EDIT.
package getstream_test

import (
	"context"
	"testing"

	"github.com/GetStream/getstream-go"
	"github.com/stretchr/testify/require"
)

func TestCommonGetApp(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.GetApp(context.Background(), &getstream.GetAppRequest{})
	require.NoError(t, err)
}
func TestCommonUpdateApp(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.UpdateApp(context.Background(), &getstream.UpdateAppRequest{})
	require.NoError(t, err)
}
func TestCommonListBlockLists(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ListBlockLists(context.Background(), &getstream.ListBlockListsRequest{})
	require.NoError(t, err)
}
func TestCommonCreateBlockList(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CreateBlockList(context.Background(), &getstream.CreateBlockListRequest{})
	require.NoError(t, err)
}
func TestCommonDeleteBlockList(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.DeleteBlockList(context.Background(), "", &getstream.DeleteBlockListRequest{})
	require.NoError(t, err)
}
func TestCommonGetBlockList(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.GetBlockList(context.Background(), "", &getstream.GetBlockListRequest{})
	require.NoError(t, err)
}
func TestCommonUpdateBlockList(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.UpdateBlockList(context.Background(), "", &getstream.UpdateBlockListRequest{})
	require.NoError(t, err)
}
func TestCommonCheckPush(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CheckPush(context.Background(), &getstream.CheckPushRequest{})
	require.NoError(t, err)
}
func TestCommonCheckSNS(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CheckSNS(context.Background(), &getstream.CheckSNSRequest{})
	require.NoError(t, err)
}
func TestCommonCheckSQS(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CheckSQS(context.Background(), &getstream.CheckSQSRequest{})
	require.NoError(t, err)
}
func TestCommonDeleteDevice(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.DeleteDevice(context.Background(), &getstream.DeleteDeviceRequest{})
	require.NoError(t, err)
}
func TestCommonListDevices(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ListDevices(context.Background(), &getstream.ListDevicesRequest{})
	require.NoError(t, err)
}
func TestCommonCreateDevice(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CreateDevice(context.Background(), &getstream.CreateDeviceRequest{})
	require.NoError(t, err)
}
func TestCommonExportUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ExportUsers(context.Background(), &getstream.ExportUsersRequest{})
	require.NoError(t, err)
}
func TestCommonListExternalStorage(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ListExternalStorage(context.Background(), &getstream.ListExternalStorageRequest{})
	require.NoError(t, err)
}
func TestCommonCreateExternalStorage(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CreateExternalStorage(context.Background(), &getstream.CreateExternalStorageRequest{})
	require.NoError(t, err)
}
func TestCommonDeleteExternalStorage(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.DeleteExternalStorage(context.Background(), "", &getstream.DeleteExternalStorageRequest{})
	require.NoError(t, err)
}
func TestCommonUpdateExternalStorage(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.UpdateExternalStorage(context.Background(), "", &getstream.UpdateExternalStorageRequest{})
	require.NoError(t, err)
}
func TestCommonCheckExternalStorage(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CheckExternalStorage(context.Background(), "", &getstream.CheckExternalStorageRequest{})
	require.NoError(t, err)
}
func TestCommonCreateGuest(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CreateGuest(context.Background(), &getstream.CreateGuestRequest{})
	require.NoError(t, err)
}
func TestCommonCreateImportURL(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CreateImportURL(context.Background(), &getstream.CreateImportURLRequest{})
	require.NoError(t, err)
}
func TestCommonListImports(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ListImports(context.Background(), &getstream.ListImportsRequest{})
	require.NoError(t, err)
}
func TestCommonCreateImport(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CreateImport(context.Background(), &getstream.CreateImportRequest{})
	require.NoError(t, err)
}
func TestCommonGetImport(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.GetImport(context.Background(), "", &getstream.GetImportRequest{})
	require.NoError(t, err)
}
func TestCommonGetOG(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.GetOG(context.Background(), &getstream.GetOGRequest{})
	require.NoError(t, err)
}
func TestCommonListPermissions(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ListPermissions(context.Background(), &getstream.ListPermissionsRequest{})
	require.NoError(t, err)
}
func TestCommonGetPermission(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.GetPermission(context.Background(), "", &getstream.GetPermissionRequest{})
	require.NoError(t, err)
}
func TestCommonListPushProviders(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ListPushProviders(context.Background(), &getstream.ListPushProvidersRequest{})
	require.NoError(t, err)
}
func TestCommonUpsertPushProvider(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.UpsertPushProvider(context.Background(), &getstream.UpsertPushProviderRequest{})
	require.NoError(t, err)
}
func TestCommonDeletePushProvider(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.DeletePushProvider(context.Background(), "", "", &getstream.DeletePushProviderRequest{})
	require.NoError(t, err)
}
func TestCommonGetRateLimits(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.GetRateLimits(context.Background(), &getstream.GetRateLimitsRequest{})
	require.NoError(t, err)
}
func TestCommonListRoles(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ListRoles(context.Background(), &getstream.ListRolesRequest{})
	require.NoError(t, err)
}
func TestCommonCreateRole(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.CreateRole(context.Background(), &getstream.CreateRoleRequest{})
	require.NoError(t, err)
}
func TestCommonDeleteRole(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.DeleteRole(context.Background(), "", &getstream.DeleteRoleRequest{})
	require.NoError(t, err)
}
func TestCommonGetTask(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.GetTask(context.Background(), "", &getstream.GetTaskRequest{})
	require.NoError(t, err)
}
func TestCommonQueryUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.QueryUsers(context.Background(), &getstream.QueryUsersRequest{})
	require.NoError(t, err)
}
func TestCommonUpdateUsersPartial(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.UpdateUsersPartial(context.Background(), &getstream.UpdateUsersPartialRequest{})
	require.NoError(t, err)
}
func TestCommonUpdateUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.UpdateUsers(context.Background(), &getstream.UpdateUsersRequest{})
	require.NoError(t, err)
}
func TestCommonGetBlockedUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.GetBlockedUsers(context.Background(), &getstream.GetBlockedUsersRequest{})
	require.NoError(t, err)
}
func TestCommonBlockUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.BlockUsers(context.Background(), &getstream.BlockUsersRequest{})
	require.NoError(t, err)
}
func TestCommonDeactivateUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.DeactivateUsers(context.Background(), &getstream.DeactivateUsersRequest{})
	require.NoError(t, err)
}
func TestCommonDeleteUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.DeleteUsers(context.Background(), &getstream.DeleteUsersRequest{})
	require.NoError(t, err)
}
func TestCommonReactivateUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ReactivateUsers(context.Background(), &getstream.ReactivateUsersRequest{})
	require.NoError(t, err)
}
func TestCommonRestoreUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.RestoreUsers(context.Background(), &getstream.RestoreUsersRequest{})
	require.NoError(t, err)
}
func TestCommonUnblockUsers(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.UnblockUsers(context.Background(), &getstream.UnblockUsersRequest{})
	require.NoError(t, err)
}
func TestCommonDeactivateUser(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.DeactivateUser(context.Background(), "", &getstream.DeactivateUserRequest{})
	require.NoError(t, err)
}
func TestCommonExportUser(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ExportUser(context.Background(), "", &getstream.ExportUserRequest{})
	require.NoError(t, err)
}
func TestCommonReactivateUser(t *testing.T) {
	client, err := getstream.NewClient("key", "secret", getstream.WithHTTPClient(&StubHTTPClient{}))
	require.NoError(t, err)

	_, err = client.ReactivateUser(context.Background(), "", &getstream.ReactivateUserRequest{})
	require.NoError(t, err)
}