package telegram

// GetMe
// A simple method for testing your bot's authentication token.
// Requires no parameters.
// Returns basic information about the bot in form of a User object.
func (c Connection) GetMe() (*User, error) {
	var res Response[User]
	if err := request(c.TelegramToken, getMe, nil, &res); err != nil {
		return nil, err
	}
	return &res.Result, nil
}
