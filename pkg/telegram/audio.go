package telegram

// Audio
// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID       string   `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string   `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     Duration `json:"duration"`       // Duration of the audio in seconds as defined by sender
	Performer    string   `json:"performer"`      // Optional. Performer of the audio as defined by sender or by audio tags
	Title        string   `json:"title"`          // Optional. Title of the audio as defined by sender or by audio tags
	FileName     string   `json:"file_name"`      // Optional. Original filename as defined by sender
	MimeType     string   `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int      `json:"file_size"`      // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}
