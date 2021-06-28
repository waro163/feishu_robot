package eventmethod

func init() {
	RegisterEventMethod("im.message.receive_v1", HandleMsgEvent)
	RegisterEventMethod("im.message.message_read_v1", HandleMsgReadEvent)
}

func HandleMsgEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleMsgReadEvent(map[string]string, map[string]interface{}) error {
	return nil
}
