package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/99designs/gqlgen"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Quote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote,omitempty"`
	Author string `json:"author,omitempty"`
}

var db *sql.DB

func main() {
	err := connectUnixSocket()
	if err != nil {
		log.Fatalln(err)
	}
	router := gin.Default()
	router.GET("/quotes", getRandomQuote)
	router.GET("/quotes/:id", getQuoteByID)
	router.POST("/quotes", addNewQuote)
	router.DELETE("/quotes/:id", deleteQuoteByID)
	router.Run("0.0.0.0:8080")
}

func connectUnixSocket() error {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", k)
		}
		return v
	}

	var (
		dbUser         = mustGetenv("DB_USER")
		dbPwd          = mustGetenv("DB_PASS")
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET")
		dbName         = mustGetenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s", dbUser, dbPwd, dbName, unixSocketPath)

	var err error

	db, err = sql.Open("pgx", dbURI)
	if err != nil {
		return fmt.Errorf("sql.Open: %v", err)
	}

	return err
}

func handleRequest(c *gin.Context) bool {
	headers := c.Request.Header["X-Api-Key"]
	if headers != nil {
		return strings.Compare(headers[0], "COCKTAILSAUCE") == 0
	}
	return false
}

func getRandomQuote(c *gin.Context) {
	if handleRequest(c) {
		row := db.QueryRow("SELECT id, phrase, author FROM quotes ORDER BY RANDOM() LIMIT 1;")
		quote := &Quote{}
		err := row.Scan(&quote.ID, &quote.Quote, &quote.Author)
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, quote)
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
}

func getQuoteByID(c *gin.Context) {
	if handleRequest(c) {
		id := c.Param("id")
		row := db.QueryRow(fmt.Sprintf("SELECT id, phrase, author FROM quotes WHERE id = '%s';", id))
		quote := &Quote{}
		err := row.Scan(&quote.ID, &quote.Quote, &quote.Author)
		if err != nil {
			log.Println(err)
		}
		if quote.ID != "" {
			c.JSON(http.StatusOK, quote)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
}

func addNewQuote(c *gin.Context) {
	if handleRequest(c) {
		newID := uuid.New().String()
		var newQuote Quote
		newQuote.ID = newID

		if err := c.BindJSON(&newQuote); err != nil {
			return
		}
		id := newQuote.ID
		phrase := newQuote.Quote
		author := newQuote.Author
		returnId := Quote{
			ID: newQuote.ID,
		}

		if len(author) >= 3 && len(phrase) >= 3 {
			_, err := db.Exec(fmt.Sprintf("INSERT INTO quotes (id, phrase, author) VALUES ('%s', '%s', '%s');", id, phrase, author))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err})
				return
			}
			c.JSON(http.StatusCreated, returnId)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": "quote and author must be greater than 3 characters"})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
}

func deleteQuoteByID(c *gin.Context) {
	if handleRequest(c) {
		id := c.Param("id")
		_, err := db.Exec(fmt.Sprintf("DELETE FROM quotes WHERE id = '%s';", id))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{"message": "204 can't display a message"})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
}
