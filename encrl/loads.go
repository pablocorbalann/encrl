/*
* In this file we store all the configuration for the codifications, as well
* as the functions that load the codifications table from a json file
*
* All the codifications tables are inside the codifications/ file. For example
* codifications/caesar.json
*/
package main 

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
)

type JsonCipherData struct {
  CharacterA string `json:"a"` 
  CharacterB string `json:"b"` 
  CharacterC string `json:"c"` 
  CharacterD string `json:"d"` 
  CharacterE string `json:"e"` 
  CharacterF string `json:"f"` 
  CharacterG string `json:"g"` 
  CharacterH string `json:"h"` 
  CharacterI string `json:"i"` 
  CharacterJ string `json:"j"` 
  CharacterK string `json:"k"` 
  CharacterL string `json:"l"` 
  CharacterM string `json:"m"` 
  CharacterN string `json:"n"` 
  CharacterO string `json:"o"` 
  CharacterP string `json:"p"` 
  CharacterQ string `json:"q"` 
  CharacterR string `json:"r"` 
  CharacterS string `json:"s"` 
  CharacterT string `json:"t"` 
  CharacterU string `json:"u"` 
  CharacterV string `json:"v"` 
  CharacterW string `json:"w"` 
  CharacterX string `json:"x"` 
  CharacterY string `json:"y"` 
  CharacterZ string `json:"z"` 
  Character0 string `json:"0"` 
  Character1 string `json:"1"` 
  Character2 string `json:"2"` 
  Character3 string `json:"3"` 
  Character4 string `json:"4"` 
  Character5 string `json:"5"` 
  Character6 string `json:"6"` 
  Character7 string `json:"7"` 
  Character8 string `json:"8"` 
  Character9 string `json:"9"` 
}

func loadCipher (cipher string) JsonCipherData {
  /*
  * This function loads a cipher from the codifications/ folder to then return it.
  * The default cipher that is loaded is the caesar cipher.
  */
  cipherToLoad, pathErr := filepath.Abs(fmt.Sprintf("../codifications/%s.json", cipher))
  if pathErr != nil {
    fmt.Printf(errorColor, "[FATAL ERROR]:      Can not get the absolute path of codification\n")
    os.Exit(1)
  }
  jsonCipher, err := ioutil.ReadFile(cipherToLoad)
  if err != nil {
    fmt.Printf(errorColor, "[FATAL ERROR]:      Can not load the json cipher file\n")
    os.Exit(1)
  }
  // Inform the user about the loading cipher
  fmt.Printf("[CIPHER]:          %s\n", cipherToLoad)
  // Parse the json file to a cipher
  var jsonCodificationData JsonCipherData
  err = json.Unmarshal(jsonCipher, &jsonCodificationData)
  if err != nil {
    fmt.Println(errorColor, "[FATAL ERROR]:      Can not Unmarshal the json cipher file\n")
    os.Exit(1)
  }
  fmt.Printf(noticeColor, "[CIPHER]:          ")
  fmt.Println("Loaded correctly")
  return jsonCodificationData
}
