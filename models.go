package getstream

import (
	"time"
)

type APIError struct {
	Code int `json:"code"`

	Duration string `json:"duration"`

	Message string `json:"message"`

	MoreInfo string `json:"more_info"`

	StatusCode int `json:"StatusCode"`

	Details []int `json:"details"`

	ExceptionFields *map[string]string `json:"exception_fields"`
}

type APNConfig struct {
	AuthKey *string `json:"auth_key"`

	AuthType *string `json:"auth_type"`

	BundleId *string `json:"bundle_id"`

	Development *bool `json:"development"`

	Disabled *bool `json:"Disabled"`

	Host *string `json:"host"`

	KeyId *string `json:"key_id"`

	NotificationTemplate *string `json:"notification_template"`

	P12Cert *string `json:"p12_cert"`

	TeamId *string `json:"team_id"`
}

type APNConfigFields struct {
	Development bool `json:"development"`

	Enabled bool `json:"enabled"`

	NotificationTemplate string `json:"notification_template"`

	AuthKey *string `json:"auth_key"`

	AuthType *string `json:"auth_type"`

	BundleId *string `json:"bundle_id"`

	Host *string `json:"host"`

	KeyId *string `json:"key_id"`

	P12Cert *string `json:"p12_cert"`

	TeamId *string `json:"team_id"`
}

type APNS struct {
	Body string `json:"body"`

	Title string `json:"title"`
}

type Action struct {
	Name string `json:"name"`

	Text string `json:"text"`

	Type string `json:"type"`

	Style *string `json:"style"`

	Value *string `json:"value"`
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

	BeforeMessageSendHookUrl *string `json:"before_message_send_hook_url"`

	RevokeTokensIssuedBefore *time.Time `json:"revoke_tokens_issued_before"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons"`

	Geofences *[]*GeofenceResponse `json:"geofences"`

	ImageModerationLabels *[]string `json:"image_moderation_labels"`

	AgoraOptions *Config `json:"agora_options"`

	DatadogInfo *DataDogInfo `json:"datadog_info"`

	HmsOptions *Config `json:"hms_options"`
}

type AsyncModerationCallbackConfig struct {
	Mode *string `json:"mode"`

	ServerUrl *string `json:"server_url"`
}

type AsyncModerationConfiguration struct {
	TimeoutMs *int `json:"timeout_ms"`

	Callback *AsyncModerationCallbackConfig `json:"callback"`
}

type Attachment struct {
	Custom map[string]any `json:"custom"`

	AssetUrl *string `json:"asset_url"`

	AuthorIcon *string `json:"author_icon"`

	AuthorLink *string `json:"author_link"`

	AuthorName *string `json:"author_name"`

	Color *string `json:"color"`

	Fallback *string `json:"fallback"`

	Footer *string `json:"footer"`

	FooterIcon *string `json:"footer_icon"`

	ImageUrl *string `json:"image_url"`

	OgScrapeUrl *string `json:"og_scrape_url"`

	OriginalHeight *int `json:"original_height"`

	OriginalWidth *int `json:"original_width"`

	Pretext *string `json:"pretext"`

	Text *string `json:"text"`

	ThumbUrl *string `json:"thumb_url"`

	Title *string `json:"title"`

	TitleLink *string `json:"title_link"`

	Type *string `json:"type"`

	Actions *[]*Action `json:"actions"`

	Fields *[]*Field `json:"fields"`

	Giphy *Images `json:"giphy"`
}

type AudioSettings struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	DefaultDevice string `json:"default_device"`

	MicDefaultOn bool `json:"mic_default_on"`

	OpusDtxEnabled bool `json:"opus_dtx_enabled"`

	RedundantCodingEnabled bool `json:"redundant_coding_enabled"`

	SpeakerDefaultOn bool `json:"speaker_default_on"`

	NoiseCancellation *NoiseCancellationSettings `json:"noise_cancellation"`
}

type AudioSettingsRequest struct {
	DefaultDevice string `json:"default_device"`

	AccessRequestEnabled *bool `json:"access_request_enabled"`

	MicDefaultOn *bool `json:"mic_default_on"`

	OpusDtxEnabled *bool `json:"opus_dtx_enabled"`

	RedundantCodingEnabled *bool `json:"redundant_coding_enabled"`

	SpeakerDefaultOn *bool `json:"speaker_default_on"`

	NoiseCancellation *NoiseCancellationSettings `json:"noise_cancellation"`
}

type AudioSettingsResponse struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	DefaultDevice string `json:"default_device"`

	MicDefaultOn bool `json:"mic_default_on"`

	OpusDtxEnabled bool `json:"opus_dtx_enabled"`

	RedundantCodingEnabled bool `json:"redundant_coding_enabled"`

	SpeakerDefaultOn bool `json:"speaker_default_on"`

	NoiseCancellation *NoiseCancellationSettings `json:"noise_cancellation"`
}

type AutomodDetails struct {
	Action *string `json:"action"`

	OriginalMessageType *string `json:"original_message_type"`

	ImageLabels *[]string `json:"image_labels"`

	MessageDetails *FlagMessageDetails `json:"message_details"`

	Result *MessageModerationResult `json:"result"`
}

type AzureRequest struct {
	AbsAccountName string `json:"abs_account_name"`

	AbsClientId string `json:"abs_client_id"`

	AbsClientSecret string `json:"abs_client_secret"`

	AbsTenantId string `json:"abs_tenant_id"`
}

type BackstageSettings struct {
	Enabled bool `json:"enabled"`
}

type BackstageSettingsRequest struct {
	Enabled *bool `json:"enabled"`
}

type BackstageSettingsResponse struct {
	Enabled bool `json:"enabled"`
}

type BanRequest struct {
	TargetUserId string `json:"target_user_id"`

	BannedById *string `json:"banned_by_id"`

	ChannelCid *string `json:"channel_cid"`

	IpBan *bool `json:"ip_ban"`

	Reason *string `json:"reason"`

	Shadow *bool `json:"shadow"`

	Timeout *int `json:"timeout"`

	UserId *string `json:"user_id"`

	BannedBy *UserRequest `json:"banned_by"`

	User *UserRequest `json:"user"`
}

type BanResponse struct {
	CreatedAt time.Time `json:"created_at"`

	Expires *time.Time `json:"expires"`

	Reason *string `json:"reason"`

	Shadow *bool `json:"shadow"`

	BannedBy *UserObject `json:"banned_by"`

	Channel *ChannelResponse `json:"channel"`

	User *UserObject `json:"user"`
}

type BlockList struct {
	Name string `json:"name"`

	Type string `json:"type"`

	Words []string `json:"words"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

type BlockListOptions struct {
	Behavior string `json:"behavior"`

	Blocklist string `json:"blocklist"`
}

type BlockUserRequest struct {
	UserId string `json:"user_id"`
}

type BlockUserResponse struct {
	Duration string `json:"duration"`
}

type BroadcastSettings struct {
	Enabled bool `json:"enabled"`

	Hls HLSSettings `json:"hls"`
}

type BroadcastSettingsRequest struct {
	Enabled *bool `json:"enabled"`

	Hls *HLSSettingsRequest `json:"hls"`
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
	JoinedAt time.Time `json:"joined_at"`

	Role string `json:"role"`

	UserSessionId string `json:"user_session_id"`

	User UserResponse `json:"user"`
}

type CallRecording struct {
	EndTime time.Time `json:"end_time"`

	Filename string `json:"filename"`

	StartTime time.Time `json:"start_time"`

	Url string `json:"url"`
}

type CallRequest struct {
	CreatedById *string `json:"created_by_id"`

	StartsAt *time.Time `json:"starts_at"`

	Team *string `json:"team"`

	Members *[]MemberRequest `json:"members"`

	CreatedBy *UserRequest `json:"created_by"`

	Custom *map[string]any `json:"custom"`

	SettingsOverride *CallSettingsRequest `json:"settings_override"`
}

type CallResponse struct {
	Backstage bool `json:"backstage"`

	Cid string `json:"cid"`

	CreatedAt time.Time `json:"created_at"`

	CurrentSessionId string `json:"current_session_id"`

	Id string `json:"id"`

	Recording bool `json:"recording"`

	Transcribing bool `json:"transcribing"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	BlockedUserIds []string `json:"blocked_user_ids"`

	CreatedBy UserResponse `json:"created_by"`

	Custom map[string]any `json:"custom"`

	Egress EgressResponse `json:"egress"`

	Ingress CallIngressResponse `json:"ingress"`

	Settings CallSettingsResponse `json:"settings"`

	EndedAt *time.Time `json:"ended_at"`

	StartsAt *time.Time `json:"starts_at"`

	Team *string `json:"team"`

	Session *CallSessionResponse `json:"session"`

	Thumbnails *ThumbnailResponse `json:"thumbnails"`
}

type CallSessionResponse struct {
	Id string `json:"id"`

	Participants []CallParticipantResponse `json:"participants"`

	AcceptedBy map[string]time.Time `json:"accepted_by"`

	ParticipantsCountByRole map[string]int `json:"participants_count_by_role"`

	RejectedBy map[string]time.Time `json:"rejected_by"`

	EndedAt *time.Time `json:"ended_at"`

	LiveEndedAt *time.Time `json:"live_ended_at"`

	LiveStartedAt *time.Time `json:"live_started_at"`

	StartedAt *time.Time `json:"started_at"`
}

type CallSettings struct {
	Audio *AudioSettings `json:"audio"`

	Backstage *BackstageSettings `json:"backstage"`

	Broadcasting *BroadcastSettings `json:"broadcasting"`

	Geofencing *GeofenceSettings `json:"geofencing"`

	Recording *RecordSettings `json:"recording"`

	Ring *RingSettings `json:"ring"`

	Screensharing *ScreensharingSettings `json:"screensharing"`

	Thumbnails *ThumbnailsSettings `json:"thumbnails"`

	Transcription *TranscriptionSettings `json:"transcription"`

	Video *VideoSettings `json:"video"`
}

type CallSettingsRequest struct {
	Audio *AudioSettingsRequest `json:"audio"`

	Backstage *BackstageSettingsRequest `json:"backstage"`

	Broadcasting *BroadcastSettingsRequest `json:"broadcasting"`

	Geofencing *GeofenceSettingsRequest `json:"geofencing"`

	Recording *RecordSettingsRequest `json:"recording"`

	Ring *RingSettingsRequest `json:"ring"`

	Screensharing *ScreensharingSettingsRequest `json:"screensharing"`

	Thumbnails *ThumbnailsSettingsRequest `json:"thumbnails"`

	Transcription *TranscriptionSettingsRequest `json:"transcription"`

	Video *VideoSettingsRequest `json:"video"`
}

type CallSettingsResponse struct {
	Audio AudioSettingsResponse `json:"audio"`

	Backstage BackstageSettingsResponse `json:"backstage"`

	Broadcasting BroadcastSettingsResponse `json:"broadcasting"`

	Geofencing GeofenceSettingsResponse `json:"geofencing"`

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

	CallSessionId string `json:"call_session_id"`

	CallStatus string `json:"call_status"`

	FirstStatsTime time.Time `json:"first_stats_time"`

	CreatedAt *time.Time `json:"created_at"`

	QualityScore *int `json:"quality_score"`
}

type CallTimeline struct {
	Events []*CallEvent `json:"events"`
}

type CallTranscription struct {
	EndTime time.Time `json:"end_time"`

	Filename string `json:"filename"`

	StartTime time.Time `json:"start_time"`

	Url string `json:"url"`
}

type CallType struct {
	AppPK int `json:"AppPK"`

	CreatedAt time.Time `json:"CreatedAt"`

	ExternalStorage string `json:"ExternalStorage"`

	Name string `json:"Name"`

	PK int `json:"PK"`

	UpdatedAt time.Time `json:"UpdatedAt"`

	NotificationSettings *NotificationSettings `json:"NotificationSettings"`

	Settings *CallSettings `json:"Settings"`
}

type CallTypeResponse struct {
	CreatedAt time.Time `json:"created_at"`

	Name string `json:"name"`

	UpdatedAt time.Time `json:"updated_at"`

	Grants map[string][]string `json:"grants"`

	NotificationSettings NotificationSettings `json:"notification_settings"`

	Settings CallSettingsResponse `json:"settings"`

	ExternalStorage *string `json:"external_storage"`
}

type CastPollVoteRequest struct {
	UserId *string `json:"user_id"`

	User *UserRequest `json:"user"`

	Vote *VoteData `json:"vote"`
}

type Channel struct {
	AutoTranslationLanguage string `json:"auto_translation_language"`

	Cid string `json:"cid"`

	CreatedAt time.Time `json:"created_at"`

	Disabled bool `json:"disabled"`

	Frozen bool `json:"frozen"`

	Id string `json:"id"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	AutoTranslationEnabled *bool `json:"auto_translation_enabled"`

	Cooldown *int `json:"cooldown"`

	DeletedAt *time.Time `json:"deleted_at"`

	LastMessageAt *time.Time `json:"last_message_at"`

	MemberCount *int `json:"member_count"`

	Team *string `json:"team"`

	Invites *[]*ChannelMember `json:"invites"`

	Members *[]*ChannelMember `json:"members"`

	Config *ChannelConfig `json:"config"`

	ConfigOverrides *ChannelConfig `json:"config_overrides"`

	CreatedBy *UserObject `json:"created_by"`

	TruncatedBy *UserObject `json:"truncated_by"`
}

type ChannelConfig struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt time.Time `json:"created_at"`

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

	UpdatedAt time.Time `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []string `json:"commands"`

	Blocklist *string `json:"blocklist"`

	BlocklistBehavior *string `json:"blocklist_behavior"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons"`

	Blocklists *[]BlockListOptions `json:"blocklists"`

	AutomodThresholds *Thresholds `json:"automod_thresholds"`
}

