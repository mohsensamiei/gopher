package telegram

type SetWebhook struct {
	URL                string       `json:"url"`                            // HTTPS URL to send updates to. Use an empty string to remove webhook integration
	Certificate        *InputFile   `json:"certificate,omitempty"`          // Optional. Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	IpAddress          string       `json:"ip_address,omitempty"`           // Optional. The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	MaxConnections     int          `json:"max_connections,omitempty"`      // Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	AllowedUpdates     []UpdateType `json:"allowed_updates,omitempty"`      // Optional. A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	DropPendingUpdates bool         `json:"drop_pending_updates,omitempty"` // Optional. Pass True to drop all pending updates
	SecretToken        string       `json:"secret_token,omitempty"`         // Optional. A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
}

// SetWebhook
// Use this method to specify a URL and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request (a request with response HTTP status code different from 2XY), we will repeat the request and give up after a reasonable amount of attempts. Returns True on success.
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.
func (c Connection) SetWebhook(req SetWebhook) (bool, error) {
	var res Response[bool]
	if err := request(c.TelegramToken, setWebhook, req, &res); err != nil {
		return false, err
	}
	return res.Result, nil
}

type DeleteWebhook struct {
	DropPendingUpdates bool `json:"drop_pending_updates"` // Optional. Pass True to drop all pending updates
}

// DeleteWebhook
// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
func (c Connection) DeleteWebhook(req DeleteWebhook) (bool, error) {
	var res Response[bool]
	if err := request(c.TelegramToken, deleteWebhook, req, &res); err != nil {
		return false, err
	}
	return res.Result, nil
}
