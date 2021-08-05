package main

import (
	"fmt"
	"time"
)

func main() {
	future := time.Unix(9223372036854755807, 0) //64bit time will last a really long time
	fmt.Println(future)
}
