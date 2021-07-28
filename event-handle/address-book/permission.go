package addressbook

import eventmethod "github.com/waro163/feishu_robot/event-method"

func init() {
	// 通讯录-权限变更
	eventmethod.RegisterEventMethod("contact.scope.updated_v3", HandleScopeUpdateEvent)
}

func HandleScopeUpdateEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}
