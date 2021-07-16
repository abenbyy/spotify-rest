package main

import (
	"github.com/abenbyy/spotify-rest/resolver"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)
const defaultPort = "8080"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main(){
	r:= gin.Default()
	r.Use(CORSMiddleware())

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
		c.JSON(http.StatusOK,gin.H{"data" : res} )
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
