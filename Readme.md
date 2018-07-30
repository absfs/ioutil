# ioutil - AbsFs FileSystem Utility Functions
Package `github.com/absfs/ioutil` implements the standard library `ioutil`
functions for the absfs.FileSystem interface.

## Examples

Read directory.

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

Temp Directory.

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

Temp file.

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

Read file.

```go
  fs, _ := osfs.NewFS()

  content, err := ioutil.ReadFile(fs, "testdata/hello")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("File contents: %s", content)

  // Output:
  // File contents: Hello, Gophers!
```

## absfs
Check out the [`absfs`](https://github.com/absfs/absfs) repo for more information about the abstract filesystem interface and features like filesystem composition

## LICENSE

Because this project is derived from from the Go standard library it is governed
by the Go license which is a BSD-style License. See [LICENSE](https://github.com/absfs/ioutil/blob/master/LICENSE)



