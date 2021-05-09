package main

import (
	"fmt"

	"github.com/prayansh1996/cache/cache"
)

func main() {
	obj := cache.GetProtoObject()
	proto := cache.Protobuf(obj)
	msgpack := cache.MsgPackSnappy(obj)
	json := cache.JsonEncoding(obj)

	fmt.Println("Bytes used by Protobuf:", len(proto))
	fmt.Println("Bytes used by MsgPack serializer and Snappy encoding:", len(msgpack))
	fmt.Println("Bytes used by Json encoding:", len(json))

	savings := 1 - float64(len(proto))/float64(len(msgpack))
	fmt.Printf("\nPercentage savings from MsgPack and Snappy -> Protobuf: %.2f%%\n", 100*savings)
}
