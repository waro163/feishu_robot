package message

import (
	"fmt"

	"github.com/waro163/feishu_robot/feishu"
)

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create
// token should be value of 'tenant_access_token'
// idType should be receive_id type in body, default is "open_id", choice is "open_id"/"user_id"/"union_id"/"email"/"chat_id"
// body should contain receive_id(not necessary), content, msg_type
func SendMsg(idType, token string, body map[string]interface{}) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/im/v1/messages?" + idType
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("POST", url, nil, header, body)
}

// https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM
// token should be value of 'tenant_access_token'
// body must contain "msg_type" and "content" object, choice is "department_ids"/"open_ids"/"user_ids"
func BatchSendMsg(token string, body map[string]interface{}) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/message/v4/batch_send/"
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("POST", url, nil, header, body)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/reply
func ReplyMsg(msgID, token string, body map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/messages/%s/reply", msgID)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("POST", url, nil, header, body)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/delete
func RevokeMsg(msgID, token string, body map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/messages/%s", msgID)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("DELETE", url, nil, header, body)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch
func UpdateMsg(msgID, token string, body map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/messages/%s", msgID)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("PATCH", url, nil, header, body)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/read_users
func QueryMsgReadStatus(msgID, token string, querys map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/messages/%s/read_users", msgID)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("GET", url, querys, header, nil)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/list
func GetChatHistoryMsg(token string, querys map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/im/v1/messages"
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("GET", url, querys, header, nil)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-resource/get
// 返回文件的二进制流
func GetResourceFile(msgID, fileKey, token string, querys map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/messages/%s/resources/%s", msgID, fileKey)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("GET", url, querys, header, nil)
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/get
func GetMsgContent(msgID, token string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/messages/%s", msgID)
	header := map[string]string{"Authorization": "Bearer " + token}
	return feishu.Request("GET", url, nil, header, nil)
}
