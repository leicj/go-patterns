package main

import "time"
import "sync"
import "fmt"
import singleton1 "p9/singleton1"
import singleton2 "p9/singleton2"

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	for i:=0;i<2;i++{
		go singleton1.GetInstance(&wg)
	}
	time.Sleep(time.Second*1)
	fmt.Println()
	for i:=0;i<2;i++{
		go singleton2.GetInstance(&wg)
	}
	wg.Wait()
}