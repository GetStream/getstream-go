package getstream

type APIError struct {
	// API error code
	Code int `json:"code"`

	// Request duration
	Duration string `json:"duration"`

	// Message describing an error
	Message string `json:"message"`

	// URL with additional information
	MoreInfo string `json:"more_info"`

	// Response HTTP status code
	StatusCode int `json:"StatusCode"`

	// Additional error-specific information
	Details []int `json:"details"`

	// Additional error info
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

	// Attachment type (e.g. image, video, url, poll)
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
	// ID of user to ban
	TargetUserID string `json:"target_user_id"`

	// User ID who issued a ban
	BannedByID *string `json:"banned_by_id,omitempty"`

	// Channel CID to ban user in eg. messaging:123
	ChannelCid *string `json:"channel_cid,omitempty"`

	// Whether to perform IP ban or not
	IpBan *bool `json:"ip_ban,omitempty"`

	// Ban reason
	Reason *string `json:"reason,omitempty"`

	// Whether to perform shadow ban or not
	Shadow *bool `json:"shadow,omitempty"`

	// Timeout of ban in minutes. User will be unbanned after this period of time
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
	// Block list name
	Name string `json:"name"`

	// Block list type.
	Type string `json:"type"`

	// List of words to block
	Words []string `json:"words"`

	// Date/time of creation
	CreatedAt *Timestamp `json:"created_at,omitempty"`

	// Date/time of the last update
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
}

type BlockListOptions struct {
	Behavior string `json:"behavior"`

	Blocklist string `json:"blocklist"`
}

type BlockUserRequest struct {
	// the user to block
	UserID string `json:"user_id"`
}

type BlockUserResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

type BlockUsersRequest struct {
	// User id to block
	BlockedUserID string `json:"blocked_user_id"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type BlockUsersResponse struct {
	BlockedByUserID string `json:"blocked_by_user_id"`

	BlockedUserID string `json:"blocked_user_id"`

	CreatedAt Timestamp `json:"created_at"`

	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

type BlockedUserResponse struct {
	// ID of the user who got blocked
	BlockedUserID string `json:"blocked_user_id"`

	CreatedAt Timestamp `json:"created_at"`

	// ID of the user who blocked another user
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

	// The unique identifier for a call (<type>:<id>)
	Cid string `json:"cid"`

	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	CurrentSessionID string `json:"current_session_id"`

	// Call ID
	ID string `json:"id"`

	Recording bool `json:"recording"`

	Transcribing bool `json:"transcribing"`

	// The type of call
	Type string `json:"type"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	BlockedUserIDs []string `json:"blocked_user_ids"`

	CreatedBy UserResponse `json:"created_by"`

	// Custom data for this object
	Custom map[string]any `json:"custom"`

	Egress EgressResponse `json:"egress"`

	Ingress CallIngressResponse `json:"ingress"`

	Settings CallSettingsResponse `json:"settings"`

	// Date/time when the call ended
	EndedAt *Timestamp `json:"ended_at,omitempty"`

	// Date/time when the call will start
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
	// List of call members
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

	// List of commands that channel supports
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

	// Channel ID
	ID *string `json:"id,omitempty"`

	// Date to export messages since
	MessagesSince *Timestamp `json:"messages_since,omitempty"`

	// Date to export messages until
	MessagesUntil *Timestamp `json:"messages_until,omitempty"`

	// Channel type
	Type *string `json:"type,omitempty"`
}

type ChannelGetOrCreateRequest struct {
	// Whether this channel will be hidden for the user who created the channel or not
	HideForCreator *bool `json:"hide_for_creator,omitempty"`

	// Refresh channel state
	State *bool `json:"state,omitempty"`

	ThreadUnreadCounts *bool `json:"thread_unread_counts,omitempty"`

	Data *ChannelInput `json:"data,omitempty"`

	Members *PaginationParams `json:"members,omitempty"`

	Messages *MessagePaginationParams `json:"messages,omitempty"`

	Watchers *PaginationParams `json:"watchers,omitempty"`
}

type ChannelInput struct {
	// Enable or disable auto translation
	AutoTranslationEnabled *bool `json:"auto_translation_enabled,omitempty"`

	// Switch auto translation language
	AutoTranslationLanguage *string `json:"auto_translation_language,omitempty"`

	CreatedByID *string `json:"created_by_id,omitempty"`

	Disabled *bool `json:"disabled,omitempty"`

	// Freeze or unfreeze the channel
	Frozen *bool `json:"frozen,omitempty"`

	// Team the channel belongs to (if multi-tenant mode is enabled)
	Team *string `json:"team,omitempty"`

	TruncatedByID *string `json:"truncated_by_id,omitempty"`

	Invites *[]*ChannelMember `json:"invites,omitempty"`

	Members *[]*ChannelMember `json:"members,omitempty"`

	ConfigOverrides *ChannelConfig `json:"config_overrides,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`
}

