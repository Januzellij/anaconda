package anaconda

type ArchiveEntities struct {
	Hashtags []struct {
		Indices []int  "json:,omitempty"
		Text    string "json:,omitempty"
	}
	Urls []struct {
		Indices      []int  "json:,omitempty"
		Url          string "json:,omitempty"
		Display_url  string "json:,omitempty"
		Expanded_url string "json:,omitempty"
	}
	User_mentions []struct {
		Name        string "json:,omitempty"
		Indices     []int  "json:,omitempty"
		Screen_name string "json:,omitempty"
		Id          int64  "json:,omitempty"
		Id_str      string "json:,omitempty"
	}
	Media []struct {
		Id              int64       "json:,omitempty"
		Id_str          string      "json:,omitempty"
		Media_url       string      "json:,omitempty"
		Media_url_https string      "json:,omitempty"
		Url             string      "json:,omitempty"
		Display_url     string      "json:,omitempty"
		Expanded_url    string      "json:,omitempty"
		Sizes           []MediaSize "json:,omitempty"
		Type            string      "json:,omitempty"
		Indices         []int       "json:,omitempty"
	}
}
