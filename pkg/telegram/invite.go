package telegram

type CreateChatInviteLink struct {
	ChatID             int64  `json:"chat_id"`                        //	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Name               string `json:"name,omitempty"`                 // Optional. Invite link name; 0-32 characters
	ExpireDate         Date   `json:"expire_date,omitempty"`          // Optional. Point in time (Unix timestamp) when the link will expire
	MemberLimit        int    `json:"member_limit,omitempty"`         // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"` // Optional. True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
}

// CreateChatInviteLink
// Use this method to create an additional invite link for a chat.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// The link can be revoked using the method revokeChatInviteLink.
// Returns the new invite link as ChatInviteLink object.
func (c Connection) CreateChatInviteLink(req CreateChatInviteLink) (*ChatInviteLink, error) {
	var res Response[ChatInviteLink]
	if err := request(c.TelegramToken, createChatInviteLink, req, &res); err != nil {
		return nil, err
	}
	return &res.Result, nil
}
