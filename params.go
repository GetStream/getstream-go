package getstream

type DeleteChannelParams struct {
	HardDelete *bool `json:"hard_delete,omitempty"`
}

type DeleteDeviceParams struct {
	// Device ID to delete
	ID string `json:"id"`

	// **Server-side only**. User ID which server acts upon
	UserID *string `json:"user_id,omitempty"`
}

type DeleteFileParams struct {
	// File URL to delete
	Url *string `json:"url,omitempty"`
}

type DeleteImageParams struct {
	// File URL to delete
	Url *string `json:"url,omitempty"`
}

type DeleteMessageParams struct {
	DeletedBy *string `json:"deleted_by,omitempty"`

	// Delete all message reactions and replies as well
	Hard *bool `json:"hard,omitempty"`
}

type DeletePollOptionParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type DeletePollParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type DeleteReactionParams struct {
	// **Server-side only**. User ID which server acts upon
	UserID *string `json:"user_id,omitempty"`
}

type GetBlockedUsersParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type GetCallParams struct {
	MembersLimit *int `json:"members_limit,omitempty"`

	Notify *bool `json:"notify,omitempty"`

	Ring *bool `json:"ring,omitempty"`
}

type GetManyMessagesParams struct {
	// List of comma-separated IDs
	IDs []string `json:"ids"`
}

type GetMessageParams struct {
	ShowDeletedMessage *bool `json:"show_deleted_message,omitempty"`
}

type GetOGParams struct {
	// URL to be scraped
	Url string `json:"url"`
}

type GetPollOptionParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type GetPollParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type GetRateLimitsParams struct {
	// Whether to include Android platform limits or not
	Android *bool `json:"android,omitempty"`

	// Specific endpoints to show limits for, as a comma-separated list of values
	Endpoints *string `json:"endpoints,omitempty"`

	// Whether to include iOS platform limits or not
	Ios *bool `json:"ios,omitempty"`

	// Whether to include server-side platform limits or not
	ServerSide *bool `json:"server_side,omitempty"`

	// Whether to include web platform limits or not
	Web *bool `json:"web,omitempty"`
}

type GetReactionsParams struct {
	// Number of records to return
	Limit *int `json:"limit,omitempty"`

	// Number of records to offset
	Offset *int `json:"offset,omitempty"`
}

type GetRepliesParams struct {
	CreatedAtAfter *Timestamp `json:"created_at_after,omitempty"`

	CreatedAtAfterOrEqual *Timestamp `json:"created_at_after_or_equal,omitempty"`

	CreatedAtAround *Timestamp `json:"created_at_around,omitempty"`

	CreatedAtBefore *Timestamp `json:"created_at_before,omitempty"`

	CreatedAtBeforeOrEqual *Timestamp `json:"created_at_before_or_equal,omitempty"`

	IDAround *string `json:"id_around,omitempty"`

	IDGt *string `json:"id_gt,omitempty"`

	IDGte *string `json:"id_gte,omitempty"`

	IDLt *string `json:"id_lt,omitempty"`

	IDLte *string `json:"id_lte,omitempty"`

	Limit *int `json:"limit,omitempty"`

	Offset *int `json:"offset,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`
}

type GetThreadParams struct {
	// Limit the number of members returned per thread channel
	MemberLimit *int `json:"member_limit,omitempty"`

	// Limit the number of participants returned
	ParticipantLimit *int `json:"participant_limit,omitempty"`

	// Limit the number of replies returned
	ReplyLimit *int `json:"reply_limit,omitempty"`
}

type ListDevicesParams struct {
	// **Server-side only**. User ID which server acts upon
	UserID *string `json:"user_id,omitempty"`
}

type QueryBannedUsersParams struct {
	Payload *QueryBannedUsersRequest `json:"payload,omitempty"`
}

type QueryMembersParams struct {
	Payload *QueryMembersRequest `json:"payload,omitempty"`
}

type QueryMessageFlagsParams struct {
	Payload *QueryMessageFlagsRequest `json:"payload,omitempty"`
}

type QueryPollVotesParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type QueryPollsParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type QueryUsersParams struct {
	Payload *QueryUsersPayload `json:"payload,omitempty"`
}

type RemovePollVoteParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type SearchParams struct {
	Payload *SearchRequest `json:"payload,omitempty"`
}

type UnbanParams struct {
	TargetUserID string `json:"target_user_id"`

	ChannelCid *string `json:"channel_cid,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`
}
