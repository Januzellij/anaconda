package anaconda

import (
	"strings"
	"time"
)

type Tweet struct {
	Contributors         []int64     `json:"contributors"`
	Coordinates          interface{} `json:"coordinates"`
	CreatedAt            string      `json:"created_at"`
	Entities             Entities    `json:"entities"`
	FavoriteCount        int         `json:"favorite_count"`
	Favorited            bool        `json:"favorited"`
	Geo                  interface{} `json:"geo"`
	Id                   int64       `json:"id"`
	IdStr                string      `json:"id_str"`
	InReplyToScreenName  string      `json:"in_reply_to_screen_name"`
	InReplyToStatusID    int64       `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr string      `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64       `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   string      `json:"in_reply_to_user_id_str"`
	Place                Place       `json:"place"`
	PossiblySensitive    bool        `json:"possibly_sensitive"`
	RetweetCount         int         `json:"retweet_count"`
	Retweeted            bool        `json:"retweeted"`
	RetweetedStatus      *Tweet      `json:"retweeted_status"`
	Source               string      `json:"source"`
	Text                 string      `json:"text"`
	Truncated            bool        `json:"truncated"`
	User                 User        `json:"user"`
}

// CreatedAtTime is a convenience wrapper that returns the Created_at time, parsed as a time.Time struct
func (t Tweet) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.CreatedAt)
}

// this is neccesary because there is *one minute difference* between Twitter API and archive JSON
// See: https://github.com/Januzellij/anaconda/commit/9c9e6a07f65ea957b3107bee70c33c4f771e6577
func ConvertToArchive(tweets []Tweet) []ArchiveTweet {
	archiveTweets := make([]ArchiveTweet, len(tweets))
	for i, tweet := range tweets {
		archiveTweet := ArchiveTweet{
			tweet.Contributors,
			tweet.Coordinates,
			tweet.CreatedAt,
			ArchiveEntities{
				tweet.Entities.Hashtags,
				escapeEntityUrls(tweet.Entities.Urls),
				tweet.Entities.User_mentions,
				convertMediaToArchive(tweet.Entities.Media),
			},
			tweet.FavoriteCount,
			tweet.Favorited,
			tweet.Geo,
			tweet.Id,
			tweet.IdStr,
			tweet.InReplyToScreenName,
			tweet.InReplyToStatusID,
			tweet.InReplyToStatusIdStr,
			tweet.InReplyToUserID,
			tweet.InReplyToUserIdStr,
			tweet.Place,
			tweet.PossiblySensitive,
			tweet.RetweetCount,
			tweet.Retweeted,
			nil,
			tweet.Source,
			escapeUrl(tweet.Text),
			tweet.Truncated,
			escapeUserUrls(tweet.User),
		}
		if tweet.RetweetedStatus != nil {
			archiveTweet.RetweetedStatus = &ArchiveTweet{
				tweet.RetweetedStatus.Contributors,
				tweet.RetweetedStatus.Coordinates,
				tweet.RetweetedStatus.CreatedAt,
				ArchiveEntities{
					tweet.RetweetedStatus.Entities.Hashtags,
					escapeEntityUrls(tweet.RetweetedStatus.Entities.Urls),
					tweet.RetweetedStatus.Entities.User_mentions,
					convertMediaToArchive(tweet.RetweetedStatus.Entities.Media),
				},
				tweet.RetweetedStatus.FavoriteCount,
				tweet.RetweetedStatus.Favorited,
				tweet.RetweetedStatus.Geo,
				tweet.RetweetedStatus.Id,
				tweet.RetweetedStatus.IdStr,
				tweet.RetweetedStatus.InReplyToScreenName,
				tweet.RetweetedStatus.InReplyToStatusID,
				tweet.RetweetedStatus.InReplyToStatusIdStr,
				tweet.RetweetedStatus.InReplyToUserID,
				tweet.RetweetedStatus.InReplyToUserIdStr,
				tweet.RetweetedStatus.Place,
				tweet.RetweetedStatus.PossiblySensitive,
				tweet.RetweetedStatus.RetweetCount,
				tweet.RetweetedStatus.Retweeted,
				nil, // TODO: fix this, so it recursively fills in all retweets
				tweet.RetweetedStatus.Source,
				escapeUrl(tweet.RetweetedStatus.Text),
				tweet.RetweetedStatus.Truncated,
				escapeUserUrls(tweet.RetweetedStatus.User),
			}
		}
		archiveTweets[i] = archiveTweet
	}
	return archiveTweets
}

// yes, all of these anonymous structs are ugly. The anaconda authors chose to make the Entities struct
// slices slices of anonymous structs (that's quite a mouthful), rather than giving a name to each struct
// See: https://github.com/ChimeraCoder/anaconda/blob/master/twitter_entities.go
func convertMediaToArchive(media []struct {
	Id              int64
	Id_str          string
	Media_url       string
	Media_url_https string
	Url             string
	Display_url     string
	Expanded_url    string
	Sizes           MediaSizes
	Type            string
	Indices         []int
}) []struct {
	Id              int64
	Id_str          string
	Media_url       string
	Media_url_https string
	Url             string
	Display_url     string
	Expanded_url    string
	Sizes           []MediaSize
	Type            string
	Indices         []int
} {
	archiveMedia := make([]struct {
		Id              int64
		Id_str          string
		Media_url       string
		Media_url_https string
		Url             string
		Display_url     string
		Expanded_url    string
		Sizes           []MediaSize
		Type            string
		Indices         []int
	}, len(media))
	for i, v := range media {
		archiveMedia[i] = struct {
			Id              int64
			Id_str          string
			Media_url       string
			Media_url_https string
			Url             string
			Display_url     string
			Expanded_url    string
			Sizes           []MediaSize
			Type            string
			Indices         []int
		}{
			v.Id,
			v.Id_str,
			escapeUrl(v.Media_url),
			escapeUrl(v.Media_url_https),
			escapeUrl(v.Url),
			escapeUrl(v.Display_url),
			escapeUrl(v.Expanded_url),
			[]MediaSize{v.Sizes.Medium, v.Sizes.Thumb, v.Sizes.Small, v.Sizes.Large},
			v.Type,
			v.Indices,
		}
	}
	return archiveMedia
}

func escapeUrl(url string) string {
	return strings.Replace(url, "/", `\/`, -1)
}

func escapeEntityUrls(entities []struct {
	Indices      []int
	Url          string
	Display_url  string
	Expanded_url string
}) []struct {
	Indices      []int
	Url          string
	Display_url  string
	Expanded_url string
} {
	for i, v := range entities {
		v.Url = escapeUrl(v.Url)
		v.Display_url = escapeUrl(v.Display_url)
		v.Expanded_url = escapeUrl(v.Expanded_url)
		entities[i] = v
	}
	return entities
}

func escapeUserUrls(u User) User {
	u.ProfileBackgroundImageURL = escapeUrl(u.ProfileBackgroundImageURL)
	u.ProfileBackgroundImageUrlHttps = escapeUrl(u.ProfileBackgroundImageUrlHttps)
	u.ProfileImageURL = escapeUrl(u.ProfileImageURL)
	u.ProfileImageUrlHttps = escapeUrl(u.ProfileImageUrlHttps)
	u.URL = escapeUrl(u.URL)
	return u
}
