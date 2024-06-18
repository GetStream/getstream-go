package getstream

type APIError struct {
	Code int `json:"code"`

	Duration string `json:"duration"`

	Message string `json:"message"`

	MoreInfo string `json:"more_info"`

	StatusCode int `json:"StatusCode"`

	Details []int `json:"details"`

	ExceptionFields *map[string]string `json:"exception_fields,omitempty"`
}

type APNConfig struct {
	AuthKey *string `json:"auth_key,omitempty"`

	AuthType *string `json:"auth_type,omitempty"`

	BundleID *string `json:"bundle_id,omitempty"`

	Development *bool `json:"development,omitempty"`

	Disabled *bool `json:"Disabled,omitempty"`

	Host *string `json:"host,omitempty"`

	KeyID *string `json:"key_id,omitempty"`

	NotificationTemplate *string `json:"notification_template,omitempty"`

	P12Cert *string `json:"p12_cert,omitempty"`

	TeamID *string `json:"team_id,omitempty"`
}

type APNConfigFields struct {
	Development bool `json:"development"`

	Enabled bool `json:"enabled"`

	NotificationTemplate string `json:"notification_template"`

	AuthKey *string `json:"auth_key,omitempty"`

	AuthType *string `json:"auth_type,omitempty"`

	BundleID *string `json:"bundle_id,omitempty"`

	Host *string `json:"host,omitempty"`

	KeyID *string `json:"key_id,omitempty"`

	P12Cert *string `json:"p12_cert,omitempty"`

	TeamID *string `json:"team_id,omitempty"`
}

type APNS struct {
	Body string `json:"body"`

	Title string `json:"title"`
}

type Action struct {
	Name string `json:"name"`

	Text string `json:"text"`

	Type string `json:"type"`

	Style *string `json:"style,omitempty"`

	Value *string `json:"value,omitempty"`
}

type AppResponseFields struct {
	AsyncUrlEnrichEnabled bool `json:"async_url_enrich_enabled"`

	AutoTranslationEnabled bool `json:"auto_translation_enabled"`

	CampaignEnabled bool `json:"campaign_enabled"`

	CdnExpirationSeconds int `json:"cdn_expiration_seconds"`

	CustomActionHandlerUrl string `json:"custom_action_handler_url"`

	DisableAuthChecks bool `json:"disable_auth_checks"`

	DisablePermissionsChecks bool `json:"disable_permissions_checks"`

	EnforceUniqueUsernames string `json:"enforce_unique_usernames"`

	ImageModerationEnabled bool `json:"image_moderation_enabled"`

	MultiTenantEnabled bool `json:"multi_tenant_enabled"`

	Name string `json:"name"`

	Organization string `json:"organization"`

	PermissionVersion string `json:"permission_version"`

	PollsEnabled bool `json:"polls_enabled"`

	RemindersInterval int `json:"reminders_interval"`

	SnsKey string `json:"sns_key"`

	SnsSecret string `json:"sns_secret"`

	SnsTopicArn string `json:"sns_topic_arn"`

	SqsKey string `json:"sqs_key"`

	SqsSecret string `json:"sqs_secret"`

	SqsUrl string `json:"sqs_url"`

	Suspended bool `json:"suspended"`

	SuspendedExplanation string `json:"suspended_explanation"`

	VideoProvider string `json:"video_provider"`

	WebhookUrl string `json:"webhook_url"`

	UserSearchDisallowedRoles []string `json:"user_search_disallowed_roles"`

	WebhookEvents []string `json:"webhook_events"`

	CallTypes map[string]*CallType `json:"call_types"`

	ChannelConfigs map[string]*ChannelConfig `json:"channel_configs"`

	FileUploadConfig FileUploadConfig `json:"file_upload_config"`

	Grants map[string][]string `json:"grants"`

	ImageUploadConfig FileUploadConfig `json:"image_upload_config"`

	Policies map[string][]Policy `json:"policies"`

	PushNotifications PushNotificationFields `json:"push_notifications"`

	BeforeMessageSendHookUrl *string `json:"before_message_send_hook_url,omitempty"`

	RevokeTokensIssuedBefore *Timestamp `json:"revoke_tokens_issued_before,omitempty"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons,omitempty"`

	Geofences *[]*GeofenceResponse `json:"geofences,omitempty"`

	ImageModerationLabels *[]string `json:"image_moderation_labels,omitempty"`

	AgoraOptions *Config `json:"agora_options,omitempty"`

	DatadogInfo *DataDogInfo `json:"datadog_info,omitempty"`

	HmsOptions *Config `json:"hms_options,omitempty"`
}

type AsyncModerationCallbackConfig struct {
	Mode *string `json:"mode,omitempty"`

	ServerUrl *string `json:"server_url,omitempty"`
}

type AsyncModerationConfiguration struct {
	TimeoutMs *int `json:"timeout_ms,omitempty"`

	Callback *AsyncModerationCallbackConfig `json:"callback,omitempty"`
}

type Attachment struct {
	Custom map[string]any `json:"custom"`

	AssetUrl *string `json:"asset_url,omitempty"`

	AuthorIcon *string `json:"author_icon,omitempty"`

	AuthorLink *string `json:"author_link,omitempty"`

	AuthorName *string `json:"author_name,omitempty"`

	Color *string `json:"color,omitempty"`

	Fallback *string `json:"fallback,omitempty"`

	Footer *string `json:"footer,omitempty"`

	FooterIcon *string `json:"footer_icon,omitempty"`

	ImageUrl *string `json:"image_url,omitempty"`

	OgScrapeUrl *string `json:"og_scrape_url,omitempty"`

	OriginalHeight *int `json:"original_height,omitempty"`

	OriginalWidth *int `json:"original_width,omitempty"`

	Pretext *string `json:"pretext,omitempty"`

	Text *string `json:"text,omitempty"`

	ThumbUrl *string `json:"thumb_url,omitempty"`

	Title *string `json:"title,omitempty"`

	TitleLink *string `json:"title_link,omitempty"`

	Type *string `json:"type,omitempty"`

	Actions *[]*Action `json:"actions,omitempty"`

	Fields *[]*Field `json:"fields,omitempty"`

	Giphy *Images `json:"giphy,omitempty"`
}

type AudioSettings struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	DefaultDevice string `json:"default_device"`

	MicDefaultOn bool `json:"mic_default_on"`

	OpusDtxEnabled bool `json:"opus_dtx_enabled"`

	RedundantCodingEnabled bool `json:"redundant_coding_enabled"`

	SpeakerDefaultOn bool `json:"speaker_default_on"`

	NoiseCancellation *NoiseCancellationSettings `json:"noise_cancellation,omitempty"`
}

type AudioSettingsRequest struct {
	DefaultDevice string `json:"default_device"`

	AccessRequestEnabled *bool `json:"access_request_enabled,omitempty"`

	MicDefaultOn *bool `json:"mic_default_on,omitempty"`

	OpusDtxEnabled *bool `json:"opus_dtx_enabled,omitempty"`

	RedundantCodingEnabled *bool `json:"redundant_coding_enabled,omitempty"`

	SpeakerDefaultOn *bool `json:"speaker_default_on,omitempty"`

	NoiseCancellation *NoiseCancellationSettings `json:"noise_cancellation,omitempty"`
}

type AudioSettingsResponse struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	DefaultDevice string `json:"default_device"`

	MicDefaultOn bool `json:"mic_default_on"`

	OpusDtxEnabled bool `json:"opus_dtx_enabled"`

	RedundantCodingEnabled bool `json:"redundant_coding_enabled"`

	SpeakerDefaultOn bool `json:"speaker_default_on"`

	NoiseCancellation *NoiseCancellationSettings `json:"noise_cancellation,omitempty"`
}

type AutomodDetails struct {
	Action *string `json:"action,omitempty"`

	OriginalMessageType *string `json:"original_message_type,omitempty"`

	ImageLabels *[]string `json:"image_labels,omitempty"`

	MessageDetails *FlagMessageDetails `json:"message_details,omitempty"`

	Result *MessageModerationResult `json:"result,omitempty"`
}

type AzureRequest struct {
	AbsAccountName string `json:"abs_account_name"`

	AbsClientID string `json:"abs_client_id"`

	AbsClientSecret string `json:"abs_client_secret"`

	AbsTenantID string `json:"abs_tenant_id"`
}

type BackstageSettings struct {
	Enabled bool `json:"enabled"`
}

type BackstageSettingsRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type BackstageSettingsResponse struct {
	Enabled bool `json:"enabled"`
}

