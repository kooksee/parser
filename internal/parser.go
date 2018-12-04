package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/storyicon/graphquery"
)

// Response defines the API response format.
type Response struct {
	// Data is the carrier for returning data.
	Data interface{} `json:"data"`
	// Error records the errors in this request.
	Error string `json:"error"`
	// TimeCost recorded the time wastage of the request.
	TimeCost int64 `json:"time_cost"`
}

// Start is used to start the http server
func Start(debug bool, port string) error {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok", )
		return
	})

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
		return
	})

	router.POST("/check", func(c *gin.Context) {
		dt, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if _, e := graphquery.Compile(dt); e != nil {
			c.String(http.StatusBadRequest, e.Error())
			return
		}

		c.String(http.StatusOK, "ok")
		return
	})

	router.POST("/parse", func(c *gin.Context) {
		document := c.PostForm("document")
		expression := c.PostForm("expression")
		timeStart := time.Now().UnixNano()
		conSeq := graphquery.ParseFromString(document, expression)
		err := ""
		if len(conSeq.Errors) > 0 {
			err = conSeq.Errors[0]
		}
		c.JSON(http.StatusOK, Response{
			Data:     conSeq.Data,
			TimeCost: int64(time.Now().UnixNano() - timeStart),
			Error:    err,
		})
	})

	return router.Run(fmt.Sprintf(":%s", port))
}
