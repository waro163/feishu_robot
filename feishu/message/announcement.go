package message

import (
	"fmt"

	"github.com/waro163/feishu_robot/feishu"
)

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/get
func GetGroupAnnouncement(chatID, token string, querys ...map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s/announcement", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	if len(querys) > 0 {
		return feishu.Request("GET", url, querys[0], header, nil)
	}
	return feishu.Request("GET", url, nil, header, nil)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/patch
func UpdateGroupAnnouncement(chatID, token string, body map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s/announcement", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("PATCH", url, nil, header, body)
}