type ChannelConfigWithInfo struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt time.Time `json:"created_at"`

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

	UpdatedAt time.Time `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []*Command `json:"commands"`

	Blocklist *string `json:"blocklist"`

	BlocklistBehavior *string `json:"blocklist_behavior"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons"`

	Blocklists *[]BlockListOptions `json:"blocklists"`

	AutomodThresholds *Thresholds `json:"automod_thresholds"`

	Grants *map[string][]string `json:"grants"`
}

type ChannelExport struct {
	Cid *string `json:"cid"`

	Id *string `json:"id"`

	MessagesSince *time.Time `json:"messages_since"`

	MessagesUntil *time.Time `json:"messages_until"`

	Type *string `json:"type"`
}

type ChannelGetOrCreateRequest struct {
	HideForCreator *bool `json:"hide_for_creator"`

	State *bool `json:"state"`

	ThreadUnreadCounts *bool `json:"thread_unread_counts"`

	Data *ChannelInput `json:"data"`

	Members *PaginationParams `json:"members"`

	Messages *MessagePaginationParams `json:"messages"`

	Watchers *PaginationParams `json:"watchers"`
}

type ChannelInput struct {
	AutoTranslationEnabled *bool `json:"auto_translation_enabled"`

	AutoTranslationLanguage *string `json:"auto_translation_language"`

	CreatedById *string `json:"created_by_id"`

	Disabled *bool `json:"disabled"`

	Frozen *bool `json:"frozen"`

	Team *string `json:"team"`

	TruncatedById *string `json:"truncated_by_id"`

	Invites *[]*ChannelMember `json:"invites"`

	Members *[]*ChannelMember `json:"members"`

	ConfigOverrides *ChannelConfig `json:"config_overrides"`

	CreatedBy *UserObject `json:"created_by"`

	Custom *map[string]any `json:"custom"`
}

type ChannelMember struct {
	Banned bool `json:"banned"`

	ChannelRole string `json:"channel_role"`

	CreatedAt time.Time `json:"created_at"`

	NotificationsMuted bool `json:"notifications_muted"`

	ShadowBanned bool `json:"shadow_banned"`

	UpdatedAt time.Time `json:"updated_at"`

	BanExpires *time.Time `json:"ban_expires"`

	DeletedAt *time.Time `json:"deleted_at"`

	InviteAcceptedAt *time.Time `json:"invite_accepted_at"`

	InviteRejectedAt *time.Time `json:"invite_rejected_at"`

	Invited *bool `json:"invited"`

	IsModerator *bool `json:"is_moderator"`

	Status *string `json:"status"`

	UserId *string `json:"user_id"`

	User *UserObject `json:"user"`
}

type ChannelMute struct {
	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

	Expires *time.Time `json:"expires"`

	Channel *ChannelResponse `json:"channel"`

	User *UserObject `json:"user"`
}

type ChannelResponse struct {
	Cid string `json:"cid"`

	CreatedAt time.Time `json:"created_at"`

	Disabled bool `json:"disabled"`

	Frozen bool `json:"frozen"`

	Id string `json:"id"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	AutoTranslationEnabled *bool `json:"auto_translation_enabled"`

	AutoTranslationLanguage *string `json:"auto_translation_language"`

	Cooldown *int `json:"cooldown"`

	DeletedAt *time.Time `json:"deleted_at"`

	Hidden *bool `json:"hidden"`

	HideMessagesBefore *time.Time `json:"hide_messages_before"`

	LastMessageAt *time.Time `json:"last_message_at"`

	MemberCount *int `json:"member_count"`

	MuteExpiresAt *time.Time `json:"mute_expires_at"`

	Muted *bool `json:"muted"`

	Team *string `json:"team"`

	TruncatedAt *time.Time `json:"truncated_at"`

	Members *[]*ChannelMember `json:"members"`

	OwnCapabilities *[]string `json:"own_capabilities"`

	Config *ChannelConfigWithInfo `json:"config"`

	CreatedBy *UserObject `json:"created_by"`

	TruncatedBy *UserObject `json:"truncated_by"`
}

type ChannelStateResponse struct {
	Duration string `json:"duration"`

	Members []*ChannelMember `json:"members"`

	Messages []MessageResponse `json:"messages"`

	PinnedMessages []MessageResponse `json:"pinned_messages"`

	Threads []*ThreadState `json:"threads"`

	Hidden *bool `json:"hidden"`

	HideMessagesBefore *time.Time `json:"hide_messages_before"`

	WatcherCount *int `json:"watcher_count"`

	PendingMessages *[]*PendingMessage `json:"pending_messages"`

	Read *[]ReadStateResponse `json:"read"`

	Watchers *[]UserResponse `json:"watchers"`

	Channel *ChannelResponse `json:"channel"`

	Membership *ChannelMember `json:"membership"`
}

type ChannelStateResponseFields struct {
	Members []*ChannelMember `json:"members"`

	Messages []MessageResponse `json:"messages"`

	PinnedMessages []MessageResponse `json:"pinned_messages"`

	Threads []*ThreadState `json:"threads"`

	Hidden *bool `json:"hidden"`

	HideMessagesBefore *time.Time `json:"hide_messages_before"`

	WatcherCount *int `json:"watcher_count"`

	PendingMessages *[]*PendingMessage `json:"pending_messages"`

	Read *[]ReadStateResponse `json:"read"`

	Watchers *[]UserResponse `json:"watchers"`

	Channel *ChannelResponse `json:"channel"`

	Membership *ChannelMember `json:"membership"`
}

type ChannelTypeConfig struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt time.Time `json:"created_at"`

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

	UpdatedAt time.Time `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []*Command `json:"commands"`

	Permissions []PolicyRequest `json:"permissions"`

	Grants map[string][]string `json:"grants"`

	Blocklist *string `json:"blocklist"`

	BlocklistBehavior *string `json:"blocklist_behavior"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons"`

	Blocklists *[]BlockListOptions `json:"blocklists"`

	AutomodThresholds *Thresholds `json:"automod_thresholds"`
}

type CheckExternalStorageResponse struct {
	Duration string `json:"duration"`
}

type CheckPushRequest struct {
	ApnTemplate *string `json:"apn_template"`

	FirebaseDataTemplate *string `json:"firebase_data_template"`

	FirebaseTemplate *string `json:"firebase_template"`

	MessageId *string `json:"message_id"`

	PushProviderName *string `json:"push_provider_name"`

	PushProviderType *string `json:"push_provider_type"`

	SkipDevices *bool `json:"skip_devices"`

	UserId *string `json:"user_id"`

	User *UserRequest `json:"user"`
}

type CheckPushResponse struct {
	Duration string `json:"duration"`

	RenderedApnTemplate *string `json:"rendered_apn_template"`

	RenderedFirebaseTemplate *string `json:"rendered_firebase_template"`

	SkipDevices *bool `json:"skip_devices"`

	GeneralErrors *[]string `json:"general_errors"`

	DeviceErrors *map[string]DeviceErrorInfo `json:"device_errors"`

	RenderedMessage *map[string]string `json:"rendered_message"`
}

type CheckSNSRequest struct {
	SnsKey *string `json:"sns_key"`

	SnsSecret *string `json:"sns_secret"`

	SnsTopicArn *string `json:"sns_topic_arn"`
}

type CheckSNSResponse struct {
	Duration string `json:"duration"`

	Status string `json:"status"`

	Error *string `json:"error"`

	Data *map[string]any `json:"data"`
}

type CheckSQSRequest struct {
	SqsKey *string `json:"sqs_key"`

	SqsSecret *string `json:"sqs_secret"`

	SqsUrl *string `json:"sqs_url"`
}

type CheckSQSResponse struct {
	Duration string `json:"duration"`

	Status string `json:"status"`

	Error *string `json:"error"`

	Data *map[string]any `json:"data"`
}

type CollectUserFeedbackRequest struct {
	Rating int `json:"rating"`

	Sdk string `json:"sdk"`

	SdkVersion string `json:"sdk_version"`

	UserSessionId string `json:"user_session_id"`

	Reason *string `json:"reason"`

	Custom *map[string]any `json:"custom"`
}

type CollectUserFeedbackResponse struct {
	Duration string `json:"duration"`
}

type Command struct {
	Args string `json:"args"`

	Description string `json:"description"`

	Name string `json:"name"`

	Set string `json:"set"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

type CommitMessageRequest struct{}

type Config struct {
	AppCertificate string `json:"app_certificate"`

	AppId string `json:"app_id"`

	DefaultRole *string `json:"default_role"`

	RoleMap *map[string]string `json:"role_map"`
}

type Coordinates struct {
	Latitude float64 `json:"latitude"`

	Longitude float64 `json:"longitude"`
}

type CreateBlockListRequest struct {
	Name string `json:"name"`

	Words []string `json:"words"`

	Type *string `json:"type"`
}

type CreateCallTypeRequest struct {
	Name string `json:"name"`

	ExternalStorage *string `json:"external_storage"`

	Grants *map[string][]string `json:"grants"`

	NotificationSettings *NotificationSettings `json:"notification_settings"`

	Settings *CallSettingsRequest `json:"settings"`
}

type CreateCallTypeResponse struct {
	CreatedAt time.Time `json:"created_at"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	UpdatedAt time.Time `json:"updated_at"`

	Grants map[string][]string `json:"grants"`

	NotificationSettings NotificationSettings `json:"notification_settings"`

	Settings CallSettingsResponse `json:"settings"`

	ExternalStorage *string `json:"external_storage"`
}

type CreateChannelTypeRequest struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	MaxMessageLength int `json:"max_message_length"`

	Name string `json:"name"`

	Blocklist *string `json:"blocklist"`

	BlocklistBehavior *string `json:"blocklist_behavior"`

	ConnectEvents *bool `json:"connect_events"`

	CustomEvents *bool `json:"custom_events"`

	MarkMessagesPending *bool `json:"mark_messages_pending"`

	MessageRetention *string `json:"message_retention"`

	Mutes *bool `json:"mutes"`

	Polls *bool `json:"polls"`

	PushNotifications *bool `json:"push_notifications"`

	Reactions *bool `json:"reactions"`

	ReadEvents *bool `json:"read_events"`

	Replies *bool `json:"replies"`

	Search *bool `json:"search"`

	TypingEvents *bool `json:"typing_events"`

	Uploads *bool `json:"uploads"`

	UrlEnrichment *bool `json:"url_enrichment"`

	Blocklists *[]BlockListOptions `json:"blocklists"`

	Commands *[]string `json:"commands"`

	Permissions *[]PolicyRequest `json:"permissions"`

	Grants *map[string][]string `json:"grants"`
}

type CreateChannelTypeResponse struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt time.Time `json:"created_at"`

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

	UpdatedAt time.Time `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []string `json:"commands"`

	Permissions []PolicyRequest `json:"permissions"`

	Grants map[string][]string `json:"grants"`

	Blocklist *string `json:"blocklist"`

	BlocklistBehavior *string `json:"blocklist_behavior"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons"`

	Blocklists *[]BlockListOptions `json:"blocklists"`

	AutomodThresholds *Thresholds `json:"automod_thresholds"`
}

type CreateCommandRequest struct {
	Description string `json:"description"`

	Name string `json:"name"`

	Args *string `json:"args"`

	Set *string `json:"set"`
}

type CreateCommandResponse struct {
	Duration string `json:"duration"`

	Command *Command `json:"command"`
}

type CreateDeviceRequest struct {
	Id string `json:"id"`

	PushProvider string `json:"push_provider"`

	PushProviderName *string `json:"push_provider_name"`

	UserId *string `json:"user_id"`

	VoipToken *bool `json:"voip_token"`

	User *UserRequest `json:"user"`
}

type CreateExternalStorageRequest struct {
	Bucket string `json:"bucket"`

	Name string `json:"name"`

	StorageType string `json:"storage_type"`

	GcsCredentials *string `json:"gcs_credentials"`

	Path *string `json:"path"`

	AwsS3 *S3Request `json:"aws_s3"`

	AzureBlob *AzureRequest `json:"azure_blob"`
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

	ImportTask *ImportTask `json:"import_task"`
}

type CreateImportURLRequest struct {
	Filename *string `json:"filename"`
}

type CreateImportURLResponse struct {
	Duration string `json:"duration"`

	Path string `json:"path"`

	UploadUrl string `json:"upload_url"`
}

type CreatePollOptionRequest struct {
	Text string `json:"text"`

	Position *int `json:"position"`

	UserId *string `json:"user_id"`

	Custom *map[string]any `json:"Custom"`

	User *UserRequest `json:"user"`
}

type CreatePollRequest struct {
	Name string `json:"name"`

	AllowAnswers *bool `json:"allow_answers"`

	AllowUserSuggestedOptions *bool `json:"allow_user_suggested_options"`

	Description *string `json:"description"`

	EnforceUniqueVote *bool `json:"enforce_unique_vote"`

	Id *string `json:"id"`

	IsClosed *bool `json:"is_closed"`

	MaxVotesAllowed *int `json:"max_votes_allowed"`

	UserId *string `json:"user_id"`

	VotingVisibility *string `json:"voting_visibility"`

	Options *[]*PollOptionInput `json:"options"`

	Custom *map[string]any `json:"Custom"`

	User *UserRequest `json:"user"`
}