type ChannelMember struct {
	// Whether member is banned this channel or not
	Banned bool `json:"banned"`

	// Role of the member in the channel
	ChannelRole string `json:"channel_role"`

	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	NotificationsMuted bool `json:"notifications_muted"`

	// Whether member is shadow banned in this channel or not
	ShadowBanned bool `json:"shadow_banned"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	// Expiration date of the ban
	BanExpires *Timestamp `json:"ban_expires,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	// Date when invite was accepted
	InviteAcceptedAt *Timestamp `json:"invite_accepted_at,omitempty"`

	// Date when invite was rejected
	InviteRejectedAt *Timestamp `json:"invite_rejected_at,omitempty"`

	// Whether member was invited or not
	Invited *bool `json:"invited,omitempty"`

	// Whether member is channel moderator or not
	IsModerator *bool `json:"is_moderator,omitempty"`

	Status *string `json:"status,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type ChannelMute struct {
	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	// Date/time of mute expiration
	Expires *Timestamp `json:"expires,omitempty"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

// Represents channel in chat
type ChannelResponse struct {
	// Channel CID (<type>:<id>)
	Cid string `json:"cid"`

	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	Disabled bool `json:"disabled"`

	// Whether channel is frozen or not
	Frozen bool `json:"frozen"`

	// Channel unique ID
	ID string `json:"id"`

	// Type of the channel
	Type string `json:"type"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	// Whether auto translation is enabled or not
	AutoTranslationEnabled *bool `json:"auto_translation_enabled,omitempty"`

	// Language to translate to when auto translation is active
	AutoTranslationLanguage *string `json:"auto_translation_language,omitempty"`

	// Whether this channel is blocked by current user or not
	Blocked *bool `json:"blocked,omitempty"`

	// Cooldown period after sending each message
	Cooldown *int `json:"cooldown,omitempty"`

	// Date/time of deletion
	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	// Whether this channel is hidden by current user or not
	Hidden *bool `json:"hidden,omitempty"`

	// Date since when the message history is accessible
	HideMessagesBefore *Timestamp `json:"hide_messages_before,omitempty"`

	// Date of the last message sent
	LastMessageAt *Timestamp `json:"last_message_at,omitempty"`

	// Number of members in the channel
	MemberCount *int `json:"member_count,omitempty"`

	// Date of mute expiration
	MuteExpiresAt *Timestamp `json:"mute_expires_at,omitempty"`

	// Whether this channel is muted or not
	Muted *bool `json:"muted,omitempty"`

	// Team the channel belongs to (multi-tenant only)
	Team *string `json:"team,omitempty"`

	// Date of the latest truncation of the channel
	TruncatedAt *Timestamp `json:"truncated_at,omitempty"`

	// List of channel members (max 100)
	Members *[]*ChannelMember `json:"members,omitempty"`

	// List of channel capabilities of authenticated user
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
	// List of channel members
	Members []*ChannelMember `json:"members"`

	// List of channel messages
	Messages []MessageResponse `json:"messages"`

	// List of pinned messages in the channel
	PinnedMessages []MessageResponse `json:"pinned_messages"`

	Threads []*ThreadState `json:"threads"`

	// Whether this channel is hidden or not
	Hidden *bool `json:"hidden,omitempty"`

	// Messages before this date are hidden from the user
	HideMessagesBefore *Timestamp `json:"hide_messages_before,omitempty"`

	// Number of channel watchers
	WatcherCount *int `json:"watcher_count,omitempty"`

	// Pending messages that this user has sent
	PendingMessages *[]*PendingMessage `json:"pending_messages,omitempty"`

	// List of read states
	Read *[]ReadStateResponse `json:"read,omitempty"`

	// List of user who is watching the channel
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

	// List of commands that channel supports
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	FileUrl string `json:"file_url"`
}

type CheckPushRequest struct {
	// Push message template for APN
	ApnTemplate *string `json:"apn_template,omitempty"`

	// Push message data template for Firebase
	FirebaseDataTemplate *string `json:"firebase_data_template,omitempty"`

	// Push message template for Firebase
	FirebaseTemplate *string `json:"firebase_template,omitempty"`

	// Message ID to send push notification for
	MessageID *string `json:"message_id,omitempty"`

	// Name of push provider
	PushProviderName *string `json:"push_provider_name,omitempty"`

	// Push provider type
	PushProviderType *string `json:"push_provider_type,omitempty"`

	// Don't require existing devices to render templates
	SkipDevices *bool `json:"skip_devices,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type CheckPushResponse struct {
	Duration string `json:"duration"`

	RenderedApnTemplate *string `json:"rendered_apn_template,omitempty"`

	RenderedFirebaseTemplate *string `json:"rendered_firebase_template,omitempty"`

	// Don't require existing devices to render templates
	SkipDevices *bool `json:"skip_devices,omitempty"`

	// List of general errors
	GeneralErrors *[]string `json:"general_errors,omitempty"`

	// Object with device errors
	DeviceErrors *map[string]DeviceErrorInfo `json:"device_errors,omitempty"`

	RenderedMessage *map[string]string `json:"rendered_message,omitempty"`
}

type CheckSNSRequest struct {
	// AWS SNS access key
	SnsKey *string `json:"sns_key,omitempty"`

	// AWS SNS key secret
	SnsSecret *string `json:"sns_secret,omitempty"`

	// AWS SNS topic ARN
	SnsTopicArn *string `json:"sns_topic_arn,omitempty"`
}

type CheckSNSResponse struct {
	Duration string `json:"duration"`

	// Validation result
	Status string `json:"status"`

	// Error text
	Error *string `json:"error,omitempty"`

	// Error data
	Data *map[string]any `json:"data,omitempty"`
}

type CheckSQSRequest struct {
	// AWS SQS access key
	SqsKey *string `json:"sqs_key,omitempty"`

	// AWS SQS key secret
	SqsSecret *string `json:"sqs_secret,omitempty"`

	// AWS SQS endpoint URL
	SqsUrl *string `json:"sqs_url,omitempty"`
}

type CheckSQSResponse struct {
	Duration string `json:"duration"`

	// Validation result
	Status string `json:"status"`

	// Error text
	Error *string `json:"error,omitempty"`

	// Error data
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

// Represents custom chat command
type Command struct {
	// Arguments help text, shown in commands auto-completion
	Args string `json:"args"`

	// Description, shown in commands auto-completion
	Description string `json:"description"`

	// Unique command name
	Name string `json:"name"`

	// Set name used for grouping commands
	Set string `json:"set"`

	// Date/time of creation
	CreatedAt *Timestamp `json:"created_at,omitempty"`

	// Date/time of the last update
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
	// Block list name
	Name string `json:"name"`

	// List of words to block
	Words []string `json:"words"`

	// Block list type.
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
	// Enables automatic message moderation
	Automod string `json:"automod"`

	// Sets behavior of automatic moderation
	AutomodBehavior string `json:"automod_behavior"`

	// Number of maximum message characters
	MaxMessageLength int `json:"max_message_length"`

	// Channel type name
	Name string `json:"name"`

	// Name of the blocklist to use
	Blocklist *string `json:"blocklist,omitempty"`

	// Sets behavior of blocklist
	BlocklistBehavior *string `json:"blocklist_behavior,omitempty"`

	// Connect events support
	ConnectEvents *bool `json:"connect_events,omitempty"`

	// Enables custom events
	CustomEvents *bool `json:"custom_events,omitempty"`

	// Marks messages as pending by default
	MarkMessagesPending *bool `json:"mark_messages_pending,omitempty"`

	MessageRetention *string `json:"message_retention,omitempty"`

	// Enables mutes
	Mutes *bool `json:"mutes,omitempty"`

	// Enables polls
	Polls *bool `json:"polls,omitempty"`

	// Enables push notifications
	PushNotifications *bool `json:"push_notifications,omitempty"`

	// Enables message reactions
	Reactions *bool `json:"reactions,omitempty"`

	// Read events support
	ReadEvents *bool `json:"read_events,omitempty"`

	// Enables message replies (threads)
	Replies *bool `json:"replies,omitempty"`

	// Enables message search
	Search *bool `json:"search,omitempty"`

	// Typing events support
	TypingEvents *bool `json:"typing_events,omitempty"`

	// Enables file uploads
	Uploads *bool `json:"uploads,omitempty"`

	// Enables URL enrichment
	UrlEnrichment *bool `json:"url_enrichment,omitempty"`

	Blocklists *[]BlockListOptions `json:"blocklists,omitempty"`

	// List of commands that channel supports
	Commands *[]string `json:"commands,omitempty"`

	// List of permissions for the channel type
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
	// Description, shown in commands auto-completion
	Description string `json:"description"`

	// Unique command name
	Name string `json:"name"`

	// Arguments help text, shown in commands auto-completion
	Args *string `json:"args,omitempty"`

	// Set name used for grouping commands
	Set *string `json:"set,omitempty"`
}

type CreateCommandResponse struct {
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

type CreateGuestRequest struct {
	User UserRequest `json:"user"`
}

type CreateGuestResponse struct {
	// the access token to authenticate the user
	AccessToken string `json:"access_token"`

	Duration string `json:"duration"`

	User UserResponse `json:"user"`
}

type CreateImportRequest struct {
	Mode string `json:"mode"`

	Path string `json:"path"`
}

type CreateImportResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	ImportTask *ImportTask `json:"import_task,omitempty"`
}

type CreateImportURLRequest struct {
	Filename *string `json:"filename,omitempty"`
}

type CreateImportURLResponse struct {
	// Duration of the request in human-readable format
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
	// The name of the poll
	Name string `json:"name"`

	// Indicates whether users can suggest user defined answers
	AllowAnswers *bool `json:"allow_answers,omitempty"`

	AllowUserSuggestedOptions *bool `json:"allow_user_suggested_options,omitempty"`

	// A description of the poll
	Description *string `json:"description,omitempty"`

	// Indicates whether users can cast multiple votes
	EnforceUniqueVote *bool `json:"enforce_unique_vote,omitempty"`

	ID *string `json:"id,omitempty"`

	// Indicates whether the poll is open for voting
	IsClosed *bool `json:"is_closed,omitempty"`

	// Indicates the maximum amount of votes a user can cast
	MaxVotesAllowed *int `json:"max_votes_allowed,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	VotingVisibility *string `json:"voting_visibility,omitempty"`

	Options *[]*PollOptionInput `json:"options,omitempty"`

	Custom *map[string]any `json:"Custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type CreateRoleRequest struct {
	// Role name
	Name string `json:"name"`
}

type CreateRoleResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Role Role `json:"role"`
}

type DataDogInfo struct {
	ApiKey *string `json:"api_key,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	Site *string `json:"site,omitempty"`
}

type DeactivateUserRequest struct {
	// ID of the user who deactivated the user
	CreatedByID *string `json:"created_by_id,omitempty"`

	// Makes messages appear to be deleted
	MarkMessagesDeleted *bool `json:"mark_messages_deleted,omitempty"`
}

type DeactivateUserResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	User *UserObject `json:"user,omitempty"`
}

type DeactivateUsersRequest struct {
	// User IDs to deactivate
	UserIDs []string `json:"user_ids"`

	// ID of the user who deactivated the users
	CreatedByID *string `json:"created_by_id,omitempty"`

	MarkChannelsDeleted *bool `json:"mark_channels_deleted,omitempty"`

	// Makes messages appear to be deleted
	MarkMessagesDeleted *bool `json:"mark_messages_deleted,omitempty"`
}

type DeactivateUsersResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	TaskID string `json:"task_id"`
}

type DeleteCallRequest struct {
	// if true the call will be hard deleted along with all related data
	Hard *bool `json:"hard,omitempty"`
}

type DeleteCallResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Call CallResponse `json:"call"`

	TaskID *string `json:"task_id,omitempty"`
}

type DeleteChannelResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Channel *ChannelResponse `json:"channel,omitempty"`
}

type DeleteChannelsRequest struct {
	// All channels that should be deleted
	Cids []string `json:"cids"`

	// Specify if channels and all ressources should be hard deleted
	HardDelete *bool `json:"hard_delete,omitempty"`
}

type DeleteChannelsResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	TaskID *string `json:"task_id,omitempty"`

	Result *map[string]*DeleteChannelsResult `json:"result,omitempty"`
}

type DeleteChannelsResult struct {
	Status string `json:"status"`

	Error *string `json:"error,omitempty"`
}

type DeleteCommandResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Name string `json:"name"`
}

type DeleteExternalStorageResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

type DeleteMessageResponse struct {
	// Duration of the request in human-readable format
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
	// IDs of users to delete
	UserIDs []string `json:"user_ids"`

	// Calls delete mode.
	// Affected calls are those that include exactly two members, one of whom is the user being deleted.
	// * null or empty string - doesn't delete any calls
	// * soft - marks user's calls and their related data as deleted (soft-delete)
	// * hard - deletes user's calls and their data completely (hard-delete)
	Calls *string `json:"calls,omitempty"`

	// Conversation channels delete mode.
	// Conversation channel is any channel which only has two members one of which is the user being deleted.
	// * null or empty string - doesn't delete any conversation channels
	// * soft - marks all conversation channels as deleted (same effect as Delete Channels with 'hard' option disabled)
	// * hard - deletes channel and all its data completely including messages (same effect as Delete Channels with 'hard' option enabled)
	Conversations *string `json:"conversations,omitempty"`

	// Message delete mode.
	// * null or empty string - doesn't delete user messages
	// * soft - marks all user messages as deleted without removing any related message data
	// * pruning - marks all user messages as deleted, nullifies message information and removes some message data such as reactions and flags
	// * hard - deletes messages completely with all related information
	Messages *string `json:"messages,omitempty"`

	NewCallOwnerID *string `json:"new_call_owner_id,omitempty"`

	NewChannelOwnerID *string `json:"new_channel_owner_id,omitempty"`

	// User delete mode.
	// * soft - marks user as deleted and retains all user data
	// * pruning - marks user as deleted and nullifies user information
	// * hard - deletes user completely. Requires 'hard' option for messages and conversations as well
	User *string `json:"user,omitempty"`
}

type DeleteUsersResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	TaskID string `json:"task_id"`
}

type Device struct {
	CreatedAt Timestamp `json:"created_at"`

	// Device ID
	ID string `json:"id"`

	PushProvider string `json:"push_provider"`

	UserID string `json:"user_id"`

	Disabled *bool `json:"disabled,omitempty"`

	DisabledReason *string `json:"disabled_reason,omitempty"`

	// Name of the push provider configuration
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Event WSEvent `json:"event"`
}

type ExportChannelsRequest struct {
	// Export options for channels
	Channels []ChannelExport `json:"channels"`

	// Set if deleted message text should be cleared
	ClearDeletedMessageText *bool `json:"clear_deleted_message_text,omitempty"`

	ExportUsers *bool `json:"export_users,omitempty"`

	// Set if you want to include deleted channels
	IncludeSoftDeletedChannels *bool `json:"include_soft_deleted_channels,omitempty"`

	// Set if you want to include truncated messages
	IncludeTruncatedMessages *bool `json:"include_truncated_messages,omitempty"`

	Version *string `json:"version,omitempty"`
}

type ExportChannelsResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	TaskID string `json:"task_id"`
}

type ExportChannelsResult struct {
	// URL of result
	Url string `json:"url"`

	// S3 path of result
	Path *string `json:"path,omitempty"`

	// S3 bucket name result
	S3BucketName *string `json:"s3_bucket_name,omitempty"`
}

type ExportUserResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Messages *[]*Message `json:"messages,omitempty"`

	Reactions *[]*Reaction `json:"reactions,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type ExportUsersRequest struct {
	UserIDs []string `json:"user_ids"`
}

type ExportUsersResponse struct {
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
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
	// file field
	File *string `json:"file,omitempty"`

	User *OnlyUserID `json:"user,omitempty"`
}

type FileUploadResponse struct {
	Duration string `json:"duration"`

	// URL to the uploaded asset. Should be used to put to `asset_url` attachment field
	File *string `json:"file,omitempty"`

	// URL of the file thumbnail for supported file formats. Should be put to `thumb_url` attachment field
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
	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	CreatedByAutomod bool `json:"created_by_automod"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	// Date of the approval
	ApprovedAt *Timestamp `json:"approved_at,omitempty"`

	Reason *string `json:"reason,omitempty"`

	// Date of the rejection
	RejectedAt *Timestamp `json:"rejected_at,omitempty"`

	// Date of the review
	ReviewedAt *Timestamp `json:"reviewed_at,omitempty"`

	ReviewedBy *string `json:"reviewed_by,omitempty"`

	// ID of flagged message
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

	// ID of the message when reporting a message
	TargetMessageID *string `json:"target_message_id,omitempty"`

	// ID of the user when reporting a user
	TargetUserID *string `json:"target_user_id,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type FlagResponse struct {
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	App AppResponseFields `json:"app"`
}

type GetBlockListResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Blocklist *BlockList `json:"blocklist,omitempty"`
}

