# encoding-compare
Compare MsgPack serialization and Snappy encoding with Protobuf encoding. This test checks the final length output of both the encodings, as well as benchmarks the performance of each.

Steps:
`make run`

You can add your own base.proto file and return an object in proto.go's GetProtoObject function.

Example run:
![alt text](https://github.com/prayansh1996/encoding-compare/blob/master/result.png)
