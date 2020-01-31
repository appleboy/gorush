# gRPC in 3 minutes (Node.js)

## PREREQUISITES

`node`: This requires Node 12.x or greater.

## INSTALL

```sh
npm install
npm install -g grpc-tools
```

## Node gRPC protoc

```sh
cd $GOPATH/src/github.com/appleboy/gorush
protoc -I rpc/proto rpc/proto/gorush.proto --js_out=import_style=commonjs,binary:rpc/example/node/ --grpc_out=rpc/example/node/ --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin`
```
