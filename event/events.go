package event

import (
	"github.com/buhuang28/qx-go-sdk/qx"
	"qx-httpapi/cst"
	"qx-httpapi/dto"
	"qx-httpapi/utils"
)

func GroupMsgEvent(botId, str string) int64 {
	_ = utils.HttpPostWithJson(cst.CALLBACK_API, dto.Success(dto.NewQXData(botId, cst.EventGroupMsg, str)))
	return qx.EVENT_CONTINUE
}

func PrivateMsgEvent(botId, str string) int64 {
	_ = utils.HttpPostWithJson(cst.CALLBACK_API, dto.Success(dto.NewQXData(botId, cst.EventPrivateMsg, str)))
	return qx.EVENT_CONTINUE
}

func GroupReqEvent(botId, str string) int64 {
	_ = utils.HttpPostWithJson(cst.CALLBACK_API, dto.Success(dto.NewQXData(botId, cst.EventGroupReq, str)))
	return qx.EVENT_CONTINUE
}

func CommonEvent(botId, str string) int64 {
	return qx.EVENT_CONTINUE
}

func CommonEvent2() int64 {
	return qx.EVENT_AGREE
}
