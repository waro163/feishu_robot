package message

import (
	"fmt"

	"github.com/waro163/feishu_robot/feishu"
)

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create
func AddMembersToGroup(chatID, token string, body map[string]interface{}, querys ...map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s/members", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	if len(querys) > 0 {
		return feishu.Request("POST", url, querys[0], header, body)
	}
	return feishu.Request("POST", url, nil, header, body)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/delete
func RemoveMembersFromGroup(chatID, token string, body map[string]interface{}, querys ...map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s/members", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	if len(querys) > 0 {
		return feishu.Request("DELETE", url, querys[0], header, body)
	}
	return feishu.Request("DELETE", url, nil, header, body)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/me_join
func JoinGroup(chatID, token string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s/members/me_join", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("PATCH", url, nil, header, nil)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/get
func GetGroupMembersList(chatID, token string, querys ...map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s/members", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	if len(querys) > 0 {
		return feishu.Request("GET", url, querys[0], header, nil)
	}
	return feishu.Request("GET", url, nil, header, nil)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/list
func GetMyGroupsList(token string, querys ...map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/im/v1/chats"
	header := map[string]string{"Authorization": "Bearer " + token}
	if len(querys) > 0 {
		return feishu.Request("GET", url, querys[0], header, nil)
	}
	return feishu.Request("GET", url, nil, header, nil)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/search
func SearchVisibleGroupsList(token string, querys ...map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/im/v1/chats/search"
	header := map[string]string{"Authorization": "Bearer " + token}
	if len(querys) > 0 {
		return feishu.Request("GET", url, querys[0], header, nil)
	}
	return feishu.Request("GET", url, nil, header, nil)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/is_in_chat
func CheckIsInGroup(chatID, token string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/chats/%s/members/is_in_chat", chatID)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("GET", url, nil, header, nil)
}
