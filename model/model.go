package model

type TopTrackResponse struct {
	TopTracks struct {
		Track []TrackResponse `json:"track"`
	} `json:"toptracks"`
}

type TrackResponse struct {
	Name      string `json:"name"`
	Playcount string `json:"playcount"`
}

type TopArtistResponse struct {
	TopArtists struct {
		Artist []ArtistResponse `json:"artist"`
	} `json:"topartists"`
}

type ArtistResponse struct {
	Name      string `json:"name"`
	Playcount string `json:"playcount"`
}
