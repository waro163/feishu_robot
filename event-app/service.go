package eventapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/waro163/feishu_robot/event-handle/message"
	eventmethod "github.com/waro163/feishu_robot/event-method"
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
	eventType, ok := in.Header["event_type"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "lost event_type field in body header"})
		return
	}
	eventFunc := eventmethod.GetEventMethod(eventType)
	err := eventFunc(in.Header, in.Event.(map[string]interface{}))
	if err != nil {
		// TODO:
		// handle event callback function error, send message to notice dev
		c.JSON(http.StatusInternalServerError, gin.H{"message": "handle callback error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
