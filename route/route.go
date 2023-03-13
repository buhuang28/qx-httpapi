package route

import (
	"github.com/buhuang28/qx-go-sdk/qx"
	"github.com/gin-gonic/gin"
	"qx-httpapi/api"
)

func RunWebApi(port string) int64 {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	qxApi := api.QXApi{}
	router.POST("/api/SendMessage", qxApi.SendMessage)
	router.POST("/api/SendImg", qxApi.SendImg)
	router.POST("/api/GetGroups", qxApi.GetGroups)
	router.POST("/api/GetBots", qxApi.GetBots)
	router.POST("/api/GetGroupMembers", qxApi.GetGroupMembers)
	go router.Run(port)
	return qx.EVENT_AGREE
}
