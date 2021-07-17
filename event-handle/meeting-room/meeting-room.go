package meetingroom

import eventmethod "github.com/waro163/feishu_robot/event-method"

func init() {
	eventmethod.RegisterEventMethod("meeting_room.created_v1", HandleMeetingRoomCreateEvent)
	eventmethod.RegisterEventMethod("meeting_room.updated_v1", HandleMeetingRoomUpdateEvent)
	eventmethod.RegisterEventMethod("meeting_room.deleted_v1", HandleMeetingRoomDeleteEvent)
	eventmethod.RegisterEventMethod("meeting_room.status_changed_v1", HandleMeetingRoomStatusChangeEvent)
}

func HandleMeetingRoomCreateEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleMeetingRoomUpdateEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleMeetingRoomDeleteEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleMeetingRoomStatusChangeEvent(map[string]string, map[string]interface{}) error {
	return nil
}
