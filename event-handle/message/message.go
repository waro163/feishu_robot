package message

import eventmethod "github.com/waro163/feishu_robot/event-method"

func init() {
	eventmethod.RegisterEventMethod("im.message.receive_v1", HandleMsgEvent)
	eventmethod.RegisterEventMethod("im.message.message_read_v1", HandleMsgReadEvent)
}

func HandleMsgEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}

func HandleMsgReadEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}
