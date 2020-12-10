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

func loadCipher (cipher string) map[string]string {
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
  var jsonCodificationData map[string]string
  err = json.Unmarshal(jsonCipher, &jsonCodificationData)
  if err != nil {
    fmt.Println(errorColor, "[FATAL ERROR]:      Can not Unmarshal the json cipher file\n")
    os.Exit(1)
  }
  fmt.Printf(noticeColor, "[CIPHER]:          ")
  fmt.Println("Loaded correctly")
  return jsonCodificationData
}
