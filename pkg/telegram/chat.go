package telegram

import (
	"encoding/json"
)

// Chat
// This object represents a chat.
type Chat struct {
	ID                                 int              `json:"id"`                                      // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type                               string           `json:"type"`                                    // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title                              string           `json:"title"`                                   // Optional. Title, for supergroups, channels and group chats
	Username                           string           `json:"username"`                                // Optional. Username, for private chats, supergroups and channels if available
	FirstName                          string           `json:"first_name"`                              // Optional. First name of the other party in a private chat
	LastName                           string           `json:"last_name"`                               // Optional. Last name of the other party in a private chat
	Photo                              *ChatPhoto       `json:"ChatPhoto"`                               // Optional. Chat photo. Returned only in getChat.
	Bio                                string           `json:"bio"`                                     // Optional. Bio of the other party in a private chat. Returned only in getChat.
	HasPrivateForwards                 bool             `json:"has_private_forwards"`                    // Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	HasRestrictedVoiceAndVideoMessages bool             `json:"has_restricted_voice_and_video_messages"` // Optional. True, if the privacy settings of the other party restrict sending voice and video note messages in the private chat. Returned only in getChat.
	JoinToSendMessages                 bool             `json:"join_to_send_messages"`                   // Optional. True, if users need to join the supergroup before they can send messages. Returned only in getChat.
	JoinByRequest                      bool             `json:"join_by_request"`                         // Optional. True, if all users directly joining the supergroup need to be approved by supergroup administrators. Returned only in getChat.
	Description                        string           `json:"description"`                             // Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	InviteLink                         string           `json:"invite_link"`                             // Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
	PinnedMessage                      *Message         `json:"pinned_message"`                          // Optional. The most recent pinned message (by sending date). Returned only in getChat.
	Permissions                        *ChatPermissions `json:"permissions"`                             // Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	SlowModeDelay                      int              `json:"slow_mode_delay"`                         // Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user; in seconds. Returned only in getChat.
	MessageAutoDeleteTime              int              `json:"message_auto_delete_time"`                // Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
	HasProtectedContent                bool             `json:"has_protected_content"`                   // Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
	StickerSetName                     string           `json:"sticker_set_name"`                        // Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet                   bool             `json:"can_set_sticker_set"`                     // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	LinkedChatID                       int              `json:"linked_chat_id"`                          // Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
	Location                           *ChatLocation    `json:"location"`                                // Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
}

// ChatPhoto
// This object represents a chat photo.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`        // File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileUniqueID string `json:"small_file_unique_id"` // Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigFileID         string `json:"big_file_id"`          // File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileUniqueID   string `json:"big_file_unique_id"`   // Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
}

// ChatPermissions
// Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`         // Optional. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMediaMessages  bool `json:"can_send_media_messages"`   // Optional. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendPolls          bool `json:"can_send_polls"`            // Optional. True, if the user is allowed to send polls, implies can_send_messages
	CanSendOtherMessages  bool `json:"can_send_other_messages"`   // Optional. True, if the user is allowed to send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"` // Optional. True, if the user is allowed to add web page previews to their messages, implies can_send_media_messages
	CanChangeInfo         bool `json:"can_change_info"`           // Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	CanInviteUsers        bool `json:"can_invite_users"`          // Optional. True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool `json:"can_pin_messages"`          // Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
}

// ChatLocation
// Represents a location to which a chat is connected.
type ChatLocation struct {
	Location Location `json:"location"` // The location to which the supergroup is connected. Can't be a live location.
	Address  string   `json:"address"`  // Location address; 1-64 characters, as defined by the chat owner
}

// ChatMemberUpdated
// This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat          Chat            `json:"chat"`            // Chat the user belongs to
	From          User            `json:"from"`            // Performer of the action, which resulted in the change
	Date          int             `json:"date"`            // Date the change was done in Unix time
	OldChatMember ChatMember      `json:"old_chat_member"` // Previous information about the chat member
	NewChatMember ChatMember      `json:"new_chat_member"` // New information about the chat member
	InviteLink    *ChatInviteLink `json:"invite_link"`     // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
}

// ChatInviteLink
// Represents an invite link for a chat.
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`                // The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	Creator                 User   `json:"creator"`                    // Creator of the link
	CreatesJoinRequest      bool   `json:"creates_join_request"`       // True, if users joining the chat via the link need to be approved by chat administrators
	IsPrimary               bool   `json:"is_primary"`                 // True, if the link is primary
	IsRevoked               bool   `json:"is_revoked"`                 // True, if the link is revoked
	Name                    string `json:"name"`                       // Optional. Invite link name
	ExpireDate              int    `json:"expire_date"`                // Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	MemberLimit             int    `json:"member_limit"`               // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	PendingJoinRequestCount int    `json:"pending_join_request_count"` // Optional. Number of pending join requests created using this link
}

// ChatJoinRequest
// Represents a join request sent to a chat.
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`        // Chat to which the request was sent
	From       User            `json:"from"`        // User that sent the join request
	Date       int             `json:"date"`        // Date the request was sent in Unix time
	Bio        string          `json:"bio"`         // Optional. Bio of the user.
	InviteLink *ChatInviteLink `json:"invite_link"` // Optional. Chat invite link that was used by the user to send the join request
}

