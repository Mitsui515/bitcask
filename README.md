# Bitcask KV

Bitcask KV is a high-performance Key-Value storage engine based on the Bitcask model, implemented in Go.

## Features

- High-performance read and write operations 
- Support for multiple index types: B-tree, ART, and B+-tree 
- Transaction support 
- Data persistence 
- Support for Merge operations 
- Iterator support 
- Support for multiple file I/O types: standard file I/O and memory mapping

## Main Components

1. **Data file management**: implements persistent storage and reading of data.
2. **Index**: supports multiple index structures to improve query efficiency.
3. **Logging**: implements encoding and decoding of log records.
4. **Iterator**: supports forward and reverse traversal of data.
5. **Merge operations**: implements the merging of data files to optimise storage space.

## Basic Operation

```go
package main

import "github.com/Mitsui515/bitcask"

func main() {
    // specify the options
    options := bitcask.DefaultOptions
    options.DirPath = "/tmp/bitcask"

    // open a database
    db, err := bitcask.Open(options)
    if err != nil {
        panic(err)
    }
    defer func() {
        _ = db.Close()
    }()

	// set a key
	err = db.Put([]byte("name"), []byte("bitcask"))
	if err != nil {
		panic(err)
	}

	// get a key
	val, err := db.Get([]byte("name"))
	if err != nil {
		panic(err)
	}
	println(string(val))

	// delete a key
	err = db.Delete([]byte("name"))
	if err != nil {
		panic(err)
	}
}
```


## Benchmark

The project contains benchmarking tests, which can be performed by running the following command:

```go
go test -bench=. ./benchmark
```

