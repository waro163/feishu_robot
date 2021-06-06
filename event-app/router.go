package eventapp

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup) {
	r.POST("/call_back", eventCallBack)

}