type BanRequest struct {
	TargetUserID string `json:"target_user_id"`

	BannedByID *string `json:"banned_by_id,omitempty"`

	ChannelCid *string `json:"channel_cid,omitempty"`

	IpBan *bool `json:"ip_ban,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Shadow *bool `json:"shadow,omitempty"`

	Timeout *int `json:"timeout,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	BannedBy *UserRequest `json:"banned_by,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type BanResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	Expires *Timestamp `json:"expires,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Shadow *bool `json:"shadow,omitempty"`

	BannedBy *UserObject `json:"banned_by,omitempty"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

// Block list contains restricted words
type BlockList struct {
	Name string `json:"name"`

	Type string `json:"type"`

	Words []string `json:"words"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
}

type BlockListOptions struct {
	Behavior string `json:"behavior"`

	Blocklist string `json:"blocklist"`
}

type BlockUserRequest struct {
	UserID string `json:"user_id"`
}

type BlockUserResponse struct {
	Duration string `json:"duration"`
}

type BlockUsersRequest struct {
	BlockedUserID string `json:"blocked_user_id"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type BlockUsersResponse struct {
	BlockedByUserID string `json:"blocked_by_user_id"`

	BlockedUserID string `json:"blocked_user_id"`

	CreatedAt Timestamp `json:"created_at"`

	Duration string `json:"duration"`
}

type BlockedUserResponse struct {
	BlockedUserID string `json:"blocked_user_id"`

	CreatedAt Timestamp `json:"created_at"`

	UserID string `json:"user_id"`

	BlockedUser UserResponse `json:"blocked_user"`

	User UserResponse `json:"user"`
}

type BroadcastSettings struct {
	Enabled bool `json:"enabled"`

	Hls HLSSettings `json:"hls"`
}

type BroadcastSettingsRequest struct {
	Enabled *bool `json:"enabled,omitempty"`

	Hls *HLSSettingsRequest `json:"hls,omitempty"`
}

type BroadcastSettingsResponse struct {
	Enabled bool `json:"enabled"`

	Hls HLSSettingsResponse `json:"hls"`
}

type CallEvent struct {
	Description string `json:"description"`

	EndTimestamp int `json:"end_timestamp"`

	Severity int `json:"severity"`

	Timestamp int `json:"timestamp"`

	Type string `json:"type"`
}

type CallIngressResponse struct {
	Rtmp RTMPIngress `json:"rtmp"`
}

type CallParticipantResponse struct {
	JoinedAt Timestamp `json:"joined_at"`

	Role string `json:"role"`

	UserSessionID string `json:"user_session_id"`

	User UserResponse `json:"user"`
}

// CallRecording represents a recording of a call.
type CallRecording struct {
	EndTime Timestamp `json:"end_time"`

	Filename string `json:"filename"`

	StartTime Timestamp `json:"start_time"`

	Url string `json:"url"`
}

type CallRequest struct {
	CreatedByID *string `json:"created_by_id,omitempty"`

	StartsAt *Timestamp `json:"starts_at,omitempty"`

	Team *string `json:"team,omitempty"`

	Members *[]MemberRequest `json:"members,omitempty"`

	CreatedBy *UserRequest `json:"created_by,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	SettingsOverride *CallSettingsRequest `json:"settings_override,omitempty"`
}

// Represents a call
type CallResponse struct {
	Backstage bool `json:"backstage"`

	Cid string `json:"cid"`

	CreatedAt Timestamp `json:"created_at"`

	CurrentSessionID string `json:"current_session_id"`

	ID string `json:"id"`

	Recording bool `json:"recording"`

	Transcribing bool `json:"transcribing"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	BlockedUserIDs []string `json:"blocked_user_ids"`

	CreatedBy UserResponse `json:"created_by"`

	Custom map[string]any `json:"custom"`

	Egress EgressResponse `json:"egress"`

	Ingress CallIngressResponse `json:"ingress"`

	Settings CallSettingsResponse `json:"settings"`

	EndedAt *Timestamp `json:"ended_at,omitempty"`

	StartsAt *Timestamp `json:"starts_at,omitempty"`

	Team *string `json:"team,omitempty"`

	Session *CallSessionResponse `json:"session,omitempty"`

	Thumbnails *ThumbnailResponse `json:"thumbnails,omitempty"`
}

type CallSessionResponse struct {
	ID string `json:"id"`

	Participants []CallParticipantResponse `json:"participants"`

	AcceptedBy map[string]Timestamp `json:"accepted_by"`

	MissedBy map[string]Timestamp `json:"missed_by"`

	ParticipantsCountByRole map[string]int `json:"participants_count_by_role"`

	RejectedBy map[string]Timestamp `json:"rejected_by"`

	EndedAt *Timestamp `json:"ended_at,omitempty"`

	LiveEndedAt *Timestamp `json:"live_ended_at,omitempty"`

	LiveStartedAt *Timestamp `json:"live_started_at,omitempty"`

	StartedAt *Timestamp `json:"started_at,omitempty"`

	TimerEndsAt *Timestamp `json:"timer_ends_at,omitempty"`
}

type CallSettings struct {
	Audio *AudioSettings `json:"audio,omitempty"`

	Backstage *BackstageSettings `json:"backstage,omitempty"`

	Broadcasting *BroadcastSettings `json:"broadcasting,omitempty"`

	Geofencing *GeofenceSettings `json:"geofencing,omitempty"`

	Limits *LimitsSettings `json:"limits,omitempty"`

	Recording *RecordSettings `json:"recording,omitempty"`

	Ring *RingSettings `json:"ring,omitempty"`

	Screensharing *ScreensharingSettings `json:"screensharing,omitempty"`

	Thumbnails *ThumbnailsSettings `json:"thumbnails,omitempty"`

	Transcription *TranscriptionSettings `json:"transcription,omitempty"`

	Video *VideoSettings `json:"video,omitempty"`
}

type CallSettingsRequest struct {
	Audio *AudioSettingsRequest `json:"audio,omitempty"`

	Backstage *BackstageSettingsRequest `json:"backstage,omitempty"`

	Broadcasting *BroadcastSettingsRequest `json:"broadcasting,omitempty"`

	Geofencing *GeofenceSettingsRequest `json:"geofencing,omitempty"`

	Limits *LimitsSettingsRequest `json:"limits,omitempty"`

	Recording *RecordSettingsRequest `json:"recording,omitempty"`

	Ring *RingSettingsRequest `json:"ring,omitempty"`

	Screensharing *ScreensharingSettingsRequest `json:"screensharing,omitempty"`

	Thumbnails *ThumbnailsSettingsRequest `json:"thumbnails,omitempty"`

	Transcription *TranscriptionSettingsRequest `json:"transcription,omitempty"`

	Video *VideoSettingsRequest `json:"video,omitempty"`
}

type CallSettingsResponse struct {
	Audio AudioSettingsResponse `json:"audio"`

	Backstage BackstageSettingsResponse `json:"backstage"`

	Broadcasting BroadcastSettingsResponse `json:"broadcasting"`

	Geofencing GeofenceSettingsResponse `json:"geofencing"`

	Limits LimitsSettingsResponse `json:"limits"`

	Recording RecordSettingsResponse `json:"recording"`

	Ring RingSettingsResponse `json:"ring"`

	Screensharing ScreensharingSettingsResponse `json:"screensharing"`

	Thumbnails ThumbnailsSettingsResponse `json:"thumbnails"`

	Transcription TranscriptionSettingsResponse `json:"transcription"`

	Video VideoSettingsResponse `json:"video"`
}

type CallStateResponseFields struct {
	Members []MemberResponse `json:"members"`

	OwnCapabilities []OwnCapability `json:"own_capabilities"`

	Call CallResponse `json:"call"`
}

type CallStatsReportSummaryResponse struct {
	CallCid string `json:"call_cid"`

	CallDurationSeconds int `json:"call_duration_seconds"`

	CallSessionID string `json:"call_session_id"`

	CallStatus string `json:"call_status"`

	FirstStatsTime Timestamp `json:"first_stats_time"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	QualityScore *int `json:"quality_score,omitempty"`
}

type CallTimeline struct {
	Events []*CallEvent `json:"events"`
}

// CallTranscription represents a transcription of a call.
type CallTranscription struct {
	EndTime Timestamp `json:"end_time"`

	Filename string `json:"filename"`

	StartTime Timestamp `json:"start_time"`

	Url string `json:"url"`
}

type CallType struct {
	AppPK int `json:"AppPK"`

	CreatedAt Timestamp `json:"CreatedAt"`

	ExternalStorage string `json:"ExternalStorage"`

	Name string `json:"Name"`

	PK int `json:"PK"`

	UpdatedAt Timestamp `json:"UpdatedAt"`

	NotificationSettings *NotificationSettings `json:"NotificationSettings,omitempty"`

	Settings *CallSettings `json:"Settings,omitempty"`
}

type CallTypeResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	Name string `json:"name"`

	UpdatedAt Timestamp `json:"updated_at"`

	Grants map[string][]string `json:"grants"`

	NotificationSettings NotificationSettings `json:"notification_settings"`

	Settings CallSettingsResponse `json:"settings"`

	ExternalStorage *string `json:"external_storage,omitempty"`
}

type CastPollVoteRequest struct {
	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`

	Vote *VoteData `json:"vote,omitempty"`
}

type Channel struct {
	AutoTranslationLanguage string `json:"auto_translation_language"`

	Cid string `json:"cid"`

	CreatedAt Timestamp `json:"created_at"`

	Disabled bool `json:"disabled"`

	Frozen bool `json:"frozen"`

	ID string `json:"id"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	AutoTranslationEnabled *bool `json:"auto_translation_enabled,omitempty"`

	Cooldown *int `json:"cooldown,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	LastMessageAt *Timestamp `json:"last_message_at,omitempty"`

	MemberCount *int `json:"member_count,omitempty"`

	Team *string `json:"team,omitempty"`

	Invites *[]*ChannelMember `json:"invites,omitempty"`

	Members *[]*ChannelMember `json:"members,omitempty"`

	Config *ChannelConfig `json:"config,omitempty"`

	ConfigOverrides *ChannelConfig `json:"config_overrides,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	TruncatedBy *UserObject `json:"truncated_by,omitempty"`
}

type ChannelConfig struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt Timestamp `json:"created_at"`

	CustomEvents bool `json:"custom_events"`

	MarkMessagesPending bool `json:"mark_messages_pending"`

	MaxMessageLength int `json:"max_message_length"`

	Mutes bool `json:"mutes"`

	Name string `json:"name"`

	Polls bool `json:"polls"`

	PushNotifications bool `json:"push_notifications"`

	Quotes bool `json:"quotes"`

	Reactions bool `json:"reactions"`

	ReadEvents bool `json:"read_events"`

	Reminders bool `json:"reminders"`

	Replies bool `json:"replies"`

	Search bool `json:"search"`

	TypingEvents bool `json:"typing_events"`

	UpdatedAt Timestamp `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []string `json:"commands"`

	Blocklist *string `json:"blocklist,omitempty"`

	BlocklistBehavior *string `json:"blocklist_behavior,omitempty"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons,omitempty"`

	Blocklists *[]BlockListOptions `json:"blocklists,omitempty"`

	AutomodThresholds *Thresholds `json:"automod_thresholds,omitempty"`
}

type ChannelConfigWithInfo struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt Timestamp `json:"created_at"`

	CustomEvents bool `json:"custom_events"`

	MarkMessagesPending bool `json:"mark_messages_pending"`

	MaxMessageLength int `json:"max_message_length"`

	Mutes bool `json:"mutes"`

	Name string `json:"name"`

	Polls bool `json:"polls"`

	PushNotifications bool `json:"push_notifications"`

	Quotes bool `json:"quotes"`

	Reactions bool `json:"reactions"`

	ReadEvents bool `json:"read_events"`

	Reminders bool `json:"reminders"`

	Replies bool `json:"replies"`

	Search bool `json:"search"`

	TypingEvents bool `json:"typing_events"`

	UpdatedAt Timestamp `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []*Command `json:"commands"`

	Blocklist *string `json:"blocklist,omitempty"`

	BlocklistBehavior *string `json:"blocklist_behavior,omitempty"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons,omitempty"`

	Blocklists *[]BlockListOptions `json:"blocklists,omitempty"`

	AutomodThresholds *Thresholds `json:"automod_thresholds,omitempty"`

	Grants *map[string][]string `json:"grants,omitempty"`
}

type ChannelExport struct {
	Cid *string `json:"cid,omitempty"`

	ID *string `json:"id,omitempty"`

	MessagesSince *Timestamp `json:"messages_since,omitempty"`

	MessagesUntil *Timestamp `json:"messages_until,omitempty"`

	Type *string `json:"type,omitempty"`
}

type ChannelGetOrCreateRequest struct {
	HideForCreator *bool `json:"hide_for_creator,omitempty"`

	State *bool `json:"state,omitempty"`

	ThreadUnreadCounts *bool `json:"thread_unread_counts,omitempty"`

	Data *ChannelInput `json:"data,omitempty"`

	Members *PaginationParams `json:"members,omitempty"`

	Messages *MessagePaginationParams `json:"messages,omitempty"`

	Watchers *PaginationParams `json:"watchers,omitempty"`
}

type ChannelInput struct {
	AutoTranslationEnabled *bool `json:"auto_translation_enabled,omitempty"`

	AutoTranslationLanguage *string `json:"auto_translation_language,omitempty"`

	CreatedByID *string `json:"created_by_id,omitempty"`

	Disabled *bool `json:"disabled,omitempty"`

	Frozen *bool `json:"frozen,omitempty"`

	Team *string `json:"team,omitempty"`

	TruncatedByID *string `json:"truncated_by_id,omitempty"`

	Invites *[]*ChannelMember `json:"invites,omitempty"`

	Members *[]*ChannelMember `json:"members,omitempty"`

	ConfigOverrides *ChannelConfig `json:"config_overrides,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`
}

type ChannelMember struct {
	Banned bool `json:"banned"`

	ChannelRole string `json:"channel_role"`

	CreatedAt Timestamp `json:"created_at"`

	NotificationsMuted bool `json:"notifications_muted"`

	ShadowBanned bool `json:"shadow_banned"`

	UpdatedAt Timestamp `json:"updated_at"`

	BanExpires *Timestamp `json:"ban_expires,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	InviteAcceptedAt *Timestamp `json:"invite_accepted_at,omitempty"`

	InviteRejectedAt *Timestamp `json:"invite_rejected_at,omitempty"`

	Invited *bool `json:"invited,omitempty"`

	IsModerator *bool `json:"is_moderator,omitempty"`

	Status *string `json:"status,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type ChannelMute struct {
	CreatedAt Timestamp `json:"created_at"`

	UpdatedAt Timestamp `json:"updated_at"`

	Expires *Timestamp `json:"expires,omitempty"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

// Represents channel in chat
type ChannelResponse struct {
	Cid string `json:"cid"`

	CreatedAt Timestamp `json:"created_at"`

	Disabled bool `json:"disabled"`

	Frozen bool `json:"frozen"`

	ID string `json:"id"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	AutoTranslationEnabled *bool `json:"auto_translation_enabled,omitempty"`

	AutoTranslationLanguage *string `json:"auto_translation_language,omitempty"`

	Blocked *bool `json:"blocked,omitempty"`

	Cooldown *int `json:"cooldown,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	Hidden *bool `json:"hidden,omitempty"`

	HideMessagesBefore *Timestamp `json:"hide_messages_before,omitempty"`

	LastMessageAt *Timestamp `json:"last_message_at,omitempty"`

	MemberCount *int `json:"member_count,omitempty"`

	MuteExpiresAt *Timestamp `json:"mute_expires_at,omitempty"`

	Muted *bool `json:"muted,omitempty"`

	Team *string `json:"team,omitempty"`

	TruncatedAt *Timestamp `json:"truncated_at,omitempty"`

	Members *[]*ChannelMember `json:"members,omitempty"`

	OwnCapabilities *[]string `json:"own_capabilities,omitempty"`

	Config *ChannelConfigWithInfo `json:"config,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	TruncatedBy *UserObject `json:"truncated_by,omitempty"`
}

type ChannelStateResponse struct {
	Duration string `json:"duration"`

	Members []*ChannelMember `json:"members"`

	Messages []MessageResponse `json:"messages"`

	PinnedMessages []MessageResponse `json:"pinned_messages"`

	Threads []*ThreadState `json:"threads"`

	Hidden *bool `json:"hidden,omitempty"`

	HideMessagesBefore *Timestamp `json:"hide_messages_before,omitempty"`

	WatcherCount *int `json:"watcher_count,omitempty"`

	PendingMessages *[]*PendingMessage `json:"pending_messages,omitempty"`

	Read *[]ReadStateResponse `json:"read,omitempty"`

	Watchers *[]UserResponse `json:"watchers,omitempty"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	Membership *ChannelMember `json:"membership,omitempty"`
}

type ChannelStateResponseFields struct {
	Members []*ChannelMember `json:"members"`

	Messages []MessageResponse `json:"messages"`

	PinnedMessages []MessageResponse `json:"pinned_messages"`

	Threads []*ThreadState `json:"threads"`

	Hidden *bool `json:"hidden,omitempty"`

	HideMessagesBefore *Timestamp `json:"hide_messages_before,omitempty"`

	WatcherCount *int `json:"watcher_count,omitempty"`

	PendingMessages *[]*PendingMessage `json:"pending_messages,omitempty"`

	Read *[]ReadStateResponse `json:"read,omitempty"`

	Watchers *[]UserResponse `json:"watchers,omitempty"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	Membership *ChannelMember `json:"membership,omitempty"`
}

type ChannelTypeConfig struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt Timestamp `json:"created_at"`

	CustomEvents bool `json:"custom_events"`

	MarkMessagesPending bool `json:"mark_messages_pending"`

	MaxMessageLength int `json:"max_message_length"`

	Mutes bool `json:"mutes"`

	Name string `json:"name"`

	Polls bool `json:"polls"`

	PushNotifications bool `json:"push_notifications"`

	Quotes bool `json:"quotes"`

	Reactions bool `json:"reactions"`

	ReadEvents bool `json:"read_events"`

	Reminders bool `json:"reminders"`

	Replies bool `json:"replies"`

	Search bool `json:"search"`

	TypingEvents bool `json:"typing_events"`

	UpdatedAt Timestamp `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []*Command `json:"commands"`

	Permissions []PolicyRequest `json:"permissions"`

	Grants map[string][]string `json:"grants"`

	Blocklist *string `json:"blocklist,omitempty"`

	BlocklistBehavior *string `json:"blocklist_behavior,omitempty"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons,omitempty"`

	Blocklists *[]BlockListOptions `json:"blocklists,omitempty"`

	AutomodThresholds *Thresholds `json:"automod_thresholds,omitempty"`
}

type CheckExternalStorageResponse struct {
	Duration string `json:"duration"`

	FileUrl string `json:"file_url"`
}

type CheckPushRequest struct {
	ApnTemplate *string `json:"apn_template,omitempty"`

	FirebaseDataTemplate *string `json:"firebase_data_template,omitempty"`

	FirebaseTemplate *string `json:"firebase_template,omitempty"`

	MessageID *string `json:"message_id,omitempty"`

	PushProviderName *string `json:"push_provider_name,omitempty"`

	PushProviderType *string `json:"push_provider_type,omitempty"`

	SkipDevices *bool `json:"skip_devices,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type CheckPushResponse struct {
	Duration string `json:"duration"`

	RenderedApnTemplate *string `json:"rendered_apn_template,omitempty"`

	RenderedFirebaseTemplate *string `json:"rendered_firebase_template,omitempty"`

	SkipDevices *bool `json:"skip_devices,omitempty"`

	GeneralErrors *[]string `json:"general_errors,omitempty"`

	DeviceErrors *map[string]DeviceErrorInfo `json:"device_errors,omitempty"`

	RenderedMessage *map[string]string `json:"rendered_message,omitempty"`
}

type CheckSNSRequest struct {
	SnsKey *string `json:"sns_key,omitempty"`

	SnsSecret *string `json:"sns_secret,omitempty"`

	SnsTopicArn *string `json:"sns_topic_arn,omitempty"`
}

type CheckSNSResponse struct {
	Duration string `json:"duration"`

	Status string `json:"status"`

	Error *string `json:"error,omitempty"`

	Data *map[string]any `json:"data,omitempty"`
}

type CheckSQSRequest struct {
	SqsKey *string `json:"sqs_key,omitempty"`

	SqsSecret *string `json:"sqs_secret,omitempty"`

	SqsUrl *string `json:"sqs_url,omitempty"`
}

type CheckSQSResponse struct {
	Duration string `json:"duration"`

	Status string `json:"status"`

	Error *string `json:"error,omitempty"`

	Data *map[string]any `json:"data,omitempty"`
}

type CollectUserFeedbackRequest struct {
	Rating int `json:"rating"`

	Sdk string `json:"sdk"`

	SdkVersion string `json:"sdk_version"`

	UserSessionID string `json:"user_session_id"`

	Reason *string `json:"reason,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`
}

type CollectUserFeedbackResponse struct {
	Duration string `json:"duration"`
}

// Represents custom chat command
type Command struct {
	Args string `json:"args"`

	Description string `json:"description"`

	Name string `json:"name"`

	Set string `json:"set"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
}

type CommitMessageRequest struct{}

type Config struct {
	AppCertificate string `json:"app_certificate"`

	AppID string `json:"app_id"`

	DefaultRole *string `json:"default_role,omitempty"`

	RoleMap *map[string]string `json:"role_map,omitempty"`
}

type Coordinates struct {
	Latitude float64 `json:"latitude"`

	Longitude float64 `json:"longitude"`
}

// Block list contains restricted words
type CreateBlockListRequest struct {
	Name string `json:"name"`

	Words []string `json:"words"`

	Type *string `json:"type,omitempty"`
}

type CreateCallTypeRequest struct {
	Name string `json:"name"`

	ExternalStorage *string `json:"external_storage,omitempty"`

	Grants *map[string][]string `json:"grants,omitempty"`

	NotificationSettings *NotificationSettings `json:"notification_settings,omitempty"`

	Settings *CallSettingsRequest `json:"settings,omitempty"`
}

type CreateCallTypeResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	UpdatedAt Timestamp `json:"updated_at"`

	Grants map[string][]string `json:"grants"`

	NotificationSettings NotificationSettings `json:"notification_settings"`

	Settings CallSettingsResponse `json:"settings"`

	ExternalStorage *string `json:"external_storage,omitempty"`
}

type CreateChannelTypeRequest struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	MaxMessageLength int `json:"max_message_length"`

	Name string `json:"name"`

	Blocklist *string `json:"blocklist,omitempty"`

	BlocklistBehavior *string `json:"blocklist_behavior,omitempty"`

	ConnectEvents *bool `json:"connect_events,omitempty"`

	CustomEvents *bool `json:"custom_events,omitempty"`

	MarkMessagesPending *bool `json:"mark_messages_pending,omitempty"`

	MessageRetention *string `json:"message_retention,omitempty"`

	Mutes *bool `json:"mutes,omitempty"`

	Polls *bool `json:"polls,omitempty"`

	PushNotifications *bool `json:"push_notifications,omitempty"`

	Reactions *bool `json:"reactions,omitempty"`

	ReadEvents *bool `json:"read_events,omitempty"`

	Replies *bool `json:"replies,omitempty"`

	Search *bool `json:"search,omitempty"`

	TypingEvents *bool `json:"typing_events,omitempty"`

	Uploads *bool `json:"uploads,omitempty"`

	UrlEnrichment *bool `json:"url_enrichment,omitempty"`

	Blocklists *[]BlockListOptions `json:"blocklists,omitempty"`

	Commands *[]string `json:"commands,omitempty"`

	Permissions *[]PolicyRequest `json:"permissions,omitempty"`

	Grants *map[string][]string `json:"grants,omitempty"`
}

type CreateChannelTypeResponse struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt Timestamp `json:"created_at"`

	CustomEvents bool `json:"custom_events"`

	Duration string `json:"duration"`

	MarkMessagesPending bool `json:"mark_messages_pending"`

	MaxMessageLength int `json:"max_message_length"`

	Mutes bool `json:"mutes"`

	Name string `json:"name"`

	Polls bool `json:"polls"`

	PushNotifications bool `json:"push_notifications"`

	Quotes bool `json:"quotes"`

	Reactions bool `json:"reactions"`

	ReadEvents bool `json:"read_events"`

	Reminders bool `json:"reminders"`

	Replies bool `json:"replies"`

	Search bool `json:"search"`

	TypingEvents bool `json:"typing_events"`

	UpdatedAt Timestamp `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []string `json:"commands"`

	Permissions []PolicyRequest `json:"permissions"`

	Grants map[string][]string `json:"grants"`

	Blocklist *string `json:"blocklist,omitempty"`

	BlocklistBehavior *string `json:"blocklist_behavior,omitempty"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons,omitempty"`

	Blocklists *[]BlockListOptions `json:"blocklists,omitempty"`

	AutomodThresholds *Thresholds `json:"automod_thresholds,omitempty"`
}

// Represents custom chat command
type CreateCommandRequest struct {
	Description string `json:"description"`

	Name string `json:"name"`

	Args *string `json:"args,omitempty"`

	Set *string `json:"set,omitempty"`
}

