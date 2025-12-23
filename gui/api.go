package gui

import (
	"net/http"
	sqldb "go-ontn/sqldb"
)
type ID int
const(
	essay ID = iota
	image
	config
	data
	question
)
func api(c *gin.Context, id int, db *sqldb.Data) {
	switch id {
	case int(essay):
		es,err:=db.DB.GetAllEssays()
		
	case int(image):
		// Handle image API
	case int(config):
		// Handle config API
	case int(data):
		// Handle data API
	case int(question):
		// Handle question API
	default:
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid API ID"})
	}	
}