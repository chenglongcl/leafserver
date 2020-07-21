package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/name5566/leaf/gate"
	"leafserver/base"
	"leafserver/msg"
	"reflect"
	"time"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.CSBroadcastMultiple{}, handleBroadcastMultiple)
	handler(&msg.CSPing{}, handlePing)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleBroadcastMultiple(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.CSBroadcastMultiple)
	// 消息的发送者
	a := args[1].(gate.Agent)
	if player := PlayerManager.GetPlayer(a.UserData().(*base.AccountInfo).ConnUUID);
		player != nil {
		_ = player.BroadcastMultiple(&msg.SCBroadcastMultiple{
			MsgType: m.MsgType,
			Source:  player.Nickname,
			Text:    m.Text,
			Image:   m.Image,
			Video:   m.Video,
		}, PlayerManager.GetPlayersInRoom(player.RoomID))
	}
}

func handlePing(args []interface{}) {
	m := args[0].(*msg.CSPing)
	_ = m
	a := args[1].(gate.Agent)
	b, _ := proto.Marshal(&msg.SCPong{Timestamp: time.Now().Unix()})
	a.WriteMsg(&msg.SCResponse{
		Code:       0,
		Message:    "OK",
		MsgHeaders: &msg.MsgHeaders{MsgName: "SCPong"},
		Data:       &any.Any{Value: b},
	})
}