type CreateCommandResponse struct {
	Duration string `json:"duration"`

	Command *Command `json:"command,omitempty"`
}

type CreateDeviceRequest struct {
	ID string `json:"id"`

	PushProvider string `json:"push_provider"`

	PushProviderName *string `json:"push_provider_name,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	VoipToken *bool `json:"voip_token,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type CreateExternalStorageRequest struct {
	Bucket string `json:"bucket"`

	Name string `json:"name"`

	StorageType string `json:"storage_type"`

	GcsCredentials *string `json:"gcs_credentials,omitempty"`

	Path *string `json:"path,omitempty"`

	AwsS3 *S3Request `json:"aws_s3,omitempty"`

	AzureBlob *AzureRequest `json:"azure_blob,omitempty"`
}

type CreateExternalStorageResponse struct {
	Duration string `json:"duration"`
}

type CreateGuestRequest struct {
	User UserRequest `json:"user"`
}

type CreateGuestResponse struct {
	AccessToken string `json:"access_token"`

	Duration string `json:"duration"`

	User UserResponse `json:"user"`
}

type CreateImportRequest struct {
	Mode string `json:"mode"`

	Path string `json:"path"`
}

type CreateImportResponse struct {
	Duration string `json:"duration"`

	ImportTask *ImportTask `json:"import_task,omitempty"`
}

type CreateImportURLRequest struct {
	Filename *string `json:"filename,omitempty"`
}

type CreateImportURLResponse struct {
	Duration string `json:"duration"`

	Path string `json:"path"`

	UploadUrl string `json:"upload_url"`
}

type CreatePollOptionRequest struct {
	Text string `json:"text"`

	Position *int `json:"position,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Custom *map[string]any `json:"Custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

// Contains all information needed to create a new poll
type CreatePollRequest struct {
	Name string `json:"name"`

	AllowAnswers *bool `json:"allow_answers,omitempty"`

	AllowUserSuggestedOptions *bool `json:"allow_user_suggested_options,omitempty"`

	Description *string `json:"description,omitempty"`

	EnforceUniqueVote *bool `json:"enforce_unique_vote,omitempty"`

	ID *string `json:"id,omitempty"`

	IsClosed *bool `json:"is_closed,omitempty"`

	MaxVotesAllowed *int `json:"max_votes_allowed,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	VotingVisibility *string `json:"voting_visibility,omitempty"`

	Options *[]*PollOptionInput `json:"options,omitempty"`

	Custom *map[string]any `json:"Custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type CreateRoleRequest struct {
	Name string `json:"name"`
}

type CreateRoleResponse struct {
	Duration string `json:"duration"`

	Role Role `json:"role"`
}

type DataDogInfo struct {
	ApiKey *string `json:"api_key,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	Site *string `json:"site,omitempty"`
}

type DeactivateUserRequest struct {
	CreatedByID *string `json:"created_by_id,omitempty"`

	MarkMessagesDeleted *bool `json:"mark_messages_deleted,omitempty"`
}

type DeactivateUserResponse struct {
	Duration string `json:"duration"`

	User *UserObject `json:"user,omitempty"`
}

type DeactivateUsersRequest struct {
	UserIDs []string `json:"user_ids"`

	CreatedByID *string `json:"created_by_id,omitempty"`

	MarkChannelsDeleted *bool `json:"mark_channels_deleted,omitempty"`

	MarkMessagesDeleted *bool `json:"mark_messages_deleted,omitempty"`
}

type DeactivateUsersResponse struct {
	Duration string `json:"duration"`

	TaskID string `json:"task_id"`
}

type DeleteCallRequest struct {
	Hard *bool `json:"hard,omitempty"`
}

type DeleteCallResponse struct {
	Duration string `json:"duration"`

	Call CallResponse `json:"call"`

	TaskID *string `json:"task_id,omitempty"`
}

type DeleteChannelResponse struct {
	Duration string `json:"duration"`

	Channel *ChannelResponse `json:"channel,omitempty"`
}

type DeleteChannelsRequest struct {
	Cids []string `json:"cids"`

	HardDelete *bool `json:"hard_delete,omitempty"`
}

type DeleteChannelsResponse struct {
	Duration string `json:"duration"`

	TaskID *string `json:"task_id,omitempty"`

	Result *map[string]*DeleteChannelsResult `json:"result,omitempty"`
}

type DeleteChannelsResult struct {
	Status string `json:"status"`

	Error *string `json:"error,omitempty"`
}

type DeleteCommandResponse struct {
	Duration string `json:"duration"`

	Name string `json:"name"`
}

type DeleteExternalStorageResponse struct {
	Duration string `json:"duration"`
}

type DeleteMessageResponse struct {
	Duration string `json:"duration"`

	Message MessageResponse `json:"message"`
}

type DeleteRecordingResponse struct {
	Duration string `json:"duration"`
}

type DeleteTranscriptionResponse struct {
	Duration string `json:"duration"`
}

type DeleteUsersRequest struct {
	UserIDs []string `json:"user_ids"`

	Calls *string `json:"calls,omitempty"`

	Conversations *string `json:"conversations,omitempty"`

	Messages *string `json:"messages,omitempty"`

	NewCallOwnerID *string `json:"new_call_owner_id,omitempty"`

	NewChannelOwnerID *string `json:"new_channel_owner_id,omitempty"`

	User *string `json:"user,omitempty"`
}

type DeleteUsersResponse struct {
	Duration string `json:"duration"`

	TaskID string `json:"task_id"`
}

type Device struct {
	CreatedAt Timestamp `json:"created_at"`

	ID string `json:"id"`

	PushProvider string `json:"push_provider"`

	UserID string `json:"user_id"`

	Disabled *bool `json:"disabled,omitempty"`

	DisabledReason *string `json:"disabled_reason,omitempty"`

	PushProviderName *string `json:"push_provider_name,omitempty"`

	Voip *bool `json:"voip,omitempty"`
}

type DeviceErrorInfo struct {
	ErrorMessage string `json:"error_message"`

	Provider string `json:"provider"`

	ProviderName string `json:"provider_name"`
}

type EdgeResponse struct {
	ContinentCode string `json:"continent_code"`

	CountryIsoCode string `json:"country_iso_code"`

	Green int `json:"green"`

	ID string `json:"id"`

	LatencyTestUrl string `json:"latency_test_url"`

	Latitude float64 `json:"latitude"`

	Longitude float64 `json:"longitude"`

	Red int `json:"red"`

	SubdivisionIsoCode string `json:"subdivision_iso_code"`

	Yellow int `json:"yellow"`
}

type EgressHLSResponse struct {
	PlaylistUrl string `json:"playlist_url"`
}

type EgressRTMPResponse struct {
	Name string `json:"name"`

	StreamKey string `json:"stream_key"`

	Url string `json:"url"`
}

type EgressResponse struct {
	Broadcasting bool `json:"broadcasting"`

	Rtmps []EgressRTMPResponse `json:"rtmps"`

	Hls *EgressHLSResponse `json:"hls,omitempty"`
}

type EndCallRequest struct{}

type EndCallResponse struct {
	Duration string `json:"duration"`
}

type ErrorResult struct {
	Type string `json:"type"`

	Stacktrace *string `json:"stacktrace,omitempty"`

	Version *string `json:"version,omitempty"`
}

type EventNotificationSettings struct {
	Enabled bool `json:"enabled"`

	Apns APNS `json:"apns"`
}

type EventRequest struct {
	Type string `json:"type"`

	ParentID *string `json:"parent_id,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type EventResponse struct {
	Duration string `json:"duration"`

	Event WSEvent `json:"event"`
}

type ExportChannelsRequest struct {
	Channels []ChannelExport `json:"channels"`

	ClearDeletedMessageText *bool `json:"clear_deleted_message_text,omitempty"`

	ExportUsers *bool `json:"export_users,omitempty"`

	IncludeSoftDeletedChannels *bool `json:"include_soft_deleted_channels,omitempty"`

	IncludeTruncatedMessages *bool `json:"include_truncated_messages,omitempty"`

	Version *string `json:"version,omitempty"`
}

type ExportChannelsResponse struct {
	Duration string `json:"duration"`

	TaskID string `json:"task_id"`
}

type ExportChannelsResult struct {
	Url string `json:"url"`

	Path *string `json:"path,omitempty"`

	S3BucketName *string `json:"s3_bucket_name,omitempty"`
}

type ExportUserResponse struct {
	Duration string `json:"duration"`

	Messages *[]*Message `json:"messages,omitempty"`

	Reactions *[]*Reaction `json:"reactions,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type ExportUsersRequest struct {
	UserIDs []string `json:"user_ids"`
}

type ExportUsersResponse struct {
	Duration string `json:"duration"`

	TaskID string `json:"task_id"`
}

type ExternalStorageResponse struct {
	Bucket string `json:"bucket"`

	Name string `json:"name"`

	Path string `json:"path"`

	Type string `json:"type"`
}

type Field struct {
	Short bool `json:"short"`

	Title string `json:"title"`

	Value string `json:"value"`
}

type FileDeleteResponse struct {
	Duration string `json:"duration"`
}

type FileUploadConfig struct {
	SizeLimit int `json:"size_limit"`

	AllowedFileExtensions *[]string `json:"allowed_file_extensions,omitempty"`

	AllowedMimeTypes *[]string `json:"allowed_mime_types,omitempty"`

	BlockedFileExtensions *[]string `json:"blocked_file_extensions,omitempty"`

	BlockedMimeTypes *[]string `json:"blocked_mime_types,omitempty"`
}

type FileUploadRequest struct {
	File *string `json:"file,omitempty"`

	User *OnlyUserID `json:"user,omitempty"`
}

type FileUploadResponse struct {
	Duration string `json:"duration"`

	File *string `json:"file,omitempty"`

	ThumbUrl *string `json:"thumb_url,omitempty"`
}

type FirebaseConfig struct {
	ApnTemplate *string `json:"apn_template,omitempty"`

	CredentialsJson *string `json:"credentials_json,omitempty"`

	DataTemplate *string `json:"data_template,omitempty"`

	Disabled *bool `json:"Disabled,omitempty"`

	NotificationTemplate *string `json:"notification_template,omitempty"`

	ServerKey *string `json:"server_key,omitempty"`
}

type FirebaseConfigFields struct {
	ApnTemplate string `json:"apn_template"`

	DataTemplate string `json:"data_template"`

	Enabled bool `json:"enabled"`

	NotificationTemplate string `json:"notification_template"`

	CredentialsJson *string `json:"credentials_json,omitempty"`

	ServerKey *string `json:"server_key,omitempty"`
}

// Contains information about flagged user or message
type Flag struct {
	CreatedAt Timestamp `json:"created_at"`

	CreatedByAutomod bool `json:"created_by_automod"`

	UpdatedAt Timestamp `json:"updated_at"`

	ApprovedAt *Timestamp `json:"approved_at,omitempty"`

	Reason *string `json:"reason,omitempty"`

	RejectedAt *Timestamp `json:"rejected_at,omitempty"`

	ReviewedAt *Timestamp `json:"reviewed_at,omitempty"`

	ReviewedBy *string `json:"reviewed_by,omitempty"`

	TargetMessageID *string `json:"target_message_id,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	Details *FlagDetails `json:"details,omitempty"`

	TargetMessage *Message `json:"target_message,omitempty"`

	TargetUser *UserObject `json:"target_user,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type FlagDetails struct {
	OriginalText string `json:"original_text"`

	Extra map[string]any `json:"Extra"`

	Automod *AutomodDetails `json:"automod,omitempty"`
}

type FlagFeedback struct {
	CreatedAt Timestamp `json:"created_at"`

	MessageID string `json:"message_id"`

	Labels []Label `json:"labels"`
}

type FlagMessageDetails struct {
	PinChanged *bool `json:"pin_changed,omitempty"`

	ShouldEnrich *bool `json:"should_enrich,omitempty"`

	SkipPush *bool `json:"skip_push,omitempty"`

	UpdatedByID *string `json:"updated_by_id,omitempty"`
}

type FlagRequest struct {
	Reason *string `json:"reason,omitempty"`

	TargetMessageID *string `json:"target_message_id,omitempty"`

	TargetUserID *string `json:"target_user_id,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type FlagResponse struct {
	Duration string `json:"duration"`

	Flag *Flag `json:"flag,omitempty"`
}

type FullUserResponse struct {
	Banned bool `json:"banned"`

	CreatedAt Timestamp `json:"created_at"`

	ID string `json:"id"`

	Invisible bool `json:"invisible"`

	Language string `json:"language"`

	Online bool `json:"online"`

	Role string `json:"role"`

	ShadowBanned bool `json:"shadow_banned"`

	TotalUnreadCount int `json:"total_unread_count"`

	UnreadChannels int `json:"unread_channels"`

	UnreadThreads int `json:"unread_threads"`

	UpdatedAt Timestamp `json:"updated_at"`

	BlockedUserIDs []string `json:"blocked_user_ids"`

	ChannelMutes []*ChannelMute `json:"channel_mutes"`

	Devices []*Device `json:"devices"`

	Mutes []*UserMute `json:"mutes"`

	Teams []string `json:"teams"`

	Custom map[string]any `json:"custom"`

	DeactivatedAt *Timestamp `json:"deactivated_at,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	Image *string `json:"image,omitempty"`

	LastActive *Timestamp `json:"last_active,omitempty"`

	Name *string `json:"name,omitempty"`

	RevokeTokensIssuedBefore *Timestamp `json:"revoke_tokens_issued_before,omitempty"`

	LatestHiddenChannels *[]string `json:"latest_hidden_channels,omitempty"`

	PrivacySettings *PrivacySettings `json:"privacy_settings,omitempty"`

	PushNotifications *PushNotificationSettings `json:"push_notifications,omitempty"`
}

type GeofenceResponse struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Type *string `json:"type,omitempty"`

	CountryCodes *[]string `json:"country_codes,omitempty"`
}

type GeofenceSettings struct {
	Names []string `json:"names"`
}

type GeofenceSettingsRequest struct {
	Names *[]string `json:"names,omitempty"`
}

type GeofenceSettingsResponse struct {
	Names []string `json:"names"`
}

type GeolocationResult struct {
	AccuracyRadius int `json:"accuracy_radius"`

	City string `json:"city"`

	Continent string `json:"continent"`

	ContinentCode string `json:"continent_code"`

	Country string `json:"country"`

	CountryIsoCode string `json:"country_iso_code"`

	Latitude float64 `json:"latitude"`

	Longitude float64 `json:"longitude"`

	Subdivision string `json:"subdivision"`

	SubdivisionIsoCode string `json:"subdivision_iso_code"`
}

type GetApplicationResponse struct {
	Duration string `json:"duration"`

	App AppResponseFields `json:"app"`
}

type GetBlockListResponse struct {
	Duration string `json:"duration"`

	Blocklist *BlockList `json:"blocklist,omitempty"`
}

type GetBlockedUsersResponse struct {
	Duration string `json:"duration"`

	Blocks []*BlockedUserResponse `json:"blocks"`
}

type GetCallResponse struct {
	Duration string `json:"duration"`

	Members []MemberResponse `json:"members"`

	OwnCapabilities []OwnCapability `json:"own_capabilities"`

	Call CallResponse `json:"call"`
}

type GetCallStatsResponse struct {
	CallDurationSeconds int `json:"call_duration_seconds"`

	CallStatus string `json:"call_status"`

	Duration string `json:"duration"`

	MaxFreezesDurationSeconds int `json:"max_freezes_duration_seconds"`

	MaxParticipants int `json:"max_participants"`

	MaxTotalQualityLimitationDurationSeconds int `json:"max_total_quality_limitation_duration_seconds"`

	PublishingParticipants int `json:"publishing_participants"`

	QualityScore int `json:"quality_score"`

	SfuCount int `json:"sfu_count"`

	ParticipantReport []*UserStats `json:"participant_report"`

	Sfus []SFULocationResponse `json:"sfus"`

	CallTimeline *CallTimeline `json:"call_timeline,omitempty"`

	Jitter *Stats `json:"jitter,omitempty"`

	Latency *Stats `json:"latency,omitempty"`
}

type GetCallTypeResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	UpdatedAt Timestamp `json:"updated_at"`

	Grants map[string][]string `json:"grants"`

	NotificationSettings NotificationSettings `json:"notification_settings"`

	Settings CallSettingsResponse `json:"settings"`

	ExternalStorage *string `json:"external_storage,omitempty"`
}

type GetCommandResponse struct {
	Args string `json:"args"`

	Description string `json:"description"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	Set string `json:"set"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
}

type GetCustomPermissionResponse struct {
	Duration string `json:"duration"`

	Permission Permission `json:"permission"`
}

type GetEdgesResponse struct {
	Duration string `json:"duration"`

	Edges []EdgeResponse `json:"edges"`
}

type GetExportChannelsStatusResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	Duration string `json:"duration"`

	Status string `json:"status"`

	TaskID string `json:"task_id"`

	UpdatedAt Timestamp `json:"updated_at"`

	Error *ErrorResult `json:"error,omitempty"`

	Result *ExportChannelsResult `json:"result,omitempty"`
}

type GetImportResponse struct {
	Duration string `json:"duration"`

	ImportTask *ImportTask `json:"import_task,omitempty"`
}

type GetManyMessagesResponse struct {
	Duration string `json:"duration"`

	Messages []*Message `json:"messages"`
}

type GetMessageResponse struct {
	Duration string `json:"duration"`

	Message MessageWithChannelResponse `json:"message"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata,omitempty"`
}

type GetOGResponse struct {
	Duration string `json:"duration"`

	Custom map[string]any `json:"custom"`

	AssetUrl *string `json:"asset_url,omitempty"`

	AuthorIcon *string `json:"author_icon,omitempty"`

	AuthorLink *string `json:"author_link,omitempty"`

	AuthorName *string `json:"author_name,omitempty"`

	Color *string `json:"color,omitempty"`

	Fallback *string `json:"fallback,omitempty"`

	Footer *string `json:"footer,omitempty"`

	FooterIcon *string `json:"footer_icon,omitempty"`

	ImageUrl *string `json:"image_url,omitempty"`

	OgScrapeUrl *string `json:"og_scrape_url,omitempty"`

	OriginalHeight *int `json:"original_height,omitempty"`

	OriginalWidth *int `json:"original_width,omitempty"`

	Pretext *string `json:"pretext,omitempty"`

	Text *string `json:"text,omitempty"`

	ThumbUrl *string `json:"thumb_url,omitempty"`

	Title *string `json:"title,omitempty"`

	TitleLink *string `json:"title_link,omitempty"`

	Type *string `json:"type,omitempty"`

	Actions *[]*Action `json:"actions,omitempty"`

	Fields *[]*Field `json:"fields,omitempty"`

	Giphy *Images `json:"giphy,omitempty"`
}

type GetOrCreateCallRequest struct {
	MembersLimit *int `json:"members_limit,omitempty"`

	Notify *bool `json:"notify,omitempty"`

	Ring *bool `json:"ring,omitempty"`

	Data *CallRequest `json:"data,omitempty"`
}

type GetOrCreateCallResponse struct {
	Created bool `json:"created"`

	Duration string `json:"duration"`

	Members []MemberResponse `json:"members"`

	OwnCapabilities []OwnCapability `json:"own_capabilities"`

	Call CallResponse `json:"call"`
}

type GetRateLimitsResponse struct {
	Duration string `json:"duration"`

	Android *map[string]LimitInfo `json:"android,omitempty"`

	Ios *map[string]LimitInfo `json:"ios,omitempty"`

	ServerSide *map[string]LimitInfo `json:"server_side,omitempty"`

	Web *map[string]LimitInfo `json:"web,omitempty"`
}

type GetReactionsResponse struct {
	Duration string `json:"duration"`

	Reactions []*Reaction `json:"reactions"`
}

type GetRepliesResponse struct {
	Duration string `json:"duration"`

	Messages []MessageResponse `json:"messages"`
}

type GetTaskResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	Duration string `json:"duration"`

	Status string `json:"status"`

	TaskID string `json:"task_id"`

	UpdatedAt Timestamp `json:"updated_at"`

	Error *ErrorResult `json:"error,omitempty"`

	Result *map[string]any `json:"result,omitempty"`
}

