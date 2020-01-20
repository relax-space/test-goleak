package main

import (
	"fmt"
	"time"
)

func Succ() {
	go func() {
		fmt.Println("hello1")
		time.Sleep(1 * time.Second)
		fmt.Println("hello2")
	}()
	time.Sleep(2 * time.Second)
}

func Fail() {
	go func() {
		fmt.Println("hello1")
		time.Sleep(1 * time.Minute)
		fmt.Println("hello2")
	}()
	time.Sleep(2 * time.Second)
}
