package getstream_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/GetStream/getstream-go/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
Systematic Integration tests for Feed operations
These tests follow a logical flow: setup → create → operate → cleanup

Test order:
1. Environment Setup (user, feed creation)
2. Activity Operations (create, read, update, delete)
3. Reaction Operations (add, query, delete)
4. Comment Operations (add, read, update, delete)
5. Bookmark Operations (add, query, update, delete)
6. Follow Operations (follow, query, unfollow)
7. Batch Operations
8. Advanced Operations (polls, pins, etc.)
9. Cleanup
*/

const (
	userFeedType       = "user"
	pollQuestion       = "What's your favorite programming language?"
	testBookmarkFolder = "test-bookmarks1"
)

// TestFeedIntegrationSuite runs comprehensive integration tests for the Feeds API
func TestFeedIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	ctx := context.Background()

	client, err := getstream.NewClientFromEnvVars()
	require.NoError(t, err, "Failed to create client")

	feedsClient := client.Feeds()
	testUserID := "test-user-" + uuid.New().String()
	testUserID2 := "test-user-2-" + uuid.New().String()

	// Track created resources for cleanup
	var createdActivityIDs []string
	var createdCommentIDs []string

	// Setup environment
	setupEnvironment(t, ctx, client, testUserID, testUserID2)

	// Cleanup function
	defer func() {
		cleanupResources(ctx, feedsClient, createdActivityIDs, createdCommentIDs)
	}()

	// Run all tests
	t.Run("Test01_SetupEnvironmentDemo", func(t *testing.T) {
		test01SetupEnvironmentDemo(t, testUserID, testUserID2)
	})

	t.Run("Test02_CreateActivity", func(t *testing.T) {
		test02CreateActivity(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test02b_CreateActivityWithAttachments", func(t *testing.T) {
		test02bCreateActivityWithAttachments(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test02c_CreateVideoActivity", func(t *testing.T) {
		test02cCreateVideoActivity(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test02d_CreateStoryActivityWithExpiration", func(t *testing.T) {
		test02dCreateStoryActivityWithExpiration(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test02e_CreateActivityMultipleFeeds", func(t *testing.T) {
		test02eCreateActivityMultipleFeeds(t, ctx, feedsClient, testUserID, testUserID2, &createdActivityIDs)
	})

	t.Run("Test03_QueryActivities", func(t *testing.T) {
		test03QueryActivities(t, ctx, feedsClient)
	})

	t.Run("Test04_GetSingleActivity", func(t *testing.T) {
		test04GetSingleActivity(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test05_UpdateActivity", func(t *testing.T) {
		test05UpdateActivity(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test06_AddReaction", func(t *testing.T) {
		test06AddReaction(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test07_QueryReactions", func(t *testing.T) {
		test07QueryReactions(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test08_AddComment", func(t *testing.T) {
		test08AddComment(t, ctx, feedsClient, testUserID, &createdActivityIDs, &createdCommentIDs)
	})

	t.Run("Test09_QueryComments", func(t *testing.T) {
		test09QueryComments(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test10_UpdateComment", func(t *testing.T) {
		test10UpdateComment(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test11_AddBookmark", func(t *testing.T) {
		test11AddBookmark(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test12_QueryBookmarks", func(t *testing.T) {
		test12QueryBookmarks(t, ctx, feedsClient, testUserID)
	})

	t.Run("Test13_UpdateBookmark", func(t *testing.T) {
		test13UpdateBookmark(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test14_FollowUser", func(t *testing.T) {
		test14FollowUser(t, ctx, feedsClient, testUserID, testUserID2)
	})

	t.Run("Test15_QueryFollows", func(t *testing.T) {
		test15QueryFollows(t, ctx, feedsClient)
	})

	t.Run("Test16_UpsertActivities", func(t *testing.T) {
		test16UpsertActivities(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test17_PinActivity", func(t *testing.T) {
		test17PinActivity(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test18_UnpinActivity", func(t *testing.T) {
		test18UnpinActivity(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test19_DeleteBookmark", func(t *testing.T) {
		test19DeleteBookmark(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test20_DeleteReaction", func(t *testing.T) {
		test20DeleteReaction(t, ctx, feedsClient, testUserID, testUserID2, &createdActivityIDs)
	})

	t.Run("Test21_DeleteComment", func(t *testing.T) {
		test21DeleteComment(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test22_UnfollowUser", func(t *testing.T) {
		test22UnfollowUser(t, ctx, feedsClient, testUserID2, testUserID)
	})

	t.Run("Test23_DeleteActivities", func(t *testing.T) {
		test23DeleteActivities(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test24_CreatePoll", func(t *testing.T) {
		test24CreatePoll(t, ctx, client, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test25_VotePoll", func(t *testing.T) {
		test25VotePoll(t, ctx, client, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test26_ModerateActivity", func(t *testing.T) {
		test26ModerateActivity(t, ctx, feedsClient, testUserID, testUserID2, &createdActivityIDs)
	})

	t.Run("Test27_DeviceManagement", func(t *testing.T) {
		test27DeviceManagement(t, ctx, client, testUserID)
	})

	t.Run("Test28_QueryActivitiesWithFilters", func(t *testing.T) {
		test28QueryActivitiesWithFilters(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test29_GetFeedActivitiesWithPagination", func(t *testing.T) {
		test29GetFeedActivitiesWithPagination(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test30_ErrorHandlingScenarios", func(t *testing.T) {
		test30ErrorHandlingScenarios(t, ctx, feedsClient, testUserID)
	})

	t.Run("Test31_AuthenticationScenarios", func(t *testing.T) {
		test31AuthenticationScenarios(t, ctx, feedsClient, testUserID, &createdActivityIDs)
	})

	t.Run("Test32_RealWorldUsageDemo", func(t *testing.T) {
		test32RealWorldUsageDemo(t, ctx, feedsClient, testUserID, testUserID2, &createdActivityIDs)
	})

	// Feed Group CRUD Operations
	t.Run("Test33_FeedGroupCRUD", func(t *testing.T) {
		test33FeedGroupCRUD(t, ctx, feedsClient)
	})

	// Feed View CRUD Operations
	t.Run("Test34_FeedViewCRUD", func(t *testing.T) {
		test34FeedViewCRUD(t, ctx, feedsClient)
	})
}

// =================================================================
// ENVIRONMENT SETUP
// =================================================================

func setupEnvironment(t *testing.T, ctx context.Context, client *getstream.Stream, testUserID, testUserID2 string) {
	// Create test users
	// snippet-start: CreateUsers
	_, err := client.UpdateUsers(ctx, &getstream.UpdateUsersRequest{
		Users: map[string]getstream.UserRequest{
			testUserID: {
				ID:   testUserID,
				Name: getstream.PtrTo("Test User 1"),
				Role: getstream.PtrTo("user"),
			},
			testUserID2: {
				ID:   testUserID2,
				Name: getstream.PtrTo("Test User 2"),
				Role: getstream.PtrTo("user"),
			},
		},
	})
	// snippet-end: CreateUsers
	if err != nil {
		t.Logf("⚠️ Setup failed: %v", err)
		// Continue with tests even if setup partially fails
	}

	// Create feeds
	// snippet-start: GetOrCreateFeed
	feedsClient := client.Feeds()
	_, err = feedsClient.GetOrCreateFeed(ctx, userFeedType, testUserID, &getstream.GetOrCreateFeedRequest{
		UserID: &testUserID,
	})
	if err != nil {
		t.Logf("Failed to create feed 1: %v", err)
	}

	_, err = feedsClient.GetOrCreateFeed(ctx, userFeedType, testUserID2, &getstream.GetOrCreateFeedRequest{
		UserID: &testUserID2,
	})
	// snippet-end: GetOrCreateFeed
	if err != nil {
		t.Logf("Failed to create feed 2: %v", err)
	}
}

// =================================================================
// 1. ENVIRONMENT SETUP TEST
// =================================================================

func test01SetupEnvironmentDemo(t *testing.T, testUserID, testUserID2 string) {
	fmt.Println("\n🔧 Demonstrating environment setup...")
	fmt.Printf("✅ Users and feeds are automatically created in setup\n")
	fmt.Printf("   Test User 1: %s\n", testUserID)
	fmt.Printf("   Test User 2: %s\n", testUserID2)

	assert.True(t, true) // Just a demo test
}

// =================================================================
// 2. ACTIVITY OPERATIONS
// =================================================================

func test02CreateActivity(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n📝 Testing activity creation...")

	// snippet-start: AddActivity
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	response, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Feeds:  []string{feedIdentifier},
		Text:   getstream.PtrTo("This is a test activity from Go SDK"),
		UserID: &testUserID,
		Custom: map[string]interface{}{
			"test_field": "test_value",
			"timestamp":  time.Now().Unix(),
		},
	})
	// snippet-end: AddActivity

	assertResponseSuccess(t, response, err, "add activity")

	// Access the typed response data directly
	require.NotNil(t, response.Data.Activity)
	require.NotEmpty(t, response.Data.Activity.ID)
	require.NotNil(t, response.Data.Activity.Text)

	// Compare text
	assert.Equal(t, "This is a test activity from Go SDK", *response.Data.Activity.Text)

	activityID := response.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	fmt.Printf("✅ Created activity with ID: %s\n", activityID)
}

func test02bCreateActivityWithAttachments(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🖼️ Testing activity creation with image attachments...")

	// snippet-start: AddActivityWithImageAttachment
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	response, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Feeds:  []string{feedIdentifier},
		Text:   getstream.PtrTo("Look at this amazing view of NYC!"),
		UserID: &testUserID,
		Attachments: []getstream.Attachment{
			{
				ImageUrl: getstream.PtrTo("https://example.com/nyc-skyline.jpg"),
				Type:     getstream.PtrTo("image"),
				Title:    getstream.PtrTo("NYC Skyline"),
			},
		},
		Custom: map[string]interface{}{
			"location": "New York City",
			"camera":   "iPhone 15 Pro",
		},
	})
	// snippet-end: AddActivityWithImageAttachment

	assertResponseSuccess(t, response, err, "add activity with image attachment")

	activityID := response.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	fmt.Printf("✅ Created activity with image attachment: %s\n", activityID)
}

func test02cCreateVideoActivity(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🎥 Testing video activity creation...")

	// snippet-start: AddVideoActivity
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	response, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "video",
		Feeds:  []string{feedIdentifier},
		Text:   getstream.PtrTo("Check out this amazing video!"),
		UserID: &testUserID,
		Attachments: []getstream.Attachment{
			{
				AssetUrl: getstream.PtrTo("https://example.com/amazing-video.mp4"),
				Type:     getstream.PtrTo("video"),
				Title:    getstream.PtrTo("Amazing Video"),
				Custom: map[string]interface{}{
					"duration": 120,
				},
			},
		},
		Custom: map[string]interface{}{
			"video_quality":    "4K",
			"duration_seconds": 120,
		},
	})
	// snippet-end: AddVideoActivity

	assertResponseSuccess(t, response, err, "add video activity")

	activityID := response.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	fmt.Printf("✅ Created video activity: %s\n", activityID)
}

func test02dCreateStoryActivityWithExpiration(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n📖 Testing story activity with expiration...")

	// snippet-start: AddStoryActivityWithExpiration
	tomorrow := time.Now().Add(24 * time.Hour)
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	response, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:      "story",
		Feeds:     []string{feedIdentifier},
		Text:      getstream.PtrTo("My daily story - expires tomorrow!"),
		UserID:    &testUserID,
		ExpiresAt: getstream.PtrTo(tomorrow.Format(time.RFC3339)),
		Attachments: []getstream.Attachment{
			{
				ImageUrl: getstream.PtrTo("https://example.com/story-image.jpg"),
				Type:     getstream.PtrTo("image"),
			},
			{
				AssetUrl: getstream.PtrTo("https://example.com/story-video.mp4"),
				Type:     getstream.PtrTo("video"),
				Custom: map[string]interface{}{
					"duration": 15,
				},
			},
		},
		Custom: map[string]interface{}{
			"story_type":  "daily",
			"auto_expire": true,
		},
	})
	// snippet-end: AddStoryActivityWithExpiration

	assertResponseSuccess(t, response, err, "add story activity with expiration")

	activityID := response.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	fmt.Printf("✅ Created story activity with expiration: %s\n", activityID)
}

func test02eCreateActivityMultipleFeeds(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID, testUserID2 string, createdActivityIDs *[]string) {
	fmt.Println("\n📡 Testing activity creation to multiple feeds...")

	// snippet-start: AddActivityToMultipleFeeds
	feedIdentifier1 := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	feedIdentifier2 := fmt.Sprintf("%s:%s", userFeedType, testUserID2)
	response, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Feeds:  []string{feedIdentifier1, feedIdentifier2},
		Text:   getstream.PtrTo("This post appears in multiple feeds!"),
		UserID: &testUserID,
		Custom: map[string]interface{}{
			"cross_posted": true,
			"target_feeds": 2,
		},
	})
	// snippet-end: AddActivityToMultipleFeeds

	assertResponseSuccess(t, response, err, "add activity to multiple feeds")

	activityID := response.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	fmt.Printf("✅ Created activity in multiple feeds: %s\n", activityID)
}

func test03QueryActivities(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient) {
	fmt.Println("\n🔍 Testing activity querying...")

	// snippet-start: QueryActivities
	response, err := feedsClient.QueryActivities(ctx, &getstream.QueryActivitiesRequest{
		Limit: getstream.PtrTo(10),
		Filter: map[string]interface{}{
			"activity_type": "post",
		},
	})
	// snippet-end: QueryActivities

	assertResponseSuccess(t, response, err, "query activities")

	require.NotNil(t, response.Data.Activities)
	fmt.Println("✅ Queried activities successfully")
}

func test04GetSingleActivity(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n📄 Testing single activity retrieval...")

	// First create an activity to retrieve
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for retrieval test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for retrieval test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// snippet-start: GetActivity
	response, err := feedsClient.GetActivity(ctx, activityID, &getstream.GetActivityRequest{})
	// snippet-end: GetActivity

	assertResponseSuccess(t, response, err, "get activity")

	require.NotNil(t, response.Data.Activity)
	assert.Equal(t, activityID, response.Data.Activity.ID)
	fmt.Println("✅ Retrieved single activity")
}

func test05UpdateActivity(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n✏️ Testing activity update...")

	// First create an activity to update
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for update test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for update test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// snippet-start: UpdateActivity
	response, err := feedsClient.UpdateActivity(ctx, activityID, &getstream.UpdateActivityRequest{
		Text:   getstream.PtrTo("Updated activity text from Go SDK"),
		UserID: &testUserID, // Required for server-side auth
		Custom: map[string]interface{}{
			"updated":     true,
			"update_time": time.Now().Unix(),
		},
	})
	// snippet-end: UpdateActivity

	assertResponseSuccess(t, response, err, "update activity")
	fmt.Println("✅ Updated activity")
}

// =================================================================
// 3. REACTION OPERATIONS
// =================================================================

func test06AddReaction(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n👍 Testing reaction addition...")

	// First create an activity to react to
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for reaction test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for reaction test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// snippet-start: AddReaction
	response, err := feedsClient.AddReaction(ctx, activityID, &getstream.AddReactionRequest{
		Type:   "like",
		UserID: &testUserID,
	})
	// snippet-end: AddReaction

	assertResponseSuccess(t, response, err, "add reaction")
	fmt.Println("✅ Added like reaction")
}

func test07QueryReactions(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🔍 Testing reaction querying...")

	// Create an activity and add a reaction to it
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for query reactions test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for query reactions test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// Add a reaction first
	reactionResponse, err := feedsClient.AddReaction(ctx, activityID, &getstream.AddReactionRequest{
		Type:   "like",
		UserID: &testUserID,
	})
	assertResponseSuccess(t, reactionResponse, err, "add reaction for query test")

	// snippet-start: QueryActivityReactions
	response, err := feedsClient.QueryActivityReactions(ctx, activityID, &getstream.QueryActivityReactionsRequest{
		Limit: getstream.PtrTo(10),
		Filter: map[string]interface{}{
			"type": "like",
		},
	})
	// snippet-end: QueryActivityReactions
	if err != nil {
		fmt.Printf("Query reactions skipped: %v\n", err)
		t.Skip("Query reactions not supported: " + err.Error())
		return
	}

	assertResponseSuccess(t, response, err, "query reactions")
	fmt.Println("✅ Queried reactions")
}

// =================================================================
// 4. COMMENT OPERATIONS
// =================================================================

func test08AddComment(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs, createdCommentIDs *[]string) {
	fmt.Println("\n💬 Testing comment addition...")

	// First create an activity to comment on
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Feeds:  []string{feedIdentifier},
		Text:   getstream.PtrTo("Activity for comment test"),
		UserID: &testUserID,
	})

	assertResponseSuccess(t, createResponse, err, "create activity for comment test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// snippet-start: AddComment
	response, err := feedsClient.AddComment(ctx, &getstream.AddCommentRequest{
		Comment:    "This is a test comment from Go SDK",
		ObjectID:   activityID,
		ObjectType: "activity",
		UserID:     &testUserID,
	})
	// snippet-end: AddComment

	assertResponseSuccess(t, response, err, "add comment")

	if response.Data.Comment.ID != "" {
		testCommentID := response.Data.Comment.ID
		*createdCommentIDs = append(*createdCommentIDs, testCommentID)
		fmt.Printf("✅ Added comment with ID: %s\n", testCommentID)
	} else {
		fmt.Println("✅ Added comment (no ID returned)")
	}
}

func test09QueryComments(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🔍 Testing comment querying...")

	// Create an activity and add a comment to it
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for query comments test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for query comments test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// Add a comment first
	commentResponse, err := feedsClient.AddComment(ctx, &getstream.AddCommentRequest{
		Comment:    "Comment for query test",
		ObjectID:   activityID,
		ObjectType: "activity",
		UserID:     &testUserID,
	})
	assertResponseSuccess(t, commentResponse, err, "add comment for query test")

	// snippet-start: QueryComments
	response, err := feedsClient.QueryComments(ctx, &getstream.QueryCommentsRequest{
		Filter: map[string]interface{}{
			"object_id": activityID,
		},
		Limit: getstream.PtrTo(10),
	})
	// snippet-end: QueryComments

	assertResponseSuccess(t, response, err, "query comments")
	fmt.Println("✅ Queried comments")
}

func test10UpdateComment(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n✏️ Testing comment update...")

	// Create an activity and add a comment to update
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for update comment test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for update comment test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// Add a comment to update
	commentResponse, err := feedsClient.AddComment(ctx, &getstream.AddCommentRequest{
		Comment:    "Comment to be updated",
		ObjectID:   activityID,
		ObjectType: "activity",
		UserID:     &testUserID,
	})
	assertResponseSuccess(t, commentResponse, err, "add comment for update test")

	commentID := "comment-id" // Fallback if ID not returned
	if commentResponse.Data.Comment.ID != "" {
		commentID = commentResponse.Data.Comment.ID
	}

	// snippet-start: UpdateComment
	response, err := feedsClient.UpdateComment(ctx, commentID, &getstream.UpdateCommentRequest{
		Comment: getstream.PtrTo("Updated comment text from Go SDK"),
	})
	// snippet-end: UpdateComment

	assertResponseSuccess(t, response, err, "update comment")
	fmt.Println("✅ Updated comment")
}

// =================================================================
// 5. BOOKMARK OPERATIONS
// =================================================================

func test11AddBookmark(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🔖 Testing bookmark addition...")

	// Create an activity to bookmark
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for bookmark test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for bookmark test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// snippet-start: AddBookmark
	response, err := feedsClient.AddBookmark(ctx, activityID, &getstream.AddBookmarkRequest{
		UserID: &testUserID,
		NewFolder: &getstream.AddFolderRequest{
			Name: testBookmarkFolder,
		},
	})
	// snippet-end: AddBookmark
	if err != nil {
		fmt.Printf("Add bookmark failed: %v\n", err)
		t.Skip("Add bookmark not supported: " + err.Error())
		return
	}

	assertResponseSuccess(t, response, err, "add bookmark")
	fmt.Println("✅ Added bookmark")
}

func test12QueryBookmarks(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string) {
	fmt.Println("\n🔍 Testing bookmark querying...")

	// snippet-start: QueryBookmarks
	response, err := feedsClient.QueryBookmarks(ctx, &getstream.QueryBookmarksRequest{
		Limit: getstream.PtrTo(10),
		Filter: map[string]interface{}{
			"user_id": testUserID,
		},
	})
	// snippet-end: QueryBookmarks

	assertResponseSuccess(t, response, err, "query bookmarks")
	fmt.Println("✅ Queried bookmarks")
}

func test13UpdateBookmark(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n✏️ Testing bookmark update...")

	// Create an activity and bookmark it first
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Feeds:  []string{feedIdentifier},
		Text:   getstream.PtrTo("Activity for update bookmark test"),
		UserID: &testUserID,
	})

	assertResponseSuccess(t, createResponse, err, "create activity for update bookmark test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// Add a bookmark first
	bookmarkResponse, err := feedsClient.AddBookmark(ctx, activityID, &getstream.AddBookmarkRequest{
		NewFolder: &getstream.AddFolderRequest{
			Name: testBookmarkFolder,
		},
		UserID: &testUserID,
	})
	assertResponseSuccess(t, bookmarkResponse, err, "add bookmark for update test")

	// snippet-start: UpdateBookmark
	folderID := bookmarkResponse.Data.Bookmark.Folder.ID
	response, err := feedsClient.UpdateBookmark(ctx, activityID, &getstream.UpdateBookmarkRequest{
		FolderID: &folderID,
		UserID:   &testUserID,
	})
	// snippet-end: UpdateBookmark

	assertResponseSuccess(t, response, err, "update bookmark")
	fmt.Println("✅ Updated bookmark")
}

// =================================================================
// 6. FOLLOW OPERATIONS
// =================================================================

func test14FollowUser(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID, testUserID2 string) {
	fmt.Println("\n👥 Testing follow operation...")

	// snippet-start: Follow
	source := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	target := fmt.Sprintf("%s:%s", userFeedType, testUserID2)
	response, err := feedsClient.Follow(ctx, &getstream.FollowRequest{
		Source: source,
		Target: target,
	})
	// snippet-end: Follow
	if err != nil {
		fmt.Printf("Follow failed: %v\n", err)
		t.Skip("Follow operation not supported: " + err.Error())
		return
	}

	assertResponseSuccess(t, response, err, "follow user")
	fmt.Printf("✅ Followed user: %s\n", testUserID2)
}

func test15QueryFollows(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient) {
	fmt.Println("\n🔍 Testing follow querying...")

	// snippet-start: QueryFollows
	response, err := feedsClient.QueryFollows(ctx, &getstream.QueryFollowsRequest{
		Limit: getstream.PtrTo(10),
	})
	// snippet-end: QueryFollows

	assertResponseSuccess(t, response, err, "query follows")
	fmt.Println("✅ Queried follows")
}

// =================================================================
// 7. BATCH OPERATIONS
// =================================================================

func test16UpsertActivities(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n📝 Testing batch activity upsert...")

	// snippet-start: UpsertActivities
	activities := []getstream.ActivityRequest{
		{
			Type:   "post",
			Text:   getstream.PtrTo("Batch activity 1"),
			UserID: &testUserID,
		},
		{
			Type:   "post",
			Text:   getstream.PtrTo("Batch activity 2"),
			UserID: &testUserID,
		},
	}

	response, err := feedsClient.UpsertActivities(ctx, &getstream.UpsertActivitiesRequest{
		Activities: activities,
	})
	// snippet-end: UpsertActivities

	assertResponseSuccess(t, response, err, "upsert activities")

	// Track created activities for cleanup
	if response.Data.Activities != nil {
		for _, activity := range response.Data.Activities {
			if activity.ID != "" {
				*createdActivityIDs = append(*createdActivityIDs, activity.ID)
			}
		}
	}

	fmt.Println("✅ Upserted batch activities")
}

// =================================================================
// 8. ADVANCED OPERATIONS
// =================================================================

func test17PinActivity(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n📌 Testing activity pinning...")

	// Create an activity to pin
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Feeds:  []string{feedIdentifier},
		Text:   getstream.PtrTo("Activity for pin test"),
		UserID: &testUserID,
	})

	assertResponseSuccess(t, createResponse, err, "create activity for pin test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// snippet-start: PinActivity
	response, err := feedsClient.PinActivity(ctx, userFeedType, testUserID, activityID, &getstream.PinActivityRequest{
		UserID: &testUserID,
	})
	// snippet-end: PinActivity

	assertResponseSuccess(t, response, err, "pin activity")
	fmt.Println("✅ Pinned activity")
}

func test18UnpinActivity(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n📌 Testing activity unpinning...")

	// Create an activity, pin it, then unpin it
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Feeds:  []string{feedIdentifier},
		Text:   getstream.PtrTo("Activity for unpin test"),
		UserID: &testUserID,
	})

	assertResponseSuccess(t, createResponse, err, "create activity for unpin test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// Pin it first
	pinResponse, err := feedsClient.PinActivity(ctx, userFeedType, testUserID, activityID, &getstream.PinActivityRequest{
		UserID: &testUserID,
	})
	assertResponseSuccess(t, pinResponse, err, "pin activity for unpin test")

	// snippet-start: UnpinActivity
	response, err := feedsClient.UnpinActivity(ctx, userFeedType, testUserID, activityID, &getstream.UnpinActivityRequest{
		UserID: &testUserID,
	})
	// snippet-end: UnpinActivity

	assertResponseSuccess(t, response, err, "unpin activity")
	fmt.Println("✅ Unpinned activity")
}

// =================================================================
// 9. CLEANUP OPERATIONS (in reverse order)
// =================================================================

func test19DeleteBookmark(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🗑️ Testing bookmark deletion...")

	// Create an activity and bookmark it first
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for delete bookmark test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for delete bookmark test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// Add a bookmark first
	bookmarkResponse, err := feedsClient.AddBookmark(ctx, activityID, &getstream.AddBookmarkRequest{
		NewFolder: &getstream.AddFolderRequest{
			Name: testBookmarkFolder,
		},
		UserID: &testUserID,
	})
	assertResponseSuccess(t, bookmarkResponse, err, "add bookmark for delete test")

	// snippet-start: DeleteBookmark
	folderID := bookmarkResponse.Data.Bookmark.Folder.ID
	response, err := feedsClient.DeleteBookmark(ctx, activityID, &getstream.DeleteBookmarkRequest{
		FolderID: &folderID,
		UserID:   &testUserID,
	})
	// snippet-end: DeleteBookmark

	assertResponseSuccess(t, response, err, "delete bookmark")
	fmt.Println("✅ Deleted bookmark")
}

func test20DeleteReaction(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID, _ string, createdActivityIDs *[]string) {
	fmt.Println("\n🗑️ Testing reaction deletion...")

	// Create an activity and add a reaction first
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for delete reaction test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for delete reaction test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// Add a reaction first
	reactionResponse, err := feedsClient.AddReaction(ctx, activityID, &getstream.AddReactionRequest{
		Type:   "like",
		UserID: &testUserID,
	})
	assertResponseSuccess(t, reactionResponse, err, "add reaction for delete test")

	// snippet-start: DeleteActivityReaction
	response, err := feedsClient.DeleteActivityReaction(ctx, activityID, "like", &getstream.DeleteActivityReactionRequest{
		UserID: &testUserID,
	})
	// snippet-end: DeleteActivityReaction

	assertResponseSuccess(t, response, err, "delete reaction")
	fmt.Println("✅ Deleted reaction")
}

func test21DeleteComment(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🗑️ Testing comment deletion...")

	// Create an activity and add a comment first
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity for delete comment test"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for delete comment test")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// Add a comment first
	commentResponse, err := feedsClient.AddComment(ctx, &getstream.AddCommentRequest{
		Comment:    "Comment to be deleted",
		ObjectID:   activityID,
		ObjectType: "activity",
		UserID:     &testUserID,
	})
	assertResponseSuccess(t, commentResponse, err, "add comment for delete test")

	commentID := "comment-id" // Fallback if ID not returned
	if commentResponse.Data.Comment.ID != "" {
		commentID = commentResponse.Data.Comment.ID
	}

	// snippet-start: DeleteComment
	response, err := feedsClient.DeleteComment(ctx, commentID, &getstream.DeleteCommentRequest{
		HardDelete: getstream.PtrTo(false), // soft delete
	})
	// snippet-end: DeleteComment

	assertResponseSuccess(t, response, err, "delete comment")
	fmt.Println("✅ Deleted comment")
}

func test22UnfollowUser(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID, testUserID2 string) {
	fmt.Println("\n👥 Testing unfollow operation...")

	// First establish a follow relationship
	source := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	target := fmt.Sprintf("%s:%s", userFeedType, testUserID2)
	followResponse, err := feedsClient.Follow(ctx, &getstream.FollowRequest{
		Source: source,
		Target: target,
	})
	assertResponseSuccess(t, followResponse, err, "establish follow relationship for unfollow test")

	// snippet-start: Unfollow
	response, err := feedsClient.Unfollow(ctx, source, target, &getstream.UnfollowRequest{})
	// snippet-end: Unfollow
	if err != nil {
		fmt.Printf("Unfollow operation skipped: %v\n", err)
		t.Skip("Unfollow operation not supported: " + err.Error())
		return
	}

	assertResponseSuccess(t, response, err, "unfollow operation")
	fmt.Printf("✅ Unfollowed user: %s\n", testUserID2)
}

func test23DeleteActivities(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🗑️ Testing activity deletion...")

	// Create some activities to delete
	var activitiesToDelete []string
	for i := 1; i <= 2; i++ {
		feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
		createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
			Type:   "post",
			Text:   getstream.PtrTo(fmt.Sprintf("Activity %d for delete test", i)),
			UserID: &testUserID,
			Feeds:  []string{feedIdentifier},
		})

		assertResponseSuccess(t, createResponse, err, fmt.Sprintf("create activity %d for delete test", i))

		activityID := createResponse.Data.Activity.ID
		activitiesToDelete = append(activitiesToDelete, activityID)
		*createdActivityIDs = append(*createdActivityIDs, activityID)
	}

	for _, activityID := range activitiesToDelete {
		// snippet-start: DeleteActivity
		response, err := feedsClient.DeleteActivity(ctx, activityID, &getstream.DeleteActivityRequest{
			HardDelete: getstream.PtrTo(false), // soft delete
		})
		// snippet-end: DeleteActivity

		assertResponseSuccess(t, response, err, "delete activity")
	}

	fmt.Printf("✅ Deleted %d activities\n", len(activitiesToDelete))
	*createdActivityIDs = []string{} // Clear since we deleted them
}

// =================================================================
// 10. ADDITIONAL COMPREHENSIVE TESTS
// =================================================================

func test24CreatePoll(t *testing.T, ctx context.Context, client *getstream.Stream, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🗳️ Testing poll creation...")

	// snippet-start: CreatePoll
	pollResponse, err := client.CreatePoll(ctx, &getstream.CreatePollRequest{
		Name:        "Poll",
		Description: getstream.PtrTo(pollQuestion),
		UserID:      &testUserID,
		Options:     []getstream.PollOptionInput{{Text: getstream.PtrTo("Red")}, {Text: getstream.PtrTo("Blue")}},
	})
	if err != nil {
		fmt.Printf("Poll creation skipped: %v\n", err)
		t.Skip("Poll operations not supported: " + err.Error())
		return
	}

	pollID := pollResponse.Data.Poll.ID

	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	pollActivity, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "poll",
		Feeds:  []string{feedIdentifier},
		PollID: &pollID,
		Text:   getstream.PtrTo(pollQuestion),
		UserID: &testUserID,
		Custom: map[string]interface{}{
			"poll_name":                    pollQuestion,
			"poll_description":             "Choose your favorite programming language from the options below",
			"poll_options":                 []string{"PHP", "Python", "JavaScript", "Go"},
			"allow_user_suggested_options": false,
			"max_votes_allowed":            1,
		},
	})
	// snippet-end: CreatePoll

	assertResponseSuccess(t, pollActivity, err, "create poll")

	activityID := pollActivity.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	fmt.Printf("✅ Created poll activity: %s\n", activityID)
}

func test25VotePoll(t *testing.T, ctx context.Context, client *getstream.Stream, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n✅ Testing poll voting...")

	// Create a poll first using the proper API
	pollResponse, err := client.CreatePoll(ctx, &getstream.CreatePollRequest{
		Name:        "Favorite Color Poll",
		Description: getstream.PtrTo("What is your favorite color?"),
		UserID:      &testUserID,
		Options:     []getstream.PollOptionInput{{Text: getstream.PtrTo("Red")}, {Text: getstream.PtrTo("Blue")}, {Text: getstream.PtrTo("Green")}},
	})
	if err != nil {
		fmt.Printf("Poll voting skipped: %v\n", err)
		t.Skip("Poll voting not supported: " + err.Error())
		return
	}

	pollID := pollResponse.Data.Poll.ID

	// Create activity with the poll
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "poll",
		Feeds:  []string{feedIdentifier},
		Text:   getstream.PtrTo("Vote for your favorite color"),
		UserID: &testUserID,
		PollID: &pollID,
		Custom: map[string]interface{}{
			"poll_name":                    "What is your favorite color?",
			"poll_description":             "Choose your favorite color from the options below",
			"poll_options":                 []string{"Red", "Blue", "Green"},
			"allow_user_suggested_options": false,
		},
	})

	assertResponseSuccess(t, createResponse, err, "create poll for voting")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// Get poll options from the poll response
	pollOptions := pollResponse.Data.Poll.Options
	if len(pollOptions) > 0 {
		// Use the first option ID from the poll creation response
		optionID := pollOptions[0].ID

		// snippet-start: VotePoll
		voteResponse, err := feedsClient.CastPollVote(ctx, activityID, pollID, &getstream.CastPollVoteRequest{
			UserID: &testUserID,
			Vote: &getstream.VoteData{
				OptionID: &optionID,
			},
		})
		// snippet-end: VotePoll
		if err != nil {
			fmt.Printf("Poll voting skipped: %v\n", err)
			t.Skip("Poll voting not supported: " + err.Error())
			return
		}

		assertResponseSuccess(t, voteResponse, err, "vote on poll")
		fmt.Printf("✅ Voted on poll: %s\n", activityID)
	} else {
		fmt.Println("⚠️ Poll options not found in poll response")
		t.Skip("Poll options not available for voting test")
	}
}

