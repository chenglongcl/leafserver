package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&CSPing{})
	Processor.Register(&SCResponse{})
	Processor.Register(&CSUserLogin{})
	Processor.Register(&CSBroadcastMultiple{})
}
