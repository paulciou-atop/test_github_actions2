package main

import (

"fmt"
"time"
)

func main() {
	for {
		time.Sleep(0.5 * time.Second) 
		fmt.Println("hello world1")
	}

}
