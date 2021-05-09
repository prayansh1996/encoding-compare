package cache

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/vmihailenco/msgpack"
)

func Protobuf(obj proto.Message) []byte {
	bytes, _ := proto.Marshal(obj)
	return bytes
}

func MsgPackSnappy(obj proto.Message) []byte {
	bytes, _ := msgpack.Marshal(obj)
	compressed := snappy.Encode(nil, bytes)
	return compressed
}

func JsonEncoding(obj proto.Message) []byte {
	var buffer bytes.Buffer
	m := jsonpb.Marshaler{}
	m.Marshal(&buffer, obj)
	return buffer.Bytes()
}