type CreateRoleRequest struct {
	Name string `json:"name"`
}

type CreateRoleResponse struct {
	Duration string `json:"duration"`

	Role Role `json:"role"`
}

type DataDogInfo struct {
	ApiKey *string `json:"api_key"`

	Enabled *bool `json:"enabled"`

	Site *string `json:"site"`
}

type DeactivateUserRequest struct {
	CreatedById *string `json:"created_by_id"`

	MarkMessagesDeleted *bool `json:"mark_messages_deleted"`
}

type DeactivateUserResponse struct {
	Duration string `json:"duration"`

	User *UserObject `json:"user"`
}

type DeactivateUsersRequest struct {
	UserIds []string `json:"user_ids"`

	CreatedById *string `json:"created_by_id"`

	MarkChannelsDeleted *bool `json:"mark_channels_deleted"`

	MarkMessagesDeleted *bool `json:"mark_messages_deleted"`
}

type DeactivateUsersResponse struct {
	Duration string `json:"duration"`

	TaskId string `json:"task_id"`
}

type DeleteChannelResponse struct {
	Duration string `json:"duration"`

	Channel *ChannelResponse `json:"channel"`
}

type DeleteChannelsRequest struct {
	Cids []string `json:"cids"`

	HardDelete *bool `json:"hard_delete"`
}

type DeleteChannelsResponse struct {
	Duration string `json:"duration"`

	TaskId *string `json:"task_id"`

	Result *map[string]*DeleteChannelsResult `json:"result"`
}

type DeleteChannelsResult struct {
	Status string `json:"status"`

	Error *string `json:"error"`
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
	UserIds []string `json:"user_ids"`

	Calls *string `json:"calls"`

	Conversations *string `json:"conversations"`

	Messages *string `json:"messages"`

	NewCallOwnerId *string `json:"new_call_owner_id"`

	NewChannelOwnerId *string `json:"new_channel_owner_id"`

	User *string `json:"user"`
}

type DeleteUsersResponse struct {
	Duration string `json:"duration"`

	TaskId string `json:"task_id"`
}

type Device struct {
	CreatedAt time.Time `json:"created_at"`

	Id string `json:"id"`

	PushProvider string `json:"push_provider"`

	UserId string `json:"user_id"`

	Disabled *bool `json:"disabled"`

	DisabledReason *string `json:"disabled_reason"`

	PushProviderName *string `json:"push_provider_name"`

	Voip *bool `json:"voip"`
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

	Id string `json:"id"`

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

	Hls *EgressHLSResponse `json:"hls"`
}

type EndCallRequest struct{}

type EndCallResponse struct {
	Duration string `json:"duration"`
}

type ErrorResult struct {
	Type string `json:"type"`

	Stacktrace *string `json:"stacktrace"`

	Version *string `json:"version"`
}

type EventNotificationSettings struct {
	Enabled bool `json:"enabled"`

	Apns APNS `json:"apns"`
}

type EventRequest struct {
	Type string `json:"type"`

	ParentId *string `json:"parent_id"`

	UserId *string `json:"user_id"`

	Custom *map[string]any `json:"custom"`

	User *UserRequest `json:"user"`
}

type EventResponse struct {
	Duration string `json:"duration"`

	Event WSEvent `json:"event"`
}

type ExportChannelsRequest struct {
	Channels []ChannelExport `json:"channels"`

	ClearDeletedMessageText *bool `json:"clear_deleted_message_text"`

	ExportUsers *bool `json:"export_users"`

	IncludeSoftDeletedChannels *bool `json:"include_soft_deleted_channels"`

	IncludeTruncatedMessages *bool `json:"include_truncated_messages"`

	Version *string `json:"version"`
}

type ExportChannelsResponse struct {
	Duration string `json:"duration"`

	TaskId string `json:"task_id"`
}

type ExportChannelsResult struct {
	Url string `json:"url"`

	Path *string `json:"path"`

	S3BucketName *string `json:"s3_bucket_name"`
}

type ExportUserResponse struct {
	Duration string `json:"duration"`

	Messages *[]*Message `json:"messages"`

	Reactions *[]*Reaction `json:"reactions"`

	User *UserObject `json:"user"`
}

type ExportUsersRequest struct {
	UserIds []string `json:"user_ids"`
}

type ExportUsersResponse struct {
	Duration string `json:"duration"`

	TaskId string `json:"task_id"`
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

	AllowedFileExtensions *[]string `json:"allowed_file_extensions"`

	AllowedMimeTypes *[]string `json:"allowed_mime_types"`

	BlockedFileExtensions *[]string `json:"blocked_file_extensions"`

	BlockedMimeTypes *[]string `json:"blocked_mime_types"`
}

type FileUploadRequest struct {
	File *string `json:"file"`

	User *OnlyUserID `json:"user"`
}

type FileUploadResponse struct {
	Duration string `json:"duration"`

	File *string `json:"file"`

	ThumbUrl *string `json:"thumb_url"`
}

type FirebaseConfig struct {
	ApnTemplate *string `json:"apn_template"`

	CredentialsJson *string `json:"credentials_json"`

	DataTemplate *string `json:"data_template"`

	Disabled *bool `json:"Disabled"`

	NotificationTemplate *string `json:"notification_template"`

	ServerKey *string `json:"server_key"`
}

type FirebaseConfigFields struct {
	ApnTemplate string `json:"apn_template"`

	DataTemplate string `json:"data_template"`

	Enabled bool `json:"enabled"`

	NotificationTemplate string `json:"notification_template"`

	CredentialsJson *string `json:"credentials_json"`

	ServerKey *string `json:"server_key"`
}

type Flag struct {
	CreatedAt time.Time `json:"created_at"`

	CreatedByAutomod bool `json:"created_by_automod"`

	UpdatedAt time.Time `json:"updated_at"`

	ApprovedAt *time.Time `json:"approved_at"`

	Reason *string `json:"reason"`

	RejectedAt *time.Time `json:"rejected_at"`

	ReviewedAt *time.Time `json:"reviewed_at"`

	ReviewedBy *string `json:"reviewed_by"`

	TargetMessageId *string `json:"target_message_id"`

	Custom *map[string]any `json:"custom"`

	Details *FlagDetails `json:"details"`

	TargetMessage *Message `json:"target_message"`

	TargetUser *UserObject `json:"target_user"`

	User *UserObject `json:"user"`
}

type FlagDetails struct {
	OriginalText string `json:"original_text"`

	Extra map[string]any `json:"Extra"`

	Automod *AutomodDetails `json:"automod"`
}

type FlagFeedback struct {
	CreatedAt time.Time `json:"created_at"`

	MessageId string `json:"message_id"`

	Labels []Label `json:"labels"`
}

type FlagMessageDetails struct {
	PinChanged *bool `json:"pin_changed"`

	ShouldEnrich *bool `json:"should_enrich"`

	SkipPush *bool `json:"skip_push"`

	UpdatedById *string `json:"updated_by_id"`
}

type FlagRequest struct {
	Reason *string `json:"reason"`

	TargetMessageId *string `json:"target_message_id"`

	TargetUserId *string `json:"target_user_id"`

	UserId *string `json:"user_id"`

	Custom *map[string]any `json:"custom"`

	User *UserRequest `json:"user"`
}

type FlagResponse struct {
	Duration string `json:"duration"`

	Flag *Flag `json:"flag"`
}

type FullUserResponse struct {
	Banned bool `json:"banned"`

	CreatedAt time.Time `json:"created_at"`

	Id string `json:"id"`

	Invisible bool `json:"invisible"`

	Language string `json:"language"`

	Online bool `json:"online"`

	Role string `json:"role"`

	ShadowBanned bool `json:"shadow_banned"`

	TotalUnreadCount int `json:"total_unread_count"`

	UnreadChannels int `json:"unread_channels"`

	UnreadThreads int `json:"unread_threads"`

	UpdatedAt time.Time `json:"updated_at"`

	ChannelMutes []*ChannelMute `json:"channel_mutes"`

	Devices []*Device `json:"devices"`

	Mutes []*UserMute `json:"mutes"`

	Teams []string `json:"teams"`

	Custom map[string]any `json:"custom"`

	DeactivatedAt *time.Time `json:"deactivated_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	Image *string `json:"image"`

	LastActive *time.Time `json:"last_active"`

	Name *string `json:"name"`

	RevokeTokensIssuedBefore *time.Time `json:"revoke_tokens_issued_before"`

	LatestHiddenChannels *[]string `json:"latest_hidden_channels"`

	PrivacySettings *PrivacySettings `json:"privacy_settings"`

	PushNotifications *PushNotificationSettings `json:"push_notifications"`
}

type GeofenceResponse struct {
	Name string `json:"name"`

	Description *string `json:"description"`

	Type *string `json:"type"`

	CountryCodes *[]string `json:"country_codes"`
}

type GeofenceSettings struct {
	Names []string `json:"names"`
}

type GeofenceSettingsRequest struct {
	Names *[]string `json:"names"`
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

	Blocklist *BlockList `json:"blocklist"`
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

	CallTimeline *CallTimeline `json:"call_timeline"`

	Jitter *Stats `json:"jitter"`

	Latency *Stats `json:"latency"`
}

type GetCallTypeResponse struct {
	CreatedAt time.Time `json:"created_at"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	UpdatedAt time.Time `json:"updated_at"`

	Grants map[string][]string `json:"grants"`

	NotificationSettings NotificationSettings `json:"notification_settings"`

	Settings CallSettingsResponse `json:"settings"`

	ExternalStorage *string `json:"external_storage"`
}

type GetCommandResponse struct {
	Args string `json:"args"`

	Description string `json:"description"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	Set string `json:"set"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
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
	CreatedAt time.Time `json:"created_at"`

	Duration string `json:"duration"`

	Status string `json:"status"`

	TaskId string `json:"task_id"`

	UpdatedAt time.Time `json:"updated_at"`

	Error *ErrorResult `json:"error"`

	Result *ExportChannelsResult `json:"result"`
}

type GetImportResponse struct {
	Duration string `json:"duration"`

	ImportTask *ImportTask `json:"import_task"`
}

type GetManyMessagesResponse struct {
	Duration string `json:"duration"`

	Messages []*Message `json:"messages"`
}

type GetMessageResponse struct {
	Duration string `json:"duration"`

	Message MessageWithChannelResponse `json:"message"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata"`
}

type GetOGResponse struct {
	Duration string `json:"duration"`

	Custom map[string]any `json:"custom"`

	AssetUrl *string `json:"asset_url"`

	AuthorIcon *string `json:"author_icon"`

	AuthorLink *string `json:"author_link"`

	AuthorName *string `json:"author_name"`

	Color *string `json:"color"`

	Fallback *string `json:"fallback"`

	Footer *string `json:"footer"`

	FooterIcon *string `json:"footer_icon"`

	ImageUrl *string `json:"image_url"`

	OgScrapeUrl *string `json:"og_scrape_url"`

	OriginalHeight *int `json:"original_height"`

	OriginalWidth *int `json:"original_width"`

	Pretext *string `json:"pretext"`

	Text *string `json:"text"`

	ThumbUrl *string `json:"thumb_url"`

	Title *string `json:"title"`

	TitleLink *string `json:"title_link"`

	Type *string `json:"type"`

	Actions *[]*Action `json:"actions"`

	Fields *[]*Field `json:"fields"`

	Giphy *Images `json:"giphy"`
}

type GetOrCreateCallRequest struct {
	MembersLimit *int `json:"members_limit"`

	Notify *bool `json:"notify"`

	Ring *bool `json:"ring"`

	Data *CallRequest `json:"data"`
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

	Android *map[string]LimitInfo `json:"android"`

	Ios *map[string]LimitInfo `json:"ios"`

	ServerSide *map[string]LimitInfo `json:"server_side"`

	Web *map[string]LimitInfo `json:"web"`
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
	CreatedAt time.Time `json:"created_at"`

	Duration string `json:"duration"`

	Status string `json:"status"`

	TaskId string `json:"task_id"`

	UpdatedAt time.Time `json:"updated_at"`

	Error *ErrorResult `json:"error"`

	Result *map[string]any `json:"result"`
}

type GetThreadResponse struct {
	Duration string `json:"duration"`

	Thread ThreadStateResponse `json:"thread"`
}

type GoLiveRequest struct {
	RecordingStorageName *string `json:"recording_storage_name"`

	StartHls *bool `json:"start_hls"`

	StartRecording *bool `json:"start_recording"`

	StartTranscription *bool `json:"start_transcription"`

	TranscriptionStorageName *string `json:"transcription_storage_name"`
}

type GoLiveResponse struct {
	Duration string `json:"duration"`

	Call CallResponse `json:"call"`
}

type HLSSettings struct {
	AutoOn bool `json:"auto_on"`

	Enabled bool `json:"enabled"`

	QualityTracks []string `json:"quality_tracks"`

	Layout *LayoutSettings `json:"layout"`
}

type HLSSettingsRequest struct {
	QualityTracks []string `json:"quality_tracks"`

	AutoOn *bool `json:"auto_on"`

	Enabled *bool `json:"enabled"`

	Layout *LayoutSettingsRequest `json:"layout"`
}

type HLSSettingsResponse struct {
	AutoOn bool `json:"auto_on"`

	Enabled bool `json:"enabled"`

	QualityTracks []string `json:"quality_tracks"`

	Layout LayoutSettingsResponse `json:"layout"`
}

type HideChannelRequest struct {
	ClearHistory *bool `json:"clear_history"`

	UserId *string `json:"user_id"`

	User *UserRequest `json:"user"`
}

