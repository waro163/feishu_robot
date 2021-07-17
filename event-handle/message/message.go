package message

import eventmethod "github.com/waro163/feishu_robot/event-method"

func init() {
	eventmethod.RegisterEventMethod("im.message.receive_v1", HandleMsgEvent)
	eventmethod.RegisterEventMethod("im.message.message_read_v1", HandleMsgReadEvent)
}

func HandleMsgEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleMsgReadEvent(map[string]string, map[string]interface{}) error {
	return nil
}
