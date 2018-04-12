package main

import "fmt"
import "time"

const epoch = 1234567890

func main() {
	fmt.Println(time.Unix(epoch, 0).Format("02-01-2006"))
}
