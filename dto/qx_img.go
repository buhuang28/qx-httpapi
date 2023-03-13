package dto

type SendImageReq struct {
	BotId string `json:"robot_wxid"`
	ToId  string `json:"to_wxid"`
	Path  string `json:"path"`
}