type GetThreadResponse struct {
	Duration string `json:"duration"`

	Thread ThreadStateResponse `json:"thread"`
}

type GoLiveRequest struct {
	RecordingStorageName *string `json:"recording_storage_name,omitempty"`

	StartHls *bool `json:"start_hls,omitempty"`

	StartRecording *bool `json:"start_recording,omitempty"`

	StartTranscription *bool `json:"start_transcription,omitempty"`

	TranscriptionStorageName *string `json:"transcription_storage_name,omitempty"`
}

type GoLiveResponse struct {
	Duration string `json:"duration"`

	Call CallResponse `json:"call"`
}

type HLSSettings struct {
	AutoOn bool `json:"auto_on"`

	Enabled bool `json:"enabled"`

	QualityTracks []string `json:"quality_tracks"`

	Layout *LayoutSettings `json:"layout,omitempty"`
}

type HLSSettingsRequest struct {
	QualityTracks []string `json:"quality_tracks"`

	AutoOn *bool `json:"auto_on,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	Layout *LayoutSettingsRequest `json:"layout,omitempty"`
}

type HLSSettingsResponse struct {
	AutoOn bool `json:"auto_on"`

	Enabled bool `json:"enabled"`

	QualityTracks []string `json:"quality_tracks"`

	Layout LayoutSettingsResponse `json:"layout"`
}

type HideChannelRequest struct {
	ClearHistory *bool `json:"clear_history,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type HideChannelResponse struct {
	Duration string `json:"duration"`
}

type HuaweiConfig struct {
	Disabled *bool `json:"Disabled,omitempty"`

	ID *string `json:"id,omitempty"`

	Secret *string `json:"secret,omitempty"`
}

type HuaweiConfigFields struct {
	Enabled bool `json:"enabled"`

	ID *string `json:"id,omitempty"`

	Secret *string `json:"secret,omitempty"`
}

type ImageData struct {
	Frames string `json:"frames"`

	Height string `json:"height"`

	Size string `json:"size"`

	Url string `json:"url"`

	Width string `json:"width"`
}

type ImageSize struct {
	Crop *string `json:"crop,omitempty"`

	Height *int `json:"height,omitempty"`

	Resize *string `json:"resize,omitempty"`

	Width *int `json:"width,omitempty"`
}

type ImageUploadRequest struct {
	File *string `json:"file,omitempty"`

	UploadSizes *[]ImageSize `json:"upload_sizes,omitempty"`

	User *OnlyUserID `json:"user,omitempty"`
}

type ImageUploadResponse struct {
	Duration string `json:"duration"`

	File *string `json:"file,omitempty"`

	ThumbUrl *string `json:"thumb_url,omitempty"`

	UploadSizes *[]ImageSize `json:"upload_sizes,omitempty"`
}

type Images struct {
	FixedHeight ImageData `json:"fixed_height"`

	FixedHeightDownsampled ImageData `json:"fixed_height_downsampled"`

	FixedHeightStill ImageData `json:"fixed_height_still"`

	FixedWidth ImageData `json:"fixed_width"`

	FixedWidthDownsampled ImageData `json:"fixed_width_downsampled"`

	FixedWidthStill ImageData `json:"fixed_width_still"`

	Original ImageData `json:"original"`
}

type ImportTask struct {
	CreatedAt Timestamp `json:"created_at"`

	ID string `json:"id"`

	Mode string `json:"mode"`

	Path string `json:"path"`

	State string `json:"state"`

	UpdatedAt Timestamp `json:"updated_at"`

	History []*ImportTaskHistory `json:"history"`

	Size *int `json:"size,omitempty"`
}

type ImportTaskHistory struct {
	CreatedAt Timestamp `json:"created_at"`

	NextState string `json:"next_state"`

	PrevState string `json:"prev_state"`
}

type Label struct {
	Name string `json:"name"`

	HarmLabels *[]string `json:"harm_labels,omitempty"`

	PhraseListIDs *[]int `json:"phrase_list_ids,omitempty"`
}

type LabelThresholds struct {
	Block *float64 `json:"block,omitempty"`

	Flag *float64 `json:"flag,omitempty"`
}

type LayoutSettings struct {
	ExternalAppUrl string `json:"external_app_url"`

	ExternalCssUrl string `json:"external_css_url"`

	Name string `json:"name"`

	Options *map[string]any `json:"options,omitempty"`
}

type LayoutSettingsRequest struct {
	Name string `json:"name"`

	ExternalAppUrl *string `json:"external_app_url,omitempty"`

	ExternalCssUrl *string `json:"external_css_url,omitempty"`

	Options *map[string]any `json:"options,omitempty"`
}

type LayoutSettingsResponse struct {
	ExternalAppUrl string `json:"external_app_url"`

	ExternalCssUrl string `json:"external_css_url"`

	Name string `json:"name"`

	Options *map[string]any `json:"options,omitempty"`
}

type LimitInfo struct {
	Limit int `json:"limit"`

	Remaining int `json:"remaining"`

	Reset int `json:"reset"`
}

type LimitsSettings struct {
	MaxDurationSeconds *int `json:"max_duration_seconds,omitempty"`

	MaxParticipants *int `json:"max_participants,omitempty"`
}

type LimitsSettingsRequest struct {
	MaxDurationSeconds *int `json:"max_duration_seconds,omitempty"`

	MaxParticipants *int `json:"max_participants,omitempty"`
}

type LimitsSettingsResponse struct {
	MaxDurationSeconds *int `json:"max_duration_seconds,omitempty"`

	MaxParticipants *int `json:"max_participants,omitempty"`
}

type ListBlockListResponse struct {
	Duration string `json:"duration"`

	Blocklists []*BlockList `json:"blocklists"`
}

type ListCallTypeResponse struct {
	Duration string `json:"duration"`

	CallTypes map[string]CallTypeResponse `json:"call_types"`
}

type ListChannelTypesResponse struct {
	Duration string `json:"duration"`

	ChannelTypes map[string]*ChannelTypeConfig `json:"channel_types"`
}

type ListCommandsResponse struct {
	Duration string `json:"duration"`

	Commands []*Command `json:"commands"`
}

type ListDevicesResponse struct {
	Duration string `json:"duration"`

	Devices []*Device `json:"devices"`
}

type ListExternalStorageResponse struct {
	Duration string `json:"duration"`

	ExternalStorages map[string]ExternalStorageResponse `json:"external_storages"`
}

type ListImportsResponse struct {
	Duration string `json:"duration"`

	ImportTasks []ImportTask `json:"import_tasks"`
}

type ListPermissionsResponse struct {
	Duration string `json:"duration"`

	Permissions []Permission `json:"permissions"`
}

type ListPushProvidersResponse struct {
	Duration string `json:"duration"`

	PushProviders []PushProviderResponse `json:"push_providers"`
}

type ListRecordingsResponse struct {
	Duration string `json:"duration"`

	Recordings []CallRecording `json:"recordings"`
}

type ListRolesResponse struct {
	Duration string `json:"duration"`

	Roles []Role `json:"roles"`
}

type ListTranscriptionsResponse struct {
	Duration string `json:"duration"`

	Transcriptions []CallTranscription `json:"transcriptions"`
}

type Location struct {
	ContinentCode string `json:"continent_code"`

	CountryIsoCode string `json:"country_iso_code"`

	SubdivisionIsoCode string `json:"subdivision_iso_code"`
}

type MOSStats struct {
	AverageScore float64 `json:"average_score"`

	MaxScore float64 `json:"max_score"`

	MinScore float64 `json:"min_score"`

	HistogramDurationSeconds []float64 `json:"histogram_duration_seconds"`
}

type MarkChannelsReadRequest struct {
	UserID *string `json:"user_id,omitempty"`

	ReadByChannel *map[string]string `json:"read_by_channel,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MarkReadRequest struct {
	MessageID *string `json:"message_id,omitempty"`

	ThreadID *string `json:"thread_id,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MarkReadResponse struct {
	Duration string `json:"duration"`

	Event *MessageReadEvent `json:"event,omitempty"`
}

type MarkUnreadRequest struct {
	MessageID *string `json:"message_id,omitempty"`

	ThreadID *string `json:"thread_id,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MediaPubSubHint struct {
	AudioPublished bool `json:"audio_published"`

	AudioSubscribed bool `json:"audio_subscribed"`

	VideoPublished bool `json:"video_published"`

	VideoSubscribed bool `json:"video_subscribed"`
}

type MemberRequest struct {
	UserID string `json:"user_id"`

	Role *string `json:"role,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`
}

type MemberResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	UpdatedAt Timestamp `json:"updated_at"`

	UserID string `json:"user_id"`

	Custom map[string]any `json:"custom"`

	User UserResponse `json:"user"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	Role *string `json:"role,omitempty"`
}

type MembersResponse struct {
	Duration string `json:"duration"`

	Members []*ChannelMember `json:"members"`
}

// Represents any chat message
type Message struct {
	Cid string `json:"cid"`

	CreatedAt Timestamp `json:"created_at"`

	DeletedReplyCount int `json:"deleted_reply_count"`

	Html string `json:"html"`

	ID string `json:"id"`

	Pinned bool `json:"pinned"`

	ReplyCount int `json:"reply_count"`

	Shadowed bool `json:"shadowed"`

	Silent bool `json:"silent"`

	Text string `json:"text"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	Attachments []*Attachment `json:"attachments"`

	LatestReactions []*Reaction `json:"latest_reactions"`

	MentionedUsers []UserObject `json:"mentioned_users"`

	OwnReactions []*Reaction `json:"own_reactions"`

	Custom map[string]any `json:"custom"`

	ReactionCounts map[string]int `json:"reaction_counts"`

	ReactionGroups map[string]*ReactionGroupResponse `json:"reaction_groups"`

	ReactionScores map[string]int `json:"reaction_scores"`

	BeforeMessageSendFailed *bool `json:"before_message_send_failed,omitempty"`

	Command *string `json:"command,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	MessageTextUpdatedAt *Timestamp `json:"message_text_updated_at,omitempty"`

	Mml *string `json:"mml,omitempty"`

	ParentID *string `json:"parent_id,omitempty"`

	PinExpires *Timestamp `json:"pin_expires,omitempty"`

	PinnedAt *Timestamp `json:"pinned_at,omitempty"`

	PollID *string `json:"poll_id,omitempty"`

	QuotedMessageID *string `json:"quoted_message_id,omitempty"`

	ShowInChannel *bool `json:"show_in_channel,omitempty"`

	ThreadParticipants *[]UserObject `json:"thread_participants,omitempty"`

	I18n *map[string]string `json:"i18n,omitempty"`

	ImageLabels *map[string][]string `json:"image_labels,omitempty"`

	PinnedBy *UserObject `json:"pinned_by,omitempty"`

	Poll *Poll `json:"poll,omitempty"`

	QuotedMessage *Message `json:"quoted_message,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type MessageActionRequest struct {
	FormData map[string]string `json:"form_data"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MessageChangeSet struct {
	Attachments bool `json:"attachments"`

	Custom bool `json:"custom"`

	Html bool `json:"html"`

	MentionedUserIDs bool `json:"mentioned_user_ids"`

	Mml bool `json:"mml"`

	Pin bool `json:"pin"`

	QuotedMessageID bool `json:"quoted_message_id"`

	Silent bool `json:"silent"`

	Text bool `json:"text"`
}

type MessageFlag struct {
	CreatedAt Timestamp `json:"created_at"`

	CreatedByAutomod bool `json:"created_by_automod"`

	UpdatedAt Timestamp `json:"updated_at"`

	ApprovedAt *Timestamp `json:"approved_at,omitempty"`

	Reason *string `json:"reason,omitempty"`

	RejectedAt *Timestamp `json:"rejected_at,omitempty"`

	ReviewedAt *Timestamp `json:"reviewed_at,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	Details *FlagDetails `json:"details,omitempty"`

	Message *Message `json:"message,omitempty"`

	ModerationFeedback *FlagFeedback `json:"moderation_feedback,omitempty"`

	ModerationResult *MessageModerationResult `json:"moderation_result,omitempty"`

	ReviewedBy *UserObject `json:"reviewed_by,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type MessageHistoryEntry struct {
	MessageID string `json:"message_id"`

	MessageUpdatedAt Timestamp `json:"message_updated_at"`

	MessageUpdatedByID string `json:"message_updated_by_id"`

	Text string `json:"text"`

	Attachments []*Attachment `json:"attachments"`

	Custom map[string]any `json:"Custom"`
}

type MessageModerationResult struct {
	Action string `json:"action"`

	CreatedAt Timestamp `json:"created_at"`

	MessageID string `json:"message_id"`

	UpdatedAt Timestamp `json:"updated_at"`

	UserBadKarma bool `json:"user_bad_karma"`

	UserKarma float64 `json:"user_karma"`

	BlockedWord *string `json:"blocked_word,omitempty"`

	BlocklistName *string `json:"blocklist_name,omitempty"`

	ModeratedBy *string `json:"moderated_by,omitempty"`

	AiModerationResponse *ModerationResponse `json:"ai_moderation_response,omitempty"`

	ModerationThresholds *Thresholds `json:"moderation_thresholds,omitempty"`
}

type MessagePaginationParams struct{}

type MessageReadEvent struct {
	ChannelID string `json:"channel_id"`

	ChannelType string `json:"channel_type"`

	Cid string `json:"cid"`

	CreatedAt Timestamp `json:"created_at"`

	Type string `json:"type"`

	LastReadMessageID *string `json:"last_read_message_id,omitempty"`

	Team *string `json:"team,omitempty"`

	Thread *Thread `json:"thread,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type MessageRequest struct {
	Html *string `json:"html,omitempty"`

	ID *string `json:"id,omitempty"`

	Mml *string `json:"mml,omitempty"`

	ParentID *string `json:"parent_id,omitempty"`

	PinExpires *Timestamp `json:"pin_expires,omitempty"`

	Pinned *bool `json:"pinned,omitempty"`

	PinnedAt *Timestamp `json:"pinned_at,omitempty"`

	PollID *string `json:"poll_id,omitempty"`

	QuotedMessageID *string `json:"quoted_message_id,omitempty"`

	ShowInChannel *bool `json:"show_in_channel,omitempty"`

	Silent *bool `json:"silent,omitempty"`

	Text *string `json:"text,omitempty"`

	Type *string `json:"type,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Attachments *[]*Attachment `json:"attachments,omitempty"`

	MentionedUsers *[]string `json:"mentioned_users,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MessageResponse struct {
	Cid string `json:"cid"`

	CreatedAt Timestamp `json:"created_at"`

	DeletedReplyCount int `json:"deleted_reply_count"`

	Html string `json:"html"`

	ID string `json:"id"`

	Pinned bool `json:"pinned"`

	ReplyCount int `json:"reply_count"`

	Shadowed bool `json:"shadowed"`

	Silent bool `json:"silent"`

	Text string `json:"text"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	Attachments []*Attachment `json:"attachments"`

	LatestReactions []ReactionResponse `json:"latest_reactions"`

	MentionedUsers []UserResponse `json:"mentioned_users"`

	OwnReactions []ReactionResponse `json:"own_reactions"`

	Custom map[string]any `json:"custom"`

	ReactionCounts map[string]int `json:"reaction_counts"`

	ReactionScores map[string]int `json:"reaction_scores"`

	User UserResponse `json:"user"`

	Command *string `json:"command,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	MessageTextUpdatedAt *Timestamp `json:"message_text_updated_at,omitempty"`

	Mml *string `json:"mml,omitempty"`

	ParentID *string `json:"parent_id,omitempty"`

	PinExpires *Timestamp `json:"pin_expires,omitempty"`

	PinnedAt *Timestamp `json:"pinned_at,omitempty"`

	PollID *string `json:"poll_id,omitempty"`

	QuotedMessageID *string `json:"quoted_message_id,omitempty"`

	ShowInChannel *bool `json:"show_in_channel,omitempty"`

	ThreadParticipants *[]UserResponse `json:"thread_participants,omitempty"`

	I18n *map[string]string `json:"i18n,omitempty"`

	ImageLabels *map[string][]string `json:"image_labels,omitempty"`

	PinnedBy *UserResponse `json:"pinned_by,omitempty"`

	Poll *Poll `json:"poll,omitempty"`

	QuotedMessage *Message `json:"quoted_message,omitempty"`

	ReactionGroups *map[string]*ReactionGroupResponse `json:"reaction_groups,omitempty"`
}

