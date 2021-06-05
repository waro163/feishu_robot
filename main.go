package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/waro163/feishu_robot/settings"
)

func main() {
	// fmt.Println(viper.GetString("APP_ID"))
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.Run()
}
