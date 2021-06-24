package eventapp

type CallBackInput struct {
	Challenge string `json:"challenge" binding:"required"`
	Token     string `json:"token" binding:"required"`
	Type      string `json:"type" binding:"required"`
}

type CallBackEvent struct {
	Schema string      `json:"schema" binding:"required"`
	Header interface{} `json:"header" binding:"required"`
	Event  interface{} `json:"event" binding:"required"`
}
