package models

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/set"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

// Message 用户信息
type Message struct {
	gorm.Model
	From    int64  // 发送用户id
	To      int64  // 接收用户id
	MsgType int    // 聊天类型 1-私聊 2-群聊 3-广播
	Media   int    // 消息类型 1-文本 2-图片 3-语音 4-视频
	Content string // 消息内容
	Pic     string // 图片内容
	Url     string // 链接内容
	Desc    string // 消息描述
	Amount  int    // 其他数字内容 发送频率|文件大小等
}

func (table *Message) TableName() string {
	return "t_msg"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

//clientMap 客户端映射
var clientMap = make(map[int64]*Node, 0)

// cmLocker 客户端映射map读写锁
var cmLocker sync.RWMutex

func init() {
	go udpSendMsg()
	go udpRecvMsg()
}

/*
	Chat 实现
	需要: 发送者ID 接收者ID 消息类型-文本|图片|语音|视频 消息内容 聊天类型-私聊|群聊|广播
*/
func Chat(ctx *gin.Context) {
	// 1-获取参数 并校验token合法性
	//token := ctx.Query("token")
	// 校验token需要与数据库中的token做对比 最好可以封装个方法 checkToken todo
	isValid := true // 暂定为true
	id := ctx.Query("uid")
	uid, _ := strconv.ParseInt(id, 10, 64)
	//recvId := ctx.Query("recvId")
	//media := ctx.Query("media")
	//content := ctx.Query("content")
	//msgType := ctx.Query("msgType")
	// 2-创建链接
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValid
		},
	}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 3-初始化一个Node
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 4-用户关系
	// 5-uid与node绑定并加锁
	cmLocker.Lock()
	clientMap[uid] = node
	cmLocker.Unlock()
	// 6-完成发送逻辑
	go sendMsg(node)
	// 7-完成接收逻辑
	go recvMsg(node)
	sendPrivateMsg(uid, []byte("欢迎进入聊天系统"))
}

// sendMsg 发送消息逻辑
func sendMsg(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// recvMsg 接收消息逻辑
func recvMsg(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data) // 广播到局域网
		fmt.Println("[ws] <<<<<", data)
	}
}

var udpSendChan = make(chan []byte)

func broadMsg(data []byte) {
	udpSendChan <- data
}

// udpSendMsg 完成udp数据发送流程
func udpSendMsg() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 7, 114),
		Port: 5000,
	})
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpSendChan:
			_, err = conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// udpRecvMsg 完成udp数据接收流程
func udpRecvMsg() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(192, 168, 7, 114),
		Port: 5000,
	})
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		//var buf [512]byte
		var buf []byte
		n, err := conn.Read(buf) // read 方法中往往都会传一个 []byte 用来接收消息数据的 刚开始看到这边还有点懵，怎么莫名其妙要传个参数呢?
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

// dispatch 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println("json解析失败|err=", err)
		return
	}
	switch msg.MsgType {
	case 1:
		sendPrivateMsg(msg.From, data) // 单聊
	//case 2:sendGroupMsg() // 群聊
	//case 3: sendAllMsg() // 广播消息
	case 4:
	}
}

// sendPrivateMsg 发送私聊消息
func sendPrivateMsg(from int64, msg []byte) {
	cmLocker.RLock()
	node, ok := clientMap[from]
	cmLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
