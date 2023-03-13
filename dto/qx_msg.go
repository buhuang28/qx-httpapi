package dto

type SendMessageReq struct {
	BotId string `json:"robot_wxid"`
	ToId  string `json:"to_wxid"`
	Msg   string `json:"msg"`
}
