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
	}
}

type MediaSizes struct {
	Medium MediaSize
	Thumb  MediaSize
	Small  MediaSize
	Large  MediaSize
}

type MediaSize struct {
	W      int    `json:",omitempty"`
	H      int    `json:",omitempty"`
	Resize string `json:",omitempty"`
}