type HideChannelResponse struct {
	Duration string `json:"duration"`
}

type HuaweiConfig struct {
	Disabled *bool `json:"Disabled"`

	Id *string `json:"id"`

	Secret *string `json:"secret"`
}

type HuaweiConfigFields struct {
	Enabled bool `json:"enabled"`

	Id *string `json:"id"`

	Secret *string `json:"secret"`
}

type ImageData struct {
	Frames string `json:"frames"`

	Height string `json:"height"`

	Size string `json:"size"`

	Url string `json:"url"`

	Width string `json:"width"`
}

type ImageSize struct {
	Crop *string `json:"crop"`

	Height *int `json:"height"`

	Resize *string `json:"resize"`

	Width *int `json:"width"`
}

type ImageUploadRequest struct {
	File *string `json:"file"`

	UploadSizes *[]ImageSize `json:"upload_sizes"`

	User *OnlyUserID `json:"user"`
}

type ImageUploadResponse struct {
	Duration string `json:"duration"`

	File *string `json:"file"`

	ThumbUrl *string `json:"thumb_url"`

	UploadSizes *[]ImageSize `json:"upload_sizes"`
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
	CreatedAt time.Time `json:"created_at"`

	Id string `json:"id"`

	Mode string `json:"mode"`

	Path string `json:"path"`

	State string `json:"state"`

	UpdatedAt time.Time `json:"updated_at"`

	History []*ImportTaskHistory `json:"history"`

	Size *int `json:"size"`
}

type ImportTaskHistory struct {
	CreatedAt time.Time `json:"created_at"`

	NextState string `json:"next_state"`

	PrevState string `json:"prev_state"`
}

type Label struct {
	Name string `json:"name"`

	HarmLabels *[]string `json:"harm_labels"`

	PhraseListIds *[]int `json:"phrase_list_ids"`
}

type LabelThresholds struct {
	Block *float64 `json:"block"`

	Flag *float64 `json:"flag"`
}

type LayoutSettings struct {
	ExternalAppUrl string `json:"external_app_url"`

	ExternalCssUrl string `json:"external_css_url"`

	Name string `json:"name"`

	Options *map[string]any `json:"options"`
}

type LayoutSettingsRequest struct {
	Name string `json:"name"`

	ExternalAppUrl *string `json:"external_app_url"`

	ExternalCssUrl *string `json:"external_css_url"`

	Options *map[string]any `json:"options"`
}

type LayoutSettingsResponse struct {
	ExternalAppUrl string `json:"external_app_url"`

	ExternalCssUrl string `json:"external_css_url"`

	Name string `json:"name"`

	Options *map[string]any `json:"options"`
}

type LimitInfo struct {
	Limit int `json:"limit"`

	Remaining int `json:"remaining"`

	Reset int `json:"reset"`
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
	UserId *string `json:"user_id"`

	ReadByChannel *map[string]string `json:"read_by_channel"`

	User *UserRequest `json:"user"`
}

type MarkReadRequest struct {
	MessageId *string `json:"message_id"`

	ThreadId *string `json:"thread_id"`

	UserId *string `json:"user_id"`

	User *UserRequest `json:"user"`
}

type MarkReadResponse struct {
	Duration string `json:"duration"`

	Event *MessageReadEvent `json:"event"`
}

type MarkUnreadRequest struct {
	MessageId *string `json:"message_id"`

	ThreadId *string `json:"thread_id"`

	UserId *string `json:"user_id"`

	User *UserRequest `json:"user"`
}

type MediaPubSubHint struct {
	AudioPublished bool `json:"audio_published"`

	AudioSubscribed bool `json:"audio_subscribed"`

	VideoPublished bool `json:"video_published"`

	VideoSubscribed bool `json:"video_subscribed"`
}

type MemberRequest struct {
	UserId string `json:"user_id"`

	Role *string `json:"role"`

	Custom *map[string]any `json:"custom"`
}

type MemberResponse struct {
	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

	UserId string `json:"user_id"`

	Custom map[string]any `json:"custom"`

	User UserResponse `json:"user"`

	DeletedAt *time.Time `json:"deleted_at"`

	Role *string `json:"role"`
}

type MembersResponse struct {
	Duration string `json:"duration"`

	Members []*ChannelMember `json:"members"`
}

type Message struct {
	Cid string `json:"cid"`

	CreatedAt time.Time `json:"created_at"`

	DeletedReplyCount int `json:"deleted_reply_count"`

	Html string `json:"html"`

	Id string `json:"id"`

	Pinned bool `json:"pinned"`

	ReplyCount int `json:"reply_count"`

	Shadowed bool `json:"shadowed"`

	Silent bool `json:"silent"`

	Text string `json:"text"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	Attachments []*Attachment `json:"attachments"`

	LatestReactions []*Reaction `json:"latest_reactions"`

	MentionedUsers []UserObject `json:"mentioned_users"`

	OwnReactions []*Reaction `json:"own_reactions"`

	Custom map[string]any `json:"custom"`

	ReactionCounts map[string]int `json:"reaction_counts"`

	ReactionGroups map[string]*ReactionGroupResponse `json:"reaction_groups"`

	ReactionScores map[string]int `json:"reaction_scores"`

	BeforeMessageSendFailed *bool `json:"before_message_send_failed"`

	Command *string `json:"command"`

	DeletedAt *time.Time `json:"deleted_at"`

	MessageTextUpdatedAt *time.Time `json:"message_text_updated_at"`

	Mml *string `json:"mml"`

	ParentId *string `json:"parent_id"`

	PinExpires *time.Time `json:"pin_expires"`

	PinnedAt *time.Time `json:"pinned_at"`

	PollId *string `json:"poll_id"`

	QuotedMessageId *string `json:"quoted_message_id"`

	ShowInChannel *bool `json:"show_in_channel"`

	ThreadParticipants *[]UserObject `json:"thread_participants"`

	I18n *map[string]string `json:"i18n"`

	ImageLabels *map[string][]string `json:"image_labels"`

	PinnedBy *UserObject `json:"pinned_by"`

	Poll *Poll `json:"poll"`

	QuotedMessage *Message `json:"quoted_message"`

	User *UserObject `json:"user"`
}

type MessageActionRequest struct {
	FormData map[string]string `json:"form_data"`

	UserId *string `json:"user_id"`

	User *UserRequest `json:"user"`
}

type MessageChangeSet struct {
	Attachments bool `json:"attachments"`

	Custom bool `json:"custom"`

	Html bool `json:"html"`

	MentionedUserIds bool `json:"mentioned_user_ids"`

	Mml bool `json:"mml"`

	Pin bool `json:"pin"`

	QuotedMessageId bool `json:"quoted_message_id"`

	Silent bool `json:"silent"`

	Text bool `json:"text"`
}

type MessageFlag struct {
	CreatedAt time.Time `json:"created_at"`

	CreatedByAutomod bool `json:"created_by_automod"`

	UpdatedAt time.Time `json:"updated_at"`

	ApprovedAt *time.Time `json:"approved_at"`

	Reason *string `json:"reason"`

	RejectedAt *time.Time `json:"rejected_at"`

	ReviewedAt *time.Time `json:"reviewed_at"`

	Custom *map[string]any `json:"custom"`

	Details *FlagDetails `json:"details"`

	Message *Message `json:"message"`

	ModerationFeedback *FlagFeedback `json:"moderation_feedback"`

	ModerationResult *MessageModerationResult `json:"moderation_result"`

	ReviewedBy *UserObject `json:"reviewed_by"`

	User *UserObject `json:"user"`
}

type MessageHistoryEntry struct {
	MessageId string `json:"message_id"`

	MessageUpdatedAt time.Time `json:"message_updated_at"`

	MessageUpdatedById string `json:"message_updated_by_id"`

	Text string `json:"text"`

	Attachments []*Attachment `json:"attachments"`

	Custom map[string]any `json:"Custom"`
}

type MessageModerationResult struct {
	Action string `json:"action"`

	CreatedAt time.Time `json:"created_at"`

	MessageId string `json:"message_id"`

	UpdatedAt time.Time `json:"updated_at"`

	UserBadKarma bool `json:"user_bad_karma"`

	UserKarma float64 `json:"user_karma"`

	BlockedWord *string `json:"blocked_word"`

	BlocklistName *string `json:"blocklist_name"`

	ModeratedBy *string `json:"moderated_by"`

	AiModerationResponse *ModerationResponse `json:"ai_moderation_response"`

	ModerationThresholds *Thresholds `json:"moderation_thresholds"`
}

type MessagePaginationParams struct{}

type MessageReadEvent struct {
	ChannelId string `json:"channel_id"`

	ChannelType string `json:"channel_type"`

	Cid string `json:"cid"`

	CreatedAt time.Time `json:"created_at"`

	Type string `json:"type"`

	LastReadMessageId *string `json:"last_read_message_id"`

	Team *string `json:"team"`

	Thread *Thread `json:"thread"`

	User *UserObject `json:"user"`
}

type MessageRequest struct {
	Html *string `json:"html"`

	Id *string `json:"id"`

	Mml *string `json:"mml"`

	ParentId *string `json:"parent_id"`

	PinExpires *time.Time `json:"pin_expires"`

	Pinned *bool `json:"pinned"`

	PinnedAt *time.Time `json:"pinned_at"`

	PollId *string `json:"poll_id"`

	QuotedMessageId *string `json:"quoted_message_id"`

	ShowInChannel *bool `json:"show_in_channel"`

	Silent *bool `json:"silent"`

	Text *string `json:"text"`

	Type *string `json:"type"`

	UserId *string `json:"user_id"`

	Attachments *[]*Attachment `json:"attachments"`

	MentionedUsers *[]string `json:"mentioned_users"`

	Custom *map[string]any `json:"custom"`

	User *UserRequest `json:"user"`
}

type MessageResponse struct {
	Cid string `json:"cid"`

	CreatedAt time.Time `json:"created_at"`

	DeletedReplyCount int `json:"deleted_reply_count"`

	Html string `json:"html"`

	Id string `json:"id"`

	Pinned bool `json:"pinned"`

	ReplyCount int `json:"reply_count"`

	Shadowed bool `json:"shadowed"`

	Silent bool `json:"silent"`

	Text string `json:"text"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	Attachments []*Attachment `json:"attachments"`

	LatestReactions []ReactionResponse `json:"latest_reactions"`

	MentionedUsers []UserResponse `json:"mentioned_users"`

	OwnReactions []ReactionResponse `json:"own_reactions"`

	Custom map[string]any `json:"custom"`

	ReactionCounts map[string]int `json:"reaction_counts"`

	ReactionScores map[string]int `json:"reaction_scores"`

	User UserResponse `json:"user"`

	Command *string `json:"command"`

	DeletedAt *time.Time `json:"deleted_at"`

	MessageTextUpdatedAt *time.Time `json:"message_text_updated_at"`

	Mml *string `json:"mml"`

	ParentId *string `json:"parent_id"`

	PinExpires *time.Time `json:"pin_expires"`

	PinnedAt *time.Time `json:"pinned_at"`

	PollId *string `json:"poll_id"`

	QuotedMessageId *string `json:"quoted_message_id"`

	ShowInChannel *bool `json:"show_in_channel"`

	ThreadParticipants *[]UserResponse `json:"thread_participants"`

	I18n *map[string]string `json:"i18n"`

	ImageLabels *map[string][]string `json:"image_labels"`

	PinnedBy *UserResponse `json:"pinned_by"`

	Poll *Poll `json:"poll"`

	QuotedMessage *Message `json:"quoted_message"`

	ReactionGroups *map[string]*ReactionGroupResponse `json:"reaction_groups"`
}

type MessageUpdate struct {
	OldText *string `json:"old_text"`

	ChangeSet *MessageChangeSet `json:"change_set"`
}

