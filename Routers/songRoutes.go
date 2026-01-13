package routers

import (
	controllers "week2/Controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/songs", controllers.GetSongs)
		api.POST("/songs", controllers.UploadSong)
		api.GET("/songs/:id/stream", controllers.StreamSong)
		api.GET("/songs/:id/cover", controllers.CoverSong)

		api.GET("/artists", controllers.GetArtists)
	}
}