type GetBlockedUsersResponse struct {
	Duration string `json:"duration"`

	// Array of blocked user object
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

	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Permission Permission `json:"permission"`
}

type GetEdgesResponse struct {
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	ImportTask *ImportTask `json:"import_task,omitempty"`
}

type GetManyMessagesResponse struct {
	Duration string `json:"duration"`

	// List of messages
	Messages []*Message `json:"messages"`
}

type GetMessageResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Message MessageWithChannelResponse `json:"message"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata,omitempty"`
}

type GetOGResponse struct {
	Duration string `json:"duration"`

	Custom map[string]any `json:"custom"`

	// URL of detected video or audio
	AssetUrl *string `json:"asset_url,omitempty"`

	AuthorIcon *string `json:"author_icon,omitempty"`

	// og:site
	AuthorLink *string `json:"author_link,omitempty"`

	// og:site_name
	AuthorName *string `json:"author_name,omitempty"`

	Color *string `json:"color,omitempty"`

	Fallback *string `json:"fallback,omitempty"`

	Footer *string `json:"footer,omitempty"`

	FooterIcon *string `json:"footer_icon,omitempty"`

	// URL of detected image
	ImageUrl *string `json:"image_url,omitempty"`

	// extracted url from the text
	OgScrapeUrl *string `json:"og_scrape_url,omitempty"`

	OriginalHeight *int `json:"original_height,omitempty"`

	OriginalWidth *int `json:"original_width,omitempty"`

	Pretext *string `json:"pretext,omitempty"`

	// og:description
	Text *string `json:"text,omitempty"`

	// URL of detected thumb image
	ThumbUrl *string `json:"thumb_url,omitempty"`

	// og:title
	Title *string `json:"title,omitempty"`

	// og:url
	TitleLink *string `json:"title_link,omitempty"`

	// Attachment type, could be empty, image, audio or video
	Type *string `json:"type,omitempty"`

	Actions *[]*Action `json:"actions,omitempty"`

	Fields *[]*Field `json:"fields,omitempty"`

	Giphy *Images `json:"giphy,omitempty"`
}

