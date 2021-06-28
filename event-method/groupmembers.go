package eventmethod

func init() {
	RegisterEventMethod("im.chat.member.bot.added_v1", HandleAddBot)
	RegisterEventMethod("im.chat.member.bot.deleted_v1", HandleRemoveBot)
	RegisterEventMethod("im.chat.member.user.added_v1", HandleAddUser)
	RegisterEventMethod("im.chat.member.user.deleted_v1", HandleRemoveUser)
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