// ChatMember
// This object contains information about one member of a chat. Currently, the following 6 types of chat members are supported:
// ChatMemberOwner
// ChatMemberAdministrator
// ChatMemberMember
// ChatMemberRestricted
// ChatMemberLeft
// ChatMemberBanned
type ChatMember struct {
	Owner         *ChatMemberOwner         `json:"-"`
	Administrator *ChatMemberAdministrator `json:"-"`
	Member        *ChatMemberMember        `json:"-"`
	Restricted    *ChatMemberRestricted    `json:"-"`
	Left          *ChatMemberLeft          `json:"-"`
	Banned        *ChatMemberBanned        `json:"-"`
}

type chatMember struct {
	Status string `json:"status"`
}

func (c *ChatMember) UnmarshalJSON(b []byte) error {
	cm := new(chatMember)
	if err := json.Unmarshal(b, cm); err != nil {
		return err
	}
	switch cm.Status {
	case "creator":
		c.Owner = new(ChatMemberOwner)
		if err := json.Unmarshal(b, c.Owner); err != nil {
			return err
		}
	case "administrator":
		c.Administrator = new(ChatMemberAdministrator)
		if err := json.Unmarshal(b, c.Administrator); err != nil {
			return err
		}
	case "member":
		c.Member = new(ChatMemberMember)
		if err := json.Unmarshal(b, c.Member); err != nil {
			return err
		}
	case "restricted":
		c.Restricted = new(ChatMemberRestricted)
		if err := json.Unmarshal(b, c.Restricted); err != nil {
			return err
		}
	case "left":
		c.Left = new(ChatMemberLeft)
		if err := json.Unmarshal(b, c.Left); err != nil {
			return err
		}
	case "kicked":
		c.Banned = new(ChatMemberBanned)
		if err := json.Unmarshal(b, c.Banned); err != nil {
			return err
		}
	}
	return nil
}

// ChatMemberOwner
// Represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	Status      string `json:"status"`       // The member's status in the chat, always “creator”
	User        User   `json:"user"`         // Information about the user
	IsAnonymous bool   `json:"is_anonymous"` // True, if the user's presence in the chat is hidden
	CustomTitle string `json:"custom_title"` // Optional. Custom title for this user
}

// ChatMemberAdministrator
// Represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	Status              string `json:"status"`                 // The member's status in the chat, always “administrator”
	User                User   `json:"user"`                   // Information about the user
	CanBeEdited         bool   `json:"can_be_edited"`          // True, if the bot is allowed to edit administrator privileges of that user
	IsAnonymous         bool   `json:"is_anonymous"`           // True, if the user's presence in the chat is hidden
	CanManageChat       bool   `json:"can_manage_chat"`        // True, if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanDeleteMessages   bool   `json:"can_delete_messages"`    // True, if the administrator can delete messages of other users
	CanManageVideoChats bool   `json:"can_manage_video_chats"` // True, if the administrator can manage video chats
	CanRestrictMembers  bool   `json:"can_restrict_members"`   // True, if the administrator can restrict, ban or unban chat members
	CanPromoteMembers   bool   `json:"can_promote_members"`    // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool   `json:"can_change_info"`        // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users"`       // True, if the user is allowed to invite new users to the chat
	CanPostMessages     bool   `json:"can_post_messages"`      // Optional. True, if the administrator can post in the channel; channels only
	CanEditMessages     bool   `json:"can_edit_messages"`      // Optional. True, if the administrator can edit messages of other users and can pin messages; channels only
	CanPinMessages      bool   `json:"can_pin_messages"`       // Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CustomTitle         string `json:"custom_title"`           // Optional. Custom title for this user
}

// ChatMemberMember
// Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	Status string `json:"status"` // The member's status in the chat, always “member”
	User   User   `json:"user"`   // Information about the user
}

// ChatMemberRestricted
// Represents a chat member that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	Status                string `json:"status"`                    // The member's status in the chat, always “restricted”
	User                  User   `json:"user"`                      // Information about the user
	IsMember              bool   `json:"is_member"`                 // True, if the user is a member of the chat at the moment of the request
	CanChangeInfo         bool   `json:"can_change_info"`           // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers        bool   `json:"can_invite_users"`          // True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool   `json:"can_pin_messages"`          // True, if the user is allowed to pin messages
	CanSendMessages       bool   `json:"can_send_messages"`         // True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`   // True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes
	CanSendPolls          bool   `json:"can_send_polls"`            // True, if the user is allowed to send polls
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`   // True, if the user is allowed to send animations, games, stickers and use inline bots
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"` // True, if the user is allowed to add web page previews to their messages
	UntilDate             int    `json:"until_date"`                // Date when restrictions will be lifted for this user; unix time. If 0, then the user is restricted forever
}

// ChatMemberLeft
// Represents a chat member that isn't currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
	Status string `json:"status"` // The member's status in the chat, always “left”
	User   User   `json:"user"`   // Information about the user
}

// ChatMemberBanned
// Represents a chat member that was banned in the chat and can't return to the chat or view chat messages.
type ChatMemberBanned struct {
	Status    string `json:"status"`     // The member's status in the chat, always “kicked”
	User      User   `json:"user"`       // Information about the user
	UntilDate int    `json:"until_date"` // Date when restrictions will be lifted for this user; unix time. If 0, then the user is banned forever
}