type MessageWithChannelResponse struct {
	Cid string `json:"cid"`

	CreatedAt time.Time `json:"created_at"`

	DeletedReplyCount int `json:"deleted_reply_count"`

	Html string `json:"html"`

	Id string `json:"id"`

	Pinned bool `json:"pinned"`

	ReplyCount int `json:"reply_count"`

	Shadowed bool `json:"shadowed"`

	Silent bool `json:"silent"`

	Text string `json:"text"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	Attachments []*Attachment `json:"attachments"`

	LatestReactions []ReactionResponse `json:"latest_reactions"`

	MentionedUsers []UserResponse `json:"mentioned_users"`

	OwnReactions []ReactionResponse `json:"own_reactions"`

	Channel ChannelResponse `json:"channel"`

	Custom map[string]any `json:"custom"`

	ReactionCounts map[string]int `json:"reaction_counts"`

	ReactionScores map[string]int `json:"reaction_scores"`

	User UserResponse `json:"user"`

	Command *string `json:"command"`

	DeletedAt *time.Time `json:"deleted_at"`

	MessageTextUpdatedAt *time.Time `json:"message_text_updated_at"`

	Mml *string `json:"mml"`

	ParentId *string `json:"parent_id"`

	PinExpires *time.Time `json:"pin_expires"`

	PinnedAt *time.Time `json:"pinned_at"`

	PollId *string `json:"poll_id"`

	QuotedMessageId *string `json:"quoted_message_id"`

	ShowInChannel *bool `json:"show_in_channel"`

	ThreadParticipants *[]UserResponse `json:"thread_participants"`

	I18n *map[string]string `json:"i18n"`

	ImageLabels *map[string][]string `json:"image_labels"`

	PinnedBy *UserResponse `json:"pinned_by"`

	Poll *Poll `json:"poll"`

	QuotedMessage *Message `json:"quoted_message"`

	ReactionGroups *map[string]*ReactionGroupResponse `json:"reaction_groups"`
}

type ModerationResponse struct {
	Action string `json:"action"`

	Explicit float64 `json:"explicit"`

	Spam float64 `json:"spam"`

	Toxic float64 `json:"toxic"`
}

type MuteChannelRequest struct {
	Expiration *int `json:"expiration"`

	UserId *string `json:"user_id"`

	ChannelCids *[]string `json:"channel_cids"`

	User *UserRequest `json:"user"`
}

type MuteChannelResponse struct {
	Duration string `json:"duration"`

	ChannelMutes *[]*ChannelMute `json:"channel_mutes"`

	ChannelMute *ChannelMute `json:"channel_mute"`

	OwnUser *OwnUser `json:"own_user"`
}

type MuteUserRequest struct {
	Timeout int `json:"timeout"`

	UserId *string `json:"user_id"`

	TargetIds *[]string `json:"target_ids"`

	User *UserRequest `json:"user"`
}

type MuteUserResponse struct {
	Duration string `json:"duration"`

	Mutes *[]*UserMute `json:"mutes"`

	NonExistingUsers *[]string `json:"non_existing_users"`

	Mute *UserMute `json:"mute"`

	OwnUser *OwnUser `json:"own_user"`
}

type MuteUsersRequest struct {
	Audio *bool `json:"audio"`

	MuteAllUsers *bool `json:"mute_all_users"`

	MutedById *string `json:"muted_by_id"`

	Screenshare *bool `json:"screenshare"`

	ScreenshareAudio *bool `json:"screenshare_audio"`

	Video *bool `json:"video"`

	UserIds *[]string `json:"user_ids"`

	MutedBy *UserRequest `json:"muted_by"`
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

	CallNotification EventNotificationSettings `json:"call_notification"`

	CallRing EventNotificationSettings `json:"call_ring"`

	SessionStarted EventNotificationSettings `json:"session_started"`
}

type NullBool struct {
	HasValue *bool `json:"HasValue"`

	Value *bool `json:"Value"`
}

type NullTime struct {
	HasValue *bool `json:"HasValue"`

	Value *time.Time `json:"Value"`
}

type OnlyUserID struct {
	Id string `json:"id"`
}

type OwnCapability string

const (
	BLOCK_USERS               OwnCapability = "block-users"
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

type OwnUser struct {
	Banned bool `json:"banned"`

	CreatedAt time.Time `json:"created_at"`

	Id string `json:"id"`

	Language string `json:"language"`

	Online bool `json:"online"`

	Role string `json:"role"`

	TotalUnreadCount int `json:"total_unread_count"`

	UnreadChannels int `json:"unread_channels"`

	UnreadCount int `json:"unread_count"`

	UnreadThreads int `json:"unread_threads"`

	UpdatedAt time.Time `json:"updated_at"`

	ChannelMutes []*ChannelMute `json:"channel_mutes"`

	Devices []*Device `json:"devices"`

	Mutes []*UserMute `json:"mutes"`

	Custom map[string]any `json:"custom"`

	DeactivatedAt *time.Time `json:"deactivated_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	Invisible *bool `json:"invisible"`

	LastActive *time.Time `json:"last_active"`

	LatestHiddenChannels *[]string `json:"latest_hidden_channels"`

	Teams *[]string `json:"teams"`

	PrivacySettings *PrivacySettings `json:"privacy_settings"`

	PushNotifications *PushNotificationSettings `json:"push_notifications"`
}

type PaginationParams struct {
	Limit *int `json:"limit"`

	Offset *int `json:"offset"`
}

type PendingMessage struct {
	Channel *Channel `json:"channel"`

	Message *Message `json:"message"`

	Metadata *map[string]string `json:"metadata"`

	User *UserObject `json:"user"`
}

type Permission struct {
	Action string `json:"action"`

	Custom bool `json:"custom"`

	Description string `json:"description"`

	Id string `json:"id"`

	Level string `json:"level"`

	Name string `json:"name"`

	Owner bool `json:"owner"`

	SameTeam bool `json:"same_team"`

	Tags []string `json:"tags"`

	Condition *map[string]any `json:"condition"`
}

type PinRequest struct {
	SessionId string `json:"session_id"`

	UserId string `json:"user_id"`
}

type PinResponse struct {
	Duration string `json:"duration"`
}

type Policy struct {
	Action int `json:"action"`

	CreatedAt time.Time `json:"created_at"`

	Name string `json:"name"`

	Owner bool `json:"owner"`

	Priority int `json:"priority"`

	UpdatedAt time.Time `json:"updated_at"`

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

	CreatedAt time.Time `json:"created_at"`

	CreatedById string `json:"created_by_id"`

	Description string `json:"description"`

	EnforceUniqueVote bool `json:"enforce_unique_vote"`

	Id string `json:"id"`

	Name string `json:"name"`

	UpdatedAt time.Time `json:"updated_at"`

	VoteCount int `json:"vote_count"`

	LatestAnswers []*PollVote `json:"latest_answers"`

	Options []*PollOption `json:"options"`

	OwnVotes []*PollVote `json:"own_votes"`

	Custom map[string]any `json:"Custom"`

	LatestVotesByOption map[string][]*PollVote `json:"latest_votes_by_option"`

	VoteCountsByOption map[string]int `json:"vote_counts_by_option"`

	IsClosed *bool `json:"is_closed"`

	MaxVotesAllowed *int `json:"max_votes_allowed"`

	VotingVisibility *string `json:"voting_visibility"`

	CreatedBy *UserObject `json:"created_by"`
}

type PollOption struct {
	Id string `json:"id"`

	Text string `json:"text"`

	Custom map[string]any `json:"custom"`
}

type PollOptionInput struct {
	Text *string `json:"text"`

	Custom *map[string]any `json:"custom"`
}

type PollOptionResponse struct {
	Duration string `json:"duration"`

	PollOption PollOptionResponseData `json:"poll_option"`
}

type PollOptionResponseData struct {
	Id string `json:"id"`

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

	CreatedAt time.Time `json:"created_at"`

	CreatedById string `json:"created_by_id"`

	Description string `json:"description"`

	EnforceUniqueVote bool `json:"enforce_unique_vote"`

	Id string `json:"id"`

	Name string `json:"name"`

	UpdatedAt time.Time `json:"updated_at"`

	VoteCount int `json:"vote_count"`

	VotingVisibility string `json:"voting_visibility"`

	Options []*PollOptionResponseData `json:"options"`

	OwnVotes []*PollVoteResponseData `json:"own_votes"`

	Custom map[string]any `json:"Custom"`

	LatestVotesByOption map[string][]*PollVoteResponseData `json:"latest_votes_by_option"`

	VoteCountsByOption map[string]int `json:"vote_counts_by_option"`

	IsClosed *bool `json:"is_closed"`

	MaxVotesAllowed *int `json:"max_votes_allowed"`

	CreatedBy *UserObject `json:"created_by"`
}

type PollVote struct {
	CreatedAt time.Time `json:"created_at"`

	Id string `json:"id"`

	OptionId string `json:"option_id"`

	PollId string `json:"poll_id"`

	UpdatedAt time.Time `json:"updated_at"`

	AnswerText *string `json:"answer_text"`

	IsAnswer *bool `json:"is_answer"`

	UserId *string `json:"user_id"`

	User *UserObject `json:"user"`
}

type PollVoteResponse struct {
	Duration string `json:"duration"`

	Vote *PollVoteResponseData `json:"vote"`
}

type PollVoteResponseData struct {
	CreatedAt time.Time `json:"created_at"`

	Id string `json:"id"`

	OptionId string `json:"option_id"`

	PollId string `json:"poll_id"`

	UpdatedAt time.Time `json:"updated_at"`

	AnswerText *string `json:"answer_text"`

	IsAnswer *bool `json:"is_answer"`

	UserId *string `json:"user_id"`

	User *UserObject `json:"user"`
}

type PollVotesResponse struct {
	Duration string `json:"duration"`

	Votes []*PollVoteResponseData `json:"votes"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`
}

type PrivacySettings struct {
	ReadReceipts *ReadReceipts `json:"read_receipts"`

	TypingIndicators *TypingIndicators `json:"typing_indicators"`
}

type PublishedTrackInfo struct {
	CodecMimeType *string `json:"codec_mime_type"`

	DurationSeconds *int `json:"duration_seconds"`

	TrackType *string `json:"track_type"`
}

type PushConfig struct {
	Version string `json:"version"`

	OfflineOnly *bool `json:"offline_only"`
}

type PushNotificationFields struct {
	OfflineOnly bool `json:"offline_only"`

	Version string `json:"version"`

	Apn APNConfigFields `json:"apn"`

	Firebase FirebaseConfigFields `json:"firebase"`

	Huawei HuaweiConfigFields `json:"huawei"`

	Xiaomi XiaomiConfigFields `json:"xiaomi"`

	Providers *[]*PushProvider `json:"providers"`
}

type PushNotificationSettings struct {
	Disabled *bool `json:"disabled"`

	DisabledUntil *time.Time `json:"disabled_until"`
}

type PushNotificationSettingsInput struct {
	Disabled *NullBool `json:"disabled"`

	DisabledUntil *NullTime `json:"disabled_until"`
}

type PushProvider struct {
	CreatedAt time.Time `json:"created_at"`

	Name string `json:"name"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	ApnAuthKey *string `json:"apn_auth_key"`

	ApnAuthType *string `json:"apn_auth_type"`

	ApnDevelopment *bool `json:"apn_development"`

	ApnHost *string `json:"apn_host"`

	ApnKeyId *string `json:"apn_key_id"`

	ApnNotificationTemplate *string `json:"apn_notification_template"`

	ApnP12Cert *string `json:"apn_p12_cert"`

	ApnTeamId *string `json:"apn_team_id"`

	ApnTopic *string `json:"apn_topic"`

	Description *string `json:"description"`

	DisabledAt *time.Time `json:"disabled_at"`

	DisabledReason *string `json:"disabled_reason"`

	FirebaseApnTemplate *string `json:"firebase_apn_template"`

	FirebaseCredentials *string `json:"firebase_credentials"`

	FirebaseDataTemplate *string `json:"firebase_data_template"`

	FirebaseHost *string `json:"firebase_host"`

	FirebaseNotificationTemplate *string `json:"firebase_notification_template"`

	FirebaseServerKey *string `json:"firebase_server_key"`

	HuaweiAppId *string `json:"huawei_app_id"`

	HuaweiAppSecret *string `json:"huawei_app_secret"`

	XiaomiAppSecret *string `json:"xiaomi_app_secret"`

	XiaomiPackageName *string `json:"xiaomi_package_name"`
}

type PushProviderResponse struct {
	CreatedAt time.Time `json:"created_at"`

	Name string `json:"name"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	ApnAuthKey *string `json:"apn_auth_key"`

	ApnAuthType *string `json:"apn_auth_type"`

	ApnDevelopment *bool `json:"apn_development"`

	ApnHost *string `json:"apn_host"`

	ApnKeyId *string `json:"apn_key_id"`

	ApnP12Cert *string `json:"apn_p12_cert"`

	ApnSandboxCertificate *bool `json:"apn_sandbox_certificate"`

	ApnSupportsRemoteNotifications *bool `json:"apn_supports_remote_notifications"`

	ApnSupportsVoipNotifications *bool `json:"apn_supports_voip_notifications"`

	ApnTeamId *string `json:"apn_team_id"`

	ApnTopic *string `json:"apn_topic"`

	Description *string `json:"description"`

	DisabledAt *time.Time `json:"disabled_at"`

	DisabledReason *string `json:"disabled_reason"`

	FirebaseApnTemplate *string `json:"firebase_apn_template"`

	FirebaseCredentials *string `json:"firebase_credentials"`

	FirebaseDataTemplate *string `json:"firebase_data_template"`

	FirebaseHost *string `json:"firebase_host"`

	FirebaseNotificationTemplate *string `json:"firebase_notification_template"`

	FirebaseServerKey *string `json:"firebase_server_key"`

	HuaweiAppId *string `json:"huawei_app_id"`

	HuaweiAppSecret *string `json:"huawei_app_secret"`

	XiaomiAppSecret *string `json:"xiaomi_app_secret"`

	XiaomiPackageName *string `json:"xiaomi_package_name"`
}

type QueryBannedUsersRequest struct {
	FilterConditions map[string]any `json:"filter_conditions"`

	ExcludeExpiredBans *bool `json:"exclude_expired_bans"`

	Limit *int `json:"limit"`

	Offset *int `json:"offset"`

	UserId *string `json:"user_id"`

	Sort *[]*SortParam `json:"sort"`

	User *UserRequest `json:"user"`
}

type QueryBannedUsersResponse struct {
	Duration string `json:"duration"`

	Bans []*BanResponse `json:"bans"`
}

type QueryCallMembersRequest struct {
	Id string `json:"id"`

	Type string `json:"type"`

	Limit *int `json:"limit"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`

	Sort *[]*SortParam `json:"sort"`

	FilterConditions *map[string]any `json:"filter_conditions"`
}

type QueryCallMembersResponse struct {
	Duration string `json:"duration"`

	Members []MemberResponse `json:"members"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`
}

type QueryCallStatsRequest struct {
	Limit *int `json:"limit"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`

	Sort *[]*SortParam `json:"sort"`

	FilterConditions *map[string]any `json:"filter_conditions"`
}

