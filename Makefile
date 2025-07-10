gen:
	protoc --proto_path=proto proto/*.proto --go_out=framework/pb --go-grpc_out=framework/pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

clean:
	del framework\pb\*.go

evans:
	docker run --rm -it -v "C:/Users/LUISFP/go/src/luisfp/plataforma-desafio:/mount:ro" ghcr.io/ktr0731/evans:latest --path ./proto/ --proto user_message.proto --host host.docker.internal --port 50051 repl