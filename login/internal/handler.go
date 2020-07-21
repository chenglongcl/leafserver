package internal

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"google.golang.org/grpc/metadata"
	"leafserver/base"
	"leafserver/base/errno"
	"leafserver/game"
	"leafserver/msg"
	"leafserver/util"
	"leafserver/util/snowflake"
	"reflect"
)

func init() {
	handleMsg(&msg.CSUserLogin{}, handleUserLogin)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleUserLogin(args []interface{}) {
	// 收到的消息
	m := args[0].(*msg.CSUserLogin)
	// 消息的发送者
	a := args[1].(gate.Agent)
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("p-id", "ws-"))
	uuid := snowflake.NextID(ctx)
	accountInfo := &base.AccountInfo{
		RoomID:   m.RoomID,
		ConnUUID: uuid,
		Nickname: "nickname-" + util.UUID(),
	}
	a.SetUserData(accountInfo)
	skeleton.AsynCall(game.ChanRPC, "loginSuccess", a, func(err error) {
		if nil != err {
			log.Error("login failed:", accountInfo.ConnUUID, " ", err.Error())
			done(a, errno.LoginFail, "SCUserLogin", nil)
			return
		}
	})
}

func done(agent gate.Agent, err error, msgName string, data proto.Message) {
	code, message := errno.DecodeErr(err)
	b, _ := proto.Marshal(data)
	agent.WriteMsg(&msg.SCResponse{
		Code:       int32(code),
		Message:    message,
		MsgHeaders: &msg.MsgHeaders{MsgName: msgName},
		Data:       &any.Any{Value: b},
	})
}
