package eventmethod

func init() {
	RegisterEventMethod("contact.department.updated_v3", HandleDepartmentUpdate)
	RegisterEventMethod("contact.department.deleted_v3", HandleDepartmentDelete)
	RegisterEventMethod("contact.department.created_v3", HandleDepartmentCreate)
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
