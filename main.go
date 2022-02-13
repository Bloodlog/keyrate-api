package main

import (
    "net/http"
    "time"
	  "log"
    "encoding/json"
  	"key-rate-api/client"
    "github.com/gin-gonic/gin"
)

func main() {
    r := setupRouter()
	  r.Run(":8080")
}

func setupRouter() *gin.Engine {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
  v1 := router.Group("/v1")
  {
      v1.GET("/keyrate", keyrate)
  }
  router.Run("localhost:8080")

	return router
}

func keyrate(c *gin.Context)  {
  layout := "2006-01-02"
  fromDate := c.DefaultQuery("fromDate", time.Now().Format(layout))

  t, err := time.Parse(layout, fromDate)
  if err != nil {
      log.Println(err)
      c.IndentedJSON(http.StatusUnprocessableEntity, err)
  }

    data, err := client.KeyRateByDate(t, time.Now())
    if err != nil {
      log.Println("error KeyRateByDate client")
      log.Println(err)
      jsonData, _ := json.Marshal(err)
      c.IndentedJSON(http.StatusInternalServerError, jsonData)
    }

    c.IndentedJSON(http.StatusOK, data.Body.Response.Result.Rows[0].KeyRates)
}
