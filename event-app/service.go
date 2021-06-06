package eventapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func eventCallBack(c *gin.Context) {
	var in CallBackInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if in.Token != viper.GetString("APP_VERIFICATION_TOKEN") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong APP_VERIFICATION_TOKEN",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"challenge": in.Challenge,
	})
}
