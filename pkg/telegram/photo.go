package telegram

// PhotoSize
// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int    `json:"width"`          // Photo width
	Height       int    `json:"height"`         // Photo height
	FileSize     int    `json:"file_size"`      // Optional. File size in bytes
}
