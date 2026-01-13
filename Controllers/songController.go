package controllers

import (
	"path/filepath"
	"strconv"

	models "week2/Models"

	"github.com/gin-gonic/gin"
)

var songs = []models.Song{
	{
		ID:       1,
		Title:    "Laskar Pelangi",
		ArtistID: 0,
		FilePath: "storage/songs/laskar_pelangi.mp3",
		Cover:    "storage/covers/laskar.jpg",
	},
	{
		ID:       2,
		Title:    "Separuh Aku",
		ArtistID: 0,
		FilePath: "storage/songs/separuh_aku.mp3",
		Cover:    "storage/covers/noah.jpg",
	},
}
var songID = 3

func GetSongs(c *gin.Context) {
	c.JSON(200, songs)
}

func UploadSong(c *gin.Context) {
	title := c.PostForm("title")
	artistID, _ := strconv.Atoi(c.PostForm("artist_id"))

	// ===== Upload SONG =====
	songFile, err := c.FormFile("song")
	if err != nil {
		c.JSON(400, gin.H{"error": "song file required"})
		return
	}

	songFilename := strconv.Itoa(songID) + "_" + filepath.Base(songFile.Filename)
	songPath := "storage/songs/" + songFilename
	c.SaveUploadedFile(songFile, songPath)

	// ===== Upload COVER =====
	coverFile, err := c.FormFile("cover")
	coverPath := ""
	if err == nil {
		coverFilename := strconv.Itoa(songID) + "_" + filepath.Base(coverFile.Filename)
		coverPath = "storage/covers/" + coverFilename
		c.SaveUploadedFile(coverFile, coverPath)
	}

	song := models.Song{
		ID:       songID,
		Title:    title,
		ArtistID: artistID,
		FilePath: songPath,
		Cover:    coverPath,
	}

	songs = append(songs, song)
	songID++

	c.JSON(201, song)
}

func StreamSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, song := range songs {
		if song.ID == id {
			c.File(song.FilePath)
			return
		}
	}

	c.JSON(404, gin.H{"error": "song not found"})
}

func CoverSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, song := range songs {
		if song.ID == id {
			c.File(song.Cover)
			return
		}
	}

	c.JSON(404, gin.H{"error": "song not found"})
}
