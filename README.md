# ww

Command line implementation of "Are you a werewolf?"


## Usage

Compile protobuf and gRPC

```bash
protoc --go_out=plugins=grpc:. proto/*.proto
```

Compile binary 

```bash
go get -u
go build
```

Start server

```bash
./ww server
```

Connect clients to server

```bash
./ww client
```
