package main

import (
	"fmt"

	"github.com/ashwanthkumar/licensd-server/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".") // always look for config.{json,yml,toml,hcl} in the current working directory
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	initDatabase()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	r.POST("/payload", handler.AddPayloadToDB)
	r.Run() // listen and serve on 0.0.0.0:8080
}
