package message

import eventmethod "github.com/waro163/feishu_robot/event-method"

func init() {
	eventmethod.RegisterEventMethod("im.chat.member.bot.added_v1", HandleAddBot)
	eventmethod.RegisterEventMethod("im.chat.member.bot.deleted_v1", HandleRemoveBot)
	eventmethod.RegisterEventMethod("im.chat.member.user.added_v1", HandleAddUser)
	eventmethod.RegisterEventMethod("im.chat.member.user.deleted_v1", HandleRemoveUser)
	eventmethod.RegisterEventMethod("im.chat.member.user.withdrawn_v1", HandleWithdrawnUser)
	eventmethod.RegisterEventMethod("p2p_chat_create", HandleChatFirstCreateUser)
}

func HandleAddBot(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleRemoveBot(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleAddUser(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleRemoveUser(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleWithdrawnUser(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleChatFirstCreateUser(map[string]string, map[string]interface{}) error {
	return nil
}
