package message

import (
	"fmt"

	"github.com/waro163/feishu_robot/feishu"
)

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/create
func CreateGroup(token string, body map[string]interface{}) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/im/v1/chats"
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("POST", url, nil, header, body)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get
func GetGroupMsg(chatID, token string, querys ...map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	if len(querys) > 0 {
		return feishu.Request("GET", url, querys[0], header, nil)
	}
	return feishu.Request("GET", url, nil, header, nil)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/update
func UpdateGroupMsg(chatID, token string, body map[string]interface{}, querys ...map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	if len(querys) > 0 {
		return feishu.Request("PUT", url, querys[0], header, body)
	}
	return feishu.Request("PUT", url, nil, header, body)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/delete
func DisbandGroup(chatID, token string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("DELETE", url, nil, header, nil)
}
