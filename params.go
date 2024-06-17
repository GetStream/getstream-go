package getstream

type DeleteChannelParams struct {
	HardDelete *bool `json:"hard_delete,omitempty"`
}

type DeleteDeviceParams struct {
	ID string `json:"id"`

	UserID *string `json:"user_id,omitempty"`
}

type DeleteFileParams struct {
	Url *string `json:"url,omitempty"`
}

type DeleteImageParams struct {
	Url *string `json:"url,omitempty"`
}

type DeleteMessageParams struct {
	DeletedBy *string `json:"deleted_by,omitempty"`

	Hard *bool `json:"hard,omitempty"`
}

type DeletePollOptionParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type DeletePollParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type DeleteReactionParams struct {
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
	IDs []string `json:"ids"`
}

type GetMessageParams struct {
	ShowDeletedMessage *bool `json:"show_deleted_message,omitempty"`
}

type GetOGParams struct {
	Url string `json:"url"`
}

type GetPollOptionParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type GetPollParams struct {
	UserID *string `json:"user_id,omitempty"`
}

type GetRateLimitsParams struct {
	Android *bool `json:"android,omitempty"`

	Endpoints *string `json:"endpoints,omitempty"`

	Ios *bool `json:"ios,omitempty"`

	ServerSide *bool `json:"server_side,omitempty"`

	Web *bool `json:"web,omitempty"`
}

type GetReactionsParams struct {
	Limit *int `json:"limit,omitempty"`

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
	MemberLimit *int `json:"member_limit,omitempty"`

	ParticipantLimit *int `json:"participant_limit,omitempty"`

	ReplyLimit *int `json:"reply_limit,omitempty"`
}

type ListDevicesParams struct {
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
