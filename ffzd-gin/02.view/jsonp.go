package main

import (
	"github.com/gin-gonic/gin"
)

func jsonp() {
	router := gin.Default()

	// JSONP?callback=x
	router.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// x({"foo": "bar"})
		c.JSONP(200, data)
	})
	router.Run(":8080")
}
