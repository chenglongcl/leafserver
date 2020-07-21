package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"io"
	"leafserver/msg"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	data, err := msg.Processor.Marshal(&msg.CSPing{})

	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// len + id + data
	m := make([]byte, 4+len(data[1]))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(2+len(data[1])))

	copy(m[2:], data[0])
	copy(m[4:], data[1])

	go func() {
		for {
			time.Sleep(1 * time.Second)
			// 发送消息
			conn.Write(m)
		}
	}()

	for {
		fmt.Println("<------")
		headData := make([]byte, 4)
		_, _ = io.ReadFull(conn, headData)
		//创建一个从输入二进制数据的ioReader
		dataBuff := bytes.NewReader(headData)
		var a, b uint16
		//读Len
		if err := binary.Read(dataBuff, binary.BigEndian, &a); err != nil {
			fmt.Println(err)
			return
		}
		//读msgid
		if err := binary.Read(dataBuff, binary.BigEndian, &b); err != nil {
			fmt.Println(err)
			return
		}
		msgData := make([]byte, a-2)
		_, err = io.ReadFull(conn, msgData)
		if err != nil {
			fmt.Println(err)
		}
		response := &msg.SCResponse{}
		_ = proto.Unmarshal(msgData, response)
		fmt.Println(response.GetMsgHeaders().GetMsgName())
		scPong := &msg.SCPong{}
		_ = proto.Unmarshal(response.GetData().Value, scPong)
		fmt.Println(scPong.GetTimestamp())
	}

}