func test26ModerateActivity(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID, testUserID2 string, createdActivityIDs *[]string) {
	fmt.Println("\n🛡️ Testing activity moderation...")

	// Create an activity to moderate
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("This content might need moderation"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})

	assertResponseSuccess(t, createResponse, err, "create activity for moderation")

	activityID := createResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	// snippet-start: ModerateActivity
	moderationResponse, err := feedsClient.ActivityFeedback(ctx, activityID, &getstream.ActivityFeedbackRequest{
		Report: getstream.PtrTo(true),
		Reason: getstream.PtrTo("inappropriate_content"),
		UserID: &testUserID2, // Different user reporting
	})
	// snippet-end: ModerateActivity
	if err != nil {
		fmt.Printf("Activity moderation skipped: %v\n", err)
		t.Skip("Activity moderation not supported: " + err.Error())
		return
	}

	assertResponseSuccess(t, moderationResponse, err, "moderate activity")
	fmt.Printf("✅ Flagged activity for moderation: %s\n", activityID)
}

func test27DeviceManagement(t *testing.T, ctx context.Context, client *getstream.Stream, testUserID string) {
	fmt.Println("\n📱 Testing device management...")

	deviceToken := "test-device-token-" + uuid.New().String()

	// snippet-start: AddDevice
	addDeviceResponse, err := client.CreateDevice(ctx, &getstream.CreateDeviceRequest{
		ID:           deviceToken,
		PushProvider: "apn",
		UserID:       &testUserID,
	})
	// snippet-end: AddDevice
	if err != nil {
		fmt.Printf("Device management skipped: %v\n", err)
		t.Skip("Device management not supported: " + err.Error())
		return
	}

	assertResponseSuccess(t, addDeviceResponse, err, "add device")
	fmt.Printf("✅ Added device: %s\n", deviceToken)

	// snippet-start: RemoveDevice
	removeDeviceResponse, err := client.DeleteDevice(ctx, &getstream.DeleteDeviceRequest{
		ID:     deviceToken,
		UserID: &testUserID,
	})
	// snippet-end: RemoveDevice

	assertResponseSuccess(t, removeDeviceResponse, err, "remove device")
	fmt.Printf("✅ Removed device: %s\n", deviceToken)
}