type MessageUpdate struct {
	OldText *string `json:"old_text,omitempty"`

	ChangeSet *MessageChangeSet `json:"change_set,omitempty"`
}

type MessageWithChannelResponse struct {
	Cid string `json:"cid"`

	CreatedAt Timestamp `json:"created_at"`

	DeletedReplyCount int `json:"deleted_reply_count"`

	Html string `json:"html"`

	ID string `json:"id"`

	Pinned bool `json:"pinned"`

	ReplyCount int `json:"reply_count"`

	Shadowed bool `json:"shadowed"`

	Silent bool `json:"silent"`

	Text string `json:"text"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	Attachments []*Attachment `json:"attachments"`

	LatestReactions []ReactionResponse `json:"latest_reactions"`

	MentionedUsers []UserResponse `json:"mentioned_users"`

	OwnReactions []ReactionResponse `json:"own_reactions"`

	Channel ChannelResponse `json:"channel"`

	Custom map[string]any `json:"custom"`

	ReactionCounts map[string]int `json:"reaction_counts"`

	ReactionScores map[string]int `json:"reaction_scores"`

	User UserResponse `json:"user"`

	Command *string `json:"command,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	MessageTextUpdatedAt *Timestamp `json:"message_text_updated_at,omitempty"`

	Mml *string `json:"mml,omitempty"`

	ParentID *string `json:"parent_id,omitempty"`

	PinExpires *Timestamp `json:"pin_expires,omitempty"`

	PinnedAt *Timestamp `json:"pinned_at,omitempty"`

	PollID *string `json:"poll_id,omitempty"`

	QuotedMessageID *string `json:"quoted_message_id,omitempty"`

	ShowInChannel *bool `json:"show_in_channel,omitempty"`

	ThreadParticipants *[]UserResponse `json:"thread_participants,omitempty"`

	I18n *map[string]string `json:"i18n,omitempty"`

	ImageLabels *map[string][]string `json:"image_labels,omitempty"`

	PinnedBy *UserResponse `json:"pinned_by,omitempty"`

	Poll *Poll `json:"poll,omitempty"`

	QuotedMessage *Message `json:"quoted_message,omitempty"`

	ReactionGroups *map[string]*ReactionGroupResponse `json:"reaction_groups,omitempty"`
}

type ModerationResponse struct {
	Action string `json:"action"`

	Explicit float64 `json:"explicit"`

	Spam float64 `json:"spam"`

	Toxic float64 `json:"toxic"`
}

type MuteChannelRequest struct {
	Expiration *int `json:"expiration,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	ChannelCids *[]string `json:"channel_cids,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MuteChannelResponse struct {
	Duration string `json:"duration"`

	ChannelMutes *[]*ChannelMute `json:"channel_mutes,omitempty"`

	ChannelMute *ChannelMute `json:"channel_mute,omitempty"`

	OwnUser *OwnUser `json:"own_user,omitempty"`
}

type MuteUserRequest struct {
	Timeout int `json:"timeout"`

	UserID *string `json:"user_id,omitempty"`

	TargetIDs *[]string `json:"target_ids,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MuteUserResponse struct {
	Duration string `json:"duration"`

	Mutes *[]*UserMute `json:"mutes,omitempty"`

	NonExistingUsers *[]string `json:"non_existing_users,omitempty"`

	Mute *UserMute `json:"mute,omitempty"`

	OwnUser *OwnUser `json:"own_user,omitempty"`
}

type MuteUsersRequest struct {
	Audio *bool `json:"audio,omitempty"`

	MuteAllUsers *bool `json:"mute_all_users,omitempty"`

	MutedByID *string `json:"muted_by_id,omitempty"`

	Screenshare *bool `json:"screenshare,omitempty"`

	ScreenshareAudio *bool `json:"screenshare_audio,omitempty"`

	Video *bool `json:"video,omitempty"`

	UserIDs *[]string `json:"user_ids,omitempty"`

	MutedBy *UserRequest `json:"muted_by,omitempty"`
}

type MuteUsersResponse struct {
	Duration string `json:"duration"`
}

type NoiseCancellationSettings struct {
	Mode string `json:"mode"`
}

type NotificationSettings struct {
	Enabled bool `json:"enabled"`

	CallLiveStarted EventNotificationSettings `json:"call_live_started"`

	CallMissed EventNotificationSettings `json:"call_missed"`

	CallNotification EventNotificationSettings `json:"call_notification"`

	CallRing EventNotificationSettings `json:"call_ring"`

	SessionStarted EventNotificationSettings `json:"session_started"`
}

type NullBool struct {
	HasValue *bool `json:"HasValue,omitempty"`

	Value *bool `json:"Value,omitempty"`
}

type NullTime struct {
	HasValue *bool `json:"HasValue,omitempty"`

	Value *Timestamp `json:"Value,omitempty"`
}

type OnlyUserID struct {
	ID string `json:"id"`
}

type OwnCapability string

const (
	BLOCK_USERS               OwnCapability = "block-users"
	CHANGE_MAX_DURATION       OwnCapability = "change-max-duration"
	CREATE_CALL               OwnCapability = "create-call"
	CREATE_REACTION           OwnCapability = "create-reaction"
	ENABLE_NOISE_CANCELLATION OwnCapability = "enable-noise-cancellation"
	END_CALL                  OwnCapability = "end-call"
	JOIN_BACKSTAGE            OwnCapability = "join-backstage"
	JOIN_CALL                 OwnCapability = "join-call"
	JOIN_ENDED_CALL           OwnCapability = "join-ended-call"
	MUTE_USERS                OwnCapability = "mute-users"
	PIN_FOR_EVERYONE          OwnCapability = "pin-for-everyone"
	READ_CALL                 OwnCapability = "read-call"
	REMOVE_CALL_MEMBER        OwnCapability = "remove-call-member"
	SCREENSHARE               OwnCapability = "screenshare"
	SEND_AUDIO                OwnCapability = "send-audio"
	SEND_VIDEO                OwnCapability = "send-video"
	START_BROADCAST_CALL      OwnCapability = "start-broadcast-call"
	START_RECORD_CALL         OwnCapability = "start-record-call"
	START_TRANSCRIPTION_CALL  OwnCapability = "start-transcription-call"
	STOP_BROADCAST_CALL       OwnCapability = "stop-broadcast-call"
	STOP_RECORD_CALL          OwnCapability = "stop-record-call"
	STOP_TRANSCRIPTION_CALL   OwnCapability = "stop-transcription-call"
	UPDATE_CALL               OwnCapability = "update-call"
	UPDATE_CALL_MEMBER        OwnCapability = "update-call-member"
	UPDATE_CALL_PERMISSIONS   OwnCapability = "update-call-permissions"
	UPDATE_CALL_SETTINGS      OwnCapability = "update-call-settings"
)

func (c OwnCapability) String() string {
	return string(c)
}

type OwnUser struct {
	Banned bool `json:"banned"`

	CreatedAt Timestamp `json:"created_at"`

	ID string `json:"id"`

	Language string `json:"language"`

	Online bool `json:"online"`

	Role string `json:"role"`

	TotalUnreadCount int `json:"total_unread_count"`

	UnreadChannels int `json:"unread_channels"`

	UnreadCount int `json:"unread_count"`

	UnreadThreads int `json:"unread_threads"`

	UpdatedAt Timestamp `json:"updated_at"`

	ChannelMutes []*ChannelMute `json:"channel_mutes"`

	Devices []*Device `json:"devices"`

	Mutes []*UserMute `json:"mutes"`

	Custom map[string]any `json:"custom"`

	DeactivatedAt *Timestamp `json:"deactivated_at,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	Invisible *bool `json:"invisible,omitempty"`

	LastActive *Timestamp `json:"last_active,omitempty"`

	BlockedUserIDs *[]string `json:"blocked_user_ids,omitempty"`

	LatestHiddenChannels *[]string `json:"latest_hidden_channels,omitempty"`

	Teams *[]string `json:"teams,omitempty"`

	PrivacySettings *PrivacySettings `json:"privacy_settings,omitempty"`

	PushNotifications *PushNotificationSettings `json:"push_notifications,omitempty"`
}

type PaginationParams struct {
	Limit *int `json:"limit,omitempty"`

	Offset *int `json:"offset,omitempty"`
}

type PendingMessage struct {
	Channel *Channel `json:"channel,omitempty"`

	Message *Message `json:"message,omitempty"`

	Metadata *map[string]string `json:"metadata,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type Permission struct {
	Action string `json:"action"`

	Custom bool `json:"custom"`

	Description string `json:"description"`

	ID string `json:"id"`

	Level string `json:"level"`

	Name string `json:"name"`

	Owner bool `json:"owner"`

	SameTeam bool `json:"same_team"`

	Tags []string `json:"tags"`

	Condition *map[string]any `json:"condition,omitempty"`
}

type PinRequest struct {
	SessionID string `json:"session_id"`

	UserID string `json:"user_id"`
}

type PinResponse struct {
	Duration string `json:"duration"`
}

type Policy struct {
	Action int `json:"action"`

	CreatedAt Timestamp `json:"created_at"`

	Name string `json:"name"`

	Owner bool `json:"owner"`

	Priority int `json:"priority"`

	UpdatedAt Timestamp `json:"updated_at"`

	Resources []string `json:"resources"`

	Roles []string `json:"roles"`
}

type PolicyRequest struct {
	Action string `json:"action"`

	Name string `json:"name"`

	Owner bool `json:"owner"`

	Priority int `json:"priority"`

	Resources []string `json:"resources"`

	Roles []string `json:"roles"`
}

type Poll struct {
	AllowAnswers bool `json:"allow_answers"`

	AllowUserSuggestedOptions bool `json:"allow_user_suggested_options"`

	AnswersCount int `json:"answers_count"`

	CreatedAt Timestamp `json:"created_at"`

	CreatedByID string `json:"created_by_id"`

	Description string `json:"description"`

	EnforceUniqueVote bool `json:"enforce_unique_vote"`

	ID string `json:"id"`

	Name string `json:"name"`

	UpdatedAt Timestamp `json:"updated_at"`

	VoteCount int `json:"vote_count"`

	LatestAnswers []*PollVote `json:"latest_answers"`

	Options []*PollOption `json:"options"`

	OwnVotes []*PollVote `json:"own_votes"`

	Custom map[string]any `json:"Custom"`

	LatestVotesByOption map[string][]*PollVote `json:"latest_votes_by_option"`

	VoteCountsByOption map[string]int `json:"vote_counts_by_option"`

	IsClosed *bool `json:"is_closed,omitempty"`

	MaxVotesAllowed *int `json:"max_votes_allowed,omitempty"`

	VotingVisibility *string `json:"voting_visibility,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`
}

type PollOption struct {
	ID string `json:"id"`

	Text string `json:"text"`

	Custom map[string]any `json:"custom"`
}

type PollOptionInput struct {
	Text *string `json:"text,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`
}

type PollOptionResponse struct {
	Duration string `json:"duration"`

	PollOption PollOptionResponseData `json:"poll_option"`
}

type PollOptionResponseData struct {
	ID string `json:"id"`

	Text string `json:"text"`

	Custom map[string]any `json:"custom"`
}

type PollResponse struct {
	Duration string `json:"duration"`

	Poll PollResponseData `json:"poll"`
}

type PollResponseData struct {
	AllowAnswers bool `json:"allow_answers"`

	AllowUserSuggestedOptions bool `json:"allow_user_suggested_options"`

	AnswersCount int `json:"answers_count"`

	CreatedAt Timestamp `json:"created_at"`

	CreatedByID string `json:"created_by_id"`

	Description string `json:"description"`

	EnforceUniqueVote bool `json:"enforce_unique_vote"`

	ID string `json:"id"`

	Name string `json:"name"`

	UpdatedAt Timestamp `json:"updated_at"`

	VoteCount int `json:"vote_count"`

	VotingVisibility string `json:"voting_visibility"`

	Options []*PollOptionResponseData `json:"options"`

	OwnVotes []*PollVoteResponseData `json:"own_votes"`

	Custom map[string]any `json:"Custom"`

	LatestVotesByOption map[string][]*PollVoteResponseData `json:"latest_votes_by_option"`

	VoteCountsByOption map[string]int `json:"vote_counts_by_option"`

	IsClosed *bool `json:"is_closed,omitempty"`

	MaxVotesAllowed *int `json:"max_votes_allowed,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`
}

type PollVote struct {
	CreatedAt Timestamp `json:"created_at"`

	ID string `json:"id"`

	OptionID string `json:"option_id"`

	PollID string `json:"poll_id"`

	UpdatedAt Timestamp `json:"updated_at"`

	AnswerText *string `json:"answer_text,omitempty"`

	IsAnswer *bool `json:"is_answer,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type PollVoteResponse struct {
	Duration string `json:"duration"`

	Vote *PollVoteResponseData `json:"vote,omitempty"`
}

type PollVoteResponseData struct {
	CreatedAt Timestamp `json:"created_at"`

	ID string `json:"id"`

	OptionID string `json:"option_id"`

	PollID string `json:"poll_id"`

	UpdatedAt Timestamp `json:"updated_at"`

	AnswerText *string `json:"answer_text,omitempty"`

	IsAnswer *bool `json:"is_answer,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type PollVotesResponse struct {
	Duration string `json:"duration"`

	Votes []*PollVoteResponseData `json:"votes"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`
}

type PrivacySettings struct {
	ReadReceipts *ReadReceipts `json:"read_receipts,omitempty"`

	TypingIndicators *TypingIndicators `json:"typing_indicators,omitempty"`
}

type PublishedTrackInfo struct {
	CodecMimeType *string `json:"codec_mime_type,omitempty"`

	DurationSeconds *int `json:"duration_seconds,omitempty"`

	TrackType *string `json:"track_type,omitempty"`
}

type PushConfig struct {
	Version string `json:"version"`

	OfflineOnly *bool `json:"offline_only,omitempty"`
}

type PushNotificationFields struct {
	OfflineOnly bool `json:"offline_only"`

	Version string `json:"version"`

	Apn APNConfigFields `json:"apn"`

	Firebase FirebaseConfigFields `json:"firebase"`

	Huawei HuaweiConfigFields `json:"huawei"`

	Xiaomi XiaomiConfigFields `json:"xiaomi"`

	Providers *[]*PushProvider `json:"providers,omitempty"`
}

