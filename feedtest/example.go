package main

import (
	"context"
	"math/rand"
	"os"

	"github.com/GetStream/getstream-go"
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

	//read API key and secret from environment
	apiKey := os.Getenv("STREAM_API_KEY")
	apiSecret := os.Getenv("STREAM_API_SECRET")
	client, err := getstream.NewClient(apiKey, apiSecret)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	//random generate userID
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
	feedOrigin, err := client.Feeds().GetOrCreateFeed(ctx, "user", userID1, &getstream.GetOrCreateFeedRequest{
		UserID: getstream.PtrTo(userID1),
	})
	if err != nil {
		panic(err)
	}

	feedFollower, err := client.Feeds().GetOrCreateFeed(ctx, "user", userID2, &getstream.GetOrCreateFeedRequest{
		UserID: getstream.PtrTo(userID2),
	})
	if err != nil {
		panic(err)
	}
	_, err = client.Feeds().Follow(ctx, &getstream.FollowRequest{
		Source: feedFollower.Data.Feed.Fid,
		Target: feedOrigin.Data.Feed.Fid,
	})
	if err != nil {
		panic(err)
	}
	activity := &getstream.AddActivityRequest{
		Type:   "post1",
		Fids:   []string{feedOrigin.Data.Feed.Fid},
		Text:   getstream.PtrTo(getRandomString(10)),
		UserID: getstream.PtrTo(userID1),
	}

	_, err = client.Feeds().AddActivity(ctx, activity)
	if err != nil {
		panic(err)
	}

	// fetch both feeds
	originActivities, err := client.Feeds().GetOrCreateFeed(ctx, "user", userID1, &getstream.GetOrCreateFeedRequest{
		UserID: getstream.PtrTo(userID1),
	})
	if err != nil {
		panic(err)
	}

	followerActivities, err := client.Feeds().GetOrCreateFeed(ctx, "user", userID2, &getstream.GetOrCreateFeedRequest{
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
