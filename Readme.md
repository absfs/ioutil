# ioutil - AbsFs FileSystem Utility Functions

[![Go Reference](https://pkg.go.dev/badge/github.com/absfs/ioutil.svg)](https://pkg.go.dev/github.com/absfs/ioutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/absfs/ioutil)](https://goreportcard.com/report/github.com/absfs/ioutil)
[![CI](https://github.com/absfs/ioutil/actions/workflows/ci.yml/badge.svg)](https://github.com/absfs/ioutil/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

Package `github.com/absfs/ioutil` implements the standard library `ioutil` functions for the [absfs.FileSystem](https://github.com/absfs/absfs) interface, providing a familiar API for working with abstract filesystems.

## Features

- 🔄 Drop-in replacement for `io/ioutil` functions adapted for abstract filesystems
- 🧪 Fully tested across Linux, macOS, and Windows
- 📦 Zero dependencies beyond the absfs ecosystem
- ⚡ Efficient implementations derived from Go's standard library

## Installation

```bash
go get github.com/absfs/ioutil
```

## Usage

Import the package:

```go
import (
    "github.com/absfs/ioutil"
    "github.com/absfs/osfs"
)
```

## Examples

### Read Directory

```go
fs, _ := osfs.NewFS()

files, err := ioutil.ReadDir(fs, ".")
if err != nil {
    log.Fatal(err)
}

for _, file := range files {
    fmt.Println(file.Name())
}
```

### Read File

```go
fs, _ := osfs.NewFS()

content, err := ioutil.ReadFile(fs, "testdata/hello")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("File contents: %s", content)
```

### Write File

```go
fs, _ := osfs.NewFS()

data := []byte("Hello, World!")
err := ioutil.WriteFile(fs, "output.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}
```

### Temporary Directory

```go
fs, _ := osfs.NewFS()
content := []byte("temporary file's content")

dir, err := ioutil.TempDir(fs, "", "example")
if err != nil {
    log.Fatal(err)
}
defer fs.RemoveAll(dir) // clean up

tmpfn := filepath.Join(dir, "tmpfile")
if err := ioutil.WriteFile(fs, tmpfn, content, 0666); err != nil {
    log.Fatal(err)
}
```

### Temporary File

```go
fs, _ := osfs.NewFS()
content := []byte("temporary file's content")

tmpfile, err := ioutil.TempFile(fs, "", "example")
if err != nil {
    log.Fatal(err)
}
defer fs.Remove(tmpfile.Name()) // clean up

if _, err := tmpfile.Write(content); err != nil {
    log.Fatal(err)
}
if err := tmpfile.Close(); err != nil {
    log.Fatal(err)
}
```

## API Reference

The package provides the following functions:

- `ReadAll(r io.Reader) ([]byte, error)` - Read all data from a reader
- `ReadFile(fs absfs.FileSystem, filename string) ([]byte, error)` - Read entire file
- `WriteFile(fs absfs.FileSystem, filename string, data []byte, perm os.FileMode) error` - Write data to file
- `ReadDir(fs absfs.FileSystem, dirname string) ([]os.FileInfo, error)` - Read directory entries
- `TempFile(fs absfs.FileSystem, dir, prefix string) (absfs.File, error)` - Create temporary file
- `TempDir(fs absfs.FileSystem, dir, prefix string) (string, error)` - Create temporary directory
- `NopCloser(r io.Reader) io.ReadCloser` - Wrap reader with no-op closer
- `Discard` - io.Writer that discards all data

See the [full documentation](https://pkg.go.dev/github.com/absfs/ioutil) for detailed usage.

## absfs Ecosystem

This package is part of the absfs ecosystem:

- [absfs](https://github.com/absfs/absfs) - Abstract filesystem interface with composition features
- [osfs](https://github.com/absfs/osfs) - OS filesystem implementation
- [memfs](https://github.com/absfs/memfs) - In-memory filesystem implementation

Check out the [`absfs`](https://github.com/absfs/absfs) repo for more information about the abstract filesystem interface and features like filesystem composition.

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

Because this project is derived from the Go standard library, it is governed by the Go license, which is a BSD-style License. See [LICENSE](LICENSE) for the full license text.

Copyright (c) 2009 The Go Authors. All rights reserved.
