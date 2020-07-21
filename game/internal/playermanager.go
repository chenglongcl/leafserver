package internal

import (
	"sync"
)

type PlayerMgr struct {
	PlayerMap       map[string]*Player
	PlayerMapInRoom map[int32][]*Player
	Lock            *sync.RWMutex
}

func NewPlayerManager() *PlayerMgr {
	return &PlayerMgr{
		PlayerMap:       make(map[string]*Player),
		PlayerMapInRoom: make(map[int32][]*Player),
		Lock:            new(sync.RWMutex),
	}
}

func (a *PlayerMgr) AddPlayer(player *Player) {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	a.PlayerMap[player.ConnUUID] = player
	a.PlayerMapInRoom[player.RoomID] = append(a.PlayerMapInRoom[player.RoomID], player)
}

func (a *PlayerMgr) GetPlayer(connUUID string) *Player {
	a.Lock.RLock()
	defer a.Lock.RUnlock()
	if _, ok := a.PlayerMap[connUUID]; !ok {
		return nil
	}
	return a.PlayerMap[connUUID]
}

func (a *PlayerMgr) GetPlayerInRoom(roomID int32, connUUID string) *Player {
	a.Lock.RLock()
	defer a.Lock.RUnlock()
	if players := a.GetPlayersInRoom(roomID); len(players) > 0 {
		for _, p := range players {
			if p.ConnUUID == connUUID {
				return p
			}
		}
	}
	return nil
}

func (a *PlayerMgr) GetPlayersInRoom(roomID int32) []*Player {
	a.Lock.RLock()
	defer a.Lock.RUnlock()
	if _, ok := a.PlayerMapInRoom[roomID]; !ok {
		return nil
	}
	return a.PlayerMapInRoom[roomID]
}

func (a *PlayerMgr) DelPlayer(player *Player) {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	delete(a.PlayerMap, player.ConnUUID)
	playersInRoom := a.PlayerMapInRoom[player.RoomID]
	for k, p := range playersInRoom {
		if player == p {
			a.PlayerMapInRoom[player.RoomID] =
				sliceRemovePlayer(a.PlayerMapInRoom[player.RoomID], k)
			if len(a.PlayerMapInRoom[player.RoomID]) == 0 {
				delete(a.PlayerMapInRoom, player.RoomID)
			}
			break
		}
	}
}

func sliceRemovePlayer(s []*Player, index int) []*Player {
	return append(s[:index], s[index+1:]...)
}
