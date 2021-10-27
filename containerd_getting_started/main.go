package main

import "fmt"
import "log"
import "github.com/containerd/containerd"

func main() {{
    if err := redisExample(); err != nil {
        log.Fatal(err)
    } // Close if statement.
} // Close main function block.

func redisExample() error {
    
  client, err := containerd.New("run/containerd/containerd.sock")
     
  if err != nil {
    return err
  } // Close if statement
   
  defer client.Close()
    
  return nil

} // Close redisExample function block.
