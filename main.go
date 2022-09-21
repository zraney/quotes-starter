package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type quote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func main() {
	router := gin.Default()
	router.GET("/quotes", getRandomQuote)
	router.GET("/quotes/:id", getQuoteByID)
	router.POST("/quotes", addNewQuote)
	router.Run("0.0.0.0:8080")

}

// TODO break this func into smaller, more focused funcs
func getRandomQuote(c *gin.Context) {
	keyArray := []string{}
	for k, _ := range quotes {
		keyArray = append(keyArray, k)
	}
	randomIndex := rand.Intn(len(keyArray))
	randomPick := keyArray[randomIndex]
	randomQuote := quotes[randomPick]
	c.JSON(http.StatusOK, randomQuote)
}

func getQuoteByID(c *gin.Context) {
	id := c.Param("id")
	quote, exists := quotes[id]

	if exists {
		c.JSON(http.StatusOK, quote)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "quote not found"})
}

func addNewQuote(c *gin.Context) {
	newID := uuid.New().String()
	var newQuote quote
	newQuote.ID = newID

	if err := c.BindJSON(&newQuote); err != nil {
		return
	}
	if len(newQuote.Author) < 3 || len(newQuote.Quote) < 3 {
		quotes[newQuote.ID] = newQuote
		c.JSON(http.StatusNotFound, gin.H{"message": "quote and author must be greater than 3 characters"})
		return
	}
	c.JSON(http.StatusCreated, newQuote.ID)
}

var quotes = map[string]quote{
	"b37c9ded-d176-4fe5-a9b9-1427ebf92ed1": {ID: "b37c9ded-d176-4fe5-a9b9-1427ebf92ed1", Quote: "Errors are values.", Author: "Rob Pike"},
	"0d95d2d8-28b0-4278-960d-cbdd16beab02": {ID: "0d95d2d8-28b0-4278-960d-cbdd16beab02", Quote: "Errors are values.", Author: "Rob Pike"},
	"0329b963-004d-4add-bb5e-cfe7defd0c6d": {ID: "0329b963-004d-4add-bb5e-cfe7defd0c6d", Quote: "Don't panic.", Author: "Go Code Review Comments"},
	"2e774b8c-672e-46bf-8b6f-38d6889edee7": {ID: "2e774b8c-672e-46bf-8b6f-38d6889edee7", Quote: "A little copying is better than a little dependency.", Author: "Rob Pike"},
	"a2ad7811-22ea-4ba4-8691-a88b5f89a475": {ID: "a2ad7811-22ea-4ba4-8691-a88b5f89a475", Quote: "Concurrency is not parallelism.", Author: "Rob Pike"},
	"ba9b4b54-3070-4665-bd29-de3e99c991d2": {ID: "ba9b4b54-3070-4665-bd29-de3e99c991d2", Quote: "interface{} says nothing.", Author: "Rob Pike"},
	"ca17bd05-4c0b-41ae-9496-518371e245f2": {ID: "ca17bd05-4c0b-41ae-9496-518371e245f2", Quote: "Make the zero value useful.", Author: "Rob Pike"},
}
