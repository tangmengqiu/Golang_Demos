<center>GPRC-PROTOBUF</center>

# Intro:
well grpc-protobuf is a lib for process communication and it's a c/s model.

# How To Use:

firstly,you need to define what's the data you want to send and recive between client and server `message`,then define the remote call function `rpc` by `message`,and serveral `rpc` consist of `service` ,all in the .proto file

secondly,you need to install `protoc`,the proto compiler and `protoc-gen-go` to fit golang ,and you must make sure these two tools are in your PATH

finally,complie the .proto file: `protoc --go_out=. *.proto`,make sure the .proto files are uner the current dir.Then it will generate the xx.pb.go files,put  them both in client project and server project and implement the rpc method. Done! 