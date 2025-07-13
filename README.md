# MisuDB

MisuDB is a distributed database engine designed for high performance and scalability. At its core, MisuDB leverages a skiplist-based architecture for in-memory key-value storage For persistent, on-disk storage, MisuDB utilizes Log-Structured Merge (LSM) trees. Consensus is ensured using the Raft protocol.

MisuDB engine components (kv and sql database) can be run independently using the command line executables in [cmd](./cmd/) folder

## Installation

You can install MisuDB by cloning the repository and running Goreleaser to build the binaries locally.

### Using Goreleaser

First, clone the repository:

```sh
git clone https://github.com/misudb/misudb.git
cd misudb
```

Then, build the binaries using [Goreleaser](https://goreleaser.com/):

```sh
# Run Goreleaser in local mode
goreleaser build --clean
```

The built binaries will be available in the `dist/` directory.

### Building from Source

Alternatively, you can build the executables directly with Go:

```sh
go install github.com/misudb/misudb/cmd/misudb@latest
go install github.com/misudb/misudb/cmd/misusql@latest
go install github.com/misudb/misudb/cmd/misukv@latest
```

## License

This project is licensed under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0).