type QueryCallStatsResponse struct {
	Duration string `json:"duration"`

	Reports []CallStatsReportSummaryResponse `json:"reports"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`
}

type QueryCallsRequest struct {
	Limit *int `json:"limit"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`

	Sort *[]*SortParam `json:"sort"`

	FilterConditions *map[string]any `json:"filter_conditions"`
}

type QueryCallsResponse struct {
	Duration string `json:"duration"`

	Calls []CallStateResponseFields `json:"calls"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`
}

type QueryChannelsRequest struct {
	Limit *int `json:"limit"`

	MemberLimit *int `json:"member_limit"`

	MessageLimit *int `json:"message_limit"`

	Offset *int `json:"offset"`

	State *bool `json:"state"`

	UserId *string `json:"user_id"`

	Sort *[]*SortParam `json:"sort"`

	FilterConditions *map[string]any `json:"filter_conditions"`

	User *UserRequest `json:"user"`
}

type QueryChannelsResponse struct {
	Duration string `json:"duration"`

	Channels []ChannelStateResponseFields `json:"channels"`
}

type QueryMembersRequest struct {
	Type string `json:"type"`

	FilterConditions map[string]any `json:"filter_conditions"`

	Id *string `json:"id"`

	Limit *int `json:"limit"`

	Offset *int `json:"offset"`

	UserId *string `json:"user_id"`

	Members *[]*ChannelMember `json:"members"`

	Sort *[]*SortParam `json:"sort"`

	User *UserRequest `json:"user"`
}

type QueryMessageFlagsRequest struct {
	Limit *int `json:"limit"`

	Offset *int `json:"offset"`

	ShowDeletedMessages *bool `json:"show_deleted_messages"`

	UserId *string `json:"user_id"`

	Sort *[]*SortParam `json:"sort"`

	FilterConditions *map[string]any `json:"filter_conditions"`

	User *UserRequest `json:"user"`
}

type QueryMessageFlagsResponse struct {
	Duration string `json:"duration"`

	Flags []*MessageFlag `json:"flags"`
}

type QueryMessageHistoryRequest struct {
	Filter map[string]any `json:"filter"`

	Limit *int `json:"limit"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`

	Sort *[]*SortParam `json:"sort"`
}

type QueryMessageHistoryResponse struct {
	Duration string `json:"duration"`

	MessageHistory []*MessageHistoryEntry `json:"message_history"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`
}

type QueryPollVotesRequest struct {
	Limit *int `json:"limit"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`

	Sort *[]*SortParam `json:"sort"`

	Filter *map[string]any `json:"filter"`
}

type QueryPollsRequest struct {
	Limit *int `json:"limit"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`

	Sort *[]*SortParam `json:"sort"`

	Filter *map[string]any `json:"filter"`
}

type QueryPollsResponse struct {
	Duration string `json:"duration"`

	Polls []PollResponseData `json:"polls"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`
}

type QueryReactionsRequest struct {
	Limit *int `json:"limit"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`

	UserId *string `json:"user_id"`

	Sort *[]*SortParam `json:"sort"`

	Filter *map[string]any `json:"filter"`

	User *UserRequest `json:"user"`
}

type QueryReactionsResponse struct {
	Duration string `json:"duration"`

	Reactions []ReactionResponse `json:"reactions"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`
}

type QueryThreadsRequest struct {
	Limit *int `json:"limit"`

	MemberLimit *int `json:"member_limit"`

	Next *string `json:"next"`

	ParticipantLimit *int `json:"participant_limit"`

	Prev *string `json:"prev"`

	ReplyLimit *int `json:"reply_limit"`

	UserId *string `json:"user_id"`

	User *UserRequest `json:"user"`
}

type QueryThreadsResponse struct {
	Duration string `json:"duration"`

	Threads []ThreadStateResponse `json:"threads"`

	Next *string `json:"next"`

	Prev *string `json:"prev"`
}

type QueryUsersPayload struct {
	FilterConditions map[string]any `json:"filter_conditions"`

	IncludeDeactivatedUsers *bool `json:"include_deactivated_users"`

	Limit *int `json:"limit"`

	Offset *int `json:"offset"`

	Presence *bool `json:"presence"`

	UserId *string `json:"user_id"`

	Sort *[]*SortParam `json:"sort"`

	User *UserRequest `json:"user"`
}

type QueryUsersResponse struct {
	Duration string `json:"duration"`

	Users []FullUserResponse `json:"users"`
}

type RTMPIngress struct {
	Address string `json:"address"`
}

type Reaction struct {
	CreatedAt time.Time `json:"created_at"`

	MessageId string `json:"message_id"`

	Score int `json:"score"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	UserId *string `json:"user_id"`

	User *UserObject `json:"user"`
}

type ReactionGroupResponse struct {
	Count int `json:"count"`

	FirstReactionAt time.Time `json:"first_reaction_at"`

	LastReactionAt time.Time `json:"last_reaction_at"`

	SumScores int `json:"sum_scores"`
}

type ReactionRemovalResponse struct {
	Duration string `json:"duration"`

	Message *Message `json:"message"`

	Reaction *Reaction `json:"reaction"`
}

type ReactionRequest struct {
	Type string `json:"type"`

	CreatedAt *time.Time `json:"created_at"`

	Score *int `json:"score"`

	UpdatedAt *time.Time `json:"updated_at"`

	UserId *string `json:"user_id"`

	Custom *map[string]any `json:"custom"`

	User *UserRequest `json:"user"`
}

type ReactionResponse struct {
	CreatedAt time.Time `json:"created_at"`

	MessageId string `json:"message_id"`

	Score int `json:"score"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	UserId string `json:"user_id"`

	Custom map[string]any `json:"custom"`

	User UserResponse `json:"user"`
}

type ReactivateUserRequest struct {
	CreatedById *string `json:"created_by_id"`

	Name *string `json:"name"`

	RestoreMessages *bool `json:"restore_messages"`
}

type ReactivateUserResponse struct {
	Duration string `json:"duration"`

	User *UserObject `json:"user"`
}

type ReactivateUsersRequest struct {
	UserIds []string `json:"user_ids"`

	CreatedById *string `json:"created_by_id"`

	RestoreChannels *bool `json:"restore_channels"`

	RestoreMessages *bool `json:"restore_messages"`
}

type ReactivateUsersResponse struct {
	Duration string `json:"duration"`

	TaskId string `json:"task_id"`
}

type Read struct {
	LastRead time.Time `json:"last_read"`

	UnreadMessages int `json:"unread_messages"`

	LastReadMessageId *string `json:"last_read_message_id"`

	User *UserObject `json:"user"`
}

type ReadReceipts struct {
	Enabled *bool `json:"enabled"`
}

type ReadStateResponse struct {
	LastRead time.Time `json:"last_read"`

	UnreadMessages int `json:"unread_messages"`

	User UserResponse `json:"user"`

	LastReadMessageId *string `json:"last_read_message_id"`
}

type RecordSettings struct {
	AudioOnly bool `json:"audio_only"`

	Mode string `json:"mode"`

	Quality string `json:"quality"`

	Layout *LayoutSettings `json:"layout"`
}

type RecordSettingsRequest struct {
	Mode string `json:"mode"`

	AudioOnly *bool `json:"audio_only"`

	Quality *string `json:"quality"`

	Layout *LayoutSettingsRequest `json:"layout"`
}

type RecordSettingsResponse struct {
	AudioOnly bool `json:"audio_only"`

	Mode string `json:"mode"`

	Quality string `json:"quality"`

	Layout LayoutSettingsResponse `json:"layout"`
}

// type Response struct {
// 	Duration string `json:"duration"`
// }

type RestoreUsersRequest struct {
	UserIds []string `json:"user_ids"`
}

type RingSettings struct {
	AutoCancelTimeoutMs int `json:"auto_cancel_timeout_ms"`

	IncomingCallTimeoutMs int `json:"incoming_call_timeout_ms"`
}

type RingSettingsRequest struct {
	AutoCancelTimeoutMs int `json:"auto_cancel_timeout_ms"`

	IncomingCallTimeoutMs int `json:"incoming_call_timeout_ms"`
}

type RingSettingsResponse struct {
	AutoCancelTimeoutMs int `json:"auto_cancel_timeout_ms"`

	IncomingCallTimeoutMs int `json:"incoming_call_timeout_ms"`
}

type Role struct {
	CreatedAt time.Time `json:"created_at"`

	Custom bool `json:"custom"`

	Name string `json:"name"`

	UpdatedAt time.Time `json:"updated_at"`

	Scopes []string `json:"scopes"`
}

type S3Request struct {
	S3Region string `json:"s3_region"`

	S3ApiKey *string `json:"s3_api_key"`

	S3Secret *string `json:"s3_secret"`
}

type SFULocationResponse struct {
	Datacenter string `json:"datacenter"`

	Id string `json:"id"`

	Coordinates Coordinates `json:"coordinates"`

	Location Location `json:"location"`
}

type ScreensharingSettings struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	Enabled bool `json:"enabled"`

	TargetResolution *TargetResolution `json:"target_resolution"`
}

type ScreensharingSettingsRequest struct {
	AccessRequestEnabled *bool `json:"access_request_enabled"`

	Enabled *bool `json:"enabled"`

	TargetResolution *TargetResolution `json:"target_resolution"`
}

type ScreensharingSettingsResponse struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	Enabled bool `json:"enabled"`

	TargetResolution *TargetResolution `json:"target_resolution"`
}

type SearchRequest struct {
	FilterConditions map[string]any `json:"filter_conditions"`

	Limit *int `json:"limit"`

	Next *string `json:"next"`

	Offset *int `json:"offset"`

	Query *string `json:"query"`

	Sort *[]*SortParam `json:"sort"`

	MessageFilterConditions *map[string]any `json:"message_filter_conditions"`
}

type SearchResponse struct {
	Duration string `json:"duration"`

	Results []SearchResult `json:"results"`

	Next *string `json:"next"`

	Previous *string `json:"previous"`

	ResultsWarning *SearchWarning `json:"results_warning"`
}

type SearchResult struct {
	Message *SearchResultMessage `json:"message"`
}

type SearchResultMessage struct {
	Cid string `json:"cid"`

	CreatedAt time.Time `json:"created_at"`

	DeletedReplyCount int `json:"deleted_reply_count"`

	Html string `json:"html"`

	Id string `json:"id"`

	Pinned bool `json:"pinned"`

	ReplyCount int `json:"reply_count"`

	Shadowed bool `json:"shadowed"`

	Silent bool `json:"silent"`

	Text string `json:"text"`

	Type string `json:"type"`

	UpdatedAt time.Time `json:"updated_at"`

	Attachments []*Attachment `json:"attachments"`

	LatestReactions []*Reaction `json:"latest_reactions"`

	MentionedUsers []UserObject `json:"mentioned_users"`

	OwnReactions []*Reaction `json:"own_reactions"`

	Custom map[string]any `json:"custom"`

	ReactionCounts map[string]int `json:"reaction_counts"`

	ReactionGroups map[string]*ReactionGroupResponse `json:"reaction_groups"`

	ReactionScores map[string]int `json:"reaction_scores"`

	BeforeMessageSendFailed *bool `json:"before_message_send_failed"`

	Command *string `json:"command"`

	DeletedAt *time.Time `json:"deleted_at"`

	MessageTextUpdatedAt *time.Time `json:"message_text_updated_at"`

	Mml *string `json:"mml"`

	ParentId *string `json:"parent_id"`

	PinExpires *time.Time `json:"pin_expires"`

	PinnedAt *time.Time `json:"pinned_at"`

	PollId *string `json:"poll_id"`

	QuotedMessageId *string `json:"quoted_message_id"`

	ShowInChannel *bool `json:"show_in_channel"`

	ThreadParticipants *[]UserObject `json:"thread_participants"`

	Channel *ChannelResponse `json:"channel"`

	I18n *map[string]string `json:"i18n"`

	ImageLabels *map[string][]string `json:"image_labels"`

	PinnedBy *UserObject `json:"pinned_by"`

	Poll *Poll `json:"poll"`

	QuotedMessage *Message `json:"quoted_message"`

	User *UserObject `json:"user"`
}

type SearchWarning struct {
	WarningCode int `json:"warning_code"`

	WarningDescription string `json:"warning_description"`

	ChannelSearchCount *int `json:"channel_search_count"`

	ChannelSearchCids *[]string `json:"channel_search_cids"`
}

type SendCallEventRequest struct {
	UserId *string `json:"user_id"`

	Custom *map[string]any `json:"custom"`

	User *UserRequest `json:"user"`
}

type SendCallEventResponse struct {
	Duration string `json:"duration"`
}

type SendEventRequest struct {
	Event EventRequest `json:"event"`
}

type SendMessageRequest struct {
	Message MessageRequest `json:"message"`

	ForceModeration *bool `json:"force_moderation"`

	KeepChannelHidden *bool `json:"keep_channel_hidden"`

	Pending *bool `json:"pending"`

	SkipEnrichUrl *bool `json:"skip_enrich_url"`

	SkipPush *bool `json:"skip_push"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata"`
}

type SendMessageResponse struct {
	Duration string `json:"duration"`

	Message MessageResponse `json:"message"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata"`
}

type SendReactionRequest struct {
	Reaction ReactionRequest `json:"reaction"`

	EnforceUnique *bool `json:"enforce_unique"`

	SkipPush *bool `json:"skip_push"`
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
	UserId *string `json:"user_id"`

	User *UserRequest `json:"user"`
}

type ShowChannelResponse struct {
	Duration string `json:"duration"`
}

