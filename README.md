# firehose-sui
Firehose on Sui Blockchain

## Re-generate Protobuf Definitions

1. Ensure that `protoc` is installed:
   ```
   brew install protoc
   ```

2. Ensure that `protoc-gen-go` and `protoc-gen-go-grpc` are installed and at the correct version
    ```
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

3. Generate go file from modified protobuf

   ```bash
   ./types/pb/generate.sh
  ```
