package telegram

// Poll
// This object contains information about a poll.
type Poll struct {
	ID                    string          `json:"id"`                      // Unique poll identifier
	Question              string          `json:"question"`                // Poll question, 1-300 characters
	Options               []PollOption    `json:"options"`                 // List of poll options
	TotalVoterCount       int             `json:"total_voter_count"`       // Total number of users that voted in the poll
	IsClosed              bool            `json:"is_closed"`               // True, if the poll is closed
	IsAnonymous           bool            `json:"is_anonymous"`            // True, if the poll is anonymous
	Type                  string          `json:"type"`                    // type, currently can be “regular” or “quiz”
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"` // True, if the poll allows multiple answers
	CorrectOptionID       int             `json:"correct_option_id"`       // Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	Explanation           string          `json:"explanation"`             // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	ExplanationEntities   []MessageEntity `json:"explanation_entities"`    // Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	OpenPeriod            int             `json:"open_period"`             // Optional. Amount of time in seconds the poll will be active after creation
	CloseDate             int             `json:"close_date"`              // Optional. Point in time (Unix timestamp) when the poll will be automatically closed
}

// PollOption
// This object contains information about one answer option in a poll.
type PollOption struct {
	Text       string `json:"text"`        // Option text, 1-100 characters
	VoterCount int    `json:"voter_count"` // Number of users that voted for this option
}

// PollAnswer
// This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID    string `json:"poll_id"`    // Unique poll identifier
	User      User   `json:"user"`       // The user, who changed the answer to the poll
	OptionIDs []int  `json:"option_ids"` // 0-based identifiers of answer options, chosen by the user. May be empty if the user retracted their vote.
}