type SortParam struct {
	Direction *int `json:"direction"`

	Field *string `json:"field"`
}

type StartHLSBroadcastingRequest struct{}

type StartHLSBroadcastingResponse struct {
	Duration string `json:"duration"`

	PlaylistUrl string `json:"playlist_url"`
}

type StartRecordingRequest struct {
	RecordingExternalStorage *string `json:"recording_external_storage"`
}

type StartRecordingResponse struct {
	Duration string `json:"duration"`
}

type StartTranscriptionRequest struct {
	TranscriptionExternalStorage *string `json:"transcription_external_storage"`
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

	SfuId string `json:"sfu_id"`

	PubSubHint *MediaPubSubHint `json:"pub_sub_hint"`
}

type TargetResolution struct {
	Bitrate int `json:"bitrate"`

	Height int `json:"height"`

	Width int `json:"width"`
}

type Thread struct {
	ChannelCid string `json:"channel_cid"`

	CreatedAt time.Time `json:"created_at"`

	ParentMessageId string `json:"parent_message_id"`

	Title string `json:"title"`

	UpdatedAt time.Time `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	DeletedAt *time.Time `json:"deleted_at"`

	LastMessageAt *time.Time `json:"last_message_at"`

	ParticipantCount *int `json:"participant_count"`

	ReplyCount *int `json:"reply_count"`

	ThreadParticipants *[]*ThreadParticipant `json:"thread_participants"`

	Channel *Channel `json:"channel"`

	CreatedBy *UserObject `json:"created_by"`

	ParentMessage *Message `json:"parent_message"`
}

type ThreadParticipant struct {
	AppPk int `json:"app_pk"`

	ChannelCid string `json:"channel_cid"`

	CreatedAt time.Time `json:"created_at"`

	LastReadAt time.Time `json:"last_read_at"`

	Custom map[string]any `json:"custom"`

	LastThreadMessageAt *time.Time `json:"last_thread_message_at"`

	LeftThreadAt *time.Time `json:"left_thread_at"`

	ThreadId *string `json:"thread_id"`

	UserId *string `json:"user_id"`

	User *UserObject `json:"user"`
}

type ThreadResponse struct {
	ChannelCid string `json:"channel_cid"`

	CreatedAt time.Time `json:"created_at"`

	CreatedByUserId string `json:"created_by_user_id"`

	ParentMessageId string `json:"parent_message_id"`

	Title string `json:"title"`

	UpdatedAt time.Time `json:"updated_at"`

	Custom map[string]any `json:"custom"`

	DeletedAt *time.Time `json:"deleted_at"`

	LastMessageAt *time.Time `json:"last_message_at"`

	ParticipantCount *int `json:"participant_count"`

	ReplyCount *int `json:"reply_count"`

	ThreadParticipants *[]*ThreadParticipant `json:"thread_participants"`

	Channel *ChannelResponse `json:"channel"`

	CreatedBy *UserObject `json:"created_by"`

	ParentMessage *Message `json:"parent_message"`
}

type ThreadState struct {
	ChannelCid string `json:"channel_cid"`

	CreatedAt time.Time `json:"created_at"`

	ParentMessageId string `json:"parent_message_id"`

	Title string `json:"title"`

	UpdatedAt time.Time `json:"updated_at"`

	LatestReplies []*Message `json:"latest_replies"`

	Custom map[string]any `json:"custom"`

	DeletedAt *time.Time `json:"deleted_at"`

	LastMessageAt *time.Time `json:"last_message_at"`

	ParticipantCount *int `json:"participant_count"`

	ReplyCount *int `json:"reply_count"`

	Read *[]*Read `json:"read"`

	ThreadParticipants *[]*ThreadParticipant `json:"thread_participants"`

	Channel *Channel `json:"channel"`

	CreatedBy *UserObject `json:"created_by"`

	ParentMessage *Message `json:"parent_message"`
}

type ThreadStateResponse struct {
	ChannelCid string `json:"channel_cid"`

	CreatedAt time.Time `json:"created_at"`

	CreatedByUserId string `json:"created_by_user_id"`

	ParentMessageId string `json:"parent_message_id"`

	Title string `json:"title"`

	UpdatedAt time.Time `json:"updated_at"`

	LatestReplies []*Message `json:"latest_replies"`

	Custom map[string]any `json:"custom"`

	DeletedAt *time.Time `json:"deleted_at"`

	LastMessageAt *time.Time `json:"last_message_at"`

	ParticipantCount *int `json:"participant_count"`

	ReplyCount *int `json:"reply_count"`

	Read *[]*Read `json:"read"`

	ThreadParticipants *[]*ThreadParticipant `json:"thread_participants"`

	Channel *ChannelResponse `json:"channel"`

	CreatedBy *UserObject `json:"created_by"`

	ParentMessage *Message `json:"parent_message"`
}

type Thresholds struct {
	Explicit *LabelThresholds `json:"explicit"`

	Spam *LabelThresholds `json:"spam"`

	Toxic *LabelThresholds `json:"toxic"`
}

type ThumbnailResponse struct {
	ImageUrl string `json:"image_url"`
}

type ThumbnailsSettings struct {
	Enabled bool `json:"enabled"`
}

type ThumbnailsSettingsRequest struct {
	Enabled *bool `json:"enabled"`
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

	ClosedCaptionMode *string `json:"closed_caption_mode"`

	Languages *[]string `json:"languages"`
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
	HardDelete *bool `json:"hard_delete"`

	SkipPush *bool `json:"skip_push"`

	TruncatedAt *time.Time `json:"truncated_at"`

	UserId *string `json:"user_id"`

	Message *MessageRequest `json:"message"`

	User *UserRequest `json:"user"`
}

type TruncateChannelResponse struct {
	Duration string `json:"duration"`

	Channel *ChannelResponse `json:"channel"`

	Message *Message `json:"message"`
}

type TypingIndicators struct {
	Enabled *bool `json:"enabled"`
}

type UnblockUserRequest struct {
	UserId string `json:"user_id"`
}

type UnblockUserResponse struct {
	Duration string `json:"duration"`
}

type UnmuteChannelRequest struct {
	Expiration *int `json:"expiration"`

	UserId *string `json:"user_id"`

	ChannelCids *[]string `json:"channel_cids"`

	User *UserRequest `json:"user"`
}

type UnmuteResponse struct {
	Duration string `json:"duration"`

	NonExistingUsers *[]string `json:"non_existing_users"`
}

type UnmuteUserRequest struct {
	Timeout int `json:"timeout"`

	UserId *string `json:"user_id"`

	TargetIds *[]string `json:"target_ids"`

	User *UserRequest `json:"user"`
}

type UnpinRequest struct {
	SessionId string `json:"session_id"`

	UserId string `json:"user_id"`
}

type UnpinResponse struct {
	Duration string `json:"duration"`
}

type UnreadCountsBatchRequest struct {
	UserIds []string `json:"user_ids"`
}

type UnreadCountsBatchResponse struct {
	Duration string `json:"duration"`

	CountsByUser map[string]*UnreadCountsResponse `json:"counts_by_user"`
}

type UnreadCountsChannel struct {
	ChannelId string `json:"channel_id"`

	LastRead time.Time `json:"last_read"`

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
	LastRead time.Time `json:"last_read"`

	LastReadMessageId string `json:"last_read_message_id"`

	ParentMessageId string `json:"parent_message_id"`

	UnreadCount int `json:"unread_count"`
}

type UpdateAppRequest struct {
	AsyncUrlEnrichEnabled *bool `json:"async_url_enrich_enabled"`

	AutoTranslationEnabled *bool `json:"auto_translation_enabled"`

	BeforeMessageSendHookUrl *string `json:"before_message_send_hook_url"`

	CdnExpirationSeconds *int `json:"cdn_expiration_seconds"`

	ChannelHideMembersOnly *bool `json:"channel_hide_members_only"`

	CustomActionHandlerUrl *string `json:"custom_action_handler_url"`

	DisableAuthChecks *bool `json:"disable_auth_checks"`

	DisablePermissionsChecks *bool `json:"disable_permissions_checks"`

	EnforceUniqueUsernames *string `json:"enforce_unique_usernames"`

	ImageModerationEnabled *bool `json:"image_moderation_enabled"`

	MigratePermissionsToV2 *bool `json:"migrate_permissions_to_v2"`

	MultiTenantEnabled *bool `json:"multi_tenant_enabled"`

	PermissionVersion *string `json:"permission_version"`

	RemindersInterval *int `json:"reminders_interval"`

	RemindersMaxMembers *int `json:"reminders_max_members"`

	RevokeTokensIssuedBefore *time.Time `json:"revoke_tokens_issued_before"`

	SnsKey *string `json:"sns_key"`

	SnsSecret *string `json:"sns_secret"`

	SnsTopicArn *string `json:"sns_topic_arn"`

	SqsKey *string `json:"sqs_key"`

	SqsSecret *string `json:"sqs_secret"`

	SqsUrl *string `json:"sqs_url"`

	VideoProvider *string `json:"video_provider"`

	WebhookUrl *string `json:"webhook_url"`

	ImageModerationBlockLabels *[]string `json:"image_moderation_block_labels"`

	ImageModerationLabels *[]string `json:"image_moderation_labels"`

	UserSearchDisallowedRoles *[]string `json:"user_search_disallowed_roles"`

	WebhookEvents *[]string `json:"webhook_events"`

	AgoraOptions *Config `json:"agora_options"`

	ApnConfig *APNConfig `json:"apn_config"`

	AsyncModerationConfig *AsyncModerationConfiguration `json:"async_moderation_config"`

	DatadogInfo *DataDogInfo `json:"datadog_info"`

	FileUploadConfig *FileUploadConfig `json:"file_upload_config"`

	FirebaseConfig *FirebaseConfig `json:"firebase_config"`

	Grants *map[string][]string `json:"grants"`

	HmsOptions *Config `json:"hms_options"`

	HuaweiConfig *HuaweiConfig `json:"huawei_config"`

	ImageUploadConfig *FileUploadConfig `json:"image_upload_config"`

	PushConfig *PushConfig `json:"push_config"`

	XiaomiConfig *XiaomiConfig `json:"xiaomi_config"`
}

type UpdateBlockListRequest struct {
	Words *[]string `json:"words"`
}

type UpdateCallMembersRequest struct {
	RemoveMembers *[]string `json:"remove_members"`

	UpdateMembers *[]MemberRequest `json:"update_members"`
}

type UpdateCallMembersResponse struct {
	Duration string `json:"duration"`

	Members []MemberResponse `json:"members"`
}

type UpdateCallRequest struct {
	StartsAt *time.Time `json:"starts_at"`

	Custom *map[string]any `json:"custom"`

	SettingsOverride *CallSettingsRequest `json:"settings_override"`
}

type UpdateCallResponse struct {
	Duration string `json:"duration"`

	Members []MemberResponse `json:"members"`

	OwnCapabilities []OwnCapability `json:"own_capabilities"`

	Call CallResponse `json:"call"`
}

type UpdateCallTypeRequest struct {
	ExternalStorage *string `json:"external_storage"`

	Grants *map[string][]string `json:"grants"`

	NotificationSettings *NotificationSettings `json:"notification_settings"`

	Settings *CallSettingsRequest `json:"settings"`
}

type UpdateCallTypeResponse struct {
	CreatedAt time.Time `json:"created_at"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	UpdatedAt time.Time `json:"updated_at"`

	Grants map[string][]string `json:"grants"`

	NotificationSettings NotificationSettings `json:"notification_settings"`

	Settings CallSettingsResponse `json:"settings"`

	ExternalStorage *string `json:"external_storage"`
}

type UpdateChannelPartialRequest struct {
	UserId *string `json:"user_id"`

	Unset *[]string `json:"unset"`

	Set *map[string]any `json:"set"`

	User *UserRequest `json:"user"`
}

type UpdateChannelPartialResponse struct {
	Duration string `json:"duration"`

	Members []*ChannelMember `json:"members"`

	Channel *ChannelResponse `json:"channel"`
}

type UpdateChannelRequest struct {
	AcceptInvite *bool `json:"accept_invite"`

	Cooldown *int `json:"cooldown"`

	HideHistory *bool `json:"hide_history"`

	RejectInvite *bool `json:"reject_invite"`

	SkipPush *bool `json:"skip_push"`

	UserId *string `json:"user_id"`

	AddMembers *[]*ChannelMember `json:"add_members"`

	AddModerators *[]string `json:"add_moderators"`

	AssignRoles *[]*ChannelMember `json:"assign_roles"`

	DemoteModerators *[]string `json:"demote_moderators"`

	Invites *[]*ChannelMember `json:"invites"`

	RemoveMembers *[]string `json:"remove_members"`

	Data *ChannelInput `json:"data"`

	Message *MessageRequest `json:"message"`

	User *UserRequest `json:"user"`
}

type UpdateChannelResponse struct {
	Duration string `json:"duration"`

	Members []*ChannelMember `json:"members"`

	Channel *ChannelResponse `json:"channel"`

	Message *Message `json:"message"`
}