type GetOrCreateCallRequest struct {
	MembersLimit *int `json:"members_limit,omitempty"`

	// if provided it sends a notification event to the members for this call
	Notify *bool `json:"notify,omitempty"`

	// if provided it sends a ring event to the members for this call
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

	// Map of endpoint rate limits for the Android platform
	Android *map[string]LimitInfo `json:"android,omitempty"`

	// Map of endpoint rate limits for the iOS platform
	Ios *map[string]LimitInfo `json:"ios,omitempty"`

	// Map of endpoint rate limits for the server-side platform
	ServerSide *map[string]LimitInfo `json:"server_side,omitempty"`

	// Map of endpoint rate limits for the web platform
	Web *map[string]LimitInfo `json:"web,omitempty"`
}

type GetReactionsResponse struct {
	Duration string `json:"duration"`

	// List of reactions
	Reactions []*Reaction `json:"reactions"`
}

type GetRepliesResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Messages []MessageResponse `json:"messages"`
}

type GetTaskResponse struct {
	CreatedAt Timestamp `json:"created_at"`

	Duration string `json:"duration"`

	// Current status of task
	Status string `json:"status"`

	// ID of task
	TaskID string `json:"task_id"`

	UpdatedAt Timestamp `json:"updated_at"`

	Error *ErrorResult `json:"error,omitempty"`

	// Result produced by task after completion
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
	// Duration of the request in human-readable format
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
	// Whether to clear message history of the channel or not
	ClearHistory *bool `json:"clear_history,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type HideChannelResponse struct {
	// Duration of the request in human-readable format
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
	// Crop mode
	Crop *string `json:"crop,omitempty"`

	// Target image height
	Height *int `json:"height,omitempty"`

	// Resize method
	Resize *string `json:"resize,omitempty"`

	// Target image width
	Width *int `json:"width,omitempty"`
}

type ImageUploadRequest struct {
	File *string `json:"file,omitempty"`

	// field with JSON-encoded array of image size configurations
	UploadSizes *[]ImageSize `json:"upload_sizes,omitempty"`

	User *OnlyUserID `json:"user,omitempty"`
}

type ImageUploadResponse struct {
	Duration string `json:"duration"`

	// URL to the uploaded asset. Should be used to put to `asset_url` attachment field
	File *string `json:"file,omitempty"`

	// URL of the file thumbnail for supported file formats. Should be put to `thumb_url` attachment field
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
	// Threshold for automatic message block
	Block *float64 `json:"block,omitempty"`

	// Threshold for automatic message flag
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
	// The maximum number of calls allowed for the time window
	Limit int `json:"limit"`

	// The number of remaining calls in the current window
	Remaining int `json:"remaining"`

	// The Unix timestamp of the next window
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Blocklists []*BlockList `json:"blocklists"`
}

type ListCallTypeResponse struct {
	Duration string `json:"duration"`

	CallTypes map[string]CallTypeResponse `json:"call_types"`
}

type ListChannelTypesResponse struct {
	Duration string `json:"duration"`

	// Object with all channel types
	ChannelTypes map[string]*ChannelTypeConfig `json:"channel_types"`
}

type ListCommandsResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Commands []*Command `json:"commands"`
}

type ListDevicesResponse struct {
	Duration string `json:"duration"`

	// List of devices
	Devices []*Device `json:"devices"`
}

type ListExternalStorageResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	ExternalStorages map[string]ExternalStorageResponse `json:"external_storages"`
}

type ListImportsResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	ImportTasks []ImportTask `json:"import_tasks"`
}

type ListPermissionsResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Permissions []Permission `json:"permissions"`
}

type ListPushProvidersResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	PushProviders []PushProviderResponse `json:"push_providers"`
}

type ListRecordingsResponse struct {
	Duration string `json:"duration"`

	Recordings []CallRecording `json:"recordings"`
}

type ListRolesResponse struct {
	// Duration of the request in human-readable format
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
	// ID of the message that is considered last read by client
	MessageID *string `json:"message_id,omitempty"`

	// Optional Thread ID to specifically mark a given thread as read
	ThreadID *string `json:"thread_id,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MarkReadResponse struct {
	Duration string `json:"duration"`

	Event *MessageReadEvent `json:"event,omitempty"`
}

