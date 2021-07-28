package clouddoc

import eventmethod "github.com/waro163/feishu_robot/event-method"

func init() {
	eventmethod.RegisterEventMethod("drive.file.read_v1", HandleFileReadEvent)
	eventmethod.RegisterEventMethod("drive.file.title_updated_v1", HandleFileTitleUpdateEvent)
	eventmethod.RegisterEventMethod("drive.file.permission_member_added_v1", HandleFileAddMemberEvent)
	eventmethod.RegisterEventMethod("drive.file.permission_member_removed_v1", HandleFileRemoveMemberEvent)
	eventmethod.RegisterEventMethod("drive.file.trashed_v1", HandleFileMoveTrashEvent)
	eventmethod.RegisterEventMethod("drive.file.deleted_v1", HandleFileDeleteEvent)
}

func HandleFileReadEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}

func HandleFileTitleUpdateEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}

func HandleFileAddMemberEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}

func HandleFileRemoveMemberEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}

func HandleFileMoveTrashEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}

func HandleFileDeleteEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}
