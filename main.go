package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var ginList = []ginItem{
	{ID: "1", Name: "Ólafsson Gin", Price: 39.99},
	{ID: "2", Name: "New Amsterdam Gin", Price: 39.99},
	{ID: "3", Name: "Kuro", Price: 39.99},
}

func getGins(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ginList)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/images", "./images")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Gin & Gonic Mixing",
		})
	})
	r.GET("/gins", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gins.tmpl", gin.H{
			"title": "Gin & Gonic Mixing",
		})
	})
	r.GET("/gonics", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gonics.tmpl", gin.H{
			"title": "Gin & Gonic Mixing",
		})
	})
	r.GET("/mixes", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mixes.tmpl", gin.H{
			"title": "Gin & Gonic Mixing",
		})
	})
	r.GET("/prices", getGins)

	r.Run(":8080")
}
