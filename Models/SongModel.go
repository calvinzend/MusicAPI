package models

type Song struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ArtistID int    `json:"artist_id"`
	FilePath string `json:"filepath"`
	Cover    string `json:"cover"`
}
