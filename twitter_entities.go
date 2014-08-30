package anaconda

type Entities struct {
	Hashtags []struct {
		Indices []int  `json:"indices"`
		Text    string `json:"text"`
	}
	Urls []struct {
		Indices      []int  `json:"indices"`
		Url          string `json:"url"`
		Display_url  string `json:"display_url"`
		Expanded_url string `json:"expanded_url"`
	}
	User_mentions []struct {
		Name        string `json:"name"`
		Indices     []int  `json:"indices"`
		Screen_name string `json:"screen_name"`
		Id          int64  `json:"id"`
		Id_str      string `json:"id_str"`
	}
	Media []struct {
		Id              int64      `json:"id"`
		Id_str          string     `json:"id_str"`
		Media_url       string     `json:"media_url"`
		Media_url_https string     `json:"media_url_https"`
		Url             string     `json:"url"`
		Display_url     string     `json:"display_url"`
		Expanded_url    string     `json:"expanded_url"`
		Sizes           MediaSizes `json:"sizes"`
		Type            string     `json:"type"`
		Indices         []int      `json:"indices"`
	}
}

type MediaSizes struct {
	Medium MediaSize `json:"medium"`
	Thumb  MediaSize `json:"thumb"`
	Small  MediaSize `json:"small"`
	Large  MediaSize `json:"large"`
}

type MediaSize struct {
	W      int    `json:"w"`
	H      int    `json:"h"`
	Resize string `json:"resize"`
}
