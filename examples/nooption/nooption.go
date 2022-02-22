package nooption

//go:generate go install ../../cmd/protoc-gen-nrpc
//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate protoc --go_out . --go_opt=paths=source_relative --nrpc_out . --nrpc_opt=paths=source_relative nooption.proto
