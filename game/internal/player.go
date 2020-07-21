package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/name5566/leaf/gate"
	"leafserver/base"
	"leafserver/base/errno"
	"leafserver/msg"
)

type Player struct {
	base.AccountInfo
	agent gate.Agent
}

func NewPlayer(info base.AccountInfo, agent gate.Agent) *Player {
	return &Player{
		AccountInfo: info,
		agent:       agent,
	}
}

// Broadcast broadcasts a text message to all sessions.
func (a *Player) Broadcast() error {
	return nil
}

// BroadcastFilter broadcasts a text message to all sessions that fn returns true for.
func (a *Player) BroadcastFilter() error {
	return nil
}

// BroadcastMultiple broadcasts a text message to multiple sessions given in the sessions slice.
func (a *Player) BroadcastMultiple(data proto.Message, players []*Player) error {
	for _, p := range players {
		p.CallClientMsg(nil, "SCBroadcastMultiple", data)
	}
	return nil
}

func (a *Player) CallClientMsg(err error, msgName string, data proto.Message) {
	code, message := errno.DecodeErr(err)
	b, _ := proto.Marshal(data)
	a.agent.WriteMsg(&msg.SCResponse{
		Code:       int32(code),
		Message:    message,
		MsgHeaders: &msg.MsgHeaders{MsgName: msgName},
		Data:       &any.Any{Value: b},
	})
}
