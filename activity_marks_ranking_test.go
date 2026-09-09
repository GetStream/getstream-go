package getstream

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateFeedGroupRequestSerializesActivityMarksRankingIsSeen(t *testing.T) {
	t.Parallel()

	req := CreateFeedGroupRequest{
		ID: "timeline",
		ActivityMarks: &ActivityMarksConfig{
			TrackSeen: PtrTo(true),
			TrackRead: PtrTo(true),
		},
		Ranking: &RankingConfig{
			Type:  "expression",
			Score: PtrTo("is_seen ? 0 : 100"),
		},
	}

	data, err := json.Marshal(req)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))

	marks, ok := payload["activity_marks"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, true, marks["track_seen"])
	require.Equal(t, true, marks["track_read"])

	ranking, ok := payload["ranking"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "expression", ranking["type"])
	require.Equal(t, "is_seen ? 0 : 100", ranking["score"])
}
