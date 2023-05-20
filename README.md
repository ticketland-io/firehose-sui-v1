# firehose-sui
Firehose on Sui Blockchain

## Pre-requisites
You would need to download the `sui-sf-indexer` source code and build it locally.

```bash
git clone https://github.com/ticketland-io/sui-sf-indexer
cd sui-sf-indexer
cargo build --release
```
Make sure the `target/release/sui-sf-indexer` binary is moved to a folder that is in the PATH so it can be executed.

> Binaries will be available to download soon so you don't have to build from source code.

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

4. Run Firehose

```bash
./devel/mainnet/start.sh -c
```

And ensure blocks are flowing:

```
grpcurl -plaintext  -import-path ./types/proto -proto sui.proto -import-path ../proto -proto sf/firehose/v2/firehose.proto -d '{"start_block_num": 0}' localhost:18015 sf.firehose.v2.Stream.Blocks
```
