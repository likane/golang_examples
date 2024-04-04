package main

import (
	"fmt"
	"os"
)

package main

import (
 "encoding/json"
 "fmt"
 "os"
)

type ChargeSession struct {
 Id        string `json:"id"`
 Watts     int    `json:"watts"`
 Vin       string `json:"vin"`
 Timestamp string `json:"timestamp"`
}

func main() {
 filepath := "session.json"
 var cs ChargeSession

 file, err := os.Open(filepath)
 if err != nil {
  fmt.Println(err)
 }
 defer file.Close()

 decoder := json.NewDecoder(file)
 err = decoder.Decode(&cs)
 if err != nil {
  fmt.Println(err)
 }
 fmt.Println(cs)
}
