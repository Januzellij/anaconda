package anaconda

type Place struct {
	Attributes  map[string]string `json:"attributes,omitempty"`
	BoundingBox struct {
		Coordinates [][][]float64 `json:"coordinates,omitempty"`
		Type        string        `json:"type,omitempty"`
	} `json:"bounding_box,omitempty"`
	ContainedWithin []struct {
		Attributes  map[string]string `json:"attributes,omitempty"`
		BoundingBox struct {
			Coordinates [][][]float64 `json:"coordinates,omitempty"`
			Type        string        `json:"type,omitempty"`
		} `json:"bounding_box,omitempty"`
		Country     string `json:"country,omitempty"`
		CountryCode string `json:"country_code,omitempty"`
		FullName    string `json:"full_name,omitempty"`
		ID          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		PlaceType   string `json:"place_type,omitempty"`
		URL         string `json:"url,omitempty"`
	} `json:"contained_within,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Geometry    struct {
		Coordinates [][][]float64 `json:"coordinates,omitempty"`
		Type        string        `json:"type,omitempty"`
	} `json:"geometry,omitempty"`
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	PlaceType string   `json:"place_type,omitempty"`
	Polylines []string `json:"polylines,omitempty"`
	URL       string   `json:"url,omitempty"`
}
