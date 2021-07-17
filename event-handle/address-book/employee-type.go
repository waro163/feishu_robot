package addressbook

import eventmethod "github.com/waro163/feishu_robot/event-method"

func init() {
	// 通讯录-人员类型
	eventmethod.RegisterEventMethod("contact.employee_type_enum.created_v3", HandleEmployeeTypeCreateEvent)
	eventmethod.RegisterEventMethod("contact.employee_type_enum.actived_v3", HandleEmployeeTypeActiveEvent)
	eventmethod.RegisterEventMethod("contact.employee_type_enum.deactivated_v3", HandleEmployeeTypeDeactiveEvent)
	eventmethod.RegisterEventMethod("contact.employee_type_enum.updated_v3", HandleEmployeeTypeUpdateEvent)
	eventmethod.RegisterEventMethod("contact.employee_type_enum.deleted_v3", HandleEmployeeTypeDeleteEvent)
}

func HandleEmployeeTypeCreateEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleEmployeeTypeActiveEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleEmployeeTypeDeactiveEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleEmployeeTypeUpdateEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleEmployeeTypeDeleteEvent(map[string]string, map[string]interface{}) error {
	return nil
}
