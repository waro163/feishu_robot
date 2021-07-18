package feishu

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/waro163/feishu_robot/requests"
)

func Request(method, url string, querys, headers map[string]string, body map[string]interface{}) (map[string]interface{}, error) {
	var query_param []string
	for key, value := range querys {
		query_param = append(query_param, key+"="+value)
	}
	url = url + "?" + strings.Join(query_param, "&")
	resp, err := requests.Request(url, strings.ToUpper(method), body, headers)
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
