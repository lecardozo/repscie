package helper

import (
   "io/ioutil"
   "log"
)


// Receives the file name and
func ReadConfigFile(configFile string) string {
    bytes, err := ioutil.ReadFile(configFile)
    if err != nil {
        log.Fatal(err)
    }
    return string(bytes)
}
