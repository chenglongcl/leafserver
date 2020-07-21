package internal

import (
	"github.com/name5566/leaf/gate"
	"leafserver/base"
	"leafserver/msg"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("loginSuccess", rpcLoginSuccess)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	if userData := a.UserData(); userData != nil {
		accountInfo := a.UserData().(*base.AccountInfo)
		if player := PlayerManager.GetPlayer(accountInfo.ConnUUID); player != nil {
			PlayerManager.DelPlayer(player)
		}
	}
}

func rpcLoginSuccess(args []interface{}) {
	a := args[0].(gate.Agent)
	accountInfo := a.UserData().(*base.AccountInfo)
	player := NewPlayer(*accountInfo, a)
	PlayerManager.AddPlayer(player)
	player.CallClientMsg(nil, "SCUserLogin", &msg.SCUserLogin{
		RoomID:   player.RoomID,
		ConnUUID: player.ConnUUID,
	})
}
