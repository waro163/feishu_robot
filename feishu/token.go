package feishu

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/waro163/feishu_robot/requests"
)

//企业自建应用获取 tenant_access_token
func GetTenantAccessToken() (string, error) {
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
	body := map[string]string{
		"app_id":     viper.GetString("APP_ID"),
		"app_secret": viper.GetString("APP_SECRET"),
	}
	resp, err := requests.Post(url, body)
	if err != nil {
		return "", err
	}
	var res map[string]interface{}
	if err = json.Unmarshal(resp.Body, &res); err != nil {
		log.Printf("unmarshal tenant_access_token error:%v\n", resp.Body)
		return "", err
	}
	if int(res["code"].(float64)) != 0 {
		log.Printf("get tenant_access_token error:%v\n", res)
		return "", fmt.Errorf("code: %d, msg: %s", res["code"], res["msg"])
	}
	return res["tenant_access_token"].(string), nil
}

// 企业自建应用获取 app_access_token
func GetAppAccessToken() (string, error) {
	url := "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal/"
	body := map[string]string{
		"app_id":     viper.GetString("APP_ID"),
		"app_secret": viper.GetString("APP_SECRET"),
	}
	resp, err := requests.Post(url, body)
	if err != nil {
		return "", err
	}
	var res map[string]interface{}
	if err = json.Unmarshal(resp.Body, &res); err != nil {
		log.Printf("unmarshal app_access_token error:%v\n", resp.Body)
		return "", err
	}
	if int(res["code"].(float64)) != 0 {
		log.Printf("get app_access_token error:%v\n", res)
		return "", fmt.Errorf("code: %d, msg: %s", res["code"], res["msg"])
	}
	return res["app_access_token"].(string), nil
}