type MarkUnreadRequest struct {
	// ID of the message from where the channel is marked unread
	MessageID *string `json:"message_id,omitempty"`

	// Mark a thread unread, specify both the thread and message id
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

	// Custom data for this object
	Custom *map[string]any `json:"custom,omitempty"`
}

type MemberResponse struct {
	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	UserID string `json:"user_id"`

	// Custom member response data
	Custom map[string]any `json:"custom"`

	User UserResponse `json:"user"`

	// Date/time of deletion
	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	Role *string `json:"role,omitempty"`
}

type MembersResponse struct {
	Duration string `json:"duration"`

	// List of found members
	Members []*ChannelMember `json:"members"`
}

// Represents any chat message
type Message struct {
	// Channel unique identifier in <type>:<id> format
	Cid string `json:"cid"`

	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	DeletedReplyCount int `json:"deleted_reply_count"`

	// Contains HTML markup of the message. Can only be set when using server-side API
	Html string `json:"html"`

	// Message ID is unique string identifier of the message
	ID string `json:"id"`

	// Whether message is pinned or not
	Pinned bool `json:"pinned"`

	// Number of replies to this message
	ReplyCount int `json:"reply_count"`

	// Whether the message was shadowed or not
	Shadowed bool `json:"shadowed"`

	// Whether message is silent or not
	Silent bool `json:"silent"`

	// Text of the message. Should be empty if `mml` is provided
	Text string `json:"text"`

	// Contains type of the message
	Type string `json:"type"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	// Array of message attachments
	Attachments []*Attachment `json:"attachments"`

	// List of 10 latest reactions to this message
	LatestReactions []*Reaction `json:"latest_reactions"`

	// List of mentioned users
	MentionedUsers []UserObject `json:"mentioned_users"`

	// List of 10 latest reactions of authenticated user to this message
	OwnReactions []*Reaction `json:"own_reactions"`

	Custom map[string]any `json:"custom"`

	// An object containing number of reactions of each type. Key: reaction type (string), value: number of reactions (int)
	ReactionCounts map[string]int `json:"reaction_counts"`

	ReactionGroups map[string]*ReactionGroupResponse `json:"reaction_groups"`

	// An object containing scores of reactions of each type. Key: reaction type (string), value: total score of reactions (int)
	ReactionScores map[string]int `json:"reaction_scores"`

	// Whether `before_message_send webhook` failed or not. Field is only accessible in push webhook
	BeforeMessageSendFailed *bool `json:"before_message_send_failed,omitempty"`

	// Contains provided slash command
	Command *string `json:"command,omitempty"`

	// Date/time of deletion
	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	MessageTextUpdatedAt *Timestamp `json:"message_text_updated_at,omitempty"`

	// Should be empty if `text` is provided. Can only be set when using server-side API
	Mml *string `json:"mml,omitempty"`

	// ID of parent message (thread)
	ParentID *string `json:"parent_id,omitempty"`

	// Date when pinned message expires
	PinExpires *Timestamp `json:"pin_expires,omitempty"`

	// Date when message got pinned
	PinnedAt *Timestamp `json:"pinned_at,omitempty"`

	// Identifier of the poll to include in the message
	PollID *string `json:"poll_id,omitempty"`

	QuotedMessageID *string `json:"quoted_message_id,omitempty"`

	// Whether thread reply should be shown in the channel as well
	ShowInChannel *bool `json:"show_in_channel,omitempty"`

	// List of users who participate in thread
	ThreadParticipants *[]UserObject `json:"thread_participants,omitempty"`

	// Object with translations. Key `language` contains the original language key. Other keys contain translations
	I18n *map[string]string `json:"i18n,omitempty"`

	// Contains image moderation information
	ImageLabels *map[string][]string `json:"image_labels,omitempty"`

	PinnedBy *UserObject `json:"pinned_by,omitempty"`

	Poll *Poll `json:"poll,omitempty"`

	QuotedMessage *Message `json:"quoted_message,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type MessageActionRequest struct {
	// ReadOnlyData to execute command with
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
	// Duration of mute in milliseconds
	Expiration *int `json:"expiration,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	// Channel CIDs to mute (if multiple channels)
	ChannelCids *[]string `json:"channel_cids,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MuteChannelResponse struct {
	Duration string `json:"duration"`

	// Object with mutes (if multiple channels were muted)
	ChannelMutes *[]*ChannelMute `json:"channel_mutes,omitempty"`

	ChannelMute *ChannelMute `json:"channel_mute,omitempty"`

	OwnUser *OwnUser `json:"own_user,omitempty"`
}

