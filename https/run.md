## run server
go run server/main.go ./cert/root.pem ./cert/server.key ./cert/server.pem

## run client
go run client/main.go ./cert/root.pem ./cert/client.key ./cert/client.pem
