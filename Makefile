gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go 

server:
	go run cmds/server/main.go -port 8085

client:
	go run cmds/client/main.go -address 0.0.0.0:8085

run:
	go run main.go
