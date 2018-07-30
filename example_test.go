// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ioutil_test

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/absfs/ioutil"
	"github.com/absfs/osfs"
)

func ExampleReadAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

	// Output:
	// Go is a general-purpose language designed with systems programming in mind.
}

func ExampleReadDir() {
	fs, _ := osfs.NewFS()

	files, err := ioutil.ReadDir(fs, ".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func ExampleTempDir() {
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
}

func ExampleTempFile() {
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
}

func ExampleReadFile() {
	fs, _ := osfs.NewFS()

	content, err := ioutil.ReadFile(fs, "testdata/hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

	// Output:
	// File contents: Hello, Gophers!
}
