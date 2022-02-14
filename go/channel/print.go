package main

import (
	"fmt"
	"sync"
)

func main() {

	// 使⽤两个 goroutine 交替打印序列，⼀个 goroutine 打印数字， 另外⼀个 goroutine 打印字⺟， 最终效果 如下
	// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

	// 考点1 为 channel 间的通信
	// 考点2 为 主 channel 对于 子 channel 的控制

	number, char := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	go printNumber(number, char, &wg)
	wg.Add(1)
	go printChar(number, char)

	number <- true
	wg.Wait()
}

func printNumber(number, char chan bool, wg *sync.WaitGroup) {
	i := 1
	for {
		select {
		case <-number:
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			if i >= 28 {
				wg.Done()
				return
			}
			char <- true
		}
	}
}

func printChar(number, char chan bool) {
	i := 'A'
	for {
		select {
		case <-char:
			fmt.Print(string(i))
			i++
			fmt.Print(string(i))
			i++
			number <- true
		}
	}
}
