package addressbook

import (
	eventmethod "github.com/waro163/feishu_robot/event-method"
)

func init() {
	// 通讯录-用户事件
	eventmethod.RegisterEventMethod("contact.user.created_v3", HandleStaffInductionEvent)
	eventmethod.RegisterEventMethod("contact.user.deleted_v3", HandleStaffOffEvent)
	eventmethod.RegisterEventMethod("contact.user.updated_v3", HandleStaffContactUpdateEvent)
}

func HandleStaffInductionEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}

func HandleStaffOffEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}

func HandleStaffContactUpdateEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}
