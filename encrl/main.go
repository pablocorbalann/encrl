package main

import (
  "fmt"
  "os"
  "bufio"
)
const noticeColor string = "\033[1;36m%s\033[0m"

func main() {
  versionFile, err := os.Open("../version.txt") 
  if err != nil {
    fmt.Println("[FATAL ERROR]:     Can not open the version file")
  }
  defer versionFile.Close()
  scanner := bufio.NewScanner(versionFile)
  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  err = scanner.Err()
  if err != nil {
    fmt.Println("[FATAL ERROR]:     Can not read the version file using a scanner")
  }
  // format the version itself
  version := fmt.Sprintf("%v.%v.%v", lines[0], lines[1], lines[2])
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
  `, version)
  fmt.Printf(noticeColor, info)
  readingFile, writingFile, codification := loadArguments()
  e := encrypt(loadCipher(codification), loadFile(readingFile))
  dump(writingFile, e)
}
