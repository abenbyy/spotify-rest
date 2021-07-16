package main

import (
	"github.com/abenbyy/spotify-rest/resolver"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)
const defaultPort = "8080"

func main(){
	r:= gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/artist", func(c *gin.Context){
		artist := c.DefaultQuery("query", "")
		if artist == ""{
			c.JSON(http.StatusOK, gin.H{"error": "please supply an artist name"})
		}

		res, err := resolver.QueryArtist(artist)
		if err != nil{
			panic(err)
		}
		c.JSON(http.StatusOK,gin.H{"data:" : res} )
	})

	r.GET("/album", func(c*gin.Context){
		id := c.DefaultQuery("id", "")
		if id == ""{
			c.JSON(http.StatusOK, gin.H{"error": "please supply an album id"})
		}

		res := resolver.GetAlbum(id)
		c.JSON(http.StatusOK, gin.H{"data": res})
	})

	r.GET("/track", func(c *gin.Context){
		id := c.DefaultQuery("id", "")
		if id == ""{
			c.JSON(http.StatusOK, gin.H{"error": "please supply a track id"})
		}

		res := resolver.GetTrack(id)
		c.JSON(http.StatusOK, gin.H{"data": res})
	})

	r.Run(":"+port)
}
