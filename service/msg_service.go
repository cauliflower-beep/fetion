package service

import (
	"fetion/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upGrade = websocket.Upgrader{
	/*
		防止跨域站点伪造请求 CSRF(cross site request forgery)
		简单来讲，CSRF通过伪装来自受信任用户的请求来利用受信任的网站
		https://blog.csdn.net/weixin_52851967/article/details/125992627
	*/
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SendMsg
// @Summary 发送消息
// @Tags 消息模块
// @Success 200 {string} json{"code","message"}
// @Router /msg/sendMsg [get]
// 发送消息
func SendMsg(ctx *gin.Context) {
	// 将http协议升级为websocket
	ws, err := upGrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 退出前关闭ws连接
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	// 消息发送流程
	MsgHandler(ws, ctx)
}

func MsgHandler(ws *websocket.Conn, ctx *gin.Context) {
	msg, err := utils.Subscribe(ctx, utils.PublishKey)
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s", now, msg)
	fmt.Println(m)
	err = ws.WriteMessage(websocket.TextMessage, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}
