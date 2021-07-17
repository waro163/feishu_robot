package addressbook

import (
	eventmethod "github.com/waro163/feishu_robot/event-method"
)

func init() {
	// 通讯录-部门事件
	eventmethod.RegisterEventMethod("contact.department.updated_v3", HandleDepartmentUpdate)
	eventmethod.RegisterEventMethod("contact.department.deleted_v3", HandleDepartmentDelete)
	eventmethod.RegisterEventMethod("contact.department.created_v3", HandleDepartmentCreate)
}

func HandleDepartmentCreate(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleDepartmentDelete(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleDepartmentUpdate(map[string]string, map[string]interface{}) error {
	return nil
}
