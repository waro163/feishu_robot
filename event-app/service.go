package eventapp

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func eventCallBack(c *gin.Context) {
	var in CallBackEvent
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Printf("receive: %v\n", in)
	// if in.Token != viper.GetString("APP_VERIFICATION_TOKEN") {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "wrong APP_VERIFICATION_TOKEN",
	// 	})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"challenge": "ok",
	})
}
