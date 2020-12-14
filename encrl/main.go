package main

import (
  "fmt"
)

const (
  infoColor    = "\033[1;34m%s\033[0m"
  noticeColor  = "\033[1;36m%s\033[0m"
  alertColor   = "\033[1;33m%s\033[0m"
  errorColor   = "\033[1;31m%s\033[0m"
  debugColor   = "\033[0;36m%s\033[0m"
)

func main() {
  encrlIcon := `
d88888b      d8b   db       .o88b.      d8888b.      db      
88           888o  88      d8P  Y8      88   8D      88      
88ooooo      88V8o 88      8P           88oobY'      88      
88~~~~~      88 V8o88      8b           88 8b        88      
88.          88  V888      Y8b  d8      88  88.      88booo. 
Y88888P      VP   V8P       Y88P'       88   YD      Y88888P 
`
  fmt.Printf(noticeColor, encrlIcon)
  info := `
Encrl - the simplest encryption tool created in Golang
More information at: https://github.com/pblcc/encrl
v[0.0.17] Alpha - by Pablo Corbalán (@pblcc)
  `
  fmt.Printf(noticeColor, info)
  readingFile, writingFile, codification := loadArguments()
  encrypt(loadCipher(codification), loadFile(readingFile))
  fmt.Sprintf(readingFile, writingFile)
}
