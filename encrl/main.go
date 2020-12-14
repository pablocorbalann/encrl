package main

import (
  "fmt"
  "os"
  "io/ioutil"
)
const noticeColor string = "\033[1;36m%s\033[0m"

func main() {
  versionFile, err := os.Open("../version") 
  if err != nil {
    fmt.Println("[FATAL ERROR]:     Can not open the version file")
  }
  version, err:= ioutil.ReadAll(versionFile)
  // Strip the \n at the end of the file
  version = version[:len(version) - 1]
  if err != nil {
    fmt.Println("[FATAL ERROR]:     Can not read the version file using io")
  }
  encrlIcon := `
d88888b      d8b   db       .o88b.      d8888b.      db      
88           888o  88      d8P  Y8      88   8D      88      
88ooooo      88V8o 88      8P           88oobY'      88      
88~~~~~      88 V8o88      8b           88 8b        88      
88.          88  V888      Y8b  d8      88  88.      88booo. 
Y88888P      VP   V8P       Y88P'       88   YD      Y88888P 
`
  fmt.Printf(noticeColor, encrlIcon)
  info := fmt.Sprintf(`
Encrl - the simplest encryption tool created in Golang
More information at: https://github.com/pblcc/encrl
v[%s] Alpha - by Pablo Corbalán (@pblcc)
  `, string(version))
  fmt.Printf(noticeColor, info)
  readingFile, writingFile, codification := loadArguments()
  encrypt(loadCipher(codification), loadFile(readingFile))
  fmt.Sprintf(readingFile, writingFile)
}
