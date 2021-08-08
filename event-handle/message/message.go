package message

import (
	"fmt"
	"log"

	eventmethod "github.com/waro163/feishu_robot/event-method"
	"github.com/waro163/feishu_robot/feishu"
	msg "github.com/waro163/feishu_robot/feishu/message"
)

func init() {
	eventmethod.RegisterEventMethod("im.message.receive_v1", HandleMsgEvent)
	eventmethod.RegisterEventMethod("im.message.message_read_v1", HandleMsgReadEvent)
}

func HandleMsgEvent(header map[string]string, event map[string]interface{}) error {
	message, ok := event["message"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("not found message in event body")
	}
	message_id := message["message_id"].(string)
	message_type := message["message_type"].(string)
	content := message["content"]
	// get access token
	// TODO: storage token to redis and get it instead of net
	token, err := feishu.GetTenantAccessToken()
	if err != nil {
		return err
	}
	body := map[string]interface{}{"content": content, "msg_type": message_type}
	resp, err := msg.ReplyMsg(message_id, token, body)
	if err != nil {
		if resp != nil {
			log.Printf("reply message error: %v\n", resp)
		}
		return err
	}
	return nil
}

func HandleMsgReadEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}