type PushNotificationSettings struct {
	Disabled *bool `json:"disabled,omitempty"`

	DisabledUntil *Timestamp `json:"disabled_until,omitempty"`
}

type PushNotificationSettingsInput struct {
	Disabled *NullBool `json:"disabled,omitempty"`

	DisabledUntil *NullTime `json:"disabled_until,omitempty"`
}

type PushProvider struct {
	CreatedAt Timestamp `json:"created_at"`

	Name string `json:"name"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	ApnAuthKey *string `json:"apn_auth_key,omitempty"`

	ApnAuthType *string `json:"apn_auth_type,omitempty"`

	ApnDevelopment *bool `json:"apn_development,omitempty"`

	ApnHost *string `json:"apn_host,omitempty"`

	ApnKeyID *string `json:"apn_key_id,omitempty"`

	ApnNotificationTemplate *string `json:"apn_notification_template,omitempty"`

	ApnP12Cert *string `json:"apn_p12_cert,omitempty"`

	ApnTeamID *string `json:"apn_team_id,omitempty"`

	ApnTopic *string `json:"apn_topic,omitempty"`

	Description *string `json:"description,omitempty"`

	DisabledAt *Timestamp `json:"disabled_at,omitempty"`

	DisabledReason *string `json:"disabled_reason,omitempty"`

	FirebaseApnTemplate *string `json:"firebase_apn_template,omitempty"`

	FirebaseCredentials *string `json:"firebase_credentials,omitempty"`

	FirebaseDataTemplate *string `json:"firebase_data_template,omitempty"`

	FirebaseHost *string `json:"firebase_host,omitempty"`

	FirebaseNotificationTemplate *string `json:"firebase_notification_template,omitempty"`

	FirebaseServerKey *string `json:"firebase_server_key,omitempty"`

	HuaweiAppID *string `json:"huawei_app_id,omitempty"`

	HuaweiAppSecret *string `json:"huawei_app_secret,omitempty"`

	XiaomiAppSecret *string `json:"xiaomi_app_secret,omitempty"`

	XiaomiPackageName *string `json:"xiaomi_package_name,omitempty"`
}

type PushProviderResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	Name string `json:"name"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	ApnAuthKey *string `json:"apn_auth_key,omitempty"`

	ApnAuthType *string `json:"apn_auth_type,omitempty"`

	ApnDevelopment *bool `json:"apn_development,omitempty"`

	ApnHost *string `json:"apn_host,omitempty"`

	ApnKeyID *string `json:"apn_key_id,omitempty"`

	ApnP12Cert *string `json:"apn_p12_cert,omitempty"`

	ApnSandboxCertificate *bool `json:"apn_sandbox_certificate,omitempty"`

	ApnSupportsRemoteNotifications *bool `json:"apn_supports_remote_notifications,omitempty"`

	ApnSupportsVoipNotifications *bool `json:"apn_supports_voip_notifications,omitempty"`

	ApnTeamID *string `json:"apn_team_id,omitempty"`

	ApnTopic *string `json:"apn_topic,omitempty"`

	Description *string `json:"description,omitempty"`

	DisabledAt *Timestamp `json:"disabled_at,omitempty"`

	DisabledReason *string `json:"disabled_reason,omitempty"`

	FirebaseApnTemplate *string `json:"firebase_apn_template,omitempty"`

	FirebaseCredentials *string `json:"firebase_credentials,omitempty"`

	FirebaseDataTemplate *string `json:"firebase_data_template,omitempty"`

	FirebaseHost *string `json:"firebase_host,omitempty"`

	FirebaseNotificationTemplate *string `json:"firebase_notification_template,omitempty"`

	FirebaseServerKey *string `json:"firebase_server_key,omitempty"`

	HuaweiAppID *string `json:"huawei_app_id,omitempty"`

	HuaweiAppSecret *string `json:"huawei_app_secret,omitempty"`

	XiaomiAppSecret *string `json:"xiaomi_app_secret,omitempty"`

	XiaomiPackageName *string `json:"xiaomi_package_name,omitempty"`
}

type QueryBannedUsersRequest struct {
	FilterConditions map[string]any `json:"filter_conditions"`

	ExcludeExpiredBans *bool `json:"exclude_expired_bans,omitempty"`

	Limit *int `json:"limit,omitempty"`

	Offset *int `json:"offset,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type QueryBannedUsersResponse struct {
	Duration string `json:"duration"`

	Bans []*BanResponse `json:"bans"`
}

type QueryCallMembersRequest struct {
	ID string `json:"id"`

	Type string `json:"type"`

	Limit *int `json:"limit,omitempty"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	FilterConditions *map[string]any `json:"filter_conditions,omitempty"`
}

type QueryCallMembersResponse struct {
	Duration string `json:"duration"`

	Members []MemberResponse `json:"members"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`
}

type QueryCallStatsRequest struct {
	Limit *int `json:"limit,omitempty"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	FilterConditions *map[string]any `json:"filter_conditions,omitempty"`
}

type QueryCallStatsResponse struct {
	Duration string `json:"duration"`

	Reports []CallStatsReportSummaryResponse `json:"reports"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`
}

type QueryCallsRequest struct {
	Limit *int `json:"limit,omitempty"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	FilterConditions *map[string]any `json:"filter_conditions,omitempty"`
}

type QueryCallsResponse struct {
	Duration string `json:"duration"`

	Calls []CallStateResponseFields `json:"calls"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`
}

type QueryChannelsRequest struct {
	Limit *int `json:"limit,omitempty"`

	MemberLimit *int `json:"member_limit,omitempty"`

	MessageLimit *int `json:"message_limit,omitempty"`

	Offset *int `json:"offset,omitempty"`

	State *bool `json:"state,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	FilterConditions *map[string]any `json:"filter_conditions,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type QueryChannelsResponse struct {
	Duration string `json:"duration"`

	Channels []ChannelStateResponseFields `json:"channels"`
}

type QueryMembersRequest struct {
	Type string `json:"type"`

	FilterConditions map[string]any `json:"filter_conditions"`

	ID *string `json:"id,omitempty"`

	Limit *int `json:"limit,omitempty"`

	Offset *int `json:"offset,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Members *[]*ChannelMember `json:"members,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type QueryMessageFlagsRequest struct {
	Limit *int `json:"limit,omitempty"`

	Offset *int `json:"offset,omitempty"`

	ShowDeletedMessages *bool `json:"show_deleted_messages,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	FilterConditions *map[string]any `json:"filter_conditions,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type QueryMessageFlagsResponse struct {
	Duration string `json:"duration"`

	Flags []*MessageFlag `json:"flags"`
}

type QueryMessageHistoryRequest struct {
	Filter map[string]any `json:"filter"`

	Limit *int `json:"limit,omitempty"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`
}

type QueryMessageHistoryResponse struct {
	Duration string `json:"duration"`

	MessageHistory []*MessageHistoryEntry `json:"message_history"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`
}

type QueryPollVotesRequest struct {
	Limit *int `json:"limit,omitempty"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	Filter *map[string]any `json:"filter,omitempty"`
}

type QueryPollsRequest struct {
	Limit *int `json:"limit,omitempty"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	Filter *map[string]any `json:"filter,omitempty"`
}

type QueryPollsResponse struct {
	Duration string `json:"duration"`

	Polls []PollResponseData `json:"polls"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`
}

type QueryReactionsRequest struct {
	Limit *int `json:"limit,omitempty"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	Filter *map[string]any `json:"filter,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type QueryReactionsResponse struct {
	Duration string `json:"duration"`

	Reactions []ReactionResponse `json:"reactions"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`
}

type QueryThreadsRequest struct {
	Limit *int `json:"limit,omitempty"`

	MemberLimit *int `json:"member_limit,omitempty"`

	Next *string `json:"next,omitempty"`

	ParticipantLimit *int `json:"participant_limit,omitempty"`

	Prev *string `json:"prev,omitempty"`

	ReplyLimit *int `json:"reply_limit,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type QueryThreadsResponse struct {
	Duration string `json:"duration"`

	Threads []ThreadStateResponse `json:"threads"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`
}

type QueryUsersPayload struct {
	FilterConditions map[string]any `json:"filter_conditions"`

	IncludeDeactivatedUsers *bool `json:"include_deactivated_users,omitempty"`

	Limit *int `json:"limit,omitempty"`

	Offset *int `json:"offset,omitempty"`

	Presence *bool `json:"presence,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type QueryUsersResponse struct {
	Duration string `json:"duration"`

	Users []FullUserResponse `json:"users"`
}

// RTMP input settings
type RTMPIngress struct {
	Address string `json:"address"`
}

// Represents user reaction to a message
type Reaction struct {
	CreatedAt Timestamp `json:"created_at"`

	MessageID string `json:"message_id"`

	Score int `json:"score"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	UserID *string `json:"user_id,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type ReactionGroupResponse struct {
	Count int `json:"count"`

	FirstReactionAt Timestamp `json:"first_reaction_at"`

	LastReactionAt Timestamp `json:"last_reaction_at"`

	SumScores int `json:"sum_scores"`
}

type ReactionRemovalResponse struct {
	Duration string `json:"duration"`

	Message *Message `json:"message,omitempty"`

	Reaction *Reaction `json:"reaction,omitempty"`
}

type ReactionRequest struct {
	Type string `json:"type"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	Score *int `json:"score,omitempty"`

	UpdatedAt *Timestamp `json:"updated_at,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type ReactionResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	MessageID string `json:"message_id"`

	Score int `json:"score"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	UserID string `json:"user_id"`

	Custom map[string]any `json:"custom"`

	User UserResponse `json:"user"`
}

type ReactivateUserRequest struct {
	CreatedByID *string `json:"created_by_id,omitempty"`

	Name *string `json:"name,omitempty"`

	RestoreMessages *bool `json:"restore_messages,omitempty"`
}

type ReactivateUserResponse struct {
	Duration string `json:"duration"`

	User *UserObject `json:"user,omitempty"`
}

type ReactivateUsersRequest struct {
	UserIDs []string `json:"user_ids"`

	CreatedByID *string `json:"created_by_id,omitempty"`

	RestoreChannels *bool `json:"restore_channels,omitempty"`

	RestoreMessages *bool `json:"restore_messages,omitempty"`
}

type ReactivateUsersResponse struct {
	Duration string `json:"duration"`

	TaskID string `json:"task_id"`
}

type Read struct {
	LastRead Timestamp `json:"last_read"`

	UnreadMessages int `json:"unread_messages"`

	LastReadMessageID *string `json:"last_read_message_id,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type ReadReceipts struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type ReadStateResponse struct {
	LastRead Timestamp `json:"last_read"`

	UnreadMessages int `json:"unread_messages"`

	User UserResponse `json:"user"`

	LastReadMessageID *string `json:"last_read_message_id,omitempty"`
}

type RecordSettings struct {
	AudioOnly bool `json:"audio_only"`

	Mode string `json:"mode"`

	Quality string `json:"quality"`

	Layout *LayoutSettings `json:"layout,omitempty"`
}

type RecordSettingsRequest struct {
	Mode string `json:"mode"`

	AudioOnly *bool `json:"audio_only,omitempty"`

	Quality *string `json:"quality,omitempty"`

	Layout *LayoutSettingsRequest `json:"layout,omitempty"`
}

type RecordSettingsResponse struct {
	AudioOnly bool `json:"audio_only"`

	Mode string `json:"mode"`

	Quality string `json:"quality"`

	Layout LayoutSettingsResponse `json:"layout"`
}

type Response struct {
	Duration string `json:"duration"`
}

type RestoreUsersRequest struct {
	UserIDs []string `json:"user_ids"`
}

type RingSettings struct {
	AutoCancelTimeoutMs int `json:"auto_cancel_timeout_ms"`

	IncomingCallTimeoutMs int `json:"incoming_call_timeout_ms"`

	MissedCallTimeoutMs int `json:"missed_call_timeout_ms"`
}

type RingSettingsRequest struct {
	AutoCancelTimeoutMs int `json:"auto_cancel_timeout_ms"`

	IncomingCallTimeoutMs int `json:"incoming_call_timeout_ms"`

	MissedCallTimeoutMs *int `json:"missed_call_timeout_ms,omitempty"`
}

type RingSettingsResponse struct {
	AutoCancelTimeoutMs int `json:"auto_cancel_timeout_ms"`

	IncomingCallTimeoutMs int `json:"incoming_call_timeout_ms"`

	MissedCallTimeoutMs int `json:"missed_call_timeout_ms"`
}

type Role struct {
	CreatedAt Timestamp `json:"created_at"`

	Custom bool `json:"custom"`

	Name string `json:"name"`

	UpdatedAt Timestamp `json:"updated_at"`

	Scopes []string `json:"scopes"`
}

type S3Request struct {
	S3Region string `json:"s3_region"`

	S3ApiKey *string `json:"s3_api_key,omitempty"`

	S3Secret *string `json:"s3_secret,omitempty"`
}

type SFULocationResponse struct {
	Datacenter string `json:"datacenter"`

	ID string `json:"id"`

	Coordinates Coordinates `json:"coordinates"`

	Location Location `json:"location"`
}

type ScreensharingSettings struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	Enabled bool `json:"enabled"`

	TargetResolution *TargetResolution `json:"target_resolution,omitempty"`
}

type ScreensharingSettingsRequest struct {
	AccessRequestEnabled *bool `json:"access_request_enabled,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	TargetResolution *TargetResolution `json:"target_resolution,omitempty"`
}

type ScreensharingSettingsResponse struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	Enabled bool `json:"enabled"`

	TargetResolution *TargetResolution `json:"target_resolution,omitempty"`
}

type SearchRequest struct {
	FilterConditions map[string]any `json:"filter_conditions"`

	Limit *int `json:"limit,omitempty"`

	Next *string `json:"next,omitempty"`

	Offset *int `json:"offset,omitempty"`

	Query *string `json:"query,omitempty"`

	Sort *[]*SortParam `json:"sort,omitempty"`

	MessageFilterConditions *map[string]any `json:"message_filter_conditions,omitempty"`
}

type SearchResponse struct {
	Duration string `json:"duration"`

	Results []SearchResult `json:"results"`

	Next *string `json:"next,omitempty"`

	Previous *string `json:"previous,omitempty"`

	ResultsWarning *SearchWarning `json:"results_warning,omitempty"`
}

type SearchResult struct {
	Message *SearchResultMessage `json:"message,omitempty"`
}

type SearchResultMessage struct {
	Cid string `json:"cid"`

	CreatedAt Timestamp `json:"created_at"`

	DeletedReplyCount int `json:"deleted_reply_count"`

	Html string `json:"html"`

	ID string `json:"id"`

	Pinned bool `json:"pinned"`

	ReplyCount int `json:"reply_count"`

	Shadowed bool `json:"shadowed"`

	Silent bool `json:"silent"`

	Text string `json:"text"`

	Type string `json:"type"`

	UpdatedAt Timestamp `json:"updated_at"`

	Attachments []*Attachment `json:"attachments"`

	LatestReactions []*Reaction `json:"latest_reactions"`

	MentionedUsers []UserObject `json:"mentioned_users"`

	OwnReactions []*Reaction `json:"own_reactions"`

	Custom map[string]any `json:"custom"`

	ReactionCounts map[string]int `json:"reaction_counts"`

	ReactionGroups map[string]*ReactionGroupResponse `json:"reaction_groups"`

	ReactionScores map[string]int `json:"reaction_scores"`

	BeforeMessageSendFailed *bool `json:"before_message_send_failed,omitempty"`

	Command *string `json:"command,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	MessageTextUpdatedAt *Timestamp `json:"message_text_updated_at,omitempty"`

	Mml *string `json:"mml,omitempty"`

	ParentID *string `json:"parent_id,omitempty"`

	PinExpires *Timestamp `json:"pin_expires,omitempty"`

	PinnedAt *Timestamp `json:"pinned_at,omitempty"`

	PollID *string `json:"poll_id,omitempty"`

	QuotedMessageID *string `json:"quoted_message_id,omitempty"`

	ShowInChannel *bool `json:"show_in_channel,omitempty"`

	ThreadParticipants *[]UserObject `json:"thread_participants,omitempty"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	I18n *map[string]string `json:"i18n,omitempty"`

	ImageLabels *map[string][]string `json:"image_labels,omitempty"`

	PinnedBy *UserObject `json:"pinned_by,omitempty"`

	Poll *Poll `json:"poll,omitempty"`

	QuotedMessage *Message `json:"quoted_message,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type SearchWarning struct {
	WarningCode int `json:"warning_code"`

	WarningDescription string `json:"warning_description"`

	ChannelSearchCount *int `json:"channel_search_count,omitempty"`

	ChannelSearchCids *[]string `json:"channel_search_cids,omitempty"`
}

