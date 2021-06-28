package eventmethod

func init() {
	RegisterEventMethod("contact.user.created_v3", HandleStaffInductionEvent)
	RegisterEventMethod("contact.user.deleted_v3", HandleStaffOffEvent)
	RegisterEventMethod("contact.user.updated_v3", HandleStaffContactUpdateEvent)
}

func HandleStaffInductionEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleStaffOffEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleStaffContactUpdateEvent(map[string]string, map[string]interface{}) error {
	return nil
}
