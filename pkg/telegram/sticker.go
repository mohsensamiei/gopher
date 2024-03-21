package telegram

// Sticker
// This object represents a sticker.
type Sticker struct {
	FileID           string        `json:"file_id"`           // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID     string        `json:"file_unique_id"`    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Type             string        `json:"type"`              // Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
	Width            int           `json:"width"`             // Sticker width
	Height           int           `json:"height"`            // Sticker height
	IsAnimated       bool          `json:"is_animated"`       // True, if the sticker is animated
	IsVideo          bool          `json:"is_video"`          // True, if the sticker is a video sticker
	Thumb            *PhotoSize    `json:"thumb"`             // Optional. Sticker thumbnail in the .WEBP or .JPG format
	Emoji            string        `json:"emoji"`             // Optional. Emoji associated with the sticker
	SetName          string        `json:"set_name"`          // Optional. Name of the sticker set to which the sticker belongs
	PremiumAnimation *File         `json:"premium_animation"` // Optional. For premium regular stickers, premium animation for the sticker
	MaskPosition     *MaskPosition `json:"mask_position"`     // Optional. For mask stickers, the position where the mask should be placed
	CustomEmojiID    string        `json:"custom_emoji_id"`   // Optional. For custom emoji stickers, unique identifier of the custom emoji
	FileSize         int           `json:"file_size"`         // Optional. File size in bytes
}
