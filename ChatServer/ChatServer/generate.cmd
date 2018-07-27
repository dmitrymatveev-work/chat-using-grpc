protoc -I=../../protos chat.proto --csharp_out=. --grpc_out=. --plugin=protoc-gen-grpc=../packages/Grpc.Tools.1.13.1/tools/windows_x86/grpc_csharp_plugin.exe
