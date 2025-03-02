** Программа для понимания базы по gRPC:
1.Unary API
2.Client streaming
3.Server streaming
4.Bidirectional streaming


** Команда для генерации кода из proto:
```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```