package eventapp

type CallBackInput struct {
	Challenge string `json:"challenge" binding:"required"`
	Token     string `json:"token" binding:"required"`
	Type      string `json:"type" binding:"required"`
}

type CallBackEvent struct {
	Challenge string            `json:"challenge"`
	Token     string            `json:"token"`
	Type      string            `json:"type"`
	Schema    string            `json:"schema"`
	Header    map[string]string `json:"header"`
	Event     interface{}       `json:"event"`
}
