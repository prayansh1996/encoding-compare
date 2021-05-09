package cache

import (
	"testing"
)

func BenchmarkMsgPackAndSnappySerialization(b *testing.B) {
	obj := GetProtoObject()

	for n := 0; n < b.N; n++ {
		MsgPackSnappy(obj)
	}
}

func BenchmarkProtoSerialization(b *testing.B) {
	obj := GetProtoObject()

	for n := 0; n < b.N; n++ {
		Protobuf(obj)
	}
}
