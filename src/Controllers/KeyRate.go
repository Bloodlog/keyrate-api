package Controllers

import (
	"key-rate-api/src/client"

	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func KeyRate(c *gin.Context) {
	layout := "2006-01-02"
	fromDate := c.DefaultQuery("fromDate", time.Now().Format(layout))

	t, err := time.Parse(layout, fromDate)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	data, err := client.KeyRateByDate(t, time.Now())
	if err != nil {
		log.Println("error KeyRateByDate client")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data.Body.Response.Result.Rows[0].KeyRates,
	})
}
