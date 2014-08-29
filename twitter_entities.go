package anaconda

type Entities struct {
	Hashtags []struct {
		Indices []int  `json:",omitempty"`
		Text    string `json:",omitempty"`
	}
	Urls []struct {
		Indices      []int  `json:",omitempty"`
		Url          string `json:",omitempty"`
		Display_url  string `json:",omitempty"`
		Expanded_url string `json:",omitempty"`
	}
	User_mentions []struct {
		Name        string `json:",omitempty"`
		Indices     []int  `json:",omitempty"`
		Screen_name string `json:",omitempty"`
		Id          int64  `json:",omitempty"`
		Id_str      string `json:",omitempty"`
	}
	Media []struct {
		Id              int64      `json:",omitempty"`
		Id_str          string     `json:",omitempty"`
		Media_url       string     `json:",omitempty"`
		Media_url_https string     `json:",omitempty"`
		Url             string     `json:",omitempty"`
		Display_url     string     `json:",omitempty"`
		Expanded_url    string     `json:",omitempty"`
		Sizes           MediaSizes `json:",omitempty"`
		Type            string     `json:",omitempty"`
		Indices         []int      `json:",omitempty"`
	}
}

type MediaSizes struct {
	Medium MediaSize `json:",omitempty"`
	Thumb  MediaSize `json:",omitempty"`
	Small  MediaSize `json:",omitempty"`
	Large  MediaSize `json:",omitempty"`
}

type MediaSize struct {
	W      int    `json:",omitempty"`
	H      int    `json:",omitempty"`
	Resize string `json:",omitempty"`
}
