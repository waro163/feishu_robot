package addressbook

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/waro163/feishu_robot/requests"
)

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get
// querys key should only contains "user_id_type" or "department_id_type"
// "user_id_type" value should be choice of ("open_id"/"union_id"/"user_id"),default is "open_id"
// id should be consistent with "user_id_type" value in querys
// token should be value of 'tenant_access_token' or 'user_access_token'
func GetSingleUserInf(id, token string, querys ...map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/contact/v3/users/" + id
	if len(querys) > 0 {
		var query_param []string
		for key, value := range querys[0] {
			query_param = append(query_param, key+"="+value)
		}
		url = url + "?" + strings.Join(query_param, "&")
	}

	header := map[string]string{"Authorization": "Bearer " + token}
	resp, err := requests.Get(url, nil, header)
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

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create
// token should be value of 'tenant_access_token'
func CreateUser(token string, body map[string]interface{}, querys ...map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/contact/v3/users"
	if len(querys) > 0 {
		var query_param []string
		for key, value := range querys[0] {
			query_param = append(query_param, key+"="+value)
		}
		url = url + "?" + strings.Join(query_param, "&")
	}
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

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/list
// token should be value of 'tenant_access_token' or 'user_access_token'
func GetUserList(token string, querys ...map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/contact/v3/users"
	if len(querys) > 0 {
		var query_param []string
		for key, value := range querys[0] {
			query_param = append(query_param, key+"="+value)
		}
		url = url + "?" + strings.Join(query_param, "&")
	}
	header := map[string]string{"Authorization": "Bearer " + token}
	resp, err := requests.Get(url, nil, header)
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

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/patch
func UpdateUserPartInf(id, token string, body map[string]interface{}, querys ...map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/contact/v3/users/" + id
	if len(querys) > 0 {
		var query_param []string
		for key, value := range querys[0] {
			query_param = append(query_param, key+"="+value)
		}
		url = url + "?" + strings.Join(query_param, "&")
	}
	header := map[string]string{"Authorization": "Bearer " + token}
	resp, err := requests.Patch(url, body, header)
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

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/update
func UpdateUserAllInf(id, token string, body map[string]interface{}, querys ...map[string]string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/contact/v3/users/" + id
	if len(querys) > 0 {
		var query_param []string
		for key, value := range querys[0] {
			query_param = append(query_param, key+"="+value)
		}
		url = url + "?" + strings.Join(query_param, "&")
	}
	header := map[string]string{"Authorization": "Bearer " + token}
	resp, err := requests.Put(url, body, header)
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

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/delete
func DeleteUser(id, token string, body map[string]interface{}, querys ...string) (map[string]interface{}, error) {
	url := "https://open.feishu.cn/open-apis/contact/v3/users/" + id
	if len(querys) > 0 {
		url = url + "?user_id_type=" + querys[0]
	}
	header := map[string]string{"Authorization": "Bearer " + token}
	resp, err := requests.Put(url, body, header)
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
