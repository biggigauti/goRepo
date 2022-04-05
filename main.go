package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginItem struct { //Collection of fields.
	ID    string  `json:"id"` //json fields
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var ginList = []ginItem{ //List of type ginItem
	{ID: "1", Name: "Ólafsson Gin", Price: 39.99},
	{ID: "2", Name: "New Amsterdam Gin", Price: 39.99},
	{ID: "3", Name: "Kuro", Price: 39.99},
}

func getGins(c *gin.Context) { //Get function to grab gin list
	c.IndentedJSON(http.StatusOK, ginList)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl") //load html templates
	r.Static("/images", "./images") //load static files

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{ //load index.tmpl at / extension
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
