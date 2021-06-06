package eventapp

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func eventCallBack(c *gin.Context) {
	var in CallBackInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Printf("challenge is %s\n", in.Challenge)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
