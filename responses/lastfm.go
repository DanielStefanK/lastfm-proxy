package responses

type GetRecentResponse struct {
	RecentTracks struct {
		Tracks     []Track `json:"track"`
		Attributes struct {
			User       string
			TotalPages string
			Page       string
			PerPage    string
			Total      string
		} `json:"@attr"`
	} `json:"recenttracks"`
}

type Track struct {
	Artist     NamedAttribute `json:"artist"`
	Album      NamedAttribute `json:"album"`
	Name       string         `json:"name"`
	URL        string         `json:"url"`
	Attributes TrackAttribute `json:"@attr"`
	Image      []Image        `json:"image"`
}

type NamedAttribute struct {
	Mbid string `json:"mbit"`
	Text string `json:"#text"`
}

type TrackAttribute struct {
	NowPlaying string `json:"nowplaying"`
}

type Image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}