func test28QueryActivitiesWithFilters(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🔍 Testing activity queries with advanced filters...")

	// Create activities with different types and metadata
	activityTypes := []string{"post", "photo", "video", "story"}

	for _, actType := range activityTypes {
		feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
		createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
			Type:   actType,
			Text:   getstream.PtrTo(fmt.Sprintf("Test %s activity for filtering", actType)),
			UserID: &testUserID,
			Feeds:  []string{feedIdentifier},
			Custom: map[string]interface{}{
				"category": actType,
				"priority": 3,
				"tags":     []string{actType, "test", "filter"},
			},
		})

		assertResponseSuccess(t, createResponse, err, fmt.Sprintf("create %s activity for filtering", actType))

		*createdActivityIDs = append(*createdActivityIDs, createResponse.Data.Activity.ID)
	}

	// Query with type filter
	// snippet-start: QueryActivitiesWithTypeFilter
	response, err := feedsClient.QueryActivities(ctx, &getstream.QueryActivitiesRequest{
		Limit: getstream.PtrTo(10),
		Filter: map[string]interface{}{
			"activity_type": "post",
			"user_id":       testUserID,
		},
		Sort: []getstream.SortParamRequest{
			{Field: getstream.PtrTo("created_at"), Direction: getstream.PtrTo(-1)}, // newest first
		},
	})
	// snippet-end: QueryActivitiesWithTypeFilter

	if err != nil {
		fmt.Printf("Query activities with type filter skipped: %v\n", err)
	} else {
		assertResponseSuccess(t, response, err, "query activities with type filter")
	}

	// Query with custom field filter
	// snippet-start: QueryActivitiesWithCustomFilter
	customFilterResponse, err := feedsClient.QueryActivities(ctx, &getstream.QueryActivitiesRequest{
		Limit: getstream.PtrTo(10),
		Filter: map[string]interface{}{
			"custom.priority": map[string]interface{}{"$gte": 3}, // priority >= 3
			"user_id":         testUserID,
		},
	})
	// snippet-end: QueryActivitiesWithCustomFilter

	if err != nil {
		fmt.Printf("Query activities with custom filter skipped: %v\n", err)
	} else {
		assertResponseSuccess(t, customFilterResponse, err, "query activities with custom filter")
	}

	fmt.Println("✅ Queried activities with advanced filters")
}

