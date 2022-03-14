package Controllers

import (
	"key-rate-api/src/Helpers"
	"key-rate-api/src/client"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Success struct {
	Data        []client.KeyRates `json:"data"`
	Total       int               `json:"total"`
	CurrentPage int               `json:"current_page"`
	PerPage     int               `json:"per_page"`
	TotalPages  int               `json:"total_pages"`
	FromDate    string            `json:"from_date"`
}

type Error struct {
	Errors error `json:"errors"`
}

// KeyRate godoc
// @Summary      Show key rates
// @Description  Return keyrates
// @Tags         keyrate
// @Produce      json
// @Param        from_date    query     string  false  "return keyrates from date"  Format(2006-01-02)
// @Param        page    query     integer  false  "Page number"  Format(2)
// @Param        per_page    query     integer  false  "Key rates per page"  Format(10)
// @Param 			 last_key_rate query integer 0 "Force get last key rate. Max count retry 5" Format(1)
// @Success      200  {array} Success
// @Failure      400  error Error
// @Failure      404
// @Failure      500  error Error
// @Router       /keyrate [get]
func KeyRate(c *gin.Context) {
	layout := "2006-01-02"
	fromDateRequest := c.DefaultQuery("from_date", time.Now().Format(layout))
	currentPageRequest := c.DefaultQuery("page", "1")
	currentPageNumber, _ := strconv.Atoi(currentPageRequest)
	perPageRequest := c.DefaultQuery("per_page", "15")
	PerPage, _ := strconv.Atoi(perPageRequest)
	lastKeyRateRequest := c.DefaultQuery("last_key_rate", "0")
	ForceLastKeyRate, _ := strconv.Atoi(lastKeyRateRequest)

	fromDate, err := time.Parse(layout, fromDateRequest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Error{
			Errors: err,
		})
		return
	}

	dataKeyRates, err := client.GetData(fromDate, time.Now())
	if err != nil {
		log.Println("error KeyRateByDate client")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, Error{
			Errors: err,
		})
		return
	}

	if ForceLastKeyRate == 1 {
		maxTry := 5
		for i := 0; i < maxTry; i++ {
			fromDatesubDay := fromDate.AddDate(0, 0, -1)
			dataKeyRates, err = client.GetData(fromDatesubDay, time.Now())
			if err != nil {
				log.Println("error KeyRateByDate client")
				log.Println(err)
				c.JSON(http.StatusInternalServerError, Error{
					Errors: err,
				})
				return
			}
			if len(dataKeyRates) != 0 {
				break
			}
		}
	}

	var result Helpers.Pagineted

	result = Helpers.Paginate(&Helpers.Pages{
		Items:       dataKeyRates,
		Total:       len(dataKeyRates),
		PerPage:     PerPage,
		CurrentPage: currentPageNumber,
	})

	c.JSON(http.StatusOK, Success{
		Data:        result.Data,
		Total:       result.Total,
		CurrentPage: result.CurrentPage,
		PerPage:     result.PerPage,
		TotalPages:  result.TotalPages,
		FromDate:    fromDateRequest,
	})
}
