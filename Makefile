gen:
	protoc --proto_path=proto proto/*.proto --go_out=framework/pb --go-grpc_out=framework/pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

clean:
	del framework\pb\*.go