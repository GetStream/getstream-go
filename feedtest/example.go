package main

import (
	"context"
	"math/rand"

	"github.com/GetStream/getstream-go/v3"
)

func getRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Int()%len(charset)]
	}
	return string(result)
}

func main() {
	// init client, create feed, add activities, etc.
	client, err := getstream.NewClientFromEnvVars()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	// random generate userID
	userID1 := getRandomString(10)
	userID2 := getRandomString(10)

	_, err = client.UpdateUsers(ctx, &getstream.UpdateUsersRequest{
		Users: map[string]getstream.UserRequest{
			userID1: {
				ID: userID1,
			},
			userID2: {
				ID: userID2,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	feedsClient := client.Feeds()

	feedOrigin := feedsClient.Feed("user", userID1)
	feedOriginRes, err := feedOrigin.GetOrCreate(ctx, &getstream.GetOrCreateFeedRequest{
		UserID: getstream.PtrTo(userID1),
	})
	if err != nil {
		panic(err)
	}

	feedFollower := feedsClient.Feed("user", userID2)
	feedFollowerRes, err := feedFollower.GetOrCreate(ctx, &getstream.GetOrCreateFeedRequest{
		UserID: getstream.PtrTo(userID2),
	})
	if err != nil {
		panic(err)
	}
	_, err = feedsClient.Follow(ctx, &getstream.FollowRequest{
		Source: feedFollowerRes.Data.Feed.Feed,
		Target: feedOriginRes.Data.Feed.Feed,
	})
	if err != nil {
		panic(err)
	}
	activity := &getstream.AddActivityRequest{
		Type:   "post1",
		Feeds:  []string{feedOriginRes.Data.Feed.Feed},
		Text:   getstream.PtrTo(getRandomString(10)),
		UserID: getstream.PtrTo(userID1),
	}

	_, err = feedsClient.AddActivity(ctx, activity)
	if err != nil {
		panic(err)
	}

	// fetch both feeds
	originActivities, err := feedOrigin.GetOrCreate(ctx, &getstream.GetOrCreateFeedRequest{
		UserID: getstream.PtrTo(userID1),
	})
	if err != nil {
		panic(err)
	}

	followerActivities, err := feedFollower.GetOrCreate(ctx, &getstream.GetOrCreateFeedRequest{
		UserID: getstream.PtrTo(userID2),
	})
	if err != nil {
		panic(err)
	}

	// Print activities
	for _, activity := range originActivities.Data.Activities {
		println("Origin Activity:", activity.ID, *activity.Text, activity.Type)
	}
	for _, activity := range followerActivities.Data.Activities {
		println("Follower Activity:", activity.ID, *activity.Text, activity.Type)
	}
}
