package Controllers

import (
	"key-rate-api/src/client"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// KeyRate godoc
// @Summary      Show key rates
// @Description  Return keyrates
// @Tags         keyrate
// @Produce      json
// @Param        fromDate    query     string  false  "return keyrates from date"  Format(2006-01-02)
// @Param        page    query     integer  false  "Page number"  Format(2)
// @Param        per_page    query     integer  false  "Key rates per page"  Format(10)
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /keyrate [get]
func KeyRate(c *gin.Context) {
	layout := "2006-01-02"
	fromDate := c.DefaultQuery("fromDate", time.Now().Format(layout))
	currentPageNumber := c.DefaultQuery("page", "1")
	currentPageNumberInt, _ := strconv.Atoi(currentPageNumber)
	perPageNumber := c.DefaultQuery("per_page", "3")
	PerPage, _ := strconv.Atoi(perPageNumber)

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

	dataKeyRates := data.Body.Response.Result.Rows[0].KeyRates
	total := len(dataKeyRates)
	totalPagesFloat := float64(total / PerPage)
	totalPages := int(math.Ceil(totalPagesFloat))

	result := make([]client.KeyRates, 0)
	startIndex := currentPageNumberInt * PerPage
	endIndex := startIndex + PerPage

	for i := startIndex; i < endIndex; i++ {
		if currentPageNumberInt <= totalPages && i <= total-1 {
			result = append(result, dataKeyRates[i])
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":         result,
		"total":        total,
		"current_page": currentPageNumberInt,
		"per_page":     PerPage,
		"total_pages":  totalPages,
		"from_date":    fromDate,
	})
}
