# install protobuf compiler


## download & install

'''
wget https://github.com/google/protobuf/releases/download/v3.11.2/protobuf-all-3.11.2.zip
unzip protobuf-all-3.11.2.zip && cd protobuf-3.11.2/
./configure
make
make install
'''
## check version

'''
ldconfig
protoc --version
'''

## go grpc install

'''
go get -u google.golang.org/grpc
'''

## Protoc Plugin

'''
go get -u github.com/golang/protobuf/protoc-gen-go
'''

# compile proto

'''
protoc --go_out=plugins=grpc:. ./proto/*.proto
'''

# build

## client

'''
go build ./client/unary_client.go
go build ./client/server_side_stream_client.go
go build ./client/client_side_stream_client.go
'''

## server

'''
go build ./server/unary_server.go
go build ./server/server_side_stream_server.go
go build ./server/client_side_stream_server.go
'''
