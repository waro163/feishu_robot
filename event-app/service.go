package eventapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func eventCallBack(c *gin.Context) {
	var in CallBackEvent
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if in.Token != viper.GetString("APP_VERIFICATION_TOKEN") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong APP_VERIFICATION_TOKEN"})
		return
	}
	eventType := in.Type
	if eventType == "url_verification" {
		c.JSON(http.StatusOK, gin.H{"challenge": in.Challenge})
		return
	} else if eventType == "event_callback" {
		c.JSON(http.StatusOK, gin.H{"message": "version 1.0 message would be ignore"})
		return
	}
	// handle version 2.0 message according to header and event
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
