package getstream_test

import (
	"context"
	"strings"
	"testing"

	. "github.com/GetStream/getstream-go/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserGroupIntegration(t *testing.T) {
	t.Parallel()
	skipIfShort(t)
	client := initClient(t)
	ctx := context.Background()

	// deleteGroup is a best-effort cleanup helper.
	deleteGroup := func(id string) {
		_, _ = client.DeleteUserGroup(context.Background(), id, &DeleteUserGroupRequest{})
	}

	t.Run("CreateAndGetUserGroup", func(t *testing.T) {
		groupID := "test-group-" + uuid.New().String()
		t.Cleanup(func() { deleteGroup(groupID) })

		groupName := "Test Group " + groupID[:15]
		description := "A test user group"

		createResp, err := client.CreateUserGroup(ctx, &CreateUserGroupRequest{
			ID:          PtrTo(groupID),
			Name:        groupName,
			Description: PtrTo(description),
		})
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			t.Skip("User groups feature not available for this app")
		}
		require.NoError(t, err)
		require.NotNil(t, createResp.Data.UserGroup)
		assert.Equal(t, groupID, createResp.Data.UserGroup.ID)
		assert.Equal(t, groupName, createResp.Data.UserGroup.Name)
		require.NotNil(t, createResp.Data.UserGroup.Description)
		assert.Equal(t, description, *createResp.Data.UserGroup.Description)

		// Get the group by ID and verify fields
		getResp, err := client.GetUserGroup(ctx, groupID, &GetUserGroupRequest{})
		require.NoError(t, err)
		require.NotNil(t, getResp.Data.UserGroup)
		assert.Equal(t, groupID, getResp.Data.UserGroup.ID)
		assert.Equal(t, groupName, getResp.Data.UserGroup.Name)
	})

	t.Run("CreateUserGroupWithInitialMembers", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 2)
		groupID := "test-group-" + uuid.New().String()
		t.Cleanup(func() { deleteGroup(groupID) })

		createResp, err := client.CreateUserGroup(ctx, &CreateUserGroupRequest{
			ID:        PtrTo(groupID),
			Name:      "Group With Members",
			MemberIds: userIDs,
		})
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			t.Skip("User groups feature not available for this app")
		}
		require.NoError(t, err)
		require.NotNil(t, createResp.Data.UserGroup)
		assert.Equal(t, groupID, createResp.Data.UserGroup.ID)

		// Get and verify members are present
		getResp, err := client.GetUserGroup(ctx, groupID, &GetUserGroupRequest{})
		require.NoError(t, err)
		require.NotNil(t, getResp.Data.UserGroup)

		foundIDs := make(map[string]bool)
		for _, m := range getResp.Data.UserGroup.Members {
			foundIDs[m.UserID] = true
		}
		for _, uid := range userIDs {
			assert.True(t, foundIDs[uid], "User %s should be a member of the group", uid)
		}
	})

	t.Run("UpdateUserGroup", func(t *testing.T) {
		groupID := "test-group-" + uuid.New().String()
		t.Cleanup(func() { deleteGroup(groupID) })

		_, err := client.CreateUserGroup(ctx, &CreateUserGroupRequest{
			ID:   PtrTo(groupID),
			Name: "Original Name",
		})
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			t.Skip("User groups feature not available for this app")
		}
		require.NoError(t, err)

		newName := "Updated Name"
		newDesc := "Updated description"
		updateResp, err := client.UpdateUserGroup(ctx, groupID, &UpdateUserGroupRequest{
			Name:        PtrTo(newName),
			Description: PtrTo(newDesc),
		})
		require.NoError(t, err)
		require.NotNil(t, updateResp.Data.UserGroup)
		assert.Equal(t, newName, updateResp.Data.UserGroup.Name)
		require.NotNil(t, updateResp.Data.UserGroup.Description)
		assert.Equal(t, newDesc, *updateResp.Data.UserGroup.Description)

		// Confirm the update via a Get
		getResp, err := client.GetUserGroup(ctx, groupID, &GetUserGroupRequest{})
		require.NoError(t, err)
		require.NotNil(t, getResp.Data.UserGroup)
		assert.Equal(t, newName, getResp.Data.UserGroup.Name)
	})

	t.Run("ListUserGroups", func(t *testing.T) {
		groupID1 := "test-group-" + uuid.New().String()
		groupID2 := "test-group-" + uuid.New().String()
		t.Cleanup(func() {
			deleteGroup(groupID1)
			deleteGroup(groupID2)
		})

		_, err := client.CreateUserGroup(ctx, &CreateUserGroupRequest{
			ID:   PtrTo(groupID1),
			Name: "List Test Group One",
		})
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			t.Skip("User groups feature not available for this app")
		}
		require.NoError(t, err)

		_, err = client.CreateUserGroup(ctx, &CreateUserGroupRequest{
			ID:   PtrTo(groupID2),
			Name: "List Test Group Two",
		})
		require.NoError(t, err)

		listResp, err := client.ListUserGroups(ctx, &ListUserGroupsRequest{})
		require.NoError(t, err)
		assert.NotEmpty(t, listResp.Data.UserGroups)

		foundGroups := make(map[string]bool)
		for _, g := range listResp.Data.UserGroups {
			foundGroups[g.ID] = true
		}
		assert.True(t, foundGroups[groupID1] || foundGroups[groupID2],
			"At least one of the created groups should appear in the list")
	})

	t.Run("ListUserGroupsWithLimit", func(t *testing.T) {
		groupID1 := "test-group-" + uuid.New().String()
		groupID2 := "test-group-" + uuid.New().String()
		groupID3 := "test-group-" + uuid.New().String()
		t.Cleanup(func() {
			deleteGroup(groupID1)
			deleteGroup(groupID2)
			deleteGroup(groupID3)
		})

		for _, id := range []string{groupID1, groupID2, groupID3} {
			_, err := client.CreateUserGroup(ctx, &CreateUserGroupRequest{
				ID:   PtrTo(id),
				Name: "Limit Test Group " + id[:15],
			})
			if err != nil && strings.Contains(err.Error(), "Not Found") {
				t.Skip("User groups feature not available for this app")
			}
			require.NoError(t, err)
		}

		limit := 2
		listResp, err := client.ListUserGroups(ctx, &ListUserGroupsRequest{
			Limit: PtrTo(limit),
		})
		require.NoError(t, err)
		assert.LessOrEqual(t, len(listResp.Data.UserGroups), limit,
			"Response should respect the limit parameter")
	})

	t.Run("SearchUserGroups", func(t *testing.T) {
		uniquePrefix := "SearchTest-" + uuid.New().String()[:8]
		groupID := "test-group-" + uuid.New().String()
		t.Cleanup(func() { deleteGroup(groupID) })

		_, err := client.CreateUserGroup(ctx, &CreateUserGroupRequest{
			ID:   PtrTo(groupID),
			Name: uniquePrefix + " Group",
		})
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			t.Skip("User groups feature not available for this app")
		}
		require.NoError(t, err)

		searchResp, err := client.SearchUserGroups(ctx, &SearchUserGroupsRequest{
			Query: uniquePrefix,
		})
		require.NoError(t, err)

		found := false
		for _, g := range searchResp.Data.UserGroups {
			if strings.HasPrefix(g.Name, uniquePrefix) {
				found = true
				break
			}
		}
		assert.True(t, found, "Created group should appear in search results for prefix %q", uniquePrefix)
	})

	t.Run("AddUserGroupMembers", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 3)
		groupID := "test-group-" + uuid.New().String()
		t.Cleanup(func() { deleteGroup(groupID) })

		// Create group with first member only
		_, err := client.CreateUserGroup(ctx, &CreateUserGroupRequest{
			ID:        PtrTo(groupID),
			Name:      "Member Management Group",
			MemberIds: userIDs[:1],
		})
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			t.Skip("User groups feature not available for this app")
		}
		require.NoError(t, err)

		// Add remaining members
		addResp, err := client.AddUserGroupMembers(ctx, groupID, &AddUserGroupMembersRequest{
			MemberIds: userIDs[1:],
		})
		require.NoError(t, err)
		require.NotNil(t, addResp.Data.UserGroup)

		// Verify all members are present
		getResp, err := client.GetUserGroup(ctx, groupID, &GetUserGroupRequest{})
		require.NoError(t, err)
		require.NotNil(t, getResp.Data.UserGroup)

		foundIDs := make(map[string]bool)
		for _, m := range getResp.Data.UserGroup.Members {
			foundIDs[m.UserID] = true
		}
		for _, uid := range userIDs {
			assert.True(t, foundIDs[uid], "User %s should be a member after AddUserGroupMembers", uid)
		}
	})

	t.Run("RemoveUserGroupMembers", func(t *testing.T) {
		userIDs := createTestUsers(t, client, 2)
		groupID := "test-group-" + uuid.New().String()
		t.Cleanup(func() { deleteGroup(groupID) })

		// Create group with members
		_, err := client.CreateUserGroup(ctx, &CreateUserGroupRequest{
			ID:        PtrTo(groupID),
			Name:      "Remove Members Group",
			MemberIds: userIDs,
		})
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			t.Skip("User groups feature not available for this app")
		}
		require.NoError(t, err)

		// Verify members present before removal
		getResp, err := client.GetUserGroup(ctx, groupID, &GetUserGroupRequest{})
		require.NoError(t, err)
		require.NotNil(t, getResp.Data.UserGroup)
		assert.Len(t, getResp.Data.UserGroup.Members, len(userIDs))

		// Remove all members explicitly by ID (backend requires member_ids)
		_, err = client.RemoveUserGroupMembers(ctx, groupID, &RemoveUserGroupMembersRequest{
			MemberIds: userIDs,
		})
		require.NoError(t, err)

		// Verify members are removed
		getResp2, err := client.GetUserGroup(ctx, groupID, &GetUserGroupRequest{})
		require.NoError(t, err)
		require.NotNil(t, getResp2.Data.UserGroup)
		assert.Empty(t, getResp2.Data.UserGroup.Members, "All members should be removed")
	})

	t.Run("DeleteUserGroup", func(t *testing.T) {
		groupID := "test-group-" + uuid.New().String()

		_, err := client.CreateUserGroup(ctx, &CreateUserGroupRequest{
			ID:   PtrTo(groupID),
			Name: "Group To Delete",
		})
		if err != nil && strings.Contains(err.Error(), "Not Found") {
			t.Skip("User groups feature not available for this app")
		}
		require.NoError(t, err)

		_, err = client.DeleteUserGroup(ctx, groupID, &DeleteUserGroupRequest{})
		require.NoError(t, err)

		// Verify deletion — getting the deleted group should fail
		_, err = client.GetUserGroup(ctx, groupID, &GetUserGroupRequest{})
		require.Error(t, err, "Getting a deleted group should return an error")
	})
}
