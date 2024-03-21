package telegram

import (
	"time"
)

type GetUpdates struct {
	Offset         uint          `json:"offset,omitempty"`          // Optional. Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Limit          uint8         `json:"limit,omitempty"`           // Optional. Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Timeout        time.Duration `json:"timeout,omitempty"`         // Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []UpdateType  `json:"allowed_updates,omitempty"` // Optional. A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the request to the getUpdates, so unwanted updates may be received for a short period of time.
}

// GetUpdates
// Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
// Notes
// 1. This method will not work if an outgoing webhook is set up.
// 2. In order to avoid getting duplicate updates, recalculate offset after each server response.
func (c Connection) GetUpdates(req GetUpdates) ([]Update, error) {
	var res Response[[]Update]
	if err := request(c.TelegramToken, getUpdates, req, &res); err != nil {
		return nil, err
	}
	return res.Result, nil
}
