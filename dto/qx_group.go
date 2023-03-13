package dto

type GetGroupReq struct {
	BotId string `json:"robot_wxid"`
	Type  string `json:"type"`
}
