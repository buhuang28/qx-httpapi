package api

import (
	"github.com/buhuang28/qx-go-sdk/qx"
	"github.com/gin-gonic/gin"
	"net/http"
	"qx-httpapi/dto"
)

type QXApi struct {
}

func (q *QXApi) GetBots(c *gin.Context) {
	resp := qx.GetBots()
	c.JSON(http.StatusOK, resp)
	return
}

func (q *QXApi) GetGroups(c *gin.Context) {
	req := new(dto.GetGroupReq)
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, dto.Fail(err.Error()))
		return
	}
	resp := qx.GetGroups(req.BotId, req.Type)
	c.JSON(http.StatusOK, resp)
	return
}

func (q *QXApi) GetGroupMembers(c *gin.Context) {
	req := new(dto.GetGroupMemberReq)
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, dto.Fail(err.Error()))
		return
	}
	resp := qx.GetMembers(req.BotId, req.GroupId)
	c.JSON(http.StatusOK, resp)
	return
}

func (q *QXApi) SendMessage(c *gin.Context) {
	req := new(dto.SendMessageReq)
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, dto.Fail(err.Error()))
		return
	}
	resp := qx.SendText(req.BotId, req.ToId, req.Msg)
	c.JSON(http.StatusOK, resp)
	return
}

func (q *QXApi) SendImg(c *gin.Context) {
	req := new(dto.SendImageReq)
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, dto.Fail(err.Error()))
		return
	}
	resp := qx.SendImg(req.BotId, req.ToId, req.Path)
	c.JSON(http.StatusOK, resp)
	return
}
