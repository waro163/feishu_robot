package message

import (
	"encoding/json"
	"fmt"

	"github.com/waro163/feishu_robot/requests"
)

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create
// token should be value of 'tenant_access_token'
// idType should be receive_id type in body, default is "open_id", choice is "open_id"/"user_id"/"union_id"/"email"/"chat_id"
// body should contain receive_id(not necessary), content, msg_type
func SendMsg(idType, token string, body map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/im/v1/messages?" + idType
	header := map[string]string{"Authorization": "Bearer " + token}
	resp, err := requests.Post(url, body, header)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	if err = json.Unmarshal(resp.Body, &res); err != nil {
		return map[string]interface{}{"msg": string(resp.Body)}, err
	}
	if resp.Code != 200 {
		return res, fmt.Errorf("http status: %d", resp.Code)
	}
	return res, nil
}

// https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM
// token should be value of 'tenant_access_token'
// body must contain "msg_type" and "content" object, choice is "department_ids"/"open_ids"/"user_ids"
func BatchSendMsg(token string, body map[string]interface{}) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/message/v4/batch_send/"
	header := map[string]string{"Authorization": "Bearer " + token}
	resp, err := requests.Post(url, body, header)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	if err = json.Unmarshal(resp.Body, &res); err != nil {
		return map[string]interface{}{"msg": string(resp.Body)}, err
	}
	if resp.Code != 200 {
		return res, fmt.Errorf("http status: %d", resp.Code)
	}
	return res, nil
}