type MuteUserRequest struct {
	// Duration of mute in minutes
	Timeout int `json:"timeout"`

	UserID *string `json:"user_id,omitempty"`

	// User IDs to mute (if multiple users)
	TargetIDs *[]string `json:"target_ids,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type MuteUserResponse struct {
	Duration string `json:"duration"`

	// Object with mutes (if multiple users were muted)
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
	// Duration of the request in human-readable format
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

	// Additional data attached to the pending message. This data is discarded once the pending message is committed.
	Metadata *map[string]string `json:"metadata,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

type Permission struct {
	// Action name this permission is for (e.g. SendMessage)
	Action string `json:"action"`

	// Whether this is a custom permission or built-in
	Custom bool `json:"custom"`

	// Description of the permission
	Description string `json:"description"`

	// Unique permission ID
	ID string `json:"id"`

	// Level at which permission could be applied (app or channel)
	Level string `json:"level"`

	// Name of the permission
	Name string `json:"name"`

	// Whether this permission applies to resource owner or not
	Owner bool `json:"owner"`

	// Whether this permission applies to teammates (multi-tenancy mode only)
	SameTeam bool `json:"same_team"`

	// List of tags of the permission
	Tags []string `json:"tags"`

	// MongoDB style condition which decides whether or not the permission is granted
	Condition *map[string]any `json:"condition,omitempty"`
}

type PinRequest struct {
	SessionID string `json:"session_id"`

	UserID string `json:"user_id"`
}

type PinResponse struct {
	// Duration of the request in human-readable format
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

	// User-friendly policy name
	Name string `json:"name"`

	// Whether policy applies to resource owner or not
	Owner bool `json:"owner"`

	// Policy priority
	Priority int `json:"priority"`

	// List of resources to apply policy to
	Resources []string `json:"resources"`

	// List of roles to apply policy to
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	PollOption PollOptionResponseData `json:"poll_option"`
}

type PollOptionResponseData struct {
	ID string `json:"id"`

	Text string `json:"text"`

	Custom map[string]any `json:"custom"`
}

type PollResponse struct {
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
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
	// Number of channels to limit
	Limit *int `json:"limit,omitempty"`

	// Number of members to limit
	MemberLimit *int `json:"member_limit,omitempty"`

	// Number of messages to limit
	MessageLimit *int `json:"message_limit,omitempty"`

	// Channel pagination offset
	Offset *int `json:"offset,omitempty"`

	// Whether to update channel state or not
	State *bool `json:"state,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	// List of sort parameters
	Sort *[]*SortParam `json:"sort,omitempty"`

	FilterConditions *map[string]any `json:"filter_conditions,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type QueryChannelsResponse struct {
	Duration string `json:"duration"`

	// List of channels
	Channels []ChannelStateResponseFields `json:"channels"`
}

type QueryMembersRequest struct {
	// Channel type to interact with
	Type string `json:"type"`

	// Filter to apply to members
	FilterConditions map[string]any `json:"filter_conditions"`

	// Channel ID to interact with
	ID *string `json:"id,omitempty"`

	// Number of records to return
	Limit *int `json:"limit,omitempty"`

	// Number of records to offset
	Offset *int `json:"offset,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	// List of members to search in distinct channels
	Members *[]*ChannelMember `json:"members,omitempty"`

	// Array of sort parameters
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
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Reactions []ReactionResponse `json:"reactions"`

	Next *string `json:"next,omitempty"`

	Prev *string `json:"prev,omitempty"`
}

type QueryThreadsRequest struct {
	Limit *int `json:"limit,omitempty"`

	MemberLimit *int `json:"member_limit,omitempty"`

	Next *string `json:"next,omitempty"`

	// Limit the number of participants returned per each thread
	ParticipantLimit *int `json:"participant_limit,omitempty"`

	Prev *string `json:"prev,omitempty"`

	// Limit the number of replies returned per each thread
	ReplyLimit *int `json:"reply_limit,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type QueryThreadsResponse struct {
	Duration string `json:"duration"`

	// List of enriched thread states
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Users []FullUserResponse `json:"users"`
}

// RTMP input settings
type RTMPIngress struct {
	Address string `json:"address"`
}

// Represents user reaction to a message
type Reaction struct {
	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	// ID of a message user reacted to
	MessageID string `json:"message_id"`

	// Reaction score. If not specified reaction has score of 1
	Score int `json:"score"`

	// The type of reaction (e.g. 'like', 'laugh', 'wow')
	Type string `json:"type"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	// ID of a user who reacted to a message
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
	// Duration of the request in human-readable format
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
	// ID of the user who's reactivating the user
	CreatedByID *string `json:"created_by_id,omitempty"`

	// Set this field to put new name for the user
	Name *string `json:"name,omitempty"`

	// Restore previously deleted messages
	RestoreMessages *bool `json:"restore_messages,omitempty"`
}

type ReactivateUserResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	User *UserObject `json:"user,omitempty"`
}

type ReactivateUsersRequest struct {
	// User IDs to reactivate
	UserIDs []string `json:"user_ids"`

	// ID of the user who's reactivating the users
	CreatedByID *string `json:"created_by_id,omitempty"`

	RestoreChannels *bool `json:"restore_channels,omitempty"`

	// Restore previously deleted messages
	RestoreMessages *bool `json:"restore_messages,omitempty"`
}

type ReactivateUsersResponse struct {
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
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
	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	// Whether this is a custom role or built-in
	Custom bool `json:"custom"`

	// Unique role name
	Name string `json:"name"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	// List of scopes where this role is currently present. `.app` means that role is present in app-level grants
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
	// Channel filter conditions
	FilterConditions map[string]any `json:"filter_conditions"`

	// Number of messages to return
	Limit *int `json:"limit,omitempty"`

	// Pagination parameter. Cannot be used with non-zero offset.
	Next *string `json:"next,omitempty"`

	// Pagination offset. Cannot be used with sort or next.
	Offset *int `json:"offset,omitempty"`

	// Search phrase
	Query *string `json:"query,omitempty"`

	// Sort parameters. Cannot be used with non-zero offset
	Sort *[]*SortParam `json:"sort,omitempty"`

	// Message filter conditions
	MessageFilterConditions *map[string]any `json:"message_filter_conditions,omitempty"`
}

type SearchResponse struct {
	Duration string `json:"duration"`

	// Search results
	Results []SearchResult `json:"results"`

	// Value to pass to the next search query in order to paginate
	Next *string `json:"next,omitempty"`

	// Value that points to the previous page. Pass as the next value in a search query to paginate backwards
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
	// Code corresponding to the warning
	WarningCode int `json:"warning_code"`

	// Description of the warning
	WarningDescription string `json:"warning_description"`

	// Number of channels searched
	ChannelSearchCount *int `json:"channel_search_count,omitempty"`

	// Channel CIDs for the searched channels
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Message MessageResponse `json:"message"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata,omitempty"`
}

type SendReactionRequest struct {
	Reaction ReactionRequest `json:"reaction"`

	// Whether to replace all existing user reactions
	EnforceUnique *bool `json:"enforce_unique,omitempty"`

	// Skips any mobile push notifications
	SkipPush *bool `json:"skip_push,omitempty"`
}

type SendReactionResponse struct {
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

type SortParam struct {
	// Direction of sorting, -1 for descending, 1 for ascending
	Direction *int `json:"direction,omitempty"`

	// Name of field to sort by
	Field *string `json:"field,omitempty"`
}

type StartHLSBroadcastingRequest struct{}

type StartHLSBroadcastingResponse struct {
	// Duration of the request in human-readable format
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

type StopLiveRequest struct{}

type StopLiveResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Call CallResponse `json:"call"`
}

type StopRecordingRequest struct{}

type StopRecordingResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

type StopTranscriptionRequest struct{}

type StopTranscriptionResponse struct {
	// Duration of the request in human-readable format
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
	// Channel CID is unique string identifier of the channel
	ChannelCid string `json:"channel_cid"`

	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	// Parent Message ID is unique string identifier of the parent message
	ParentMessageID string `json:"parent_message_id"`

	// Title is the title of the thread
	Title string `json:"title"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	// Custom is the custom data of the thread
	Custom map[string]any `json:"custom"`

	// Date/time of deletion
	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	// Last Message At is the time of the last message in the thread
	LastMessageAt *Timestamp `json:"last_message_at,omitempty"`

	// The number of participants in the thread
	ParticipantCount *int `json:"participant_count,omitempty"`

	// The number of replies in the thread
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

	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	LastReadAt Timestamp `json:"last_read_at"`

	Custom map[string]any `json:"custom"`

	LastThreadMessageAt *Timestamp `json:"last_thread_message_at,omitempty"`

	// Left Thread At is the time when the user left the thread
	LeftThreadAt *Timestamp `json:"left_thread_at,omitempty"`

	// Thead ID is unique string identifier of the thread
	ThreadID *string `json:"thread_id,omitempty"`

	// User ID is unique string identifier of the user
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
	// Channel CID is unique string identifier of the channel
	ChannelCid string `json:"channel_cid"`

	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	// Parent Message ID is unique string identifier of the parent message
	ParentMessageID string `json:"parent_message_id"`

	// Title is the title of the thread
	Title string `json:"title"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	LatestReplies []*Message `json:"latest_replies"`

	// Custom is the custom data of the thread
	Custom map[string]any `json:"custom"`

	// Date/time of deletion
	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	// Last Message At is the time of the last message in the thread
	LastMessageAt *Timestamp `json:"last_message_at,omitempty"`

	// The number of participants in the thread
	ParticipantCount *int `json:"participant_count,omitempty"`

	// The number of replies in the thread
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

	// oneof=available disabled auto-on
	Mode string `json:"mode"`

	// the languages to transcribe to
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
	// Language to translate message to
	Language string `json:"language"`
}

type TruncateChannelRequest struct {
	// Permanently delete channel data (messages, reactions, etc.)
	HardDelete *bool `json:"hard_delete,omitempty"`

	// When `message` is set disables all push notifications for it
	SkipPush *bool `json:"skip_push,omitempty"`

	// Truncate channel data up to `truncated_at`. The system message (if provided) creation time is always greater than `truncated_at`
	TruncatedAt *Timestamp `json:"truncated_at,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	Message *MessageRequest `json:"message,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type TruncateChannelResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Channel *ChannelResponse `json:"channel,omitempty"`

	Message *Message `json:"message,omitempty"`
}

type TypingIndicators struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type UnblockUserRequest struct {
	// the user to unblock
	UserID string `json:"user_id"`
}

type UnblockUserResponse struct {
	// Duration of the request in human-readable format
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
	// Duration of mute in milliseconds
	Expiration *int `json:"expiration,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	// Channel CIDs to mute (if multiple channels)
	ChannelCids *[]string `json:"channel_cids,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UnmuteResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	NonExistingUsers *[]string `json:"non_existing_users,omitempty"`
}

type UnmuteUserRequest struct {
	// Duration of mute in minutes
	Timeout int `json:"timeout"`

	UserID *string `json:"user_id,omitempty"`

	// User IDs to mute (if multiple users)
	TargetIDs *[]string `json:"target_ids,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UnpinRequest struct {
	SessionID string `json:"session_id"`

	UserID string `json:"user_id"`
}

type UnpinResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

type UnreadCountsBatchRequest struct {
	UserIDs []string `json:"user_ids"`
}

type UnreadCountsBatchResponse struct {
	// Duration of the request in human-readable format
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
	// List of words to block
	Words *[]string `json:"words,omitempty"`
}

type UpdateCallMembersRequest struct {
	// List of userID to remove
	RemoveMembers *[]string `json:"remove_members,omitempty"`

	// List of members to update or insert
	UpdateMembers *[]MemberRequest `json:"update_members,omitempty"`
}

type UpdateCallMembersResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Members []MemberResponse `json:"members"`
}

type UpdateCallRequest struct {
	// the time the call is scheduled to start
	StartsAt *Timestamp `json:"starts_at,omitempty"`

	// Custom data for this object
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Members []*ChannelMember `json:"members"`

	Channel *ChannelResponse `json:"channel,omitempty"`
}

type UpdateChannelRequest struct {
	// Set to `true` to accept the invite
	AcceptInvite *bool `json:"accept_invite,omitempty"`

	// Sets cool down period for the channel in seconds
	Cooldown *int `json:"cooldown,omitempty"`

	// Set to `true` to hide channel's history when adding new members
	HideHistory *bool `json:"hide_history,omitempty"`

	// Set to `true` to reject the invite
	RejectInvite *bool `json:"reject_invite,omitempty"`

	// When `message` is set disables all push notifications for it
	SkipPush *bool `json:"skip_push,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	// List of user IDs to add to the channel
	AddMembers *[]*ChannelMember `json:"add_members,omitempty"`

	// List of user IDs to make channel moderators
	AddModerators *[]string `json:"add_moderators,omitempty"`

	// List of channel member role assignments. If any specified user is not part of the channel, the request will fail
	AssignRoles *[]*ChannelMember `json:"assign_roles,omitempty"`

	// List of user IDs to take away moderators status from
	DemoteModerators *[]string `json:"demote_moderators,omitempty"`

	// List of user IDs to invite to the channel
	Invites *[]*ChannelMember `json:"invites,omitempty"`

	// List of user IDs to remove from the channel
	RemoveMembers *[]string `json:"remove_members,omitempty"`

	Data *ChannelInput `json:"data,omitempty"`

	Message *MessageRequest `json:"message,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UpdateChannelResponse struct {
	// Duration of the request in human-readable format
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

	// List of commands that channel supports
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
	// Description, shown in commands auto-completion
	Description string `json:"description"`

	// Arguments help text, shown in commands auto-completion
	Args *string `json:"args,omitempty"`

	// Set name used for grouping commands
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

	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Name string `json:"name"`

	Path string `json:"path"`

	Type string `json:"type"`
}

type UpdateMessagePartialRequest struct {
	SkipEnrichUrl *bool `json:"skip_enrich_url,omitempty"`

	UserID *string `json:"user_id,omitempty"`

	// Array of field names to unset
	Unset *[]string `json:"unset,omitempty"`

	// Sets new field values
	Set *map[string]any `json:"set,omitempty"`

	User *UserRequest `json:"user,omitempty"`
}

type UpdateMessagePartialResponse struct {
	// Duration of the request in human-readable format
	Duration string `json:"duration"`

	Message *Message `json:"message,omitempty"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata,omitempty"`
}

type UpdateMessageRequest struct {
	Message MessageRequest `json:"message"`

	SkipEnrichUrl *bool `json:"skip_enrich_url,omitempty"`
}

type UpdateMessageResponse struct {
	// Duration of the request in human-readable format
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

	// Array of field names to unset
	Unset *[]string `json:"unset,omitempty"`

	// Sets new field values
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
	// User ID to update
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
	// Duration of the request in human-readable format
	Duration string `json:"duration"`
}

type UpdateUsersPartialRequest struct {
	Users []UpdateUserPartialRequest `json:"users"`
}

type UpdateUsersRequest struct {
	// Object containing users
	Users map[string]UserRequest `json:"users"`
}

type UpdateUsersResponse struct {
	Duration string `json:"duration"`

	MembershipDeletionTaskID string `json:"membership_deletion_task_id"`

	// Object containing users
	Users map[string]FullUserResponse `json:"users"`
}

type UpsertPushProviderRequest struct {
	PushProvider *PushProvider `json:"push_provider,omitempty"`
}

type UpsertPushProviderResponse struct {
	// Duration of the request in human-readable format
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
	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	// Date/time of mute expiration
	Expires *Timestamp `json:"expires,omitempty"`

	Target *UserObject `json:"target,omitempty"`

	User *UserObject `json:"user,omitempty"`
}

// Represents chat user
type UserObject struct {
	// Whether a user is banned or not
	Banned bool `json:"banned"`

	// Unique user identifier
	ID string `json:"id"`

	// Whether a user online or not
	Online bool `json:"online"`

	// Determines the set of user permissions
	Role string `json:"role"`

	Custom map[string]any `json:"custom"`

	// Expiration date of the ban
	BanExpires *Timestamp `json:"ban_expires,omitempty"`

	// Date/time of creation
	CreatedAt *Timestamp `json:"created_at,omitempty"`

	// Date of deactivation
	DeactivatedAt *Timestamp `json:"deactivated_at,omitempty"`

	// Date/time of deletion
	DeletedAt *Timestamp `json:"deleted_at,omitempty"`

	Invisible *bool `json:"invisible,omitempty"`

	// Preferred language of a user
	Language *string `json:"language,omitempty"`

	// Date of last activity
	LastActive *Timestamp `json:"last_active,omitempty"`

	// Revocation date for tokens
	RevokeTokensIssuedBefore *Timestamp `json:"revoke_tokens_issued_before,omitempty"`

	// Date/time of the last update
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`

	// List of teams user is a part of
	Teams *[]string `json:"teams,omitempty"`

	PrivacySettings *PrivacySettings `json:"privacy_settings,omitempty"`

	PushNotifications *PushNotificationSettings `json:"push_notifications,omitempty"`
}

type UserRequest struct {
	// User ID
	ID string `json:"id"`

	Image *string `json:"image,omitempty"`

	Invisible *bool `json:"invisible,omitempty"`

	Language *string `json:"language,omitempty"`

	// Optional name of user
	Name *string `json:"name,omitempty"`

	Role *string `json:"role,omitempty"`

	Teams *[]string `json:"teams,omitempty"`

	Custom *map[string]any `json:"custom,omitempty"`

	PrivacySettings *PrivacySettings `json:"privacy_settings,omitempty"`

	PushNotifications *PushNotificationSettingsInput `json:"push_notifications,omitempty"`
}

type UserResponse struct {
	Banned bool `json:"banned"`

	// Date/time of creation
	CreatedAt Timestamp `json:"created_at"`

	ID string `json:"id"`

	Invisible bool `json:"invisible"`

	Language string `json:"language"`

	Online bool `json:"online"`

	Role string `json:"role"`

	ShadowBanned bool `json:"shadow_banned"`

	// Date/time of the last update
	UpdatedAt Timestamp `json:"updated_at"`

	BlockedUserIDs []string `json:"blocked_user_ids"`

	Devices []*Device `json:"devices"`

	Teams []string `json:"teams"`

	Custom map[string]any `json:"custom"`

	DeactivatedAt *Timestamp `json:"deactivated_at,omitempty"`

	// Date/time of deletion
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
	// Duration of the request in human-readable format
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
