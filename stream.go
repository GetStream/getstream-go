package getstream

type Stream struct {
	*Client
	chat  *ChatClient
	video *VideoClient
}

func NewClientFromEnvVars(options ...ClientOption) (*Stream, error) {
	client, err := newClientFromEnvVars(options...)
	if err != nil {
		return nil, err
	}
	return &Stream{
		Client: client,
	}, nil
}

func NewClient(apiKey, apiSecret string, options ...ClientOption) (*Stream, error) {
	client, err := newClient(apiKey, apiSecret, options...)
	if err != nil {
		return nil, err
	}
	return &Stream{
		Client: client,
	}, nil
}

// CreateToken generates a token for a given user ID, with optional claims.
//
// Parameters:
// - userID (string): The unique identifier of the user for whom the token is being created.
// - claims (*Claims): A pointer to a Claims struct containing optional parameters.
//
// Returns:
// - (string): The generated JWT token.
// - (error): An error object if token creation fails.
//
// Example:
//
// expiration:= 3600, // Token expires in 1 hour
//	claims := &Claims{
//	   
//	    Role:       "admin",
//	    ChannelCIDs: []string{"channel1", "channel2"},
//	}
//
// token, err := client.CreateToken("userID", claims, 3600)
func (s *Stream) CreateToken(userID string, claims *Claims, expiration int64) (string, error) {
	return s.createToken(userID, claims, expiration)
}

// CreateCallToken generates a token for a given user ID, including optional claims specific to calls.
//
// Parameters:
// - userID (string): The unique identifier of the user for whom the token is being created.
// - claims (*Claims): A pointer to a Claims struct containing optional parameters.
//
// Returns:
// - (string): The generated JWT token.
// - (error): An error object if token creation fails.
//
// Example:
//
//	claims := &Claims{
//	    Role:       "moderator",
//	    CallCIDs:   []string{"call1", "call2"},
//	}
//
//	expiration:= 7200, // Token expires in 2 hours
// token, err := client.CreateCallToken("userID", claims, expiration)
func (s *Stream) CreateCallToken(userID string, claims *Claims, expiration int64) (string, error) {
	return s.createCallToken(userID, claims, expiration)
}

func (s *Stream) Chat() *ChatClient {
	if s.chat == nil {
		s.chat = NewChatClient(s.Client)
	}
	return s.chat
}

func (s *Stream) Video() *VideoClient {
	if s.video == nil {
		s.video = NewVideoClient(s.Client)
	}
	return s.video
}
