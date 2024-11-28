package telegram

type ApproveChatJoinRequest struct {
	ChatID int64 `json:"chat_id"` //	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID int64 `json:"user_id"` //	Unique identifier of the target user
}

// ApproveChatJoinRequest
// Use this method to approve a chat join request.
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
// Returns True on success.
func (c Connection) ApproveChatJoinRequest(req ApproveChatJoinRequest) (bool, error) {
	var res Response[bool]
	if err := request(c.TelegramToken, approveChatJoinRequest, req, &res); err != nil {
		return false, err
	}
	return res.Result, nil
}

type DeclineChatJoinRequest struct {
	ChatID int64 `json:"chat_id"` //	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID int64 `json:"user_id"` //	Unique identifier of the target user
}

// DeclineChatJoinRequest
// Use this method to decline a chat join request.
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
// Returns True on success.
func (c Connection) DeclineChatJoinRequest(req DeclineChatJoinRequest) (bool, error) {
	var res Response[bool]
	if err := request(c.TelegramToken, declineChatJoinRequest, req, &res); err != nil {
		return false, err
	}
	return res.Result, nil
}
