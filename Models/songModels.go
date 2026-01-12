package models

type Song struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	FilePath string `json:"filepath"`
	Cover    string `json:"cover"`
}