type UpdateChannelTypeRequest struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	MaxMessageLength int `json:"max_message_length"`

	Blocklist *string `json:"blocklist"`

	BlocklistBehavior *string `json:"blocklist_behavior"`

	ConnectEvents *bool `json:"connect_events"`

	CustomEvents *bool `json:"custom_events"`

	MarkMessagesPending *bool `json:"mark_messages_pending"`

	Mutes *bool `json:"mutes"`

	Polls *bool `json:"polls"`

	PushNotifications *bool `json:"push_notifications"`

	Quotes *bool `json:"quotes"`

	Reactions *bool `json:"reactions"`

	ReadEvents *bool `json:"read_events"`

	Reminders *bool `json:"reminders"`

	Replies *bool `json:"replies"`

	Search *bool `json:"search"`

	TypingEvents *bool `json:"typing_events"`

	Uploads *bool `json:"uploads"`

	UrlEnrichment *bool `json:"url_enrichment"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons"`

	Blocklists *[]BlockListOptions `json:"blocklists"`

	Commands *[]string `json:"commands"`

	Permissions *[]PolicyRequest `json:"permissions"`

	AutomodThresholds *Thresholds `json:"automod_thresholds"`

	Grants *map[string][]string `json:"grants"`
}

type UpdateChannelTypeResponse struct {
	Automod string `json:"automod"`

	AutomodBehavior string `json:"automod_behavior"`

	ConnectEvents bool `json:"connect_events"`

	CreatedAt time.Time `json:"created_at"`

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

	UpdatedAt time.Time `json:"updated_at"`

	Uploads bool `json:"uploads"`

	UrlEnrichment bool `json:"url_enrichment"`

	Commands []string `json:"commands"`

	Permissions []PolicyRequest `json:"permissions"`

	Grants map[string][]string `json:"grants"`

	Blocklist *string `json:"blocklist"`

	BlocklistBehavior *string `json:"blocklist_behavior"`

	AllowedFlagReasons *[]string `json:"allowed_flag_reasons"`

	Blocklists *[]BlockListOptions `json:"blocklists"`

	AutomodThresholds *Thresholds `json:"automod_thresholds"`
}

type UpdateCommandRequest struct {
	Description string `json:"description"`

	Args *string `json:"args"`

	Set *string `json:"set"`
}

type UpdateCommandResponse struct {
	Duration string `json:"duration"`

	Command *Command `json:"command"`
}

type UpdateExternalStorageRequest struct {
	Bucket string `json:"bucket"`

	StorageType string `json:"storage_type"`

	GcsCredentials *string `json:"gcs_credentials"`

	Path *string `json:"path"`

	AwsS3 *S3Request `json:"aws_s3"`

	AzureBlob *AzureRequest `json:"azure_blob"`
}

type UpdateExternalStorageResponse struct {
	Bucket string `json:"bucket"`

	Duration string `json:"duration"`

	Name string `json:"name"`

	Path string `json:"path"`

	Type string `json:"type"`
}

type UpdateMessagePartialRequest struct {
	SkipEnrichUrl *bool `json:"skip_enrich_url"`

	UserId *string `json:"user_id"`

	Unset *[]string `json:"unset"`

	Set *map[string]any `json:"set"`

	User *UserRequest `json:"user"`
}

type UpdateMessagePartialResponse struct {
	Duration string `json:"duration"`

	Message *Message `json:"message"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata"`
}

type UpdateMessageRequest struct {
	Message MessageRequest `json:"message"`

	SkipEnrichUrl *bool `json:"skip_enrich_url"`
}

type UpdateMessageResponse struct {
	Duration string `json:"duration"`

	Message Message `json:"message"`

	PendingMessageMetadata *map[string]string `json:"pending_message_metadata"`
}

type UpdatePollOptionRequest struct {
	Id string `json:"id"`

	Text string `json:"text"`

	UserId *string `json:"user_id"`

	Custom *map[string]any `json:"Custom"`

	User *UserRequest `json:"user"`
}

type UpdatePollPartialRequest struct {
	UserId *string `json:"user_id"`

	Unset *[]string `json:"unset"`

	Set *map[string]any `json:"set"`

	User *UserRequest `json:"user"`
}

type UpdatePollRequest struct {
	Id string `json:"id"`

	Name string `json:"name"`

	AllowAnswers *bool `json:"allow_answers"`

	AllowUserSuggestedOptions *bool `json:"allow_user_suggested_options"`

	Description *string `json:"description"`

	EnforceUniqueVote *bool `json:"enforce_unique_vote"`

	IsClosed *bool `json:"is_closed"`

	MaxVotesAllowed *int `json:"max_votes_allowed"`

	UserId *string `json:"user_id"`

	VotingVisibility *string `json:"voting_visibility"`

	Options *[]*PollOption `json:"options"`

	Custom *map[string]any `json:"Custom"`

	User *UserRequest `json:"user"`
}

type UpdateThreadPartialRequest struct {
	UserId *string `json:"user_id"`

	Unset *[]string `json:"unset"`

	Set *map[string]any `json:"set"`

	User *UserRequest `json:"user"`
}

type UpdateThreadPartialResponse struct {
	Duration string `json:"duration"`

	Thread ThreadResponse `json:"thread"`
}

type UpdateUserPartialRequest struct {
	Id string `json:"id"`

	Unset *[]string `json:"unset"`

	Set *map[string]any `json:"set"`
}

type UpdateUserPermissionsRequest struct {
	UserId string `json:"user_id"`

	GrantPermissions *[]string `json:"grant_permissions"`

	RevokePermissions *[]string `json:"revoke_permissions"`
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

	MembershipDeletionTaskId string `json:"membership_deletion_task_id"`

	Users map[string]FullUserResponse `json:"users"`
}

type UpsertPushProviderRequest struct {
	PushProvider *PushProvider `json:"push_provider"`
}

type UpsertPushProviderResponse struct {
	Duration string `json:"duration"`

	PushProvider PushProviderResponse `json:"push_provider"`
}

type UserCustomEventRequest struct {
	Type string `json:"type"`

	Custom *map[string]any `json:"custom"`
}

type UserInfoResponse struct {
	Image string `json:"image"`

	Name string `json:"name"`

	Roles []string `json:"roles"`

	Custom map[string]any `json:"custom"`
}

type UserMute struct {
	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

	Expires *time.Time `json:"expires"`

	Target *UserObject `json:"target"`

	User *UserObject `json:"user"`
}

type UserObject struct {
	Banned bool `json:"banned"`

	Id string `json:"id"`

	Online bool `json:"online"`

	Role string `json:"role"`

	Custom map[string]any `json:"custom"`

	BanExpires *time.Time `json:"ban_expires"`

	CreatedAt *time.Time `json:"created_at"`

	DeactivatedAt *time.Time `json:"deactivated_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	Invisible *bool `json:"invisible"`

	Language *string `json:"language"`

	LastActive *time.Time `json:"last_active"`

	RevokeTokensIssuedBefore *time.Time `json:"revoke_tokens_issued_before"`

	UpdatedAt *time.Time `json:"updated_at"`

	Teams *[]string `json:"teams"`

	PrivacySettings *PrivacySettings `json:"privacy_settings"`

	PushNotifications *PushNotificationSettings `json:"push_notifications"`
}

type UserRequest struct {
	Id string `json:"id"`

	Image *string `json:"image"`

	Invisible *bool `json:"invisible"`

	Language *string `json:"language"`

	Name *string `json:"name"`

	Role *string `json:"role"`

	Teams *[]string `json:"teams"`

	Custom *map[string]any `json:"custom"`

	PrivacySettings *PrivacySettings `json:"privacy_settings"`

	PushNotifications *PushNotificationSettingsInput `json:"push_notifications"`
}

type UserResponse struct {
	Banned bool `json:"banned"`

	CreatedAt time.Time `json:"created_at"`

	Id string `json:"id"`

	Invisible bool `json:"invisible"`

	Language string `json:"language"`

	Online bool `json:"online"`

	Role string `json:"role"`

	ShadowBanned bool `json:"shadow_banned"`

	UpdatedAt time.Time `json:"updated_at"`

	Devices []*Device `json:"devices"`

	Teams []string `json:"teams"`

	Custom map[string]any `json:"custom"`

	DeactivatedAt *time.Time `json:"deactivated_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	Image *string `json:"image"`

	LastActive *time.Time `json:"last_active"`

	Name *string `json:"name"`

	RevokeTokensIssuedBefore *time.Time `json:"revoke_tokens_issued_before"`

	PushNotifications *PushNotificationSettings `json:"push_notifications"`
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

	SessionId string `json:"session_id"`

	TotalPixelsIn int `json:"total_pixels_in"`

	TotalPixelsOut int `json:"total_pixels_out"`

	Browser *string `json:"browser"`

	BrowserVersion *string `json:"browser_version"`

	CurrentIp *string `json:"current_ip"`

	CurrentSfu *string `json:"current_sfu"`

	DeviceModel *string `json:"device_model"`

	DeviceVersion *string `json:"device_version"`

	DistanceToSfuKilometers *float64 `json:"distance_to_sfu_kilometers"`

	MaxFirPerSecond *float64 `json:"max_fir_per_second"`

	MaxFreezesPerSecond *float64 `json:"max_freezes_per_second"`

	MaxNackPerSecond *float64 `json:"max_nack_per_second"`

	MaxPliPerSecond *float64 `json:"max_pli_per_second"`

	Os *string `json:"os"`

	OsVersion *string `json:"os_version"`

	PublisherNoiseCancellationSeconds *float64 `json:"publisher_noise_cancellation_seconds"`

	PublisherQualityLimitationFraction *float64 `json:"publisher_quality_limitation_fraction"`

	PublishingAudioCodec *string `json:"publishing_audio_codec"`

	PublishingVideoCodec *string `json:"publishing_video_codec"`

	ReceivingAudioCodec *string `json:"receiving_audio_codec"`

	ReceivingVideoCodec *string `json:"receiving_video_codec"`

	Sdk *string `json:"sdk"`

	SdkVersion *string `json:"sdk_version"`

	SubscriberVideoQualityThrottledDurationSeconds *float64 `json:"subscriber_video_quality_throttled_duration_seconds"`

	WebrtcVersion *string `json:"webrtc_version"`

	PublishedTracks *[]PublishedTrackInfo `json:"published_tracks"`

	Subsessions *[]*Subsession `json:"subsessions"`

	Geolocation *GeolocationResult `json:"geolocation"`

	Jitter *Stats `json:"jitter"`

	Latency *Stats `json:"latency"`

	MaxPublishingVideoQuality *VideoQuality `json:"max_publishing_video_quality"`

	MaxReceivingVideoQuality *VideoQuality `json:"max_receiving_video_quality"`

	PubSubHints *MediaPubSubHint `json:"pub_sub_hints"`

	PublisherAudioMos *MOSStats `json:"publisher_audio_mos"`

	PublisherJitter *Stats `json:"publisher_jitter"`

	PublisherLatency *Stats `json:"publisher_latency"`

	PublisherVideoQualityLimitationDurationSeconds *map[string]float64 `json:"publisher_video_quality_limitation_duration_seconds"`

	SubscriberAudioMos *MOSStats `json:"subscriber_audio_mos"`

	SubscriberJitter *Stats `json:"subscriber_jitter"`

	SubscriberLatency *Stats `json:"subscriber_latency"`

	Timeline *CallTimeline `json:"timeline"`
}

type UserStats struct {
	MinEventTs int `json:"min_event_ts"`

	SessionStats []UserSessionStats `json:"session_stats"`

	Info UserInfoResponse `json:"info"`

	Rating *int `json:"rating"`
}

type VideoQuality struct {
	UsageType *string `json:"usage_type"`

	Resolution *VideoResolution `json:"resolution"`
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
	AccessRequestEnabled *bool `json:"access_request_enabled"`

	CameraDefaultOn *bool `json:"camera_default_on"`

	CameraFacing *string `json:"camera_facing"`

	Enabled *bool `json:"enabled"`

	TargetResolution *TargetResolution `json:"target_resolution"`
}

type VideoSettingsResponse struct {
	AccessRequestEnabled bool `json:"access_request_enabled"`

	CameraDefaultOn bool `json:"camera_default_on"`

	CameraFacing string `json:"camera_facing"`

	Enabled bool `json:"enabled"`

	TargetResolution TargetResolution `json:"target_resolution"`
}

type VoteData struct {
	AnswerText *string `json:"answer_text"`

	OptionId *string `json:"option_id"`

	Option *PollOption `json:"Option"`
}

type WSEvent struct {
	CreatedAt time.Time `json:"created_at"`

	Type string `json:"type"`

	Custom map[string]any `json:"custom"`

	Automoderation *bool `json:"automoderation"`

	ChannelId *string `json:"channel_id"`

	ChannelType *string `json:"channel_type"`

	Cid *string `json:"cid"`

	ConnectionId *string `json:"connection_id"`

	ParentId *string `json:"parent_id"`

	Reason *string `json:"reason"`

	Team *string `json:"team"`

	UserId *string `json:"user_id"`

	WatcherCount *int `json:"watcher_count"`

	AutomoderationScores *ModerationResponse `json:"automoderation_scores"`

	Channel *ChannelResponse `json:"channel"`

	CreatedBy *UserObject `json:"created_by"`

	Me *OwnUser `json:"me"`

	Member *ChannelMember `json:"member"`

	Message *Message `json:"message"`

	MessageUpdate *MessageUpdate `json:"message_update"`

	Poll *Poll `json:"poll"`

	PollVote *PollVote `json:"poll_vote"`

	Reaction *Reaction `json:"reaction"`

	Thread *Thread `json:"thread"`

	User *UserObject `json:"user"`
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
	Disabled *bool `json:"Disabled"`

	PackageName *string `json:"package_name"`

	Secret *string `json:"secret"`
}

type XiaomiConfigFields struct {
	Enabled bool `json:"enabled"`

	PackageName *string `json:"package_name"`

	Secret *string `json:"secret"`
}