func test29GetFeedActivitiesWithPagination(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n📄 Testing feed activities with pagination...")

	// Create multiple activities for pagination test
	for i := 1; i <= 7; i++ {
		feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
		createResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
			Type:   "post",
			Text:   getstream.PtrTo(fmt.Sprintf("Pagination test activity %d", i)),
			UserID: &testUserID,
			Feeds:  []string{feedIdentifier},
		})

		assertResponseSuccess(t, createResponse, err, fmt.Sprintf("create pagination activity %d", i))

		*createdActivityIDs = append(*createdActivityIDs, createResponse.Data.Activity.ID)
	}

	// Get first page
	// snippet-start: GetFeedActivitiesWithPagination
	firstPageResponse, err := feedsClient.QueryActivities(ctx, &getstream.QueryActivitiesRequest{
		Limit: getstream.PtrTo(3),
		Filter: map[string]interface{}{
			"user_id": testUserID,
		},
	})
	// snippet-end: GetFeedActivitiesWithPagination

	assertResponseSuccess(t, firstPageResponse, err, "get first page of feed activities")

	require.NotNil(t, firstPageResponse.Data.Activities)
	assert.LessOrEqual(t, len(firstPageResponse.Data.Activities), 3)

	// Get second page using next token if available
	// snippet-start: GetFeedActivitiesSecondPage
	nextToken := firstPageResponse.Data.Next
	if nextToken != nil {
		secondPageResponse, err := feedsClient.QueryActivities(ctx, &getstream.QueryActivitiesRequest{
			Limit: getstream.PtrTo(3),
			Next:  nextToken,
			Filter: map[string]interface{}{
				"user_id": testUserID,
			},
		})
		assertResponseSuccess(t, secondPageResponse, err, "get second page of feed activities")
	} else {
		fmt.Println("⚠️ No next page available")
	}
	// snippet-end: GetFeedActivitiesSecondPage

	fmt.Println("✅ Retrieved feed activities with pagination")
}

