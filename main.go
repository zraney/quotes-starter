package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotes = []quote{
	{Quote: "Errors are values.", Author: "Rob Pike"},
	{Quote: "Don't panic.", Author: "Go Code Review Comments"},
	{Quote: "A little copying is better than a little dependency.", Author: "Rob Pike"},
	{Quote: "Concurrency is not parallelism.", Author: "Rob Pike"},
	{Quote: "interface{} says nothing.", Author: "Rob Pike"},
}

func getRandomQuote(c *gin.Context) {
	randomIndex := rand.Intn(len(quotes))
	pick := quotes[randomIndex]
	c.IndentedJSON(http.StatusOK, pick)
}

// func getQuotes(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, quotes)
// }

func main() {
	router := gin.Default()
	router.GET("/quote", getRandomQuote)
	router.Run("localhost:8080")
}
