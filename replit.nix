{ pkgs }: {
    deps = [
        pkgs.cowsay
        pkgs.go_1_18
        pkgs.protobuf
        pkgs.protoc-gen-go
        pkgs.protoc-gen-go-grpc
    ];
}