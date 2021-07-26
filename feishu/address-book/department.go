package addressbook

import (
	"github.com/waro163/feishu_robot/feishu"
)

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/create
func CreateDepartment(token string, body map[string]interface{}, querys ...map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/contact/v3/departments"
	header := map[string]string{"Authorization": "Bearer " + token}
	if len(querys) > 0 {
		return feishu.Request("POST", url, querys[0], header, body)
	}
	return feishu.Request("POST", url, nil, header, body)
}
