package filefuncs

import (
  "os"
)


// Exists Finds If File Exsits
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
                return false
            }
    }
    return true
}