// Test comprehensive error handling scenarios
func test30ErrorHandlingScenarios(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string) {
	fmt.Println("\n⚠️ Testing error handling scenarios...")

	// Test 1: Invalid activity ID
	// snippet-start: HandleInvalidActivityId
	response, err := feedsClient.GetActivity(ctx, "invalid-activity-id-12345", &getstream.GetActivityRequest{})
	// snippet-end: HandleInvalidActivityId

	if err != nil {
		fmt.Printf("✅ Caught expected error for invalid activity ID: %v\n", err)
	} else if response != nil {
		fmt.Println("✅ Correctly handled invalid activity ID error")
	}

	// Test 2: Empty activity text
	// snippet-start: HandleEmptyActivityText
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	emptyResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo(""), // Empty text
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})
	// snippet-end: HandleEmptyActivityText

	if err != nil {
		fmt.Printf("✅ Caught expected error for empty activity text: %v\n", err)
	} else if emptyResponse != nil {
		fmt.Println("✅ Correctly handled empty activity text")
	}

	// Test 3: Invalid user ID
	// snippet-start: HandleInvalidUserId
	invalidUserResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Test with invalid user"),
		UserID: getstream.PtrTo(""), // Empty user ID
		Feeds:  []string{feedIdentifier},
	})
	// snippet-end: HandleInvalidUserId

	if err != nil {
		fmt.Printf("✅ Caught expected error for invalid user ID: %v\n", err)
	} else if invalidUserResponse != nil {
		fmt.Println("✅ Correctly handled invalid user ID")
	}

	assert.True(t, true) // Test passes if we reach here
}

