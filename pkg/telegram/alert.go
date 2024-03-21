package telegram

// ProximityAlertTriggered
// This object represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"` // User that triggered the alert
	Watcher  User `json:"watcher"`  // User that set the alert
	Distance int  `json:"distance"` // The distance between the users
}
