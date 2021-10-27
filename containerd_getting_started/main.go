package main

import "log"
import "github.com/containerd/containerd"

func main() {
    if err := redisExample(); err != nil {
        log.Fatal(err)
    } // Close if statement.
} // Close main function block.

func redisExample() error {
    client, err := containerd.New("/run/containerd/containerd.sock")
    
    if err != nil {
        return err
    }
    
    defer client.Close()

    ctx := namespace.WithNamespace(context.Background(), "example")
    image, err := client.Pull(ctx, "docker.io/libraryredis:alpine", containerd.WithPullUnpack)
    
    if err != nil {
        return err
    }
    
    log.Printf("Successfully pulled %s image\n", image.Name())
    
    return nil

} // Close redisExample function block.
