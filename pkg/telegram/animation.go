package telegram

// Animation
// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileID       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width"`          // Video width as defined by sender
	Height       int        `json:"height"`         // Video height as defined by sender
	Duration     int        `json:"duration"`       // Duration of the video in seconds as defined by sender
	Thumb        *PhotoSize `json:"thumb"`          // Optional. Animation thumbnail as defined by sender
	FileName     string     `json:"file_name"`      // Optional. Original animation filename as defined by sender
	MimeType     string     `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int        `json:"file_size"`      // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}
