package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type quote struct {
	Id     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func main() {
	router := gin.Default()
	router.GET("/quotes", getRandomQuote)
	router.Run("0.0.0.0:8080")

}

func getRandomQuote(c *gin.Context) {
	randomKey := rand.Intn(len(quotes))
	pick := quotes[randomKey]
	c.JSON(http.StatusOK, pick)
}

var quotes = map[int]quote{
	1: {Id: "b37c9ded-d176-4fe5-a9b9-1427ebf92ed1", Quote: "Errors are values.", Author: "Rob Pike"},
	2: {Id: "0d95d2d8-28b0-4278-960d-cbdd16beab02", Quote: "Errors are values.", Author: "Rob Pike"},
	3: {Id: "0329b963-004d-4add-bb5e-cfe7defd0c6d", Quote: "Don't panic.", Author: "Go Code Review Comments"},
	4: {Id: "2e774b8c-672e-46bf-8b6f-38d6889edee7", Quote: "A little copying is better than a little dependency.", Author: "Rob Pike"},
	5: {Id: "a2ad7811-22ea-4ba4-8691-a88b5f89a475", Quote: "Concurrency is not parallelism.", Author: "Rob Pike"},
	6: {Id: "ba9b4b54-3070-4665-bd29-de3e99c991d2", Quote: "interface{} says nothing.", Author: "Rob Pike"},
	7: {Id: "ca17bd05-4c0b-41ae-9496-518371e245f2", Quote: "Make the zero value useful.", Author: "Rob Pike"},
}

func Add() {
	//will use uuid.New().String() here to create the uuid only once and assign it to the quote at creation
}
