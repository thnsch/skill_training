
package main

import (
    "fmt"
)

func main() {
    var f float32
    var i int32

    fmt.Print("Enter a float number: ")
    fmt.Scan(&f)

    i = int32(f)

    fmt.Printf("\nTruncate float: %.2f \nTo integer: %d \n", f, i)
}
