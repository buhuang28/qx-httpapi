package dto

type QXData struct {
	BotId string `json:"robot_wxid"`
	Event string `json:"event"`
	Data  string `json:"data"`
}

func NewQXData(botId, event, data string) QXData {
	return QXData{
		BotId: botId,
		Event: event,
		Data:  data,
	}
}