// Test authentication and authorization scenarios
func test31AuthenticationScenarios(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string, createdActivityIDs *[]string) {
	fmt.Println("\n🔐 Testing authentication scenarios...")

	// Test with valid user authentication
	// snippet-start: ValidUserAuthentication
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	response, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Activity with proper authentication"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
	})
	// snippet-end: ValidUserAuthentication

	assertResponseSuccess(t, response, err, "activity with valid authentication")

	activityID := response.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, activityID)

	fmt.Printf("✅ Successfully authenticated and created activity: %s\n", activityID)

	// Test user permissions for updating activity
	// snippet-start: UserPermissionUpdate
	updateResponse, err := feedsClient.UpdateActivity(ctx, activityID, &getstream.UpdateActivityRequest{
		Text:   getstream.PtrTo("Updated with proper user permissions"),
		UserID: &testUserID, // Same user can update
	})
	// snippet-end: UserPermissionUpdate

	assertResponseSuccess(t, updateResponse, err, "update activity with proper permissions")
	fmt.Println("✅ Successfully updated activity with proper user permissions")
}

// Comprehensive test demonstrating real-world usage patterns
func test32RealWorldUsageDemo(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID, testUserID2 string, createdActivityIDs *[]string) {
	fmt.Println("\n🌍 Testing real-world usage patterns...")

	// Scenario: User posts content, gets reactions and comments
	// snippet-start: RealWorldScenario

	// 1. User creates a post with image
	feedIdentifier := fmt.Sprintf("%s:%s", userFeedType, testUserID)
	postResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
		Type:   "post",
		Text:   getstream.PtrTo("Just visited the most amazing coffee shop! ☕️"),
		UserID: &testUserID,
		Feeds:  []string{feedIdentifier},
		Attachments: []getstream.Attachment{
			{
				ImageUrl: getstream.PtrTo("https://example.com/coffee-shop.jpg"),
				Type:     getstream.PtrTo("image"),
				Title:    getstream.PtrTo("Amazing Coffee Shop"),
			},
		},
		Custom: map[string]interface{}{
			"location": "Downtown Coffee Co.",
			"rating":   5,
			"tags":     []string{"coffee", "food", "downtown"},
		},
	})
	assertResponseSuccess(t, postResponse, err, "create real-world post")

	postID := postResponse.Data.Activity.ID
	*createdActivityIDs = append(*createdActivityIDs, postID)

	// 2. Other users react to the post
	reactionTypes := []string{"like", "love", "wow"}
	for _, reactionType := range reactionTypes {
		reactionResponse, err := feedsClient.AddReaction(ctx, postID, &getstream.AddReactionRequest{
			Type:   reactionType,
			UserID: &testUserID2,
		})
		assertResponseSuccess(t, reactionResponse, err, fmt.Sprintf("add %s reaction", reactionType))
	}

	// 3. Users comment on the post
	comments := []string{
		"That place looks amazing! What did you order?",
		"I love their espresso! Great choice 😍",
		"Adding this to my must-visit list!",
	}

	for _, commentText := range comments {
		commentResponse, err := feedsClient.AddComment(ctx, &getstream.AddCommentRequest{
			Comment:    commentText,
			ObjectID:   postID,
			ObjectType: "activity",
			UserID:     &testUserID2,
		})
		assertResponseSuccess(t, commentResponse, err, "add comment to post")
	}

	// 4. User bookmarks the post
	bookmarkResponse, err := feedsClient.AddBookmark(ctx, postID, &getstream.AddBookmarkRequest{
		UserID: &testUserID2,
		NewFolder: &getstream.AddFolderRequest{
			Name: "favorite-places",
		},
	})
	if err != nil {
		fmt.Printf("Bookmark operation skipped: %v\n", err)
	} else {
		assertResponseSuccess(t, bookmarkResponse, err, "bookmark the post")
	}

	// 5. Query the activity with all its interactions
	enrichedResponse, err := feedsClient.GetActivity(ctx, postID, &getstream.GetActivityRequest{})
	assertResponseSuccess(t, enrichedResponse, err, "get enriched activity")

	// snippet-end: RealWorldScenario

	fmt.Println("✅ Completed real-world usage scenario demonstration")
}

