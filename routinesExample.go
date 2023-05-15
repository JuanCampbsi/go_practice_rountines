package main

import (
	"fmt"
	"time"
)

func main() {
	go cont("sem go routine")
	go cont("com go routine")

	time.Sleep(time.Second * 10)
	fmt.Println("Fim ...")
}

func cont(types string) {
	for i := 0; i < 5; i++ {
		fmt.Println(types, i)
		time.Sleep(time.Second)
	}
}
