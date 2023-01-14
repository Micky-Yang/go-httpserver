package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type HealthyStatus struct {
	Status string
}

func main() {
	r := gin.Default()

	StatusDesc := "OK"
	statusCode := http.StatusOK

	r.GET("/", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"Message": "Hello Go",
			"Version": "v2",
			"Time":    time.Now().Format("2006/01/02 15:04"),
		})
	})

	r.GET("/healthy", func(c *gin.Context) {
		if StatusDesc != "OK" {
			statusCode = http.StatusBadGateway
		} else {
			statusCode = http.StatusOK
		}
		c.PureJSON(statusCode, gin.H{
			"Status": StatusDesc,
			"Time":   time.Now().Format("2006/01/02 15:04"),
		})
	})

	r.POST("/healthy", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(body))
		h := HealthyStatus{}
		json.Unmarshal(body, &h)

		StatusDesc = h.Status

		c.PureJSON(200, gin.H{
			"ChangeStatus": StatusDesc,
			"Time":         time.Now().Format("2006/01/02 15:04"),
		})
	})

	r.Run(":8080")
}