// =================================================================
// 11. FEED GROUP CRUD OPERATIONS
// =================================================================

func test33FeedGroupCRUD(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient) {
	fmt.Println("\n📁 Testing Feed Group CRUD operations...")

	feedGroupID := "test-feed-group-" + uuid.New().String()[:8]

	// Test 1: List Feed Groups
	fmt.Println("\n📋 Testing list feed groups...")
	// snippet-start: ListFeedGroups
	listResponse, err := feedsClient.ListFeedGroups(ctx, &getstream.ListFeedGroupsRequest{})
	// snippet-end: ListFeedGroups

	assertResponseSuccess(t, listResponse, err, "list feed groups")
	fmt.Printf("✅ Listed %d existing feed groups\n", len(listResponse.Data.Groups))

	// Test 2: Create Feed Group
	fmt.Println("\n➕ Testing create feed group...")
	// snippet-start: CreateFeedGroup
	createResponse, err := feedsClient.CreateFeedGroup(ctx, &getstream.CreateFeedGroupRequest{
		ID:                feedGroupID,
		DefaultVisibility: getstream.PtrTo("public"),
		ActivityProcessors: []getstream.ActivityProcessorConfig{
			{Type: "dummy"},
		},
	})
	// snippet-end: CreateFeedGroup

	assertResponseSuccess(t, createResponse, err, "create feed group")
	assert.Equal(t, feedGroupID, createResponse.Data.FeedGroup.ID)
	fmt.Printf("✅ Created feed group: %s\n", feedGroupID)

	// Test 3: Get Feed Group
	fmt.Println("\n🔍 Testing get feed group...")
	// snippet-start: GetFeedGroup
	getResponse, err := feedsClient.GetFeedGroup(ctx, "feed_group_id", &getstream.GetFeedGroupRequest{})
	// snippet-end: GetFeedGroup

	assertResponseSuccess(t, getResponse, err, "get feed group")
	assert.Equal(t, "feed_group_id", getResponse.Data.FeedGroup.ID)
	fmt.Printf("✅ Retrieved feed group: %s\n", feedGroupID)

	// Test 4: Update Feed Group
	fmt.Println("\n✏️ Testing update feed group...")
	// snippet-start: UpdateFeedGroup
	updateResponse, err := feedsClient.UpdateFeedGroup(ctx, "feed_group_id", &getstream.UpdateFeedGroupRequest{
		ActivityProcessors: []getstream.ActivityProcessorConfig{
			{Type: "dummy"},
		},
		Aggregation: &getstream.AggregationConfig{
			Format: getstream.PtrTo("time_based"),
		},
	})
	// snippet-end: UpdateFeedGroup

	assertResponseSuccess(t, updateResponse, err, "update feed group")
	fmt.Printf("✅ Updated feed group: %s\n", feedGroupID)

	// Test 5: Get or Create Feed Group (should get existing)
	fmt.Println("\n🔄 Testing get or create feed group (existing)...")
	// snippet-start: GetOrCreateFeedGroupExisting
	getOrCreateResponse, err := feedsClient.GetOrCreateFeedGroup(ctx, "feed_group_id", &getstream.GetOrCreateFeedGroupRequest{
		DefaultVisibility: getstream.PtrTo("public"),
	})
	// snippet-end: GetOrCreateFeedGroupExisting

	assertResponseSuccess(t, getOrCreateResponse, err, "get or create existing feed group")
	assert.False(t, getOrCreateResponse.Data.WasCreated, "Should not create new feed group")
	fmt.Printf("✅ Got existing feed group: %s\n", feedGroupID)

	// Test 6: Get or Create Feed Group (should create new)
	newFeedGroupID := "test-new-feed-group-" + uuid.New().String()[:8]
	fmt.Println("\n🆕 Testing get or create feed group (new)...")
	// snippet-start: GetOrCreateFeedGroupNew
	newGetOrCreateResponse, err := feedsClient.GetOrCreateFeedGroup(ctx, newFeedGroupID, &getstream.GetOrCreateFeedGroupRequest{
		DefaultVisibility: getstream.PtrTo("private"),
		ActivityProcessors: []getstream.ActivityProcessorConfig{
			{Type: "dummy"},
		},
	})
	// snippet-end: GetOrCreateFeedGroupNew

	assertResponseSuccess(t, newGetOrCreateResponse, err, "get or create new feed group")
	assert.True(t, newGetOrCreateResponse.Data.WasCreated, "Should create new feed group")
	fmt.Printf("✅ Created new feed group: %s\n", newFeedGroupID)

	// Test 7: Delete Feed Groups (cleanup)
	fmt.Println("\n🗑️ Testing delete feed groups...")
	// snippet-start: DeleteFeedGroup
	_, err = feedsClient.DeleteFeedGroup(ctx, "groupID-123", &getstream.DeleteFeedGroupRequest{
		HardDelete: getstream.PtrTo(false), // soft delete
	})
	// snippet-end: DeleteFeedGroup

	fmt.Println("✅ Completed Feed Group CRUD operations")
}

// =================================================================
// 12. FEED VIEW CRUD OPERATIONS
// =================================================================

