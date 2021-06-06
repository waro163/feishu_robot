package eventapp

type CallBackInput struct {
	Challenge string `json:"challenge" binding:"required"`
	Token     string `json:"token" binding:"required"`
	Type      string `json:"type" binding:"required"`
}
