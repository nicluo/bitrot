package main

import (
  "fmt"
  "hash/crc32"
  "io/ioutil"
  "os"
  "path/filepath"
  "regexp"
)

func ensureDirectories() {
  os.MkdirAll("./photos", 0777)
  os.MkdirAll("./out", 0777)
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func rot(testfile string) {
  testdir := ("./photos/")

  f, err := os.Open(testdir + testfile)
  check(err)

  info, err := f.Stat()
  check(err)

  b1 := make([]byte, info.Size())
  f.Read(b1)

  for i := 0; i < int(info.Size() / 1000); i++ {
    filename := fmt.Sprintf("./out/%d%s", i, filepath.Ext(testfile))
    w, err := os.Create(filename)
    check(err)

    _, err = w.Write(b1)
    check(err)

    _, err = w.WriteAt([]byte{0}, int64(i * 1000))
    check(err)

    err = w.Close()
    check(err)

    checksum(filename)
  }
}

func checksum(path string){
  f, err := os.Open(path)
  check(err)

  info, err := f.Stat()
  check(err)

  b1 := make([]byte, info.Size())
  n1, err := f.Read(b1)
  check(err)

  err = f.Close()
  check(err)

  n2 := crc32.ChecksumIEEE(b1)
  fmt.Printf("%s\n%d bytes: %d\n", path, n1, n2)
}

func main() {
  ensureDirectories()

  files, err := ioutil.ReadDir("./photos/")
  check(err)

  r, _ := regexp.Compile("(?i)(gif|png|jpe?g)$")

  for _, fileinfo := range files {
    if r.MatchString(fileinfo.Name()) {
      rot(fileinfo.Name())
    }
  }
}
