package message

import eventmethod "github.com/waro163/feishu_robot/event-method"

func init() {
	eventmethod.RegisterEventMethod("im.chat.disbanded_v1", HandleGroupDisbandedEvent)
	eventmethod.RegisterEventMethod("iim.chat.updated_v1", HandleGroupUpdateEvent)
}

func HandleGroupDisbandedEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}

func HandleGroupUpdateEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}
