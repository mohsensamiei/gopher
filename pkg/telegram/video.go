package telegram

// Video
// This object represents a video file.
type Video struct {
	FileID       string     `json:"file_id"`        // file_id			String		Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"` // file_unique_id	String		Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width"`          // width			Integer		Video width as defined by sender
	Height       int        `json:"height"`         // height			Integer		Video height as defined by sender
	Duration     int        `json:"duration"`       // duration			Integer		Duration of the video in seconds as defined by sender
	Thumb        *PhotoSize `json:"thumb"`          // thumb			PhotoSize	Optional. Video thumbnail
	FileName     string     `json:"file_name"`      // file_name		String		Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type"`      // mime_type		String		Optional. MIME type of the file as defined by sender
	FileSize     int        `json:"file_size"`      // file_size		Integer		Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}

// VideoNote
// This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileID       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int        `json:"length"`         // Video width and height (diameter of the video message) as defined by sender
	Duration     int        `json:"duration"`       // Duration of the video in seconds as defined by sender
	Thumb        *PhotoSize `json:"thumb"`          // Optional. Video thumbnail
	FileSize     int        `json:"file_size"`      // Optional. File size in bytes
}

// VideoChatScheduled
// This object represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	StartDate int `json:"start_date"` // Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
}

// VideoChatStarted
// This object represents a service message about a video chat started in the chat. Currently holds no information.
type VideoChatStarted struct {
}

// VideoChatEnded
// This object represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	Duration int `json:"duration"` // Video chat duration in seconds
}

// VideoChatParticipantsInvited
// This object represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"` // New members that were invited to the video chat
}