type SendCallEventRequest struct {
	UserID *string `json:"user_id,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type SendCallEventResponse struct {
	Duration string `json:"duration"`
}

type SendEventRequest struct {
	Event EventRequest `json:"event"`
}

type SendMessageRequest struct {
	Message MessageRequest `json:"message"`

	ForceModeration *bool `json:"force_moderation,omitempty"`

	KeepChannelHidden *bool `json:"keep_channel_hidden,omitempty"`

	Pending *bool `json:"pending,omitempty"`

	SkipEnrichUrl *bool `json:"skip_enrich_url,omitempty"`

	SkipPush *bool `json:"skip_push,omitempty"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata,omitempty"`
}

type SendMessageResponse struct {
	Duration string `json:"duration"`

	Message MessageResponse `json:"message"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata,omitempty"`
}

type SendReactionRequest struct {
	Reaction ReactionRequest `json:"reaction"`

	EnforceUnique *bool `json:"enforce_unique,omitempty"`

	SkipPush *bool `json:"skip_push,omitempty"`
}

type SendReactionResponse struct {
	Duration string `json:"duration"`

	Message MessageResponse `json:"message"`

	Reaction ReactionResponse `json:"reaction"`
}

type SendUserCustomEventRequest struct {
	Event UserCustomEventRequest `json:"event"`
}

type ShowChannelRequest struct {
	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type ShowChannelResponse struct {
	Duration string `json:"duration"`
}

type SortParam struct {
	Direction *int `json:"direction,omitempty"`

	Field *string `json:"field,omitempty"`
}

type StartHLSBroadcastingRequest struct{}

type StartHLSBroadcastingResponse struct {
	Duration string `json:"duration"`

	PlaylistUrl string `json:"playlist_url"`
}

type StartRecordingRequest struct {
	RecordingExternalStorage *string `json:"recording_external_storage,omitempty"`
}

type StartRecordingResponse struct {
	Duration string `json:"duration"`
}

type StartTranscriptionRequest struct {
	TranscriptionExternalStorage *string `json:"transcription_external_storage,omitempty"`
}

type StartTranscriptionResponse struct {
	Duration string `json:"duration"`
}

type Stats struct {
	AverageSeconds float64 `json:"average_seconds"`

	MaxSeconds float64 `json:"max_seconds"`
}

type StopHLSBroadcastingRequest struct{}

type StopHLSBroadcastingResponse struct {
	Duration string `json:"duration"`
}

type StopLiveRequest struct{}

type StopLiveResponse struct {
	Duration string `json:"duration"`

	Call CallResponse `json:"call"`
}

type StopRecordingRequest struct{}

type StopRecordingResponse struct {
	Duration string `json:"duration"`
}

type StopTranscriptionRequest struct{}

type StopTranscriptionResponse struct {
	Duration string `json:"duration"`
}

type Subsession struct {
	EndedAt int `json:"ended_at"`

	JoinedAt int `json:"joined_at"`

	SfuID string `json:"sfu_id"`

	PubSubHint *MediaPubSubHint `json:"pub_sub_hint,omitempty"`
}

type TargetResolution struct {
	Bitrate int `json:"bitrate"`

	Height int `json:"height"`

	Width int `json:"width"`
}

// Represents a conversation thread linked to a specific message in a channel.
type Thread struct {
	ChannelCid string `json:"channel_cid"`

	CreatedAt Timestamp `json:"created_at"`

	ParentMessageID string `json:"parent_message_id"`

	Title string `json:"title"`

	UpdatedAt Timestamp `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	LastMessageAt *Timestamp `json:"last_message_at,omitempty"`

	ParticipantCount *int `json:"participant_count,omitempty"`

	ReplyCount *int `json:"reply_count,omitempty"`

	ThreadParticipants *[]*ThreadParticipant `json:"thread_participants,omitempty"`

	Channel *Channel `json:"channel,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	ParentMessage *Message `json:"parent_message,omitempty"`
}

// Represents a user that is participating in a thread.
type ThreadParticipant struct {
	AppPk int `json:"app_pk"`

	ChannelCid string `json:"channel_cid"`

	CreatedAt Timestamp `json:"created_at"`

	LastReadAt Timestamp `json:"last_read_at"`

	Custom map[string]any `json:"custom"`

	LastThreadMessageAt *Timestamp `json:"last_thread_message_at,omitempty"`

	LeftThreadAt *Timestamp `json:"left_thread_at,omitempty"`

	ThreadID *string `json:"thread_id,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type ThreadResponse struct {
	ChannelCid string `json:"channel_cid"`

	CreatedAt Timestamp `json:"created_at"`

	CreatedByUserID string `json:"created_by_user_id"`

	ParentMessageID string `json:"parent_message_id"`

	Title string `json:"title"`

	UpdatedAt Timestamp `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	LastMessageAt *Timestamp `json:"last_message_at,omitempty"`

	ParticipantCount *int `json:"participant_count,omitempty"`

	ReplyCount *int `json:"reply_count,omitempty"`

	ThreadParticipants *[]*ThreadParticipant `json:"thread_participants,omitempty"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	ParentMessage *Message `json:"parent_message,omitempty"`
}

// Represents a conversation thread linked to a specific message in a channel.
type ThreadState struct {
	ChannelCid string `json:"channel_cid"`

	CreatedAt Timestamp `json:"created_at"`

	ParentMessageID string `json:"parent_message_id"`

	Title string `json:"title"`

	UpdatedAt Timestamp `json:"updated_at"`

	LatestReplies []*Message `json:"latest_replies"`

	Custom map[string]any `json:"custom"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	LastMessageAt *Timestamp `json:"last_message_at,omitempty"`

	ParticipantCount *int `json:"participant_count,omitempty"`

	ReplyCount *int `json:"reply_count,omitempty"`

	Read *[]*Read `json:"read,omitempty"`

	ThreadParticipants *[]*ThreadParticipant `json:"thread_participants,omitempty"`

	Channel *Channel `json:"channel,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	ParentMessage *Message `json:"parent_message,omitempty"`
}

type ThreadStateResponse struct {
	ChannelCid string `json:"channel_cid"`

	CreatedAt Timestamp `json:"created_at"`

	CreatedByUserID string `json:"created_by_user_id"`

	ParentMessageID string `json:"parent_message_id"`

	Title string `json:"title"`

	UpdatedAt Timestamp `json:"updated_at"`

	LatestReplies []*Message `json:"latest_replies"`

	Custom map[string]any `json:"custom"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	LastMessageAt *Timestamp `json:"last_message_at,omitempty"`

	ParticipantCount *int `json:"participant_count,omitempty"`

	ReplyCount *int `json:"reply_count,omitempty"`

	Read *[]*Read `json:"read,omitempty"`

	ThreadParticipants *[]*ThreadParticipant `json:"thread_participants,omitempty"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	ParentMessage *Message `json:"parent_message,omitempty"`
}

// Sets thresholds for AI moderation
type Thresholds struct {
	Explicit *LabelThresholds `json:"explicit,omitempty"`

	Spam *LabelThresholds `json:"spam,omitempty"`

	Toxic *LabelThresholds `json:"toxic,omitempty"`
}

type ThumbnailResponse struct {
	ImageUrl string `json:"image_url"`
}

type ThumbnailsSettings struct {
	Enabled bool `json:"enabled"`
}

type ThumbnailsSettingsRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type ThumbnailsSettingsResponse struct {
	Enabled bool `json:"enabled"`
}

type TranscriptionSettings struct {
	ClosedCaptionMode string `json:"closed_caption_mode"`

	Mode string `json:"mode"`

	Languages []string `json:"languages"`
}

type TranscriptionSettingsRequest struct {
	Mode string `json:"mode"`

	ClosedCaptionMode *string `json:"closed_caption_mode,omitempty"`

	Languages *[]string `json:"languages,omitempty"`
}

type TranscriptionSettingsResponse struct {
	ClosedCaptionMode string `json:"closed_caption_mode"`

	Mode string `json:"mode"`

	Languages []string `json:"languages"`
}

type TranslateMessageRequest struct {
	Language string `json:"language"`
}

type TruncateChannelRequest struct {
	HardDelete *bool `json:"hard_delete,omitempty"`

	SkipPush *bool `json:"skip_push,omitempty"`

	TruncatedAt *Timestamp `json:"truncated_at,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Message *MessageRequest `json:"message,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type TruncateChannelResponse struct {
	Duration string `json:"duration"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	Message *Message `json:"message,omitempty"`
}

type TypingIndicators struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type UnblockUserRequest struct {
	UserID string `json:"user_id"`
}

type UnblockUserResponse struct {
	Duration string `json:"duration"`
}

type UnblockUsersRequest struct {
	BlockedUserID string `json:"blocked_user_id"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UnblockUsersResponse struct {
	Duration string `json:"duration"`
}

type UnmuteChannelRequest struct {
	Expiration *int `json:"expiration,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	ChannelCids *[]string `json:"channel_cids,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UnmuteResponse struct {
	Duration string `json:"duration"`

	NonExistingUsers *[]string `json:"non_existing_users,omitempty"`
}

type UnmuteUserRequest struct {
	Timeout int `json:"timeout"`

	UserID *string `json:"user_id,omitempty"`

	TargetIDs *[]string `json:"target_ids,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UnpinRequest struct {
	SessionID string `json:"session_id"`

	UserID string `json:"user_id"`
}

type UnpinResponse struct {
	Duration string `json:"duration"`
}

type UnreadCountsBatchRequest struct {
	UserIDs []string `json:"user_ids"`
}

type UnreadCountsBatchResponse struct {
	Duration string `json:"duration"`

	CountsByUser map[string]*UnreadCountsResponse `json:"counts_by_user"`
}

type UnreadCountsChannel struct {
	ChannelID string `json:"channel_id"`

	LastRead Timestamp `json:"last_read"`

	UnreadCount int `json:"unread_count"`
}

type UnreadCountsChannelType struct {
	ChannelCount int `json:"channel_count"`

	ChannelType string `json:"channel_type"`

	UnreadCount int `json:"unread_count"`
}

type UnreadCountsResponse struct {
	TotalUnreadCount int `json:"total_unread_count"`

	TotalUnreadThreadsCount int `json:"total_unread_threads_count"`

	ChannelType []UnreadCountsChannelType `json:"channel_type"`

	Channels []UnreadCountsChannel `json:"channels"`

	Threads []UnreadCountsThread `json:"threads"`
}

type UnreadCountsThread struct {
	LastRead Timestamp `json:"last_read"`

	LastReadMessageID string `json:"last_read_message_id"`

	ParentMessageID string `json:"parent_message_id"`

	UnreadCount int `json:"unread_count"`
}

type UpdateAppRequest struct {
	AsyncUrlEnrichEnabled *bool `json:"async_url_enrich_enabled,omitempty"`

	AutoTranslationEnabled *bool `json:"auto_translation_enabled,omitempty"`

	BeforeMessageSendHookUrl *string `json:"before_message_send_hook_url,omitempty"`

	CdnExpirationSeconds *int `json:"cdn_expiration_seconds,omitempty"`

	ChannelHideMembersOnly *bool `json:"channel_hide_members_only,omitempty"`

	CustomActionHandlerUrl *string `json:"custom_action_handler_url,omitempty"`

	DisableAuthChecks *bool `json:"disable_auth_checks,omitempty"`

	DisablePermissionsChecks *bool `json:"disable_permissions_checks,omitempty"`

	EnforceUniqueUsernames *string `json:"enforce_unique_usernames,omitempty"`

	ImageModerationEnabled *bool `json:"image_moderation_enabled,omitempty"`

	MigratePermissionsToV2 *bool `json:"migrate_permissions_to_v2,omitempty"`

	MultiTenantEnabled *bool `json:"multi_tenant_enabled,omitempty"`

	PermissionVersion *string `json:"permission_version,omitempty"`

	RemindersInterval *int `json:"reminders_interval,omitempty"`

	RemindersMaxMembers *int `json:"reminders_max_members,omitempty"`

	RevokeTokensIssuedBefore *Timestamp `json:"revoke_tokens_issued_before,omitempty"`

	SnsKey *string `json:"sns_key,omitempty"`

	SnsSecret *string `json:"sns_secret,omitempty"`

	SnsTopicArn *string `json:"sns_topic_arn,omitempty"`

	SqsKey *string `json:"sqs_key,omitempty"`

	SqsSecret *string `json:"sqs_secret,omitempty"`

	SqsUrl *string `json:"sqs_url,omitempty"`

	VideoProvider *string `json:"video_provider,omitempty"`

	WebhookUrl *string `json:"webhook_url,omitempty"`

	ImageModerationBlockLabels *[]string `json:"image_moderation_block_labels,omitempty"`

	ImageModerationLabels *[]string `json:"image_moderation_labels,omitempty"`

	UserSearchDisallowedRoles *[]string `json:"user_search_disallowed_roles,omitempty"`

	WebhookEvents *[]string `json:"webhook_events,omitempty"`

	AgoraOptions *Config `json:"agora_options,omitempty"`

	ApnConfig *APNConfig `json:"apn_config,omitempty"`

	AsyncModerationConfig *AsyncModerationConfiguration `json:"async_moderation_config,omitempty"`

	DatadogInfo *DataDogInfo `json:"datadog_info,omitempty"`

	FileUploadConfig *FileUploadConfig `json:"file_upload_config,omitempty"`

	FirebaseConfig *FirebaseConfig `json:"firebase_config,omitempty"`

	Grants *map[string][]string `json:"grants,omitempty"`

	HmsOptions *Config `json:"hms_options,omitempty"`

	HuaweiConfig *HuaweiConfig `json:"huawei_config,omitempty"`

	ImageUploadConfig *FileUploadConfig `json:"image_upload_config,omitempty"`

	PushConfig *PushConfig `json:"push_config,omitempty"`

	XiaomiConfig *XiaomiConfig `json:"xiaomi_config,omitempty"`
}

type UpdateBlockListRequest struct {
	Words *[]string `json:"words,omitempty"`
}

type UpdateCallMembersRequest struct {
	RemoveMembers *[]string `json:"remove_members,omitempty"`

	UpdateMembers *[]MemberRequest `json:"update_members,omitempty"`
}

type UpdateCallMembersResponse struct {
	Duration string `json:"duration"`

	Members []MemberResponse `json:"members"`
}

type UpdateCallRequest struct {
	StartsAt *Timestamp `json:"starts_at,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	SettingsOverride *CallSettingsRequest `json:"settings_override,omitempty"`
}

// Represents a call
type UpdateCallResponse struct {
	Duration string `json:"duration"`

	Members []MemberResponse `json:"members"`

	OwnCapabilities []OwnCapability `json:"own_capabilities"`

	Call CallResponse `json:"call"`
}

type UpdateCallTypeRequest struct {
	ExternalStorage *string `json:"external_storage,omitempty"`

	Grants *map[string][]string `json:"grants,omitempty"`

	NotificationSettings *NotificationSettings `json:"notification_settings,omitempty"`

	Settings *CallSettingsRequest `json:"settings,omitempty"`
}

type UpdateCallTypeResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	UpdatedAt Timestamp `json:"updated_at"`

	Grants map[string][]string `json:"grants"`

	NotificationSettings NotificationSettings `json:"notification_settings"`

	Settings CallSettingsResponse `json:"settings"`

	ExternalStorage *string `json:"external_storage,omitempty"`
}

type UpdateChannelPartialRequest struct {
	UserID *string `json:"user_id,omitempty"`

	Unset *[]string `json:"unset,omitempty"`

	Set *map[string]any `json:"set,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UpdateChannelPartialResponse struct {
	Duration string `json:"duration"`

	Members []*ChannelMember `json:"members"`

	Channel *ChannelResponse `json:"channel,omitempty"`
}

type UpdateChannelRequest struct {
	AcceptInvite *bool `json:"accept_invite,omitempty"`

	Cooldown *int `json:"cooldown,omitempty"`

	HideHistory *bool `json:"hide_history,omitempty"`

	RejectInvite *bool `json:"reject_invite,omitempty"`

	SkipPush *bool `json:"skip_push,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	AddMembers *[]*ChannelMember `json:"add_members,omitempty"`

	AddModerators *[]string `json:"add_moderators,omitempty"`

	AssignRoles *[]*ChannelMember `json:"assign_roles,omitempty"`

	DemoteModerators *[]string `json:"demote_moderators,omitempty"`

	Invites *[]*ChannelMember `json:"invites,omitempty"`

	RemoveMembers *[]string `json:"remove_members,omitempty"`

	Data *ChannelInput `json:"data,omitempty"`

	Message *MessageRequest `json:"message,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UpdateChannelResponse struct {
	Duration string `json:"duration"`

	Members []*ChannelMember `json:"members"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	Message *Message `json:"message,omitempty"`
}

type UpdateChannelTypeRequest struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	MaxMessageLength int `json:"max_message_length"`

	Blocklist *string `json:"blocklist,omitempty"`

	BlocklistBehavior *string `json:"blocklist_behavior,omitempty"`

	ConnectEvents *bool `json:"connect_events,omitempty"`

	CustomEvents *bool `json:"custom_events,omitempty"`

	MarkMessagesPending *bool `json:"mark_messages_pending,omitempty"`

	Mutes *bool `json:"mutes,omitempty"`

	Polls *bool `json:"polls,omitempty"`

	PushNotifications *bool `json:"push_notifications,omitempty"`

	Quotes *bool `json:"quotes,omitempty"`

	Reactions *bool `json:"reactions,omitempty"`

	ReadEvents *bool `json:"read_events,omitempty"`

	Reminders *bool `json:"reminders,omitempty"`

	Replies *bool `json:"replies,omitempty"`

	Search *bool `json:"search,omitempty"`

	TypingEvents *bool `json:"typing_events,omitempty"`

	Uploads *bool `json:"uploads,omitempty"`

	UrlEnrichment *bool `json:"url_enrichment,omitempty"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons,omitempty"`

	Blocklists *[]BlockListOptions `json:"blocklists,omitempty"`

	Commands *[]string `json:"commands,omitempty"`

	Permissions *[]PolicyRequest `json:"permissions,omitempty"`

	AutomodThresholds *Thresholds `json:"automod_thresholds,omitempty"`

	Grants *map[string][]string `json:"grants,omitempty"`
}

type UpdateChannelTypeResponse struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt Timestamp `json:"created_at"`

	CustomEvents bool `json:"custom_events"`

	Duration string `json:"duration"`

	MarkMessagesPending bool `json:"mark_messages_pending"`

	MaxMessageLength int `json:"max_message_length"`

	Mutes bool `json:"mutes"`

	Name string `json:"name"`

	Polls bool `json:"polls"`

	PushNotifications bool `json:"push_notifications"`

	Quotes bool `json:"quotes"`

	Reactions bool `json:"reactions"`

	ReadEvents bool `json:"read_events"`

	Reminders bool `json:"reminders"`

	Replies bool `json:"replies"`

	Search bool `json:"search"`

	TypingEvents bool `json:"typing_events"`

	UpdatedAt Timestamp `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []string `json:"commands"`

	Permissions []PolicyRequest `json:"permissions"`

	Grants map[string][]string `json:"grants"`

	Blocklist *string `json:"blocklist,omitempty"`

	BlocklistBehavior *string `json:"blocklist_behavior,omitempty"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons,omitempty"`

	Blocklists *[]BlockListOptions `json:"blocklists,omitempty"`

	AutomodThresholds *Thresholds `json:"automod_thresholds,omitempty"`
}

