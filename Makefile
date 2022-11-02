# generate graphql
gql:
	go run github.com/99designs/gqlgen generate

# generate protos
gnpu:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative app/grpc/protos/user.proto

# generate protos product
gnpp:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative app/grpc/protos/product.proto

# generate protos auth
gnpa:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative app/grpc/protos/auth.proto

# run go
r:
	go run main.go

create con:
	docker container create --name server-main --network grpc -p 8080:8080 server-main