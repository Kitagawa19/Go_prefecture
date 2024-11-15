package handlers

import (
	"log"
	"net/http"
	"regexp"
	"Go_prefecture/database"
	"github.com/gin-gonic/gin"
)

func PrefecturetocityHandler(c *gin.Context) {
	rows, err := database.DB.Query("SELECT DISTINCT field7 FROM addresses")
	if err != nil {
		log.Printf("Failed to fetch prefectures: %v", err)
		c.String(http.StatusInternalServerError, "Failed to fetch prefectures")
		return
	}
	defer rows.Close()

	var prefectures []string
	reader:=regexp.MustCompile(`^[\p{Han}]{2,3}(?:都|道|府|県)$`)

	for rows.Next() {
		var prefecture string
		rows.Scan(&prefecture)
		if(reader.MatchString(prefecture)){
			prefectures = append(prefectures, prefecture)
		}
	}
	c.HTML(http.StatusOK, "cities.html", gin.H{
		"Prefectures": prefectures,
	})
}