// Represents custom chat command
type UpdateCommandRequest struct {
	Description string `json:"description"`

	Args *string `json:"args,omitempty"`

	Set *string `json:"set,omitempty"`
}

type UpdateCommandResponse struct {
	Duration string `json:"duration"`

	Command *Command `json:"command,omitempty"`
}

type UpdateExternalStorageRequest struct {
	Bucket string `json:"bucket"`

	StorageType string `json:"storage_type"`

	GcsCredentials *string `json:"gcs_credentials,omitempty"`

	Path *string `json:"path,omitempty"`

	AwsS3 *S3Request `json:"aws_s3,omitempty"`

	AzureBlob *AzureRequest `json:"azure_blob,omitempty"`
}

type UpdateExternalStorageResponse struct {
	Bucket string `json:"bucket"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	Path string `json:"path"`

	Type string `json:"type"`
}

type UpdateMessagePartialRequest struct {
	SkipEnrichUrl *bool `json:"skip_enrich_url,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Unset *[]string `json:"unset,omitempty"`

	Set *map[string]any `json:"set,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UpdateMessagePartialResponse struct {
	Duration string `json:"duration"`

	Message *Message `json:"message,omitempty"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata,omitempty"`
}

type UpdateMessageRequest struct {
	Message MessageRequest `json:"message"`

	SkipEnrichUrl *bool `json:"skip_enrich_url,omitempty"`
}

type UpdateMessageResponse struct {
	Duration string `json:"duration"`

	Message Message `json:"message"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata,omitempty"`
}

type UpdatePollOptionRequest struct {
	ID string `json:"id"`

	Text string `json:"text"`

	UserID *string `json:"user_id,omitempty"`

	Custom *map[string]any `json:"Custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UpdatePollPartialRequest struct {
	UserID *string `json:"user_id,omitempty"`

	Unset *[]string `json:"unset,omitempty"`

	Set *map[string]any `json:"set,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UpdatePollRequest struct {
	ID string `json:"id"`

	Name string `json:"name"`

	AllowAnswers *bool `json:"allow_answers,omitempty"`

	AllowUserSuggestedOptions *bool `json:"allow_user_suggested_options,omitempty"`

	Description *string `json:"description,omitempty"`

	EnforceUniqueVote *bool `json:"enforce_unique_vote,omitempty"`

	IsClosed *bool `json:"is_closed,omitempty"`

	MaxVotesAllowed *int `json:"max_votes_allowed,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	VotingVisibility *string `json:"voting_visibility,omitempty"`

	Options *[]*PollOption `json:"options,omitempty"`

	Custom *map[string]any `json:"Custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UpdateThreadPartialRequest struct {
	UserID *string `json:"user_id,omitempty"`

	Unset *[]string `json:"unset,omitempty"`

	Set *map[string]any `json:"set,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UpdateThreadPartialResponse struct {
	Duration string `json:"duration"`

	Thread ThreadResponse `json:"thread"`
}

type UpdateUserPartialRequest struct {
	ID string `json:"id"`

	Unset *[]string `json:"unset,omitempty"`

	Set *map[string]any `json:"set,omitempty"`
}

type UpdateUserPermissionsRequest struct {
	UserID string `json:"user_id"`

	GrantPermissions *[]string `json:"grant_permissions,omitempty"`

	RevokePermissions *[]string `json:"revoke_permissions,omitempty"`
}

type UpdateUserPermissionsResponse struct {
	Duration string `json:"duration"`
}

type UpdateUsersPartialRequest struct {
	Users []UpdateUserPartialRequest `json:"users"`
}

type UpdateUsersRequest struct {
	Users map[string]UserRequest `json:"users"`
}

type UpdateUsersResponse struct {
	Duration string `json:"duration"`

	MembershipDeletionTaskID string `json:"membership_deletion_task_id"`

	Users map[string]FullUserResponse `json:"users"`
}

type UpsertPushProviderRequest struct {
	PushProvider *PushProvider `json:"push_provider,omitempty"`
}

type UpsertPushProviderResponse struct {
	Duration string `json:"duration"`

	PushProvider PushProviderResponse `json:"push_provider"`
}

type UserCustomEventRequest struct {
	Type string `json:"type"`

	Custom *map[string]any `json:"custom,omitempty"`
}

type UserInfoResponse struct {
	Image string `json:"image"`

	Name string `json:"name"`

	Roles []string `json:"roles"`

	Custom map[string]any `json:"custom"`
}

type UserMute struct {
	CreatedAt Timestamp `json:"created_at"`

	UpdatedAt Timestamp `json:"updated_at"`

	Expires *Timestamp `json:"expires,omitempty"`

	Target *UserObject `json:"target,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

// Represents chat user
type UserObject struct {
	Banned bool `json:"banned"`

	ID string `json:"id"`

	Online bool `json:"online"`

	Role string `json:"role"`

	Custom map[string]any `json:"custom"`

	BanExpires *Timestamp `json:"ban_expires,omitempty"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	DeactivatedAt *Timestamp `json:"deactivated_at,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	Invisible *bool `json:"invisible,omitempty"`

	Language *string `json:"language,omitempty"`

	LastActive *Timestamp `json:"last_active,omitempty"`

	RevokeTokensIssuedBefore *Timestamp `json:"revoke_tokens_issued_before,omitempty"`

	UpdatedAt *Timestamp `json:"updated_at,omitempty"`

	Teams *[]string `json:"teams,omitempty"`

	PrivacySettings *PrivacySettings `json:"privacy_settings,omitempty"`

	PushNotifications *PushNotificationSettings `json:"push_notifications,omitempty"`
}

type UserRequest struct {
	ID string `json:"id"`

	Image *string `json:"image,omitempty"`

	Invisible *bool `json:"invisible,omitempty"`

	Language *string `json:"language,omitempty"`

	Name *string `json:"name,omitempty"`

	Role *string `json:"role,omitempty"`

	Teams *[]string `json:"teams,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	PrivacySettings *PrivacySettings `json:"privacy_settings,omitempty"`

	PushNotifications *PushNotificationSettingsInput `json:"push_notifications,omitempty"`
}

type UserResponse struct {
	Banned bool `json:"banned"`

	CreatedAt Timestamp `json:"created_at"`

	ID string `json:"id"`

	Invisible bool `json:"invisible"`

	Language string `json:"language"`

	Online bool `json:"online"`

	Role string `json:"role"`

	ShadowBanned bool `json:"shadow_banned"`

	UpdatedAt Timestamp `json:"updated_at"`

	BlockedUserIDs []string `json:"blocked_user_ids"`

	Devices []*Device `json:"devices"`

	Teams []string `json:"teams"`

	Custom map[string]any `json:"custom"`

	DeactivatedAt *Timestamp `json:"deactivated_at,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	Image *string `json:"image,omitempty"`

	LastActive *Timestamp `json:"last_active,omitempty"`

	Name *string `json:"name,omitempty"`

	RevokeTokensIssuedBefore *Timestamp `json:"revoke_tokens_issued_before,omitempty"`

	PushNotifications *PushNotificationSettings `json:"push_notifications,omitempty"`
}

type UserSessionStats struct {
	FreezeDurationSeconds int `json:"freeze_duration_seconds"`

	MaxFreezeFraction float64 `json:"max_freeze_fraction"`

	MaxFreezesDurationSeconds int `json:"max_freezes_duration_seconds"`

	PacketLossFraction float64 `json:"packet_loss_fraction"`

	PublisherPacketLossFraction float64 `json:"publisher_packet_loss_fraction"`

	PublishingDurationSeconds int `json:"publishing_duration_seconds"`

	QualityScore float64 `json:"quality_score"`

	ReceivingDurationSeconds int `json:"receiving_duration_seconds"`

	SessionID string `json:"session_id"`

	TotalPixelsIn int `json:"total_pixels_in"`

	TotalPixelsOut int `json:"total_pixels_out"`

	Browser *string `json:"browser,omitempty"`

	BrowserVersion *string `json:"browser_version,omitempty"`

	CurrentIp *string `json:"current_ip,omitempty"`

	CurrentSfu *string `json:"current_sfu,omitempty"`

	DeviceModel *string `json:"device_model,omitempty"`

	DeviceVersion *string `json:"device_version,omitempty"`

	DistanceToSfuKilometers *float64 `json:"distance_to_sfu_kilometers,omitempty"`

	MaxFirPerSecond *float64 `json:"max_fir_per_second,omitempty"`

	MaxFreezesPerSecond *float64 `json:"max_freezes_per_second,omitempty"`

	MaxNackPerSecond *float64 `json:"max_nack_per_second,omitempty"`

	MaxPliPerSecond *float64 `json:"max_pli_per_second,omitempty"`

	Os *string `json:"os,omitempty"`

	OsVersion *string `json:"os_version,omitempty"`

	PublisherNoiseCancellationSeconds *float64 `json:"publisher_noise_cancellation_seconds,omitempty"`

	PublisherQualityLimitationFraction *float64 `json:"publisher_quality_limitation_fraction,omitempty"`

	PublishingAudioCodec *string `json:"publishing_audio_codec,omitempty"`

	PublishingVideoCodec *string `json:"publishing_video_codec,omitempty"`

	ReceivingAudioCodec *string `json:"receiving_audio_codec,omitempty"`

	ReceivingVideoCodec *string `json:"receiving_video_codec,omitempty"`

	Sdk *string `json:"sdk,omitempty"`

	SdkVersion *string `json:"sdk_version,omitempty"`

	SubscriberVideoQualityThrottledDurationSeconds *float64 `json:"subscriber_video_quality_throttled_duration_seconds,omitempty"`

	WebrtcVersion *string `json:"webrtc_version,omitempty"`

	PublishedTracks *[]PublishedTrackInfo `json:"published_tracks,omitempty"`

	Subsessions *[]*Subsession `json:"subsessions,omitempty"`

	Geolocation *GeolocationResult `json:"geolocation,omitempty"`

	Jitter *Stats `json:"jitter,omitempty"`

	Latency *Stats `json:"latency,omitempty"`

	MaxPublishingVideoQuality *VideoQuality `json:"max_publishing_video_quality,omitempty"`

	MaxReceivingVideoQuality *VideoQuality `json:"max_receiving_video_quality,omitempty"`

	PubSubHints *MediaPubSubHint `json:"pub_sub_hints,omitempty"`

	PublisherAudioMos *MOSStats `json:"publisher_audio_mos,omitempty"`

	PublisherJitter *Stats `json:"publisher_jitter,omitempty"`

	PublisherLatency *Stats `json:"publisher_latency,omitempty"`

	PublisherVideoQualityLimitationDurationSeconds *map[string]float64 `json:"publisher_video_quality_limitation_duration_seconds,omitempty"`

	SubscriberAudioMos *MOSStats `json:"subscriber_audio_mos,omitempty"`

	SubscriberJitter *Stats `json:"subscriber_jitter,omitempty"`

	SubscriberLatency *Stats `json:"subscriber_latency,omitempty"`

	Timeline *CallTimeline `json:"timeline,omitempty"`
}

type UserStats struct {
	MinEventTs int `json:"min_event_ts"`

	SessionStats []UserSessionStats `json:"session_stats"`

	Info UserInfoResponse `json:"info"`

	Rating *int `json:"rating,omitempty"`
}

type VideoQuality struct {
	UsageType *string `json:"usage_type,omitempty"`

	Resolution *VideoResolution `json:"resolution,omitempty"`
}

type VideoResolution struct {
	Height int `json:"height"`

	Width int `json:"width"`
}

type VideoSettings struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	CameraDefaultOn bool `json:"camera_default_on"`

	CameraFacing string `json:"camera_facing"`

	Enabled bool `json:"enabled"`

	TargetResolution TargetResolution `json:"target_resolution"`
}

type VideoSettingsRequest struct {
	AccessRequestEnabled *bool `json:"access_request_enabled,omitempty"`

	CameraDefaultOn *bool `json:"camera_default_on,omitempty"`

	CameraFacing *string `json:"camera_facing,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	TargetResolution *TargetResolution `json:"target_resolution,omitempty"`
}

type VideoSettingsResponse struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	CameraDefaultOn bool `json:"camera_default_on"`

	CameraFacing string `json:"camera_facing"`

	Enabled bool `json:"enabled"`

	TargetResolution TargetResolution `json:"target_resolution"`
}

type VoteData struct {
	AnswerText *string `json:"answer_text,omitempty"`

	OptionID *string `json:"option_id,omitempty"`

	Option *PollOption `json:"Option,omitempty"`
}

// Represents an BaseEvent that happened in Stream Chat
type WSEvent struct {
	CreatedAt Timestamp `json:"created_at"`

	Type string `json:"type"`

	Custom map[string]any `json:"custom"`

	Automoderation *bool `json:"automoderation,omitempty"`

	ChannelID *string `json:"channel_id,omitempty"`

	ChannelType *string `json:"channel_type,omitempty"`

	Cid *string `json:"cid,omitempty"`

	ConnectionID *string `json:"connection_id,omitempty"`

	ParentID *string `json:"parent_id,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Team *string `json:"team,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	WatcherCount *int `json:"watcher_count,omitempty"`

	AutomoderationScores *ModerationResponse `json:"automoderation_scores,omitempty"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	Me *OwnUser `json:"me,omitempty"`

	Member *ChannelMember `json:"member,omitempty"`

	Message *Message `json:"message,omitempty"`

	MessageUpdate *MessageUpdate `json:"message_update,omitempty"`

	Poll *Poll `json:"poll,omitempty"`

	PollVote *PollVote `json:"poll_vote,omitempty"`

	Reaction *Reaction `json:"reaction,omitempty"`

	Thread *Thread `json:"thread,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type WrappedUnreadCountsResponse struct {
	Duration string `json:"duration"`

	TotalUnreadCount int `json:"total_unread_count"`

	TotalUnreadThreadsCount int `json:"total_unread_threads_count"`

	ChannelType []UnreadCountsChannelType `json:"channel_type"`

	Channels []UnreadCountsChannel `json:"channels"`

	Threads []UnreadCountsThread `json:"threads"`
}

type XiaomiConfig struct {
	Disabled *bool `json:"Disabled,omitempty"`

	PackageName *string `json:"package_name,omitempty"`

	Secret *string `json:"secret,omitempty"`
}

type XiaomiConfigFields struct {
	Enabled bool `json:"enabled"`

	PackageName *string `json:"package_name,omitempty"`

	Secret *string `json:"secret,omitempty"`
}
