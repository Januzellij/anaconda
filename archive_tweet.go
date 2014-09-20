package anaconda

import "time"

type ArchiveTweet struct {
	Contributors         []int64          `json:"contributors,omitempty"`
	Coordinates          interface{}      `json:"coordinates,omitempty"`
	CreatedAt            string           `json:"created_at,omitempty"`
	Entities             *ArchiveEntities `json:"entities,omitempty"`
	FavoriteCount        int              `json:"favorite_count,omitempty"`
	Favorited            bool             `json:"favorited"`
	Geo                  interface{}      `json:"geo,omitempty"`
	Id                   int64            `json:"id,omitempty"`
	IdStr                string           `json:"id_str,omitempty"`
	InReplyToScreenName  string           `json:"in_reply_to_screen_name,omitempty"`
	InReplyToStatusID    int64            `json:"in_reply_to_status_id,omitempty"`
	InReplyToStatusIdStr string           `json:"in_reply_to_status_id_str,omitempty"`
	InReplyToUserID      int64            `json:"in_reply_to_user_id,omitempty"`
	InReplyToUserIdStr   string           `json:"in_reply_to_user_id_str,omitempty"`
	Place                *Place           `json:"place,omitempty"`
	PossiblySensitive    bool             `json:"possibly_sensitive"`
	RetweetCount         int              `json:"retweet_count,omitempty"`
	Retweeted            bool             `json:"retweeted"`
	RetweetedStatus      *ArchiveTweet    `json:"retweeted_status,omitempty"`
	Source               string           `json:"source,omitempty"`
	Text                 string           `json:"text,omitempty"`
	Truncated            bool             `json:"truncated"`
	User                 User             `json:"user,omitempty"` // User isn't a pointer or ",omitempty" because it's guarenteed to be there
}

// CreatedAtTime is a convenience wrapper that returns the Created_at time, parsed as a time.Time struct
func (t ArchiveTweet) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.CreatedAt)
}
