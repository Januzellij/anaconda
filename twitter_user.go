package anaconda

type User struct {
	ContributorsEnabled            bool   `json:"contributors_enabled,omitempty"`
	CreatedAt                      string `json:"created_at,omitempty"`
	DefaultProfile                 bool   `json:"default_profile,omitempty"`
	DefaultProfileImage            bool   `json:"default_profile_image,omitempty"`
	Description                    string `json:"description,omitempty"`
	FavouritesCount                int    `json:"favourites_count,omitempty"`
	FollowRequestSent              bool   `json:"follow_request_sent,omitempty"`
	FollowersCount                 int    `json:"followers_count,omitempty"`
	Following                      bool   `json:"following,omitempty"`
	FriendsCount                   int    `json:"friends_count,omitempty"`
	GeoEnabled                     bool   `json:"geo_enabled,omitempty"`
	Id                             int64  `json:"id,omitempty"`
	IdStr                          string `json:"id_str,omitempty"`
	IsTranslator                   bool   `json:"is_translator,omitempty"`
	Lang                           string `json:"lang,omitempty"`
	ListedCount                    int64  `json:"listed_count,omitempty"`
	Location                       string `json:"location,omitempty"`
	Name                           string `json:"name,omitempty"`
	Notifications                  bool   `json:"notifications,omitempty"`
	ProfileBackgroundColor         string `json:"profile_background_color,omitempty"`
	ProfileBackgroundImageURL      string `json:"profile_background_image_url,omitempty"`
	ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https,omitempty"`
	ProfileBackgroundTile          bool   `json:"profile_background_tile,omitempty"`
	ProfileImageURL                string `json:"profile_image_url,omitempty"`
	ProfileImageUrlHttps           string `json:"profile_image_url_https,omitempty"`
	ProfileLinkColor               string `json:"profile_link_color,omitempty"`
	ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color,omitempty"`
	ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color,omitempty"`
	ProfileTextColor               string `json:"profile_text_color,omitempty"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image,omitempty"`
	Protected                      bool   `json:"protected,omitempty"`
	ScreenName                     string `json:"screen_name,omitempty"`
	ShowAllInlineMedia             bool   `json:"show_all_inline_media,omitempty"`
	Status                         *Tweet `json:"status,omitempty"` // Only included if the user is a friend
	StatusesCount                  int64  `json:"statuses_count,omitempty"`
	TimeZone                       string `json:"time_zone,omitempty"`
	URL                            string `json:"url,omitempty"`
	UtcOffset                      int    `json:"utc_offset,omitempty"`
	Verified                       bool   `json:"verified,omitempty"`
}
