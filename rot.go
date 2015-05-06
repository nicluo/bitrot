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

  b1, err := ioutil.ReadFile(testdir + testfile)
  check(err)

  n1 := len(b1)

  for i := 0; i < int(n1 / 1000); i++ {
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
  b1, err := ioutil.ReadFile(path)
  check(err)

  n1 := len(b1)
  n2 := crc32.ChecksumIEEE(b1)
  fmt.Printf("%s\n%d bytes: %d\n", path, n1, n2)
}

func matchImageExtension(filename string) bool {
  r, _ := regexp.Compile("(?i)(gif|png|jpe?g)$")

  return r.MatchString(filename)
}

func main() {
  ensureDirectories()

  files, err := ioutil.ReadDir("./photos/")
  check(err)

  for _, fileinfo := range files {
    if matchImageExtension(fileinfo.Name()) {
      rot(fileinfo.Name())
    }
  }
}
