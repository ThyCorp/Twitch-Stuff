package main

import (
  "os"
  "bufio"
  "fmt"
)

//Finds Password and Adds it To Pass .txt File for Twitch Bot
func PassFinder() {
  filename := "twitch_pass.txt"
  in := bufio.NewReader(os.Stdin)
  fmt.Println("Enter Pass")
  u1, err := in.ReadString('\n')
  if err != nil {
    fmt.Println("Somthing Went Wrong Rerun")
    os.Exit(1)
  }
  file, err := os.Create(filename)
  if err != nil {
    fmt.Println("error At File Create")
    os.Exit(1)
  }
  file.WriteString(u1)
}
