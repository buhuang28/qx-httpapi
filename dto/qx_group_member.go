package dto

type GetGroupMemberReq struct {
	BotId   string `json:"robot_wxid"`
	GroupId string `json:"wxid"`
}