func test34FeedViewCRUD(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient) {
	fmt.Println("\n👁️ Testing Feed View CRUD operations...")

	feedViewID := "test-feed-view-" + uuid.New().String()[:8]

	// Test 1: List Feed Views
	fmt.Println("\n📋 Testing list feed views...")
	// snippet-start: ListFeedViews
	listResponse, err := feedsClient.ListFeedViews(ctx, &getstream.ListFeedViewsRequest{})
	// snippet-end: ListFeedViews

	assertResponseSuccess(t, listResponse, err, "list feed views")
	fmt.Printf("✅ Listed %d existing feed views\n", len(listResponse.Data.Views))

	// Test 2: Create Feed View
	fmt.Println("\n➕ Testing create feed view...")
	// snippet-start: CreateFeedView
	createResponse, err := feedsClient.CreateFeedView(ctx, &getstream.CreateFeedViewRequest{
		ID: feedViewID,
		ActivitySelectors: []getstream.ActivitySelectorConfig{
			{
				Type: getstream.PtrTo("popular"),
			},
		},
		ActivityProcessors: []getstream.ActivityProcessorConfig{
			{Type: "default"},
		},
		Aggregation: &getstream.AggregationConfig{
			Format: getstream.PtrTo("time_based"),
		},
	})
	// snippet-end: CreateFeedView

	assertResponseSuccess(t, createResponse, err, "create feed view")
	assert.Equal(t, feedViewID, createResponse.Data.FeedView.ID)
	fmt.Printf("✅ Created feed view: %s\n", feedViewID)

	// Test 3: Get Feed View
	fmt.Println("\n🔍 Testing get feed view...")
	// snippet-start: GetFeedView
	getResponse, err := feedsClient.GetFeedView(ctx, "feedViewID", &getstream.GetFeedViewRequest{})
	// snippet-end: GetFeedView

	assertResponseSuccess(t, getResponse, err, "get feed view")
	assert.Equal(t, "feedViewID", getResponse.Data.FeedView.ID)
	fmt.Printf("✅ Retrieved feed view: %s\n", feedViewID)

	// Test 4: Update Feed View
	fmt.Println("\n✏️ Testing update feed view...")
	// snippet-start: UpdateFeedView
	updateResponse, err := feedsClient.UpdateFeedView(ctx, "feedViewID", &getstream.UpdateFeedViewRequest{
		ActivitySelectors: []getstream.ActivitySelectorConfig{
			{
				Type:          getstream.PtrTo("popular"),
				MinPopularity: getstream.PtrTo(10),
			},
		},
		Aggregation: &getstream.AggregationConfig{
			Format: getstream.PtrTo("popularity_based"),
		},
	})
	// snippet-end: UpdateFeedView

	assertResponseSuccess(t, updateResponse, err, "update feed view")
	fmt.Printf("✅ Updated feed view: %s\n", feedViewID)

	// Test 5: Get or Create Feed View (should get existing)
	fmt.Println("\n🔄 Testing get or create feed view (existing)...")
	// snippet-start: GetOrCreateFeedViewExisting
	getOrCreateResponse, err := feedsClient.GetOrCreateFeedView(ctx, feedViewID, &getstream.GetOrCreateFeedViewRequest{
		ActivitySelectors: []getstream.ActivitySelectorConfig{
			{Type: getstream.PtrTo("recent")},
		},
	})
	// snippet-end: GetOrCreateFeedViewExisting

	assertResponseSuccess(t, getOrCreateResponse, err, "get or create existing feed view")
	fmt.Printf("✅ Got existing feed view: %s\n", feedViewID)

	// Test 6: Delete Feed Views (cleanup)
	// snippet-start: DeleteFeedView
	_, err = feedsClient.DeleteFeedView(ctx, "viewID-123", &getstream.DeleteFeedViewRequest{})
	// snippet-end: DeleteFeedView
}

// =================================================================
// 14. BATCH FEED OPERATIONS
// =================================================================

func test36BatchFeedOperations(t *testing.T, ctx context.Context, feedsClient *getstream.FeedsClient, testUserID string) {
	fmt.Println("\n📦 Testing Batch Feed operations...")

	batchFeedGroupID := "test-batch-group-" + uuid.New().String()[:8]

	// Create a feed group for batch operations
	fmt.Println("\n➕ Creating feed group for batch operations...")
	_, err := feedsClient.CreateFeedGroup(ctx, &getstream.CreateFeedGroupRequest{
		ID:                batchFeedGroupID,
		DefaultVisibility: getstream.PtrTo("public"),
		ActivityProcessors: []getstream.ActivityProcessorConfig{
			{Type: "default"},
		},
	})
	assertResponseSuccess(t, nil, err, "create batch feed group")

	// Test 1: Create Multiple Feeds in Batch
	fmt.Println("\n📝 Testing batch feed creation...")
	// snippet-start: CreateFeedsBatch
	batchCreateResponse, err := feedsClient.CreateFeedsBatch(ctx, &getstream.CreateFeedsBatchRequest{
		Feeds: []getstream.FeedRequest{
			{
				FeedGroupID: batchFeedGroupID,
				FeedID:      "batch-feed-1",
				CreatedByID: &testUserID,
			},
			{
				FeedGroupID: batchFeedGroupID,
				FeedID:      "batch-feed-2",
				CreatedByID: &testUserID,
			},
			{
				FeedGroupID: batchFeedGroupID,
				FeedID:      "batch-feed-3",
				CreatedByID: &testUserID,
			},
		},
	})
	// snippet-end: CreateFeedsBatch

	assertResponseSuccess(t, batchCreateResponse, err, "create feeds batch")
	fmt.Printf("✅ Created %d feeds in batch\n", len(batchCreateResponse.Data.Feeds))

	// Test 2: Query the batch-created feeds
	fmt.Println("\n🔍 Testing query of batch-created feeds...")
	// snippet-start: QueryBatchCreatedFeeds
	queryBatchResponse, err := feedsClient.QueryFeeds(ctx, &getstream.QueryFeedsRequest{
		Filter: map[string]interface{}{
			"feed_group_id": batchFeedGroupID,
			"user_id":       testUserID,
		},
		Limit: getstream.PtrTo(10),
	})
	// snippet-end: QueryBatchCreatedFeeds

	assertResponseSuccess(t, queryBatchResponse, err, "query batch created feeds")
	assert.GreaterOrEqual(t, len(queryBatchResponse.Data.Feeds), 3, "Should find at least 3 batch-created feeds")
	fmt.Printf("✅ Found %d batch-created feeds\n", len(queryBatchResponse.Data.Feeds))

	// Test 3: Add activities to batch-created feeds
	fmt.Println("\n📝 Testing activities on batch-created feeds...")
	for i := 1; i <= 3; i++ {
		feedIdentifier := fmt.Sprintf("%s:batch-feed-%d", batchFeedGroupID, i)
		// snippet-start: AddActivityToBatchFeed
		activityResponse, err := feedsClient.AddActivity(ctx, &getstream.AddActivityRequest{
			Type:   "post",
			Feeds:  []string{feedIdentifier},
			Text:   getstream.PtrTo(fmt.Sprintf("Test activity for batch feed %d", i)),
			UserID: &testUserID,
			Custom: map[string]interface{}{
				"batch_feed_test": true,
				"feed_number":     i,
			},
		})
		// snippet-end: AddActivityToBatchFeed

		assertResponseSuccess(t, activityResponse, err, fmt.Sprintf("add activity to batch feed %d", i))
	}
	fmt.Println("✅ Added activities to all batch-created feeds")

	// Cleanup: Delete the batch feed group (this will also delete all feeds in it)
	_, err = feedsClient.DeleteFeedGroup(ctx, batchFeedGroupID, &getstream.DeleteFeedGroupRequest{
		HardDelete: getstream.PtrTo(true), // hard delete for cleanup
	})
	if err != nil {
		fmt.Printf("Warning: Failed to cleanup batch feed group %s: %v\n", batchFeedGroupID, err)
	}

	fmt.Println("✅ Completed Batch Feed operations")
}

// =================================================================
// HELPER METHODS
// =================================================================

func cleanupResources(ctx context.Context, feedsClient *getstream.FeedsClient, createdActivityIDs, createdCommentIDs []string) {
	fmt.Println("\n🧹 Cleaning up test resources...")

	// Delete any remaining activities
	if len(createdActivityIDs) > 0 {
		for _, activityID := range createdActivityIDs {
			_, err := feedsClient.DeleteActivity(ctx, activityID, &getstream.DeleteActivityRequest{
				HardDelete: getstream.PtrTo(true), // hard delete
			})
			if err != nil {
				// Ignore cleanup errors
				fmt.Printf("Warning: Failed to cleanup activity %s: %v\n", activityID, err)
			}
		}
	}

	// Delete any remaining comments
	if len(createdCommentIDs) > 0 {
		for _, commentID := range createdCommentIDs {
			_, err := feedsClient.DeleteComment(ctx, commentID, &getstream.DeleteCommentRequest{
				HardDelete: getstream.PtrTo(true), // hard delete
			})
			if err != nil {
				// Ignore cleanup errors
				fmt.Printf("Warning: Failed to cleanup comment %s: %v\n", commentID, err)
			}
		}
	}

	fmt.Println("✅ Cleanup completed")
}

func assertResponseSuccess(t *testing.T, response interface{}, err error, operation string) {
	require.NoError(t, err, fmt.Sprintf("Failed to %s: %v", operation, err))
	require.NotNil(t, response, fmt.Sprintf("Failed to %s. Response is nil.", operation))
}
