package helpdisk

import eventmethod "github.com/waro163/feishu_robot/event-method"

func init() {
	// 服务台-工单事件
	eventmethod.RegisterEventMethod("helpdesk.ticket.created_v1", HandleTicketCreateEvent)
	eventmethod.RegisterEventMethod("helpdesk.ticket.updated_v1", HandleTicketUpdateEvent)
	eventmethod.RegisterEventMethod("helpdesk.ticket_message.created_v1", HandleTicketMsgCreateEvent)
}

func HandleTicketCreateEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleTicketUpdateEvent(map[string]string, map[string]interface{}) error {
	return nil
}

func HandleTicketMsgCreateEvent(map[string]string, map[string]interface{}) error {
	return nil
}
