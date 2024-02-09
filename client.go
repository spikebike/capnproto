package main

import (
    "net"
    // Import the generated Cap'n Proto structs 
    "github.com/your-project/schema" 
    "github.com/glycerine/go-capnproto2"
)

func main() {
    conn, _ := net.Dial("tcp", ":1234") 

    seg := capnp.NewBuffer(nil)
    err := capnp.NewDecoder(conn).Decode(seg)
    if err != nil {
        // Handle error
    }

    data := schema.ReadRootData(seg)
    matrix := data.Matrix() 

    // Process the received matrix
    for _, row := range matrix {
         for _, value := range row {
             // Do something with value
         }
    }
}
