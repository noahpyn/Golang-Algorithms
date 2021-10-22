/* program to calculate how many items will be packaged in the last box of a shipment */

package main

import "fmt"

func main() {
    items := 1081
    
    fmt.Println(items % 7)
}
