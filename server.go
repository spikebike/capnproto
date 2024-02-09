package main

import (
    "net"
    // Import the generated Cap'n Proto structs 
    "github.com/your-project/schema" 
    "github.com/glycerine/go-capnproto2"
)

func main() {
    listener, _ := net.Listen("tcp", ":1234") // Listen on some port
    conn, _ := listener.Accept()

    seg := capnp.NewBuffer(nil)
    data := schema.NewRootData(seg)

    // Create your [][]byte data (Replace with your real data)
    matrix := [][]byte{
        {1, 2, 3},
        {4, 5, 6},
    }
    data.SetMatrix(matrix) 

    // Encode the Cap'n Proto message
    err := capnp.NewEncoder(conn).Encode(seg)
    if err != nil {
        // Handle error
    }
}
