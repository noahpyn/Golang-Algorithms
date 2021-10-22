// Slices used to determine what days are weekdays. 

package main

import (
	"fmt"
)

func main() {
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekdays := days[0:5]
	fmt.Println(weekdays)
	// This returns: [Monday Tuesday Wednesday Thursday Friday]
}
