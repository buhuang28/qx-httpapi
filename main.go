package main

import (
	"encoding/json"
	"github.com/buhuang28/qx-go-sdk/qx"
	"qx-httpapi/event"
	"qx-httpapi/route"
)

type QXInfo struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Skey        string `json:"skey"`
	Sdk         string `json:"sdk"`
}

func InitQXInfo() string {
	info := QXInfo{
		Name:        "httpapi",
		Author:      "不慌",
		Description: "不慌的httpapi",
		Version:     "1.1",
		Skey:        "QXASDOG5SD7A8SD",
		Sdk:         "S1",
	}
	marshal, _ := json.Marshal(info)
	return string(marshal)
}

func init() {
	qx.Info = InitQXInfo
	qx.EventGroupMsg = event.GroupMsgEvent
	qx.EventPrivateMsg = event.PrivateMsgEvent
	qx.EventGroupReq = event.GroupReqEvent
	qx.EventEnable = func() int64 {
		return route.RunWebApi(":10066")
	}
	qx.EventDisable = event.CommonEvent2
	qx.EventUninstall = event.CommonEvent2
	qx.EventFriendReq = event.CommonEvent
	qx.EventPluginIns = func(i int64) int64 {
		return qx.EVENT_CONTINUE
	}
	qx.EventQrcPay = event.CommonEvent
	qx.EventRecall = event.CommonEvent
	qx.EventSendMsg = event.CommonEvent
	qx.EventTransfer = event.CommonEvent
	qx.EventWeChat = event.CommonEvent
	qx.Set = func() {
	}

}

func main() {
	//此处留空，别写东西
}
